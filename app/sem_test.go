package app

import "testing"

func Test_process(t *testing.T) {
	type args struct {
		r string
	}
	tests := []struct {
		name string
		args args
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			process(tt.args.r)
		})
	}
}

func TestServe(t *testing.T) {
	type args struct {
		queue chan string
	}
	tests := []struct {
		name string
		args args
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Serve(tt.args.queue)
		})
	}
}
