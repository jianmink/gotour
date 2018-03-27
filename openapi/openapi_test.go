package openapi

import (
	"testing"
	//"fmt"
	//"strings"
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
			}
		}
	}

}`

func TestSpec(t *testing.T) {
	a, err := DecodeSpec(specJSON)
	if err != nil {
		t.Error(err.Error())
	}
	for _, v := range a {
		fmt.Println(v.String())
	}
}

func TestDecodeSpecFile(t *testing.T) {
	DecodeSpecFile("udm_auth.json",
		"output", "udm", "v1")
}

func TestDecodeSpecFile2(t *testing.T) {
	DecodeSpecFile("oafish_ausf-nausf_ue_authentication_service_v3_swagger.json",
		"output", "ausf", "v1")
}


func TestDecodeSpecFile3(t *testing.T) {
	DecodeSpecFile("Nrf-nf_management_service_v1_swagger_version_04_pa3.yaml",
		"output", "nrf", "v1")
}
