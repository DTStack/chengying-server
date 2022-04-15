package rsa

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"dtstack.com/dtstack/easymatrix/go-common/log"
	"encoding/hex"
	"encoding/pem"
)

const BIT_SIZE = 1024

func NewRsaEncrypt() Cipher {
	client, err := genRsaKey(BIT_SIZE)
	if err != nil {
		log.Errorf("Init cipher failed error: %v", err)
	}
	return client
}

type Cipher interface {
	Encrypt(msg []byte) ([]byte, error)
	Decrypt(msg []byte) ([]byte, error)
	PrivateKey() *rsa.PrivateKey
	PrivateKeyToString() string
	PublicKey() *rsa.PublicKey
	PublicKeyToString() string
}

type rsaClient struct {
	privateKey *rsa.PrivateKey
	publicKey  *rsa.PublicKey
}

func (r *rsaClient) Encrypt(msg []byte) ([]byte, error) {
	return rsa.EncryptPKCS1v15(rand.Reader, r.publicKey, msg)
}

func (r *rsaClient) Decrypt(msg []byte) ([]byte, error) {
	msgHexDecode, err := hex.DecodeString(string(msg))
	if err!=nil{
		log.Errorf("[rsaClient.Decrypt]%v", err)
		return nil, err
	}
	return rsa.DecryptPKCS1v15(rand.Reader, r.privateKey, msgHexDecode)
}

func (r *rsaClient) PrivateKey() *rsa.PrivateKey {
	return r.privateKey
}

func (r *rsaClient) PublicKey() *rsa.PublicKey {
	return r.publicKey
}

func (r *rsaClient) PrivateKeyToString() string {
	b := x509.MarshalPKCS1PrivateKey(r.privateKey)
	block := &pem.Block{
		Type:  "PRIVATE KEY",
		Bytes: b,
	}
	return string(pem.EncodeToMemory(block))
}

func (r *rsaClient) PublicKeyToString() string {
	b, _ := x509.MarshalPKIXPublicKey(r.publicKey)
	block := &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: b,
	}
	return string(pem.EncodeToMemory(block))
}

func genRsaKey(bits int) (*rsaClient, error) {
	privateKey, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		return nil, err
	}
	publicKey := &privateKey.PublicKey
	return &rsaClient{privateKey: privateKey, publicKey: publicKey}, nil
}
