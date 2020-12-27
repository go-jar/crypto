package crypto

import "testing"

func TestSha1String(t *testing.T) {
	s := Sha1String([]byte("aaa"))
	if len(s) != 40 {
		t.Error(s)
	}

	t.Log(s)
}
