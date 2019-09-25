package services

import (
	"encoding/json"
	"math/rand"
	"os"
	"time"
)


type Logo struct {
	Name string `json:"name"`
	ShortName string `json:"shortname"`
	URL string `json:"url"`
	Files []string `json:"files"`
}

type Logos []Logo

var OurLogo Logos


func ReadLogos(path string) (logos Logos, err error) {
	file, err := os.Open(path)
	if err != nil {
		return logos, err
	}

	jsonParser := json.NewDecoder(file)

	err = jsonParser.Decode(&logos)
	return logos, err
}


func GetRandomLogos() (uri string, options []string) {
	length := len(OurLogo)
	rand.Seed(time.Now().UnixNano())
	dupCheck := make(map[int]bool)
	numSelected := 0
	var logos = make([]Logo, 0)

	// select 4 random logo object without duplicated
	for {
		randomIndex := rand.Intn(length)
		tempLogo := OurLogo[randomIndex]
		if numSelected == 4 {
			break
		} else {
			if _, ok := dupCheck[randomIndex]; ok {
				continue
			} else {
				dupCheck[randomIndex] = true
				logos = append(logos, tempLogo)
				numSelected += 1
				options = append(options, tempLogo.ShortName)
			}
		}
	}

	rand.Seed(time.Now().UnixNano())
	randomUrl := "https://cdn.svgporn.com/logos/" + logos[rand.Intn(4)].Files[0]
	return randomUrl, options


}