package cryptoExample

import (
	"bytes"
	"compress/gzip"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
	"math/big"
	"strings"
)

//通过随机数或种子构建公私钥，至少36位
func GetecdsaKey(randKey string) (*ecdsa.PrivateKey,ecdsa.PublicKey,error)  {
	var err error

	var priv *ecdsa.PrivateKey
	var pub ecdsa.PublicKey
	var curve elliptic.Curve

	lenth := len(randKey)
	if lenth < 224/8 {
		err = errors.New("随机数至少为36位")
		return priv,pub,err
	}
	if lenth > 521/8+8 {
		curve = elliptic.P521()
	} else if lenth > 384/8+8 {
		curve = elliptic.P384()
	} else if lenth >256/8+8 {
		curve = elliptic.P256()
	} else if lenth > 224/8+8 {
		curve = elliptic.P224()
	}

	priv, err = ecdsa.GenerateKey(curve,strings.NewReader(randKey))
	if err != nil {
		return priv,pub,err
	}
	pub = priv.PublicKey

	return priv,pub,err
}
/**
对text加密，text必须是一个hash值，例如md5、sha1等
使用私钥priv
使用随机熵增强加密安全，安全依赖于此熵，randsign
返回加密结果，结果为数字证书r、s的序列化后拼接，然后用hex转换为string
*/
func SignEcdsa(text []byte, randsign string,priv *ecdsa.PrivateKey) (string, error)  {

	r,s,err := ecdsa.Sign(strings.NewReader(randsign),priv,text)
	if err != nil {
		return "", err
	}
	rt,err := r.MarshalText()
	if err != nil {
		return "", nil
	}
	st, err := s.MarshalText()
	if err != nil {
		return "", nil
	}

	var b bytes.Buffer
	w := gzip.NewWriter(&b)
	defer w.Close()
	_,err = w.Write([]byte(string(rt)+"+"+string(st)))
	if err != nil {
		return "", nil
	}
	w.Flush()
	return hex.EncodeToString(b.Bytes()), nil
}

/**
hash加密
使用md5加密
*/
func HashText(text, salt string) []byte  {
	md5Inst := md5.New()
	md5Inst.Write([]byte(text))
	result := md5Inst.Sum([]byte(salt))
	return result
}
/**
证书分解
通过hex解码，分割成数字证书r，s
*/
func GetSignEcdsa(signature string) (rint, sint big.Int,err error) {
	byterun, err := hex.DecodeString(signature)
	if err != nil {
		err = errors.New("decrpt error, "+ err.Error())
		return
	}
	r, err := gzip.NewReader(bytes.NewBuffer(byterun))
	if err != nil {
		err = errors.New("decode error, "+err.Error())
		return
	}
	defer r.Close()

	buf := make([]byte,1024)
	count,err := r.Read(buf)
	if err != nil {
		fmt.Println("decode=",err)
		err = errors.New("decode read error, "+err.Error())
		return
	}
	rs := strings.Split(string(buf[:count]),"+")
	if len(rs) != 2 {
		err = errors.New("decode fail")
		return
	}

}

