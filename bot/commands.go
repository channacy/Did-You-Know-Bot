package bot

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/bwmarrin/discordgo"
)

type AdviceResponse struct {
	Slip struct {
		ID     int    `json:"id"`
		Advice string `json:"advice"`
	} `json:"slip"`
}

type CatFact struct {
	Fact   string `json:"fact"`
	Length string `json:"length"`
}

type FactAttributes struct {
	Body string `json:"body"`
}

type DogFact struct {
	ID         string         `json:"id"`
	Type       string         `json:"type"`
	Attributes FactAttributes `json:"attributes"`
}

type DogImg struct {
	Image  string `json:"message"`
	Status string `json:"status"`
}

type DogFactData struct {
	Data []DogFact `json:"data"`
}

type BasicQuote struct {
	Quote  string `json:"quote"`
	Author string `json:"author"`
}

type SouthParkQuote struct {
	Quote     string `json:"quote"`
	Character string `json:"character"`
}

type BreakingBadData []BasicQuote

type Quote struct {
	ID           string   `json:"id"`
	Content      string   `json:"content"`
	Author       string   `json:"author"`
	Tags         []string `json:"tags"`
	AuthorSlug   string   `json:"authorSlug"`
	Length       int16    `json:"length"`
	DateAdded    string   `json:"dateAdded"`
	DateModified string   `json:"dateModified"`
}

type House struct {
	Name string `json:"name"`
	Slug string `json:"slug"`
}

type Character struct {
	Name  string `json:"name"`
	Slug  string `json:"slug"`
	House House  `json:"house"`
}

type GameOfThronesData struct {
	Sentence  string    `json:"sentence"`
	Character Character `json:"character"`
}

type LucifierData []BasicQuote
type StrangerThingsData []BasicQuote
type SouthParkData []SouthParkQuote

type Joke struct {
	Type      string `json:"type"`
	Setup     string `json:"setup"`
	Punchline string `json:"punchline"`
	Id        string `json:"id"`
}

const adviceApiUrl string = "https://api.adviceslip.com/advice"
const catFactApiUrl string = "https://catfact.ninja/fact"
const dogFactApiUrl string = "https://dogapi.dog/api/v2/facts"
const dogImgApiUrl string = "https://dog.ceo/api/breeds/image/random"
const mathFactApiUrl string = "http://numbersapi.com/random/math"
const quoteApiUrl string = "http://api.quotable.io/random"
const breakingBadQuoteApiUrl string = "https://api.breakingbadquotes.xyz/v1/quotes"
const gameOfThronesQuoteApiUrl string = "https://api.gameofthronesquotes.xyz/v1/random"
const luciferQuoteApiUrl string = "https://luciferquotes.shadowdev.xyz/api/quotes"
const southParkQuoteApiUrl string = "http://southparkquotes.onrender.com/v1/quotes"
const strangerThingsQuoteApiUrl string = "https://strangerthings-quotes.vercel.app/api/quotes"
const jokeApiUrl string = "https://official-joke-api.appspot.com/jokes/random"

func getAdvice(message string) *discordgo.MessageSend {
	client := http.Client{Timeout: 5 * time.Second}

	response, err := client.Get(adviceApiUrl)
	if err != nil {
		return &discordgo.MessageSend{
			Content: "Sorry, there was an error trying to get a simple piece of advice.",
		}
	}

	// Open HTTP response body
	body, _ := io.ReadAll(response.Body)
	defer response.Body.Close()

	// Convert JSON
	var data AdviceResponse
	json.Unmarshal([]byte(body), &data)

	advice := data.Slip.Advice
	embed := &discordgo.MessageSend{
		Embeds: []*discordgo.MessageEmbed{{
			Type:        discordgo.EmbedTypeRich,
			Title:       "Advice",
			Description: advice,
		}}}
	return embed
}

func getCatFact(message string) *discordgo.MessageSend {
	client := http.Client{Timeout: 5 * time.Second}
	response, err := client.Get(catFactApiUrl)
	if err != nil {
		return &discordgo.MessageSend{
			Content: "Sorry, there was an error trying to get a fact about cats.",
		}
	}

	// Open HTTP response body
	body, _ := io.ReadAll(response.Body)
	defer response.Body.Close()

	// Convert JSON
	var data CatFact
	json.Unmarshal([]byte(body), &data)

	catFact := data.Fact
	embed := &discordgo.MessageSend{
		Embeds: []*discordgo.MessageEmbed{{
			Type:        discordgo.EmbedTypeRich,
			Title:       "Cat Fact",
			Description: catFact,
		}}}
	return embed

}

func getDogFact(message string) *discordgo.MessageSend {
	client := http.Client{Timeout: 5 * time.Second}
	response, err := client.Get(dogFactApiUrl)
	if err != nil {
		return &discordgo.MessageSend{
			Content: "Sorry, there was an error trying to get a fact about dogs.",
		}
	}

	// Open HTTP response body
	body, _ := ioutil.ReadAll(response.Body)
	defer response.Body.Close()

	// Convert JSON
	var data DogFactData
	var dogFact string = ""
	json.Unmarshal([]byte(body), &data)

	if len(data.Data) > 0 {
		dogFact = data.Data[0].Attributes.Body
	} else {
		fmt.Println("No data available.")
	}
	embed := &discordgo.MessageSend{
		Embeds: []*discordgo.MessageEmbed{{
			Type:        discordgo.EmbedTypeRich,
			Title:       "Dog Fact",
			Description: dogFact,
		}}}
	return embed

}

func getMathFact(message string) *discordgo.MessageSend {
	client := http.Client{Timeout: 5 * time.Second}
	response, err := client.Get(mathFactApiUrl)
	if err != nil {
		return &discordgo.MessageSend{
			Content: "Sorry, there was an error trying to get a fact about dogs.",
		}
	}

	// Open HTTP response body
	body, _ := io.ReadAll(response.Body)
	defer response.Body.Close()

	embed := &discordgo.MessageSend{
		Embeds: []*discordgo.MessageEmbed{{
			Type:        discordgo.EmbedTypeRich,
			Title:       "Math Fact",
			Description: string(body),
		}}}
	return embed

}

func getQuote(message string) *discordgo.MessageSend {
	client := http.Client{Timeout: 5 * time.Second}
	response, err := client.Get(quoteApiUrl)
	if err != nil {
		return &discordgo.MessageSend{
			Content: "Sorry, there was an error trying to get a quote.",
		}
	}

	// Open HTTP response body
	body, _ := io.ReadAll(response.Body)
	defer response.Body.Close()

	// Convert JSON
	var data Quote
	json.Unmarshal([]byte(body), &data)

	quote := data.Content
	author := data.Author
	embed := &discordgo.MessageSend{
		Embeds: []*discordgo.MessageEmbed{{
			Type:        discordgo.EmbedTypeRich,
			Title:       "Quote",
			Description: quote,
			Fields: []*discordgo.MessageEmbedField{
				{
					Name:   "Author",
					Value:  author,
					Inline: true,
				},
			}}}}

	return embed

}

func getDogImg(message string) *discordgo.MessageSend {
	client := http.Client{Timeout: 5 * time.Second}
	response, err := client.Get(dogImgApiUrl)
	if err != nil {
		return &discordgo.MessageSend{
			Content: "Sorry, there was an error trying to get a dog image.",
		}
	}

	// Open HTTP response body
	body, _ := io.ReadAll(response.Body)
	defer response.Body.Close()

	// Convert JSON
	var data DogImg
	json.Unmarshal([]byte(body), &data)

	dogImg := data.Image
	embed := &discordgo.MessageSend{
		Embeds: []*discordgo.MessageEmbed{{
			Type:  discordgo.EmbedTypeRich,
			Title: "Dog Image",
			Image: &discordgo.MessageEmbedImage{
				URL: dogImg,
			},
		}}}
	return embed
}

func getBreakingBadQuote(message string) *discordgo.MessageSend {
	client := http.Client{Timeout: 5 * time.Second}
	response, err := client.Get(breakingBadQuoteApiUrl)
	if err != nil {
		return &discordgo.MessageSend{
			Content: "Sorry, there was an error trying to get a Breaking Bad Quote.",
		}
	}

	// Open HTTP response body
	body, _ := io.ReadAll(response.Body)
	defer response.Body.Close()

	// Convert JSON
	var data BreakingBadData
	json.Unmarshal([]byte(body), &data)

	quote := data[0].Quote
	author := data[0].Author
	embed := &discordgo.MessageSend{
		Embeds: []*discordgo.MessageEmbed{{
			Type:        discordgo.EmbedTypeRich,
			Title:       "Breaking Bad",
			Description: quote,
			Fields: []*discordgo.MessageEmbedField{
				{
					Name:   "Character",
					Value:  author,
					Inline: true,
				},
			}}}}
	return embed
}

func getGameOfThronesQuote(message string) *discordgo.MessageSend {
	client := http.Client{Timeout: 5 * time.Second}
	response, err := client.Get(gameOfThronesQuoteApiUrl)
	if err != nil {
		return &discordgo.MessageSend{
			Content: "Sorry, there was an error trying to get a Game of Thrones Quote.",
		}
	}

	// Open HTTP response body
	body, _ := io.ReadAll(response.Body)
	defer response.Body.Close()

	// Convert JSON
	var data GameOfThronesData
	json.Unmarshal([]byte(body), &data)

	quote := data.Sentence
	author := data.Character.Name
	embed := &discordgo.MessageSend{
		Embeds: []*discordgo.MessageEmbed{{
			Type:        discordgo.EmbedTypeRich,
			Title:       "Game of Thrones Quote",
			Description: quote,
			Fields: []*discordgo.MessageEmbedField{
				{
					Name:   "Character",
					Value:  author,
					Inline: true,
				},
			}}}}
	return embed
}

func getLucifierQuote(message string) *discordgo.MessageSend {
	client := http.Client{Timeout: 5 * time.Second}
	response, err := client.Get(luciferQuoteApiUrl)
	if err != nil {
		return &discordgo.MessageSend{
			Content: "Sorry, there was an error trying to get a Lucifier Quote.",
		}
	}

	// Open HTTP response body
	body, _ := io.ReadAll(response.Body)
	defer response.Body.Close()

	// Convert JSON
	var data LucifierData
	json.Unmarshal([]byte(body), &data)

	quote := data[0].Quote
	author := data[0].Author
	embed := &discordgo.MessageSend{
		Embeds: []*discordgo.MessageEmbed{{
			Type:        discordgo.EmbedTypeRich,
			Title:       "Lucifier Quote",
			Description: quote,
			Fields: []*discordgo.MessageEmbedField{
				{
					Name:   "Character",
					Value:  author,
					Inline: true,
				},
			}}}}
	return embed
}

func getStrangerThingsQuote(message string) *discordgo.MessageSend {
	client := http.Client{Timeout: 5 * time.Second}
	response, err := client.Get(strangerThingsQuoteApiUrl)
	if err != nil {
		return &discordgo.MessageSend{
			Content: "Sorry, there was an error trying to get a Stranger Things quote.",
		}
	}

	// Open HTTP response body
	body, _ := io.ReadAll(response.Body)
	defer response.Body.Close()

	// Convert JSON
	var data StrangerThingsData
	json.Unmarshal([]byte(body), &data)

	quote := data[0].Quote
	author := data[0].Author
	embed := &discordgo.MessageSend{
		Embeds: []*discordgo.MessageEmbed{{
			Type:        discordgo.EmbedTypeRich,
			Title:       "Stranger Things Quote",
			Description: quote,
			Fields: []*discordgo.MessageEmbedField{
				{
					Name:   "Character",
					Value:  author,
					Inline: true,
				},
			}}}}
	return embed
}

func getSouthParkQuote(message string) *discordgo.MessageSend {
	client := http.Client{Timeout: 5 * time.Second}
	response, err := client.Get(southParkQuoteApiUrl)
	if err != nil {
		return &discordgo.MessageSend{
			Content: "Sorry, there was an error trying to get a South Park quote.",
		}
	}

	// Open HTTP response body
	body, _ := io.ReadAll(response.Body)
	defer response.Body.Close()

	// Convert JSON
	var data SouthParkData
	json.Unmarshal([]byte(body), &data)

	quote := data[0].Quote
	character := data[0].Character
	embed := &discordgo.MessageSend{
		Embeds: []*discordgo.MessageEmbed{{
			Type:        discordgo.EmbedTypeRich,
			Title:       "South Park Quote",
			Description: quote,
			Fields: []*discordgo.MessageEmbedField{
				{
					Name:   "Character",
					Value:  character,
					Inline: true,
				},
			}}}}
	return embed
}

func getJoke(message string) *discordgo.MessageSend {
	client := http.Client{Timeout: 5 * time.Second}
	response, err := client.Get(jokeApiUrl)
	if err != nil {
		return &discordgo.MessageSend{
			Content: "Sorry, there was an error trying to get a joke.",
		}
	}

	// Open HTTP response body
	body, _ := io.ReadAll(response.Body)
	defer response.Body.Close()

	// Convert JSON
	var data Joke
	json.Unmarshal([]byte(body), &data)

	setup := data.Setup
	punchline := data.Punchline
	embed := &discordgo.MessageSend{
		Embeds: []*discordgo.MessageEmbed{{
			Type:        discordgo.EmbedTypeRich,
			Title:       "Joke",
			Description: setup,
			Fields: []*discordgo.MessageEmbedField{
				{
					Name:   "Punchline",
					Value:  punchline,
					Inline: true,
				},
			}}}}
	return embed
}
