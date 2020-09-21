package main

import (
	"6-18/object"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	fmt.Println("Welcome to League Of Legends.")

	url := "https://game.gtimg.cn/images/lol/act/img/js/heroList/hero_list.js"

	client := http.Client{}

	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println(request)
		return
	}

	resp, err := client.Do(request)
	if err != nil {
		fmt.Println(request)
		return
	}
	heroData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(request)
		return
	}
		fmt.Println(string(heroData))


	var herolist object.HeroList


	 err = json.Unmarshal(heroData, &herolist)
		if err !=nil {
			fmt.Println(err.Error())
			return
		}
	fmt.Println(herolist)
	for i:=0;i<len(herolist.Hero) ;i++  {
		fmt.Println((herolist.Hero[i]))
	}
	 fmt.Println(herolist.Version)
	 fmt.Println(herolist.FileName)
	 fmt.Println(herolist.FileTime)
}
