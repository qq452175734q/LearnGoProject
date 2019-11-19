package main

import (
	"github.com/astaxie/beego"
	_ "project/models"
	_ "project/routers"
)

func main() {
	beego.Run()
}

