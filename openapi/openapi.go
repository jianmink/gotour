package openapi

import "fmt"

import (
	"encoding/json"
	"io/ioutil"
	"strings"
	"github.com/ghodss/yaml"
	"path/filepath"
	"errors"
)


func yaml2json(y []byte) []byte {
	j, err := yaml.YAMLToJSON(y)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return nil
	}

	return j
}

func DecodeSpecFile(sf string, df string) error {
	b,err := ioutil.ReadFile(sf)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	// support both .json and .yaml/.yml
	switch ext:=filepath.Ext(sf); ext {
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

	if df != "" {
		tmp := []string{"package openapi"}
		a2 := append (tmp, a1...)
		return ioutil.WriteFile(df, []byte(strings.Join(a2, "\n")), 0644)
	}
	//
	//fmt.Println(strings.Join(a, "\n"))
	return nil
}

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

	// set OpenApiStruct.IsStruct
	for i, t1 := range a {
		for j,f1 := range t1.Fields {
			for _, t2 := range a {
				if f1.Type == t2.Name && t2.Fields != nil {
					a[i].Fields[j].IsStructType = true
				}
			}
		}
	}

	return a, err
}

func addArrayMark(a string, t string ) string {


	return ""
}


func DecodeJsonMap(v  []byte) (map[string]*json.RawMessage, error){
	var data map[string]*json.RawMessage
	err := json.Unmarshal(v, &data)
	if err != nil {
		fmt.Errorf("error: %v", err.Error())
		return nil, err
	} else {
		//for k, v := range data {
		//	fmt.Printf("key[%v] value[%v]\n", k, string(*v))
		//}
	}

	return data, nil
}

//func DecodeProperty(name string, data map[string]*json.RawMessage) Property {
//	var p = Property{Name: name }
//	if v,ok := data["$ref"]; ok {
//		s := strings.Split(string(*v),"\\")
//		p.Type = s[len(s)-1]
//	}
//	return p
//}
//
//
//type Property struct {
//	Name string
//	Type string
//}

//type Schema struct {
//	Name string
//	Type string
//	Properties Property
//	Required []string
//}

