package main

import (
	"cryptoExample"
	"fmt"
)

func main() {

	////DES
	//resource := []byte("DES对称加密算法需要64比特明文被分为左、" +
	//	"右两部分处理")
	//key := []byte("13423123")
	//
	//x1 :=cryptoExample.EncryptoDES(resource,key)
	//fmt.Println(string(x1))
	//
	//x2 := cryptoExample.DecryptoDES(x1,key)
	//fmt.Println(string(x2))


	//AES
	srcdata := []byte("区块链AES算法加密解密")
	key := []byte("wangchong1234567")
	x1, _ := cryptoExample.AesEncrypto(srcdata,key)
	fmt.Printf("密文：%s\n",x1)

	x2, _ := cryptoExample.AesDecrypto(x1,key)
	fmt.Printf("明文：%s",x2)


}

