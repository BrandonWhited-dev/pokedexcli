package main

func main() {
	configuration := &config{
		Next:     "https://pokeapi.co/api/v2/location-area/",
		Previous: "",
	}
	replStart(configuration)
}
