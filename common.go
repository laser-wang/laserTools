// Package laser.tools
// 一些常用的通用函数
package laserTools

import (
	"crypto/md5"
	"fmt"
	"io"

	"github.com/satori/go.uuid"
)

// GetToken 取得token，任何时候都是唯一值
func GetToken() string {
	// Creating UUID Version 4
	u1 := uuid.NewV4()
	//	fmt.Printf("UUIDv4: %s\n", u1)

	// Parsing UUID from string input
	//	u2, err := uuid.FromString("6ba7b810-9dad-11d1-80b4-00c04fd430c8")
	//	if err != nil {
	//		fmt.Printf("Something gone wrong: %s", err)
	//	}
	//	fmt.Printf("Successfully parsed: %s", u2)
	//	fmt.Println("")
	return u1.String()
}

// Md5 对string进行md5加密
func Md5(text string) string {
	hashMd5 := md5.New()
	io.WriteString(hashMd5, text)
	return fmt.Sprintf("%x", hashMd5.Sum(nil))
}
