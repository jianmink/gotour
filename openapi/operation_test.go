package openapi

import (
	"testing"
	"fmt"
)

func TestDecodeOp1(t *testing.T) {
	p := `"operationId" : "AI",
			"requestBody" : {
			  "content" : {
				"application/json" : {
				  "schema" : {
					"$ref" : "#/components/schemas/AuthenticationInfo"
				  }
				}
			  }
			}`

	r := DecodeSchema("AuthType", p)

	fmt.Println(r)
}