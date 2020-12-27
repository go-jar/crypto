package crypto

import "testing"

func TestAesCBCCrypter(t *testing.T) {
	key := Md5String([]byte("key"))
	iv := Md5String([]byte("initialization vector"))[:AesBlockSize]
	data := []byte("data")

	a, err := NewAesCBCCrypter([]byte(key), []byte(iv))
	t.Log(err)
	t.Log(a.BlockSize())

	crypted := a.Encrypt(data)
	t.Log(crypted)

	decrypted := a.Decrypt(crypted)
	t.Log(string(decrypted))

	if string(decrypted) != string(data) {
		t.Error(decrypted, data)
	}
}
