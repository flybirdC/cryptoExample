package main

import (
	"cryptoExample"
	"fmt"
)

func main() {

	resource := []byte("DES对称加密算法需要64比特明文被分为左、" +
		"右两部分处理，右侧数据和子密钥经过轮函数f生成用于加密左侧数据的比特序列，" +
		"与左侧数据异或运算，运算结果输出为加密后的左侧，右侧数据则直接输出为右侧。" +
		"其中子密钥为本轮加密使用的密钥，每次Feistel均使用不同的子密钥")
	key := []byte("13423123")

	x1 :=cryptoExample.EncryptoDES(resource,key)
	fmt.Println(string(x1))

	x2 := cryptoExample.DecryptoDES(x1,key)
	fmt.Println(string(x2))
}

