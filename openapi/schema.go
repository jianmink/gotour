package openapi

import (
	"fmt"
	"encoding/json"
	"strings"
)

var GoType  = map[string]string{
	"string": "string",
	"integer": "int",
	"boolean": "bool",
	"array":   "array",   //todo, not support array yet
	"object":  "struct",
	"number":  "float",   //todo, not support number yet
}

var JsonType = map[string]string {
	"string": "string",
	"int":"integer",
	"bool":"boolean",
	"struct":"object",
}

// "xxx: array"
var typesNew = map[string]string {

}

type TJsonField struct{
	Object string
	Type string
	Tag string
}

type TJsonStruct struct{
	Name string
	JsonField []TJsonField
}

func DecodeSchema (name string, p string) string {
	//fmt.Printf("schema %v\n", name)
	d,err := DecodeJsonMap([]byte(p))
	if err != nil {
		fmt.Println(err.Error())
	}

	t := ""
	if v,ok := d["type"]; ok {
		t = strings.Trim(string(*v), " \"")
	} else {
		t = "object"
	}

	switch t {
	case "object":
		return decodeStruct(name, p)
	case "array":
		return decodeArray(name, p)
	default:
	}

	r := fmt.Sprintf("\ntype %v %v", renameType(name), GoType[t])

	return r
}


func decodeArray(name,p string) string{
	var data struct {
		Type       string `json:"type"`
		Items json.RawMessage `json:"items"`
	}

	err := json.Unmarshal([]byte(p), &data)
	if err != nil {
		fmt.Printf("error: %v\n", err.Error())
		return ""
	}

	r := DecodeSchema(name, string(data.Items))

	typesNew[name]=data.Type

	return r
	//return r + "\n" + goArray(name, name)
}

func decodeStruct(name string, p string) string {
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
		t := decodeType(string(*v))

		fields[k] = t
	}

	return "\n"+goStruct(name, fields, data.Required)+"\n"
}

func decodeType (s string) (string){
	var data map[string]*json.RawMessage
	err := json.Unmarshal([]byte(s), &data)
	if err != nil {
		fmt.Printf("error: %v\n", err.Error())
		return ""
	}

	for k, v := range data {
		tmp := string(*v)
		tmp = strings.Trim(tmp, "\"")
		if k == "type" {
			if t, ok := GoType[tmp]; ok {
				return t
			} else {
				return "unknown"
			}
		}

		if k == "$ref" {
			s := strings.Split(tmp,"/")
			return s[len(s)-1]
		}

	}

	return ""
}

func goArray (name string, t string) string {

	theObject := renameObject(name)
	theType := renameType(t)
	theJsonObject := name

	r := fmt.Sprintf("\t%v\t[]%v\t `json:\"%v,omitempty\"`\n", theObject, theType, theJsonObject)

	return r
}

func goStruct (name string, fields map[string]string, required []string) string {
	t := fmt.Sprintf("type %v struct { \n", renameType(name))

	for k, v := range fields {

		theObject := renameObject(k)
		theType := renameType(v)
		theJsonObject := k

		if has(k, required) {
			s := fmt.Sprintf("\t%v\t%v\t `json:\"%v\"`\n", theObject, theType, theJsonObject)
			t += s
		} else {
			s := fmt.Sprintf("\t%v\t%v\t `json:\"%v,omitempty\"`\n", theObject, theType, theJsonObject)
			t += s
		}
	}

	t += "\n}"
	return t
}

func renameObject(s string) string {

	s = strings.Trim(s," ")
	if _,ok := JsonType[s]; ok {
		return s
	}

	if _,ok := GoType[s]; ok {
		return s
	}

	r := strings.ToUpper(s[0:1]) + s[1:]

	d := r[0]
	if d>='0' && d<='9' {
		r = "A" + r
	}

	return r
}

func renameType(s string) string {

	s = strings.Trim(s," ")
	if _,ok := JsonType[s]; ok {
		return s
	}

	if _,ok := GoType[s]; ok {
		return s
	}

	return "T_" + s
}


func has(s string, list []string) bool {
	s = strings.Trim(s, " ")
	for _, v := range list {
		if v == s {
			return true
		}
	}
	return false
}