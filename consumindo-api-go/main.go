package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

type Response struct{
	Name string `json:"name"`
	Pokemon []Pokemon `json:"pokemon_entries"`
}
type Pokemon struct{
	Id uint `json:"entry_number"`
	Species PokemonSpecies `json:"pokemon_species"`
}
type PokemonSpecies struct{
	Name string `json:"name"`
}

func main(){
	res, err := http.Get("https://pokeapi.co/api/v2/pokedex/kanto")
	if err!=nil{
		log.Fatal(err)
		os.Exit(1)
	}
	resData, err := io.ReadAll(res.Body)
	if err!=nil{
		log.Fatal(err)
		os.Exit(1)
	}
	var resObj Response
	json.Unmarshal(resData,&resObj)

	for _,pokemon := range resObj.Pokemon{
		fmt.Println(pokemon.Id,"-",pokemon.Species.Name)
	}
}