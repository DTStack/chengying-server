package sm2

import (
	"crypto/rand"
	"dtstack.com/dtstack/easymatrix/matrix/log"
	"encoding/hex"
	"fmt"
	"github.com/tjfoc/gmsm/sm2"
)

// https://juejin.cn/post/6966639973340545032

var (
	mode = 0 // C1C3C2=0 C1C2C3=1
)

type SM2encrypt interface {
	Encrypt(msg []byte) ([]byte, error)          // publickey 加密
	DecryptHexString(msg []byte) ([]byte, error) // privatekey 16进制字符串解密
	DecryptByte(msg []byte) ([]byte, error)      // privatekey 字节解密
	GetPubliKey() string
	GetPrivateKey() string
	Sign(msg string) (sign []byte, ok bool) // privatekey签名
	Verify(msg, sign string) bool           // publickey 验签
}

type sm struct {
	privatekey *sm2.PrivateKey
	publickey  *sm2.PublicKey
}

func (s *sm) DecryptHexString(msg []byte) ([]byte, error) {
	msgHexDecode, _ := hex.DecodeString(string(msg))
	planiText, err := sm2.Decrypt(s.privatekey, msgHexDecode, mode)
	if err != nil {
		log.Errorf("%v", err)
	}
	return planiText, err
}

func (s *sm) DecryptByte(msg []byte) ([]byte, error) {
	planiText, err := sm2.Decrypt(s.privatekey, msg, mode)
	return planiText, err
}

func (s *sm) GetPrivateKey() string {
	d := fmt.Sprintf("%x", s.privatekey.D)
	return d
}

func (s *sm) GetPubliKey() string {
	xy := fmt.Sprintf("04%x%x", s.publickey.X, s.publickey.Y)
	return xy
}

func (s *sm) Sign(msg string) (sign []byte, ok bool) {
	sign, err := s.privatekey.Sign(rand.Reader, []byte(msg), nil)
	if err != nil {
		log.Errorf("[sm.Sign] %s", err)
		return sign, false
	}
	return sign, true
}

func (s *sm) Verify(msg, sign string) bool {
	return s.publickey.Verify([]byte(msg), []byte(sign)) //sm2验签
}

func (s *sm) Encrypt(msg []byte) ([]byte, error) {
	encrptData, err := sm2.Encrypt(&s.privatekey.PublicKey, msg, rand.Reader, mode)
	return encrptData, err
}

func NewSm2Encrypt() SM2encrypt {
	priv, _ := sm2.GenerateKey(rand.Reader) // 生成密钥对
	//fmt.Printf("GetPubliKey: 04%x%x\n", priv.PublicKey.X, priv.PublicKey.Y)
	//fmt.Printf("GetPrivateKey: %x\n", priv.D)
	return &sm{
		privatekey: priv,
		publickey:  &priv.PublicKey,
	}
}
