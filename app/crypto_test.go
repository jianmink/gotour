package app

import (
	"reflect"
	"testing"
	"fmt"
)



// HMAC_SHA256("", "") = b613679a0814d9ec772f95d778c35fc5ff1697c493715653c6c712144292c5ad
// HMAC_SHA256("key", "The quick brown fox jumps over the lazy dog") = f7bc83f430538424b13298e6aa6fb143ef4d59a14946175997479dbc2d1a3cd8
func format(s string){
	for i:= 0; i< len(s); i+=2 {
		fmt.Printf("0x%c%c,", s[i], s[i+1])
	}
}


func TestHMACSHA256(t *testing.T) {
	type args struct {
		key []byte
		s   []byte
	}

	//format("f7bc83f430538424b13298e6aa6fb143ef4d59a14946175997479dbc2d1a3cd8")

	tests := []struct {
		name string
		args args
		want []byte
	}{
		// TODO: Add test cases.
		{"empty", args{[]byte(""), []byte("")}, []byte{
			0xb6, 0x13, 0x67, 0x9a, 0x08, 0x14, 0xd9, 0xec, 0x77, 0x2f, 0x95, 0xd7, 0x78, 0xc3,0x5f, 0xc5,
			0xff, 0x16, 0x97, 0xc4, 0x93, 0x71, 0x56, 0x53, 0xc6, 0xc7, 0x12, 0x14, 0x42, 0x92, 0xc5, 0xad},
		},
		{"lazy dog", args{[]byte("key"), []byte("The quick brown fox jumps over the lazy dog")}, []byte{
			0xf7,0xbc,0x83,0xf4,0x30,0x53,0x84,0x24,0xb1,0x32,0x98,0xe6,0xaa,0x6f,0xb1,0x43,
			0xef,0x4d,0x59, 0xa1,0x49,0x46,0x17,0x59,0x97,0x47,0x9d,0xbc,0x2d,0x1a,0x3c,0xd8},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := HMACSHA256(tt.args.key, tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("HMACSHA256() = %v, want %v", got, tt.want)
			}
		})
	}
}
