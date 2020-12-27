package crypto

import (
	"testing"
)

func TestPKCS5Padding(t *testing.T) {
	padding := &PKCS5Padding{
		BlockSize: 15,
	}

	data := []byte("0123456789")
	pd := padding.Padding(data)
	t.Log(data, pd)

	upd := padding.UnPadding(pd)
	t.Log(upd)

	if string(data) != string(upd) {
		t.Error(upd)
	}
}
