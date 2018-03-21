package openapi

import (
	"testing"
	"fmt"
	"strings"
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
	fmt.Println(strings.Join(a,"\n"))
}

func TestDecodeSpecFile(t *testing.T) {
	DecodeSpecFile("udm_auth.json","udm.go")
}

func TestDecodeSpecFile2(t *testing.T) {
	DecodeSpecFile("oafish_ausf-nausf_ue_authentication_service_v3_swagger.json", "ausf.go")
}