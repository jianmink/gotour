package app

import (
	"testing"
	"fmt"
)

func TestMapLiteral(t *testing.T) {
	var wantBody = map[string]interface{}{
		"UEId": map[string]string{
			"UeType":     "SUPI",
			"UeIdentity": "123",
		},
		"AnthenticationType": "5G-AKA",
	}

	fmt.Println(wantBody)
}
