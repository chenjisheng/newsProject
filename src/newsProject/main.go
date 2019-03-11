package main

import (
	_ "newsProject/routers"
	"github.com/astaxie/beego"
	_ "newsProject/models"
)

func main() {
	beego.Run()
}

