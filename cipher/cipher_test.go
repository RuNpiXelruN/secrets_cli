package cipher

import "testing"

func TestCipher(t *testing.T) {

	key := "brooks"
	plainText := "sawyer"
	Encrypt(key, plainText)
}
