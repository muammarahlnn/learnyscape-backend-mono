package encryptutil

type Encryptor interface {
	Encrypt(data string) string
	Decrypt(data string) (string, error)
}
