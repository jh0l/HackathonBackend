package services

func VerifyAnswer(url, selection string) bool {
	//fmt.Println(url, selection)
	return "https://cdn.svgporn.com/logos/" + selection == url
}