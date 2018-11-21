package cryptoExample

import (
	"bytes"
	"crypto/cipher"
	"crypto/des"
)

func padding(src []byte, blocksize int) []byte  {

	n := len(src)
	padnum := blocksize-n%blocksize
	pad := bytes.Repeat([]byte{byte(padnum)},padnum)
	dst := append(src,pad...)
	return dst
}

func unpadding(src []byte) []byte  {

	n := len(src)
	unpadnum := int(src[n-1])
	dst := src[:n-unpadnum]
	return dst
}
//使用CBC
func EncryptoDES(src []byte, key []byte) []byte  {

	block, _ := des.NewCipher(key)
	src = padding(src,block.BlockSize())
	//这里默认iv=key,iv为偏移量
	blockmode := cipher.NewCBCEncrypter(block,key)
	blockmode.CryptBlocks(src,src)
	return src
}
//使用CBC
func DecryptoDES(src []byte, key []byte) []byte  {

	block,_ := des.NewCipher(key)
	blockmode := cipher.NewCBCDecrypter(block,key)
	blockmode.CryptBlocks(src,src)
	src = unpadding(src)
	return src
}

//使用EBC加密
func EncryptoEBC(src []byte, key []byte)  []byte {

	block, _ := des.NewCipher(key)
	bs := block.BlockSize();
	src = padding(src,bs)
	out := make([]byte, len(src))
	dst := out
	for len(src) > 0 {
		block.Encrypt(dst,src[:bs])
		src = src[bs:]
		dst = dst[bs:]
	}

	return out
}
//使用EBC解密
func DecryptoEBC(src []byte, key []byte)  []byte {

	block, _ := des.NewCipher(key)
	bs := block.BlockSize();
	out := make([]byte, len(src))
	dst := out
	for len(src) > 0 {
		block.Decrypt(dst,src[:bs])
		src = src[bs:]
		dst = dst[bs:]
	}
	out = unpadding(out)
	return out
}


