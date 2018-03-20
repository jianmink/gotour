package app

import (
	"testing"
	"time"
	"os"
	"encoding/json"
	"fmt"
)


type Address struct {
	road string
	building string
}

type MyUser struct {
	ID       int64     `json:"id"`
	Name     string    `json:"name,omitempty"`
	Addr	 *Address  `json:"address,omitempty"`
	LastSeen time.Time `json:"lastSeen"`
}

func TestJsonMarshal(t *testing.T) {
	enc := json.NewEncoder(os.Stdout)
	enc.Encode(
		&MyUser{ID: 1, LastSeen:time.Now()})

	u1 := &MyUser{
		ID: 1,
		LastSeen:time.Now(),
	}

	b1,_ := json.Marshal(u1)
	fmt.Println(string(b1))

	//dec := json.NewDecoder(os.Stdin)
	//dec.Decode()


	byt := []byte(`{"ID":2,"Name":"mike"}`)
	//var dat map[string]interface{}
	var u2 = MyUser{}
	json.Unmarshal(byt, &u2)

	fmt.Println(u2)
}



func TestOpenapi3(t *testing.T) {

	var specJSON = `{
	  "openapi" : "3.0.0",
	  "info" : {
		"version" : "v3",
		"title" : "AUSF Nausf_UEAuthentication Service",
		"description" : "AUSF Nausf_UEAuthentication Service"
	  },
	"servers" : [ {
		"description" : "SwaggerHub API Auto Mocking",
		"url" : "https://virtserver.swaggerhub.com/EriCT4/nausf-ueauth/v1"
	  } ],
	  "paths" : {
		"/ue-authentications" : {
		  "post" : {
			"summary" : "Authentication Initiation Request",
			"operationId" : "AuthenticationId",
			"requestBody" : {
			  "content" : {
				"application/json" : {
				  "schema" : {
					"$ref" : "#/components/schemas/AuthenticationInfo"
				  }
				}
			  }
			},
			"responses" : {
			  "201" : {
				"description" : "successful operation",
				"content" : {
				  "application/json" : {
					"schema" : {
					  "$ref" : "#/components/schemas/UEAuthenticationCtx"
					}
				  }
				}
			  }
			}
		  }
		}
	}
}`


	var objmap map[string]*json.RawMessage
	err := json.Unmarshal([]byte(specJSON), &objmap)
	if err != nil {
		t.Errorf("error: %v", err.Error())
	} else {
		t.Logf("%v", string(*objmap["paths"]))
	}

}
