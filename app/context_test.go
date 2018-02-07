package app

import (
	"context"
	"net/http"
	"reflect"
	"testing"
)

func Test_newContextWithRequestID(t *testing.T) {
	type args struct {
		ctx context.Context
		req *http.Request
	}
	tests := []struct {
		name string
		args args
		want context.Context
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newContextWithRequestID(tt.args.ctx, tt.args.req); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newContextWithRequestID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_requestIDFromContext(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name string
		args args
		want string
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := requestIDFromContext(tt.args.ctx); got != tt.want {
				t.Errorf("requestIDFromContext() = %v, want %v", got, tt.want)
			}
		})
	}
}
