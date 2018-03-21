package openapi

import (
	"testing"
	"fmt"
)

func TestDecodeSchema1(t *testing.T) {
	p := `{
		"type" : "object",
		"properties" : {
			"authType" : {
			"type" : "string",
			"enum" : [ "5G-AKA", "EAP-AKA-PRIME" ]
			}
		},
		"required" : [ "authType" ]
	}`

	r := DecodeSchema("AuthType", p)

	fmt.Println(r)
}


func TestDecodeSchema2(t *testing.T) {
	p := `{
			"type" : "string",
			"format" : "byte"
		  }`

	r := DecodeSchema("EapPayload", p)

	want := "\ntype EapPayload string"

	if r != want {
		t.Errorf("want %v, get %v", want, r)
	}
}

func TestDecodeSchema3(t *testing.T) {
	p := `{
		"type" : "object",
		"properties" : {
			"servingNetworkName" : { 
				"type" : "string"
			},
			"accessType" : {
			  "$ref" : "#/components/schemas/AccessType"
			}
		},
		"required" : [ "accessType" ]
	}`

	r := DecodeSchema("AuthenticationInfoRequest", p)

	fmt.Println(r)
}


func TestDecodeSchema4(t *testing.T) {
	p := `{
        "required": [
          "type"
        ],
        "properties": {
          "type": {
            "type": "string"
          },
          "title": {
            "type": "string"
          },
          "status": {
            "type": "integer"
          },
          "detail": {
            "type": "string"
          },
          "instance": {
            "type": "string"
          }
        }
      }`

	r := DecodeSchema("ProblemDetails", p)

	fmt.Println(r)
}

func TestDecodeSchema5(t *testing.T) {
	p := `{
        "type": "array",
        "items": {
          "type": "object",
          "properties": {
            "rand": {
              "$ref": "#/components/schemas/Rand"
            }
		}}}`

	r := DecodeSchema("Av5GAka", p)

	fmt.Println(r)
}