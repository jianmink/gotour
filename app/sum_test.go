package app

import "testing"

func TestSum(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name  string
		args  args
		wantR int
	}{
		{ "1+1=2", args{1,1}, 2},
		{ "1+2=3", args{1,2}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotR := Sum(tt.args.a, tt.args.b); gotR != tt.wantR {
				t.Errorf("Sum() = %v, want %v", gotR, tt.wantR)
			}
		})
	}
}
