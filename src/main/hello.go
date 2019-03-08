package main

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strings"
)

func main() {
	pattern := "123456"
	enPattern := sha256.New()
	has := enPattern.Sum([]byte(pattern))
	fmt.Println(hex.EncodeToString(has))

	pattern2 := []byte("123456")
	en := md5.New()
	en.Write(pattern2)
	ddd := en.Sum(nil)
	fmt.Println(hex.EncodeToString(ddd))
	aaa := "  abc "
	newAAA := strings.TrimRight(aaa," ")
	fmt.Print(newAAA)
	fmt.Print("hhhh")
	fmt.Printf("%p",&aaa)
}
