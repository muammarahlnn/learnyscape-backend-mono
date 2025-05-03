package encryptutil

import "encoding/base64"

type Base64Encryptor struct {
}

func NewBase64Encryptor() *Base64Encryptor {
	return &Base64Encryptor{}
}

func (e *Base64Encryptor) Encrypt(data string) string {
	return base64.StdEncoding.EncodeToString([]byte(data))
}

func (e *Base64Encryptor) Decrypt(data string) (string, error) {
	decoded, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		return "", err
	}

	return string(decoded), nil
}
