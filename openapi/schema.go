package openapi

import (
	"fmt"
	"encoding/json"
	"strings"
)


var GoType  = map[string]string{
	"string": "string",
}

func DecodeSchema (name string, p string) string {
	//fmt.Printf("schema %v\n", name)
	d,_ := DecodeJsonMap([]byte(p))

	t := strings.Trim(string(*d["type"])," \"")
	switch t {
	case "object":
		return decodeStructSchema(name, p)

	default:
		//fmt.Println("trivial schema")
	}

	r := fmt.Sprintf("\ntype %v %v", name, GoType[t])

	return r
}


func decodeStructSchema(name string, p string) string {
	var data struct {
		Type       string `json:"type"`
		Properties map[string] *json.RawMessage `json:"properties"`
		Required   []string `json:"required"`
	}
	err := json.Unmarshal([]byte(p), &data)
	if err != nil {
		fmt.Printf("error: %v\n", err.Error())
		return ""
	}

	var fields = make(map[string]string)
	for k,v := range data.Properties {
		err, t := decodeType(string(*v))
		if err != nil {
			break
		}
		//fmt.Printf("k,v  (%v,%v)",k,t)
		fields[k] = t
	}

	//fmt.Printf("\n%v\n", goStruct(name, fields))
	return "\n"+goStruct(name, fields)+"\n"
}

func decodeType (s string) (error,string){
	var data map[string]*json.RawMessage
	err := json.Unmarshal([]byte(s), &data)
	if err != nil {
		fmt.Printf("error: %v\n", err.Error())
		return err, ""
	}

	t := string(*data["type"])

	switch t{
	case "string":
		return nil, "string"
	default:
		return nil, t
	}

	return nil, ""
}


func goStruct (name string, fields map[string]string) string {
	t := fmt.Sprintf("type %v struct { \n", name)

	for k, v := range fields {
		t += "\t"+k + "\t" + v
	}

	t += "\n}"
	return t
}
