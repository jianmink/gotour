package app

import (
	//"reflect"
	"testing"
	"fmt"
)

func Test_adder(t *testing.T) {
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}

	//tests := []struct {
	//	name string
	//	want func(int) int
	//}{
	//// TODO: Add test cases.
	//	{"case1", adder()},
	//}
	//for _, tt := range tests {
	//	t.Run(tt.name, func(t *testing.T) {
	//		if got := adder(); !reflect.DeepEqual(got, tt.want) {
	//			t.Errorf("adder() = %v, want %v", got, tt.want)
	//		}
	//	})
	//}
}
