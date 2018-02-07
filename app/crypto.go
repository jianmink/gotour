package app

import (
	"crypto/hmac"
	"crypto/sha256"
)


// KDF key derivation function
// S = FC || P0 || L0 || P1 || L1 || P2 || L2 || P3 || L3 ||... || Pn || Ln
// HMAC-SHA-256 (key, S)
// Keyed-Hash Message Authentication Code (HMAC)

// HMACSHA256
// HMAC-SHA-256
func HMACSHA256(key,s [] byte) []byte{
	mac := hmac.New(sha256.New, key)
	mac.Write(s)
	return mac.Sum(nil)
}
