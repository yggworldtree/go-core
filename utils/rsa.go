package utils

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
)

//GenerateRSAKey 生成RSA私钥和公钥，保存到文件中
func GenerateRSAKey(bits int) error {
	//GenerateKey函数使用随机数据生成器random生成一对具有指定字位数的RSA密钥
	//Reader是一个全局、共享的密码用强随机数生成器
	privateKey, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		return err
	}
	//保存私钥
	//通过x509标准将得到的ras私钥序列化为ASN.1 的 DER编码字符串
	X509PrivateKey := x509.MarshalPKCS1PrivateKey(privateKey)
	//使用pem格式对x509输出的内容进行编码
	//创建文件保存私钥
	privateFile, err := os.OpenFile("private.pem", os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer privateFile.Close()
	//构建一个pem.Block结构体对象
	privateBlock := pem.Block{Type: "RSA Private Key", Bytes: X509PrivateKey}
	//将数据保存到文件
	err = pem.Encode(privateFile, &privateBlock)
	if err != nil {
		return err
	}

	//保存公钥
	//获取公钥的数据
	publicKey := privateKey.PublicKey
	//X509对公钥编码
	X509PublicKey, err := x509.MarshalPKIXPublicKey(&publicKey)
	if err != nil {
		return err
	}
	//pem格式编码
	//创建用于保存公钥的文件
	publicFile, err := os.OpenFile("public.pem", os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer publicFile.Close()
	//创建一个pem.Block结构体对象
	publicBlock := pem.Block{Type: "RSA Public Key", Bytes: X509PublicKey}
	//保存到文件
	return pem.Encode(publicFile, &publicBlock)
}

func RSALoadKey(o interface{}) (rt *pem.Block, rterr error) {
	defer func() {
		if errs := recover(); errs != nil {
			rterr = fmt.Errorf("recover:%v", errs)
		}
	}()
	var bts []byte
	switch o.(type) {
	case string:
		//打开文件
		fl, err := os.Open(o.(string))
		if err != nil {
			return nil, err
		}
		defer fl.Close()
		//读取文件的内容
		bts, err = ioutil.ReadAll(fl)
		if err != nil {
			return nil, err
		}
	case []byte:
		bts = o.([]byte)
	}
	if bts == nil {
		return nil, errors.New("bytes not found")
	}
	//pem解码
	block, _ := pem.Decode(bts)
	return block, nil
}

//RSAEncrypt RSA加密
func RSAEncrypt(plainText []byte, publicKey interface{}) ([]byte, error) {
	if publicKey == nil {
		return nil, errors.New("param err")
	}
	block, err := RSALoadKey(publicKey)
	if err != nil {
		return nil, err
	}
	//x509解码
	publicKeyInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	//类型断言
	publicKeys := publicKeyInterface.(*rsa.PublicKey)
	//对明文进行加密
	cipherText, err := rsa.EncryptPKCS1v15(rand.Reader, publicKeys, plainText)
	//返回密文
	return cipherText, err
}

//RSADecrypt RSA解密
func RSADecrypt(cipherText []byte, privateKey interface{}) ([]byte, error) {
	if privateKey == nil {
		return nil, errors.New("param err")
	}
	block, err := RSALoadKey(privateKey)
	if err != nil {
		return nil, err
	}
	//X509解码
	privateKeys, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	//对密文进行解密
	plainText, err := rsa.DecryptPKCS1v15(rand.Reader, privateKeys, cipherText)
	//返回明文
	return plainText, err
}
func RSAGetSign(data []byte, privateKey interface{}) ([]byte, error) {
	if privateKey == nil {
		return nil, errors.New("param err")
	}
	block, err := RSALoadKey(privateKey)
	if err != nil {
		return nil, err
	}
	//X509解码
	privateKeys, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	//计算散列值
	hash := sha256.New()
	hash.Write(data)
	bts := hash.Sum(nil)
	return rsa.SignPKCS1v15(rand.Reader, privateKeys, crypto.SHA256, bts)
}
func RSAVerifySign(data, sign []byte, publicKey interface{}) error {
	if publicKey == nil {
		return errors.New("param err")
	}
	block, err := RSALoadKey(publicKey)
	if err != nil {
		return err
	}
	//x509解码
	publicKeyInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return err
	}
	//类型断言
	publicKeys := publicKeyInterface.(*rsa.PublicKey)
	//计算消息散列值
	hash := sha256.New()
	hash.Write(data)
	bts := hash.Sum(nil)
	//验证数字签名
	return rsa.VerifyPKCS1v15(publicKeys, crypto.SHA256, bts, sign)
}
