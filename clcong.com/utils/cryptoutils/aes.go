package cryptoutils

import (
	"crypto/aes"
	"crypto/cipher"
)

type AesEncrypt struct {
}

func (this *AesEncrypt) getKey(strKey string) []byte {
	keyLen := len(strKey)
	if keyLen < 16 {
		panic("res key 长度不能小于16")
	}
	arrKey := []byte(strKey)
	if keyLen >= 32 {
		//取前32个字节
		return arrKey[:32]
	}
	if keyLen >= 24 {
		//取前24个字节
		return arrKey[:24]
	}
	//取前16个字节
	return arrKey[:16]
}

//加密字符串
func (this *AesEncrypt) Encrypt(key string, data []byte) ([]byte, error) {
	k := this.getKey(key)
	var iv = []byte(key)[:aes.BlockSize]
	encrypted := make([]byte, len(data))
	aesBlockEncrypter, err := aes.NewCipher(k)
	if err != nil {
		return nil, err
	}
	aesEncrypter := cipher.NewCFBEncrypter(aesBlockEncrypter, iv)
	aesEncrypter.XORKeyStream(encrypted, data)
	return encrypted, nil
}

//解密字符串
func (this *AesEncrypt) Decrypt(key string, src []byte) (data []byte, err error) {
	defer func() {
		//错误处理
		if e := recover(); e != nil {
			err = e.(error)
		}
	}()
	k := this.getKey(key)
	var iv = []byte(key)[:aes.BlockSize]
	decrypted := make([]byte, len(src))
	var aesBlockDecrypter cipher.Block
	aesBlockDecrypter, err = aes.NewCipher([]byte(k))
	if err != nil {
		return []byte{}, err
	}
	aesDecrypter := cipher.NewCFBDecrypter(aesBlockDecrypter, iv)
	aesDecrypter.XORKeyStream(decrypted, src)
	return decrypted, nil
}
