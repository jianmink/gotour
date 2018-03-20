package openapi

import "fmt"

import (
	"encoding/json"
	)

func DecodeSpec(s string) error {
	// paths and components are must in openAPI spec 3.0
	type Spec struct {
		Paths map[string]*json.RawMessage `json:"paths"`
		Components struct {
			Schemas map[string]*json.RawMessage `json:"schemas"`
		} `json:"components"`
	}

	var data Spec
	err := json.Unmarshal([]byte(s), &data)
	if err != nil {
		return err
	}

	//for k, v := range data.Paths {
	//
	//	fmt.Printf("path[%v] \n", k)
	//
	//	// decode v which is an mapping from operation {post, get, put, ...} to a structure
	//	//DecodeJsonMap(*v) //marshall path item
	//
	//}

	for k, v := range data.Components.Schemas {
		//fmt.Printf("component[%v] \n", k)

		//DecodeJsonMap(*v)

		// decode schema
		DecodeSchema(k, string(*v))

	}


	return nil
}


func decodeOperation (op,s string) error {

	fmt.Printf("op(%v), s(%v)", op, s)

	var data map[string]*json.RawMessage
	err := json.Unmarshal([]byte(s), &data)
	if err != nil {
		fmt.Errorf("error: %v", err.Error())
		return err
	}

	fmt.Printf("operationId (%v)", string(*data["operationId"]))

	//decode requestBody
	//rb := data["requestBody"]
	//
	//type Fields struct {
	//	Content struct {
	//		Format struct {
	//			Schema string `json:"schema"`
	//		} `json:"application/json"`
	//	}  `json:"content"`
	//}

	//decode responses

	return nil
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

