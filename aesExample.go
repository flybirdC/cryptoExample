package cryptoExample

import (
	"crypto/aes"
	"crypto/cipher"
)

//AES加密
func AesEncrypto(srcData []byte, key []byte) ([]byte, error) {

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	blockSize := block.BlockSize()
	srcData = padding(srcData,blockSize)
	//默认key=IV
	blockMode := cipher.NewCBCEncrypter(block,key[:blockSize])
	cryptedData := make([]byte,len(srcData))

	blockMode.CryptBlocks(cryptedData,srcData)

	return cryptedData, nil
}

//AES解密
func AesDecrypto(cryptoData []byte, key []byte) ([]byte, error) {

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	blockSize := block.BlockSize()
	//默认key=IV
	blockMode := cipher.NewCBCDecrypter(block,key[:blockSize])
	srcData := make([]byte, len(cryptoData))
	blockMode.CryptBlocks(srcData,cryptoData)
	srcData = unpadding(srcData)

	return srcData,nil

}
