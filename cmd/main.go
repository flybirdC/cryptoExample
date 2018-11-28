package main

import (
	"cryptoExample"
	"fmt"
	"io/ioutil"
	"os"
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


	////AES
	//srcdata := []byte("区块链AES算法加密解密")
	//key := []byte("wangchong1234567")
	//x1, _ := cryptoExample.AesEncrypto(srcdata,key)
	//fmt.Printf("密文：%s\n",x1)
	//
	//x2, _ := cryptoExample.AesDecrypto(x1,key)
	//fmt.Printf("明文：%s",x2)

	//rsa
	fmt.Println(cryptoExample.NewPrime(4))

	fmt.Println(cryptoExample.NewPrime(5))

	//var bits int
	//flag.IntVar(&bits,"b",1024,"密钥长度默认为1024位")
	//if err := cryptoExample.GenRsaKey(bits); err != nil {
	//	log.Fatal("密钥文件生成失败")
	//}
	//log.Println("密钥文件生成成功")

	//读取公私钥
	publickey,err:= ioutil.ReadFile("public.pem")
	if err != nil {
		os.Exit(-1)
	}
	privatekey,err := ioutil.ReadFile("private.pem")
	if err != nil {
		os.Exit(-1)
	}
	fmt.Printf("pub=%s\n,private=%s\n",publickey,privatekey)
//
////	time.Sleep(2)
//
	theMSG := "加密我的名字wangchong"
	fmt.Println("source msg:",theMSG)

	////公私钥加解密
	//enc , err := cryptoExample.RsaEncrypt([]byte(theMSG),publickey)
	//
	//fmt.Println(string(enc),err)
	//
	//dec, err := cryptoExample.RsaDecrypt(enc,privatekey)
	//
	//fmt.Println(string(dec), err)

	//公私钥签名--验签
	sig, err := cryptoExample.RsaSign([]byte(theMSG),privatekey)
	fmt.Println(string(sig),err)

	err1 := cryptoExample.RsaSignVerity([]byte(theMSG),sig,publickey)
	fmt.Println(err1)

}




