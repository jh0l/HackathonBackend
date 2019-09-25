package main

import (
	"fmt"
	_ "hackathon/routers"
	"github.com/astaxie/beego"
	"hackathon/services"
)



func init() {
	 logos, err := services.ReadLogos("./logos.json")
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

