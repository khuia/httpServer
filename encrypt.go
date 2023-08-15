package main

import (
	"errors"

	"crypto/aes"
	"crypto/cipher"

	"crypto/rand"

	"encoding/hex"

	"io"

	"strings"
)

const (
	// Nonce 的长度，单位为字节
	NonceSize = 12
	// 密钥的长度，单位为字节
	KeySize = 32
)

func Encrypt(keyStr string, srcData []byte) ([]byte, []byte, error) {
	// 检查输入参数的合法性
	if len(keyStr) == 0 {
		return nil, nil, errors.New("empty key string")
	}
	if len(srcData) == 0 {
		return nil, nil, errors.New("empty source data")
	}

	// 生成随机的 Nonce
	nonce := make([]byte, NonceSize)
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, nil, err
	}

	// 生成密钥
	key := make([]byte, KeySize)
	key = []byte(keyStr)
	if len(key) > KeySize {
		key = key[:KeySize]
	} else if len(key) < KeySize {
		padding := make([]byte, KeySize-len(key))
		key = append(key, padding...)
	}

	// 创建 AES 加密块
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, nil, err
	}

	// 创建 GCM 模式的 AES 加密器
	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return nil, nil, err
	}

	// 使用 AES-GCM 加密数据
	ciphertext := aesGCM.Seal(nil, nonce, srcData, nil)

	return ciphertext, nonce, nil
}

func Decrypt(keyStr string,nonce []byte,ciphertext []byte) ([]byte, error) {
	// 检查输入参数的合法性
	if len(keyStr) == 0 {
		return nil, errors.New("empty key string")
	}
	/*if len(data) < NonceSize {
		return nil, errors.New("invalid data")
	}

	// 分离 Nonce 和密文
	/*nonce := data[:NonceSize]
	ciphertext := data[NonceSize:]*/

	// 生成密钥
	key := make([]byte, KeySize)
	key = []byte(keyStr)
	if len(key) > KeySize {
		key = key[:KeySize]
	} else if len(key) < KeySize {
		padding := make([]byte, KeySize-len(key))
		key = append(key, padding...)
	}

	// 创建 AES 解密块
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	// 创建 GCM 模式的 AES 解密器
	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	// 使用 AES-GCM 解密数据
	plaintext, err := aesGCM.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return nil, err
	}

	return plaintext, nil
}

func getHexString(data []byte) string {
	return hex.EncodeToString(data)
}

func getByteArray(hexString string) ([]byte, error) {
	hexString = strings.TrimSpace(hexString)
	if len(hexString)%2 != 0 {
		return nil, errors.New("invalid hex string")
	}
	data, err := hex.DecodeString(hexString)
	if err != nil {
		return nil, err
	}
	return data, nil
}
