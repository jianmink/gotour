package openapi

import "fmt"

import (
	"encoding/json"
	"io/ioutil"
	"strings"
	"github.com/ghodss/yaml"
	"path/filepath"
	"errors"
	"os"
)

// yaml2json translate bytes from yaml format  to json format
func yaml2json(y []byte) []byte {
	j, err := yaml.YAMLToJSON(y)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return nil
	}

	return j
}

// DecodeSpecFile() input is the spec file, target is the dest go file
//		input:
//			openapi spec file of json or yaml format
//		output:
//		  	.../<package_name>/struct.go
//		  	package_name shall contain the service name and api version, e.g. ausfv1
func DecodeSpecFile(srcFile, destFolder, serviceName, apiVersion string) error {
	b,err := ioutil.ReadFile(srcFile)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	// support both .json and .yaml/.yml
	switch ext:=filepath.Ext(srcFile); ext {
	case ".json":
		break
	case ".yaml":
		b = yaml2json(b)
	case ".yml":
		b = yaml2json(b)
	default:
		err := errors.New(fmt.Sprintf("Unsupport file extension %v", ext))
		fmt.Println(err.Error())
		return err
	}

	a,err := DecodeSpec(string(b))
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	var a1 []string
	for _, v := range a {
		a1 = append(a1, v.String())
	}

	// output to target file
	df := destFolder + "/" + serviceName + apiVersion

	if df != "" {
		tmp := []string{"package "+ serviceName +apiVersion}
		a2 := append (tmp, a1...)

		//create folder

		os.MkdirAll(df,0777)
		return ioutil.WriteFile(df+"/struct.go", []byte(strings.Join(a2, "\n")), 0644)
	}

	return nil
}

//  DecodeSpec() decode openapi spec in json string
//
func DecodeSpec(s string) ([]OpenApiStruct, error) {

	// paths and components are must in openAPI spec 3.0
	type Spec struct {
		Paths map[string]*json.RawMessage `json:"paths,omitempty"`
		Components struct {
			Schemas map[string]*json.RawMessage `json:"schemas"`
		} `json:"components"`
	}

	var data Spec
	err := json.Unmarshal([]byte(s), &data)
	if err != nil {
		return nil, err
	}

	//for k, v := range data.Paths {
	//
	//	fmt.Printf("path[%v] \n", k)
	//
	//	// decode v which is an mapping from operation {post, get, put, ...} to a structure
	//	//DecodeJsonMap(*v) //marshall path item
	//
	//}

	var a []OpenApiStruct
	for k, v := range data.Components.Schemas {
		r := DecodeSchema(k, string(*v))
		a = append(a,r)
	}

	// set IsStruct for each field
	for i, t1 := range a {
		for j,f1 := range t1.Fields {
			for _, t2 := range a {
				if f1.Type == t2.Name && t2.Fields != nil {
					a[i].Fields[j].IsStructType = true
				}
			}
		}
	}

	// set IsArray for each field
	for i, t1 := range a {
		for j,f1 := range t1.Fields {
			for _, t2 := range a {
				if f1.Type == t2.Name && t2.IsArrayType {
					a[i].Fields[j].IsArray = true
				}
			}
		}
	}

	return a, err
}

