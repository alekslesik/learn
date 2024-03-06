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
	key := "FFDB0757E5E5A2FF2D7A297CE95BDA2D"
	secret := "F9C4980BC3166E19C9A42607358B8DA6"

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

	i, err := kandinsky.GetImage(key, secret, p)

	select {
	case <- err:
		log.Fatal(<- err)
	case <- i:
		image := <- i
		fmt.Println(image.ToByte())
	}

	// select {
	// case image := <- i:
	// 	b, err := image.ToByte()
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}

	// 	fmt.Println(b)
	// }

}
