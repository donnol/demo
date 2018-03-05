package main

import (
	"crypto/dsa"
	"crypto/rand"
	"log"
	"math/big"
)

func main() {
	// testKey()

	// 签名
	params := new(dsa.Parameters)
	size := dsa.L2048N256
	err := dsa.GenerateParameters(params, rand.Reader, size)
	if err != nil {
		log.Fatal(err)
	}
	priv := &dsa.PrivateKey{
		PublicKey: dsa.PublicKey{
			Parameters: *params,
		},
	}
	err = dsa.GenerateKey(priv, rand.Reader) // 私钥只存在服务器端
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%+v", priv)

	hash := []byte("Hello, I am jd")
	r, s, err := dsa.Sign(rand.Reader, priv, hash)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%d, %d, %s\n", r, s, hash) // r, s构成签名，将会和hash一起提供给客户

	// 校验
	for _, h := range []string{
		"Hello, I am jd",
		"Hello world",
	} {
		right := dsa.Verify(&priv.PublicKey, []byte(h), r, s) // 校验过程在软件里进行
		log.Println(right)
	}
}

func testKey() {
	// 跑这一段时，CPU占用14%左右，而且跑的时间非常长，两个小时了还没跑完
	param := dsa.Parameters{
		P: big.NewInt(1),
		Q: big.NewInt(2),
		G: big.NewInt(3),
	}
	puli := dsa.PublicKey{
		Parameters: param,
	}
	priv := &dsa.PrivateKey{
		PublicKey: puli,
	}
	err := dsa.GenerateKey(priv, rand.Reader)
	if err != nil {
		log.Fatal(err)
	}
}
