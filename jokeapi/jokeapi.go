package jokeapi

import (
	"fmt"

	joke "github.com/icelain/jokeapi"
)

func GetJoke()  {

	jokeType := "single"
	blacklist := []string{"nsfw", "religious", "political", "racist", "sexist"}


	api := joke.New()

	api.Set(joke.Params{Blacklist: blacklist, JokeType: jokeType})

	response, err := api.Fetch()

	if err != nil {
		fmt.Printf("Error getting joke %s", err)
	}

	fmt.Println(response.Joke)

}