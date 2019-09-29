package main

import (
	"fmt"
	_ "hackathon/routers"
	"hackathon/services"

	"github.com/astaxie/beego"
)

func init() {
	//logos, err := services.ReadLogos("./logos.json")
	err := services.InitDatabase()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	services.OurLogo = logos

	//fmt.Println("logos!", logos)

}

func main() {
	beego.Run()

}
