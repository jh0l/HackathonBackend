package main

import (
	"fmt"
	_ "hackathon/routers"

	_ "github.com/mattn/go-sqlite3"

	"hackathon/services"

	"github.com/astaxie/beego"
)

func init() {
	logos, err := services.ReadLogos("./logos.json")

	if err != nil {
		fmt.Println(err.Error())
	}

	err = services.InitDatabase()

	if err != nil {
		fmt.Println(err.Error())
	}

	services.OurLogo = logos

	//fmt.Println("logos!", logos)

}

func main() {
	beego.Run()

}
