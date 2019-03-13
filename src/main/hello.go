package main

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strings"
)

type student struct {
	Name string
	Age int
}

func main() {
	pattern := "123456"
	enPattern := sha256.New()
	has := enPattern.Sum([]byte(pattern))
	// sha256 加密
	fmt.Println(hex.EncodeToString(has))
	pattern2 := []byte("123456")
	en := md5.New()
	ddd := en.Sum(pattern2)
	// md5 加密
	fmt.Println(hex.EncodeToString(ddd))
	aaa := "  abc "
	newAAA := strings.TrimRight(aaa," ")
	fmt.Print(newAAA+"123"+"\n")
	fmt.Print("hhhh\n")
	fmt.Printf("变量传入函数前的地址为: %p\n",&aaa)
	// 函数参数是值传递
	func(a *string,b string){
		fmt.Printf("指针传入函数内部的内存地址不变,地址为: %p\n",a)
		fmt.Printf("值传入函数内部的内存地址会改变(引用拷贝),地址为: %p\n",&b)
		*a = "abcd"
	}(&aaa,aaa)
	fmt.Println("函数通过指针更改外部变量1:"+aaa)
	changeVar(&aaa)
	fmt.Println("函数通过指针更改外部变量2:"+aaa)
	paseStudent()
	sampleIota()
}

/*
通过引用传递改变外部变量
 */
func changeVar(a *string) {
	*a = "changeVar"
	//aaa = "cccc"
}
/*
在使用 for range 循环时是值引用
 */
func paseStudent() map[string]*student{
	m := make(map[string]*student)
	_student := []student{
		{Name:"chen",Age:21},
		{Name:"gong",Age:22},
		{Name:"li",Age:21},
	}
	for _,std := range _student {
		fmt.Printf("当前的值地址为:%p\n",&std) // 打印同一个内存地址
	}
	for i,_ := range _student{
		std := _student[i]
		fmt.Printf("更改后的内存地址为: %p\n",&std)
	}
	return m
}

/*

 */
func sampleIota(){
	const (
		sun = iota
		mon
		tue
		wes
		thr
		fri
		sat
	)
	fmt.Println(sun,mon,tue,wes,thr,fri,sat)
}
