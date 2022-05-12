// Licensed to Apache Software Foundation(ASF) under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Apache Software Foundation(ASF) licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

package sm2

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"github.com/tjfoc/gmsm/sm2"
	"math/big"
	"testing"
)

var (
	SM2 = NewSm2Encrypt()
)

func TestSm2PublicKeyInt(t *testing.T) {
	pub := SM2.GetPubliKey()
	x, _ := new(big.Int).SetString(pub[2:66], 16)
	y, _ := new(big.Int).SetString(pub[66:], 16)
	publicKey := sm2.PublicKey{
		Curve: sm2.P256Sm2(),
		X:     x,
		Y:     y,
	}
	var (
		plainText = "qwer"
	)
	cipherText, err := publicKey.EncryptAsn1([]byte(plainText), rand.Reader)
	if err != nil {
		fmt.Println(err)
	}

	pri := SM2.GetPrivateKey()
	d, _ := new(big.Int).SetString(pri[:], 16)
	privatekey := &sm2.PrivateKey{
		PublicKey: publicKey,
		D:         d,
	}
	planiText, err := privatekey.DecryptAsn1(cipherText)

	fmt.Println(string(planiText)) // qwer
}

func TestParseWithKey(t *testing.T) {
	var (
		prk      = "ab41ceea04c0a7c643c29597452c156ef871b9afdf2b82fe35e1b6f70979df17"
		pbk      = "0444b39f0a6e3c14ceefcbc283d70c20ad003e4bed20b5a8f8f207a6642b8b400b630f9c405d107b73cc8e5534efc3cab73e31e35cdaba3af07af3fecf50246298"
		password = "54777d20c35fc3feb523925630a3daa2de79a076f3e405daf95c7affa5e9a6acf4fad58a5a9bdef8f70bedbd092c072541c1b7cc17fd0497149ead5fa55ac18affb5391ad02c3b35282f6a56ef073557da00b8d864b906be766547694d107431108d6330cda80bac40deb49fd0826a"
	)
	var (
		x, _       = new(big.Int).SetString(pbk[2:66], 16)
		y, _       = new(big.Int).SetString(pbk[66:], 16)
		d, _       = new(big.Int).SetString(prk[:], 16)
		privatekey = &sm2.PrivateKey{
			PublicKey: sm2.PublicKey{
				Curve: sm2.P256Sm2(),
				X:     x,
				Y:     y,
			},
			D: d,
		}
	)
	encrptDataHex := fmt.Sprintf("%x", password)
	encrptDataHexDocode, _ := hex.DecodeString(encrptDataHex)
	planiTexbt, err := sm2.Decrypt(privatekey, encrptDataHexDocode, 0)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(planiTexbt))
}

func TestParseWithKeyHex(t *testing.T) {
	da, _ := base64.StdEncoding.DecodeString("Bu9yfbVvUWWOPV6grq25Eb0yo/DWAgigWdJ50hmSIcVd3iYMbD9UXcjdmhuE+cTnoqYbLHmbHQPya5AB0zP5s4xiferrL1MXtdA9vWRWtCEa1RG5B+ZmwowsSrIaWZKdVhStBFZylX1p5JUIHCXH")
	fmt.Println(string(da))
	fmt.Printf("%x", da)
}

func TestName(t *testing.T) {
	var msg = "1ay"
	// pub 加密
	e, _ := SM2.Encrypt([]byte(msg))
	h := fmt.Sprintf("%x", e) //041abe3975272a65189f415a93012a086646f78904e550547c21ebf61aaddd83859cf224b134a3e4e1154476d6b47eef52919e441ac62aebad16c1a268a2f72ff84ca7c73c578e653e851b320d2bfce491c68f9ff1f3e89c6c8b79fd75dd65142f2c11e2

	// pri 16进制字符串解密
	str, _ := SM2.DecryptHexString([]byte(h))
	fmt.Println(string(str))
	// pri 切片解密
	by, _ := SM2.DecryptByte(e)
	fmt.Println(string(by))

}

func TestStr(t *testing.T) {
	var (
		prk      = "ab41ceea04c0a7c643c29597452c156ef871b9afdf2b82fe35e1b6f70979df17"
		pbk      = "0444b39f0a6e3c14ceefcbc283d70c20ad003e4bed20b5a8f8f207a6642b8b400b630f9c405d107b73cc8e5534efc3cab73e31e35cdaba3af07af3fecf50246298"
		password = "0454777d20c35fc3feb523925630a3daa2de79a076f3e405daf95c7affa5e9a6acf4fad58a5a9bdef8f70bedbd092c072541c1b7cc17fd0497149ead5fa55ac18affb5391ad02c3b35282f6a56ef073557da00b8d864b906be766547694d107431108d6330cda80bac40deb49fd0826a"
	)

	var (
		x, _       = new(big.Int).SetString(pbk[2:66], 16)
		y, _       = new(big.Int).SetString(pbk[66:], 16)
		d, _       = new(big.Int).SetString(prk[:], 16)
		privatekey = &sm2.PrivateKey{
			PublicKey: sm2.PublicKey{
				Curve: sm2.P256Sm2(),
				X:     x,
				Y:     y,
			},
			D: d,
		}
	)
	msgHexDocode, _ := hex.DecodeString(string([]byte(password)))
	planiText, err := sm2.Decrypt(privatekey, msgHexDocode, mode)
	fmt.Println(string(planiText),err)
}