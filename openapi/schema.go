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
	IsStructType bool
	IsMust	bool
	IsArray bool
}

type TJsonStruct struct{
	Name string

	Fields []TJsonField  //for type xxx struct

	Type string   //for built-in type
}


func IsStuctBlob(b string ) bool {
	if strings.ContainsAny(b, "properties") {
		return true
	}

	return false
}

func DecodeSchema (name string, p string) TJsonStruct {
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
		// field or struct
		if IsStuctBlob(p) {
			return decodeArrayStruct(name, p)
		} else {
			fmt.Println("a standalone schema shal be an field")
		}

	default:
	}

	r := TJsonStruct{
		Name: renameType(name),
		Type: GoType[t],
	}

	return r
}


func (t *TJsonStruct)String()string{
	if t.Type != "" {
		return fmt.Sprintf("\ntype %v %v \n", t.Name, t.Type)
	}

	var str string
	if t.Fields != nil {
		str = fmt.Sprintf("type %v struct { \n", t.Name)

		for _,f := range t.Fields {
			theType := f.Type


			if f.IsStructType && !f.IsMust && !f.IsArray {
				theType = "*" + theType
			}

			if f.IsArray {
				theType = "[]"+theType
			}


			str += fmt.Sprintf(
				"\t%-20v%-20v %v\n", f.Object, theType, f.Tag)

		}

		str += "\n}"
	}

	return str
}

func decodeArrayStruct(name,p string) TJsonStruct{
	var data struct {
		Type       string `json:"type"`
		Items json.RawMessage `json:"items"`
	}

	err := json.Unmarshal([]byte(p), &data)
	if err != nil {
		fmt.Printf("error: %v\n", err.Error())
		return TJsonStruct{}
	}

	r := DecodeSchema(name, string(data.Items))

	typesNew[name]=data.Type

	return r
	//return r + "\n" + goArray(name, name)
}

func decodeStruct(name string, p string) TJsonStruct {
	var data struct {
		Type       string `json:"type"`
		Properties map[string] *json.RawMessage `json:"properties"`
		Required   []string `json:"required"`
	}
	err := json.Unmarshal([]byte(p), &data)
	if err != nil {
		fmt.Printf("error: %v\n", err.Error())
		return TJsonStruct{}
	}

	var fields = make(map[string]TJsonField)
	for k,v := range data.Properties {
		//t,isArray := decodeFieldType(string(*v))
		t,isArray := decodeFieldType(string(*v))

		theObject := renameObject(k)
		theType := renameType(t)
		theJsonObject := k

		var f = TJsonField{
			Object:       theObject,
			Type:         theType,
			Tag:          "",// fmt.Sprintf("`json:\"%v\"`", theJsonObject),
			IsStructType: true,
			IsMust:       false,
			IsArray: isArray,
		}

		//fmt.Println(f)
		if _,ok := JsonType[theType]; ok {
			f.IsStructType = false
		}

		if has(theJsonObject, data.Required) {
			f.IsMust = true
			f.Tag = fmt.Sprintf("`json:\"%v\"`", theJsonObject)
		} else {
			f.IsMust = false
			f.Tag = fmt.Sprintf("`json:\"%v,omitempty\"`", theJsonObject)
		}

		f.IsArray = isArray

		fields[k] = f
	}

	return goStruct2(name, fields)


}

func decodeFieldType (s string) (string, bool){
	var data map[string]*json.RawMessage
	err := json.Unmarshal([]byte(s), &data)
	if err != nil {
		fmt.Printf("error: %v\n", err.Error())
		return "", false
	}

	isArray := false
	for k, v := range data {
		tmp := string(*v)
		if k == "type" {
			tmp = strings.Trim(tmp, "\"")
			if t, ok := GoType[tmp]; ok {
				if t != "array" {
					return t, false
				} else {
					isArray = true
					continue
				}
			} else {
				return "unknown", false
			}
		}

		if k == "items" {
			if isArray {
				t0,_ :=decodeFieldType(tmp)
				// set the type as array
				return t0,true
			}
		}

		if k == "$ref" {
			tmp = strings.Trim(tmp, "\"")
			s := strings.Split(tmp,"/")
			return s[len(s)-1], false
		}

	}

	return "", false
}

func goArray (name string, t string) string {

	theObject := renameObject(name)
	theType := renameType(t)
	theJsonObject := name

	r := fmt.Sprintf("\t%v\t[]%v\t `json:\"%v,omitempty\"`\n", theObject, theType, theJsonObject)

	return r
}

func goStruct2 (name string, fields map[string]TJsonField ) TJsonStruct {
	var t TJsonStruct

	t.Name = renameType(name)

	for _, v := range fields {
		t.Fields = append(t.Fields, v)
	}

	return t
}

//func goStruct (name string, fields map[string]string, required []string) string {
//	t := fmt.Sprintf("type %v struct { \n", renameType(name))
//
//	for k, v := range fields {
//
//		theObject := renameObject(k)
//		theType := renameType(v)
//		theJsonObject := k
//
//		if has(k, required) {
//			s := fmt.Sprintf("\t%v\t%v\t `json:\"%v\"`\n", theObject, theType, theJsonObject)
//			t += s
//		} else {
//			s := fmt.Sprintf("\t%v\t%v\t `json:\"%v,omitempty\"`\n", theObject, theType, theJsonObject)
//			t += s
//		}
//	}
//
//	t += "\n}"
//	return t
//}


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