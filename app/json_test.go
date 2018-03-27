package app

import (
	"testing"
	"time"
	"os"
	"encoding/json"
	"fmt"
)


type Address1 struct {
	Road string   `json:"road"`
	Building string `json:"building,omitempty"`
}

type MyUser struct {
	ID       int64     `json:"id"`
	Name     string    `json:"name,omitempty"`
	Addr	 Address1  `json:"address"`
	LastSeen time.Time `json:"lastSeen,omitempty"`
}

func TestJsonMismatch(t *testing.T) {
	byt := []byte(`{
				"id":3,
				"name":"mike",
				"address": {"road": "road-1", "building": "B1"}
				}`)

	var u2 = MyUser{}
	err := json.Unmarshal(byt, &u2)
	if err !=nil {
		fmt.Println(err.Error())
	}

	fmt.Println(u2)
	//fmt.Println(u2.Addr)
}

func TestJsonMarsha0(t *testing.T) {
	enc := json.NewEncoder(os.Stdout)
	enc.Encode(
		&MyUser{ID: 1, LastSeen: time.Now()})

	u1 := &MyUser{
		ID:       1,
		LastSeen: time.Now(),
	}

	json.Marshal(u1)
	//fmt.Println(string(b1))
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
