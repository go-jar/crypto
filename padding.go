package crypto

import (
	"bytes"
)

type PaddingInterface interface {
	Padding(data []byte) []byte
	UnPadding(data []byte) []byte
}

type PKCS5Padding struct {
	BlockSize int
}

func (p *PKCS5Padding) Padding(data []byte) []byte {
	paddingLen := p.BlockSize - len(data)%p.BlockSize
	paddingText := bytes.Repeat([]byte{byte(paddingLen)}, paddingLen)

	return append(data, paddingText...)
}

func (p *PKCS5Padding) UnPadding(data []byte) []byte {
	l := len(data)
	nonPaddingText := int(data[l-1])

	return data[:(l - nonPaddingText)]
}
