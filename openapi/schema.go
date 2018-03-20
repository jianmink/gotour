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
	d,err := DecodeJsonMap([]byte(p))
	if err != nil {
		fmt.Println(err.Error())
	}


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
	return "\n"+goStruct(name, fields, data.Required)+"\n"
}

func decodeType (s string) (error,string){
	var data map[string]*json.RawMessage
	err := json.Unmarshal([]byte(s), &data)
	if err != nil {
		fmt.Printf("error: %v\n", err.Error())
		return err, ""
	}

	for k, v := range data {
		tmp := string(*v)
		tmp = strings.Trim(tmp, "\"")
		if k == "type" {
			switch tmp {
			case "string":
				return nil, "string"
			default:
				return nil, tmp
			}
		}

		if k == "$ref" {
			s := strings.Split(tmp,"/")
			return nil, s[len(s)-1]
		}

	}

	return nil, ""
}


func goStruct (name string, fields map[string]string, required []string) string {
	t := fmt.Sprintf("type %v struct { \n", name)

	for k, v := range fields {

		if has(k, required) {
			s := fmt.Sprintf("\t%v\t%v\t `json:\"%v\"`\n", upper(k), v, k)
			t += s
		} else {
			s := fmt.Sprintf("\t%v\t%v\t `json:\"%v,omitempty\"`\n", upper(k), v, k)
			t += s
		}
	}

	t += "\n}"
	return t
}


func upper(s string) string {
	return strings.ToUpper(s[0:1]) + s[1:]
}

func has(str string, list []string) bool {
	for _, v := range list {
		if v == str {
			return true
		}
	}
	return false
}