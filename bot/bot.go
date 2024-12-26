package bot

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"

	"github.com/bwmarrin/discordgo"
)

var (
	BotToken string
)

func Run() {
	// Create new Discord Session
	discord, err := discordgo.New("Bot " + BotToken)
	if err != nil {
		log.Fatal(err)
	}

	// Add event handler
	discord.AddHandler(newMessage)

	// Open session via websocket
	err = discord.Open()
	if err != nil {
		fmt.Println("Error opening connection:", err)
		return
	}
	defer discord.Close()

	// Run until code is terminated
	fmt.Println("Bot running...")
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c

	// Close the Discord session
	//discord.Close()
}

func newMessage(discord *discordgo.Session, message *discordgo.MessageCreate) {

	// Ignore its own bot messaage
	if message.Author.ID == discord.State.User.ID {
		return
	}

	// Respond to messages
	switch {
	// Get advice
	case strings.Contains(message.Content, "!a"):
		advice := getAdvice(message.Content)
		discord.ChannelMessageSendComplex(message.ChannelID, advice)
	// Get advice
	case strings.Contains(message.Content, "!j"):
		joke := getJoke(message.Content)
		discord.ChannelMessageSendComplex(message.ChannelID, joke)
	// Get Cat Image
	case strings.Contains(message.Content, "!cat pic"):
		catImg := getCatImg(message.Content)
		discord.ChannelMessageSendComplex(message.ChannelID, catImg)
	// Get cat fact
	case strings.Contains(message.Content, "!cat"):
		catFact := getCatFact(message.Content)
		discord.ChannelMessageSendComplex(message.ChannelID, catFact)
	// Get dog picture
	case strings.Contains(message.Content, "!dog pic"):
		dogImg := getDogImg(message.Content)
		discord.ChannelMessageSendComplex(message.ChannelID, dogImg)
	// Get dog fact
	case strings.Contains(message.Content, "!dog"):
		dogFact := getDogFact(message.Content)
		discord.ChannelMessageSendComplex(message.ChannelID, dogFact)
	// Get math fact
	case strings.Contains(message.Content, "!math"):
		mathFact := getMathFact(message.Content)
		discord.ChannelMessageSendComplex(message.ChannelID, mathFact)
	// Get Breaking Bad Quote
	case strings.Contains(message.Content, "!q bb"):
		quote := getBreakingBadQuote(message.Content)
		discord.ChannelMessageSendComplex(message.ChannelID, quote)
	// Get Game of Thrones quote
	case strings.Contains(message.Content, "!q got"):
		quote := getGameOfThronesQuote(message.Content)
		discord.ChannelMessageSendComplex(message.ChannelID, quote)
	// Get Lucifier quote
	case strings.Contains(message.Content, "!q lucifier"):
		quote := getLucifierQuote(message.Content)
		discord.ChannelMessageSendComplex(message.ChannelID, quote)
	// Get Stranger Things quote
	case strings.Contains(message.Content, "!q stranger"):
		quote := getStrangerThingsQuote(message.Content)
		discord.ChannelMessageSendComplex(message.ChannelID, quote)
	// Get South Park quote
	case strings.Contains(message.Content, "!q south"):
		quote := getSouthParkQuote(message.Content)
		discord.ChannelMessageSendComplex(message.ChannelID, quote)
	// Get quote
	case strings.Contains(message.Content, "!q"):
		quote := getQuote(message.Content)
		discord.ChannelMessageSendComplex(message.ChannelID, quote)
	// Get basic hi
	case strings.Contains(message.Content, "!help") || strings.Contains(message.Content, "!h"):
		embed := &discordgo.MessageSend{
			Embeds: []*discordgo.MessageEmbed{{
				Type:        discordgo.EmbedTypeRich,
				Title:       "How to Use Did You Know? Bot",
				Description: "Commands",
				Fields: []*discordgo.MessageEmbedField{
					{
						Name:   "!help",
						Value:  "To view instructions.",
						Inline: true,
					},
					{
						Name:   "!a",
						Value:  "To get advice.",
						Inline: true,
					},
					{
						Name:   "!j",
						Value:  "To get a joke.",
						Inline: true,
					},
					{
						Name:   "!q",
						Value:  "To get a quote.",
						Inline: true,
					},
					{
						Name:   "!math",
						Value:  "To get a math fact.",
						Inline: true,
					},
					{
						Name:   "!cat",
						Value:  "To get a cat fact.",
						Inline: true,
					},
					{
						Name:   "!dog",
						Value:  "To get a dog fact.",
						Inline: true,
					},
					{
						Name:   "!dog pic",
						Value:  "To get a dog image.",
						Inline: true,
					},
					{
						Name:   "!q bb",
						Value:  "To get a Breaking Bad quote.",
						Inline: true,
					},
					{
						Name:   "!q got",
						Value:  "To get a Game of Thrones quote.",
						Inline: true,
					},
					{
						Name:   "!q lucifier",
						Value:  "To get a Lucifier quote.",
						Inline: true,
					},
					{
						Name:   "!q stranger",
						Value:  "To get a Stranger Things quote.",
						Inline: true,
					},
					{
						Name:   "!q south",
						Value:  "To get a South Park quote.",
						Inline: true,
					},
				}}}}
		discord.ChannelMessageSendComplex(message.ChannelID, embed)
	}
}
