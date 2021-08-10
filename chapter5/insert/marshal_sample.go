package main

import (
	"encoding/json"
	"fmt"

)

type Idol struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func main() {
	kaho := Idol{
		Name: "Komiya kaho",
		ID:   "3",
	}
	s, err := json.Marshal(kaho)
	if err != nil {
		panic(err)
	}
	fmt.Println(s)
	fmt.Println(string(s))

	var idol Idol
	input := []byte(`{"id":"8","name":"tukioka kogane"}`)
	fmt.Println(input)
	fmt.Println(string(input))
	err = json.Unmarshal(input, &idol)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(idol)
}
