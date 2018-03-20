package app

import (
	"testing"
	"fmt"
)

type Writer interface {
	Write()
}

type MyWrite struct {
	a int
}

func (m *MyWrite) Write() {
	fmt.Println("my write")
}

func New() Writer {
	return &MyWrite{}
}



func TestNilInterface(t *testing.T) {
	var err error
	if err != nil {
		t.Errorf("not nil got %v", err)
	}
}

func TestInterface(t *testing.T) {
	m := MyWrite{}
	var w Writer = &m
	w.Write()

	a := w.(*MyWrite)
	fmt.Println(a.a)
}



type API interface {
	//server      string // url to the target host
	//path        string // end point
	//method      string // operation object, e.g. "post"
	//operationid string // AIR/ACR/EAP
	//req RequestBody
	//rsp Response

	SetOperationID(string)
	GetOperationID() string

	//GetRequestBody() RequestBody
	//SetRequestBody(RequestBody)
	//
	//GetResponse() Response
	//SetResponse(Response)
}


func NewAPI(t string) API {
	if t == "HTTP" {
		return &APIImpHttp{}
	}

	return nil
}


type APIImpHttp struct {
	server      string // url to the target host
	path        string // end point
	method      string // operation object, e.g. "post"
	operationID string // AIR/ACR/EAP

	//req RequestBody
	//rsp Response

	//eap *EapSession
}

//func (api *APIImpHttp) GetRequestBody() RequestBody {
//	return api.req
//}
//
//func (api *APIImpHttp) SetRequestBody(v RequestBody) {
//	api.req = v
//}
//
//func (api *APIImpHttp) GetResponse() Response {
//	return api.rsp
//}
//
//func (api *APIImpHttp) SetResponse( v Response) {
//	api.rsp = v
//}

func (api *APIImpHttp) SetOperationID(id string) {
	api.operationID = id
}

func (api *APIImpHttp) GetOperationID() string {
	return api.operationID
}
