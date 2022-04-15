package encrypt

import (
	"dtstack.com/dtstack/easymatrix/matrix/cache"
	"dtstack.com/dtstack/easymatrix/matrix/encrypt/aes"
	"dtstack.com/dtstack/easymatrix/matrix/encrypt/rsa"
	"dtstack.com/dtstack/easymatrix/matrix/encrypt/sm2"
	"errors"
	"fmt"
)

type commonEncrypt interface {
	CommonEncrypt(msg []byte) ([]byte, error)
	CommonDecrypt(msg []byte) ([]byte, error)
	SchemeDecrypt(msg string, aesPassword string) (string, error)
	CommonGetPublicKey() string
}

var (
	PlatformEncrypt commonEncrypt = newEncrypt()
)

type platFormEncrypt struct {
	sm  sm2.SM2encrypt
	rsa rsa.Cipher
}

func (a *platFormEncrypt) CommonGetPublicKey() string {
	switch cache.SysConfig.PlatFormSecurity.LoginEncrypt {
	case "sm2":
		return a.sm.GetPubliKey()
	case "rsa":
		return a.rsa.PublicKeyToString()
	}
	return ""
}

func (a *platFormEncrypt) CommonEncrypt(msg []byte) ([]byte, error) {

	panic("implement me")
}

func (a *platFormEncrypt) CommonDecrypt(msg []byte) ([]byte, error) {
	plain := []byte{}
	switch cache.SysConfig.PlatFormSecurity.LoginEncrypt {
	case "sm2":
		decryptMsg, err := a.sm.DecryptHexString(msg)
		if err != nil {
			return nil, err
		}
		plain = decryptMsg
	case "rsa":
		decryptMsg, err := a.rsa.Decrypt(msg)
		if err != nil {
			return nil, err
		}
		plain = decryptMsg
	}
	return plain, nil
}

// aesPassword == ""
// scheme 加密 rsa类型时候，默认使用aes
func (a *platFormEncrypt) SchemeDecrypt(msg string, aesPassword string) (string, error) {
	switch cache.SysConfig.PlatFormSecurity.LoginEncrypt {
	case "sm2":
		decryptMsg, err := a.sm.DecryptHexString([]byte(msg))
		fmt.Println("[platFormEncrypt.SchemeDecrypt]",msg)
		if err != nil {
			return "", err
		}
		fmt.Println("[platFormEncrypt.SchemeDecrypt]",decryptMsg)
		return string(decryptMsg), nil
	case "rsa":
		password, err := aes.AesDecryptByPassword(msg, aesPassword)
		if err != nil {
			return "", err
		}
		return password, nil
	}
	return "", errors.New("no support encrypt type")
}

func newEncrypt() commonEncrypt {
	return &platFormEncrypt{
		sm:  sm2.NewSm2Encrypt(),
		rsa: rsa.NewRsaEncrypt(),
	}
}
