package main

import (
	"fmt"
	"log"

	"githhub.com/alekslesik/learn/kandinsky"
)

// https://api-key.fusionbrain.ai/key/api/v1/text2image/run
// FFDB0757E5E5A2FF2D7A297CE95BDA2D
// F9C4980BC3166E19C9A42607358B8DA6

func main() {
	URLmodel := "https://api-key.fusionbrain.ai/key/api/v1/models"
	URLUUID := "https://api-key.fusionbrain.ai/key/api/v1/text2image/run"
	key := "FFDB0757E5E5A2FF2D7A297CE95BDA2D"
	secret := "F9C4980BC3166E19C9A42607358B8DA6"

	k, err := kandinsky.New(key, secret)
	if err != nil {
		log.Fatal("create new Kandinsky error > ", err)
	}

	err = k.SetModel(URLmodel)
	if err != nil {
		log.Fatal("set model error > ", err)
	}

	p := kandinsky.Params{
		Type:           "GENERATE",
		Style:          "string",
		Width:          1024,
		Height:         1024,
		NumImages:      1,
		NegativePrompt: "",
		GenerateParams: struct {
			Query string "json:\"query\""
		}{"Северная Пальмира"},
	}

	u, err := k.GetUUID(URLUUID, p)
	if err != nil {
		log.Fatal("get UUID error > ", err)
	}

	fmt.Println(u.ID)

}
