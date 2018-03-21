package openapi

import "fmt"

import (
	"encoding/json"
	"io/ioutil"
	"strings"
	)

func DecodeSpecFile(sf string, df string) error {

	b,err := ioutil.ReadFile(sf)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	a,err := DecodeSpec(string(b))

	if df != "" {
		tmp := []string{"package openapi"}
		a = append (tmp, a...)
		return ioutil.WriteFile(df, []byte(strings.Join(a, "\n")), 0644)
	}

	fmt.Println(strings.Join(a, "\n"))
	return nil
}

func DecodeSpec(s string) ([]string, error) {

	// reset types create from the spec
	typesNew = map[string]string {

	}


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

	var a []string
	for k, v := range data.Components.Schemas {
		r := DecodeSchema(k, string(*v))
		a = append(a,r)
		//fmt.Println(r)
	}

	// add [] for ref of array type


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

