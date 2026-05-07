package main

import (
	_ "awesomeProject1/routers"
	beego "github.com/beego/beego/v2/server/web"
)

func main() {
	beego.Run()
}

