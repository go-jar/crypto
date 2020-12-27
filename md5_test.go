package crypto

import "testing"

func TestMd5File(t *testing.T) {
	md5 := Md5String([]byte("aaa"))
	if len(md5) != 32 {
		t.Error(md5)
	}

	t.Log(md5)
}

func TestMd5FileByPath(t *testing.T) {
	md5, err := Md5FileByPath("./md5")
	if err != nil {
		t.Error(err)
	}

	t.Log(md5)
}
