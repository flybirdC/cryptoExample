package cryptoExample

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
	"math/big"
	"os"
)

//生成随机质数
func NewPrime(n int) *big.Int  {

	pq, err := rand.Prime(rand.Reader,n)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return pq
}

//生成公私钥文件
func GenRsaKey(bits int) error  {
	//生成私钥文件
	privatekey, err := rsa.GenerateKey(rand.Reader,bits)
	if err != nil {
		return err
	}
	derStream := x509.MarshalPKCS1PrivateKey(privatekey)
	block := &pem.Block{
		Type:"RSA PRIVATE KRY",
		Bytes:derStream,
	}
	file,err := os.Create("private.pem")
	if err != nil {
		return err
	}
	err = pem.Encode(file,block)
	if err != nil {
		return err
	}
	//生成公钥文件
	publickey := &privatekey.PublicKey
	defPkix, err := x509.MarshalPKIXPublicKey(publickey)
	if err != nil {
		return err
	}
	block = &pem.Block{
		Type:"RSA PUBLIC KEY",
		Bytes:defPkix,
	}
	file, err = os.Create("public.pem")
	if err != nil {
		return err
	}
	err = pem.Encode(file,block)
	if err != nil {
		return err
	}
	return nil
}



//公钥加密
func RsaEncrypt(data []byte, pubkey []byte) ([]byte, error) {
	//解密pem格式公钥
	block, _ := pem.Decode(pubkey)
	if block == nil {
		return nil, errors.New("public key error")
	}
	//解析公钥
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	//类型断言
	pub := pubInterface.(*rsa.PublicKey)
	//加密
	return rsa.EncryptPKCS1v15(rand.Reader,pub,data)
}
//私钥解密
func RsaDecrypt(ciphertext []byte, privkey []byte) ([]byte, error)  {
	//获取私钥
	block, _ := pem.Decode(privkey)
	if block == nil {
		return nil,errors.New("privatekey error")
	}
	//解析pkcs1格式的私钥
	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	//解密
	return rsa.DecryptPKCS1v15(rand.Reader,priv,ciphertext)
}

//私钥签名
func RsaSign(data []byte,privatekey []byte) ([]byte, error)  {

	h := sha256.New()
	h.Write(data)
	hashed := h.Sum(nil)

	//获取私钥
	block, _ := pem.Decode(privatekey)
	if block == nil {
		return nil, errors.New("privatekey error")
	}
	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	return rsa.SignPKCS1v15(rand.Reader,priv,crypto.SHA256,hashed)
}

//公钥验签
func RsaSignVerity(data []byte, sinature []byte, publickey []byte) error  {
	hashed := sha256.Sum256(data)
	block, _:= pem.Decode(publickey)
	if block == nil {
		return errors.New("publickey error")
	}
	//解析公钥
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return err
	}
	//类型断言
	pub := pubInterface.(*rsa.PublicKey)
	//验证签名
	return rsa.VerifyPKCS1v15(pub,crypto.SHA256,hashed[:],sinature)
}


