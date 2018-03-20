package openapi

import (
	"testing"
	//"encoding/json"
	//"fmt"
	"fmt"
)

var specJSON = `{
	 "paths" : {
		"/ue-authentications" : {
			"post" : {
				"summary" : "Authentication Initiation Request",
				"operationId" : "AI",
				"requestBody" : {
				  "content" : {
					"application/json" : {
					  "schema" : {
						"$ref" : "#/components/schemas/AuthenticationInfo"
					  }
					}
				  }
				}
			}
		},
		"/ue-authentications/{authCtxId}/5g-aka-confirmation" : {
		}
	},
	"components" : {
		"schemas" : {
		  "AuthType" : {
			"type" : "object",
			"properties" : {
			  "authType" : {
				"type" : "string",
				"enum" : [ "5G-AKA", "EAP-AKA-PRIME" ]
			  }
			},
			"required" : [ "authType" ]
		  },
			"EapPayload" : {
					"type" : "string",
					"format" : "byte"
				  },
		}
	}

}`

func TestSpec(t *testing.T) {
	err := DecodeSpec(specJSON)
	if err != nil {
		t.Error(err.Error())
	}
}

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
		}
		"required" : [ "accessType" ]
	}`

	r := DecodeSchema("AuthenticationInfoRequest", p)

	fmt.Println(r)
}
