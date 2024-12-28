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
	go runScheduler(discord)
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
	case message.Content == "!a":
		advice := getAdvice()
		discord.ChannelMessageSendComplex(message.ChannelID, advice)
	// Get advice
	case message.Content == "!j":
		joke := getJoke()
		discord.ChannelMessageSendComplex(message.ChannelID, joke)
	// Get Cat Image
	case message.Content == "!cat pic":
		catImg := getCatImg(message.Content)
		discord.ChannelMessageSendComplex(message.ChannelID, catImg)
	// Get cat fact
	case message.Content == "!cat":
		catFact := getCatFact(message.Content)
		discord.ChannelMessageSendComplex(message.ChannelID, catFact)
	// Get dog picture
	case message.Content == "!dog pic":
		dogImg := getDogImg(message.Content)
		discord.ChannelMessageSendComplex(message.ChannelID, dogImg)
	// Get dog fact
	case message.Content == "!dog":
		dogFact := getDogFact(message.Content)
		discord.ChannelMessageSendComplex(message.ChannelID, dogFact)
	// Get math fact
	case message.Content == "!math":
		mathFact := getMathFact(message.Content)
		discord.ChannelMessageSendComplex(message.ChannelID, mathFact)
		// Delete schedule for daily quote
	case message.Content == "!q daily delete":
		key := message.ChannelID + "dailyquote"
		deleteServerSchedule(key)
		discord.ChannelMessageSend(message.ChannelID, "Succesfully removed scheduled daily quote.")
	// Get quote daily
	case strings.Contains(message.Content, "!q daily"):
		words := strings.Fields(message.Content)
		if len(words) < 3 {
			discord.ChannelMessageSend(message.ChannelID, "Could not set daily message. Example usage: !q daily 5:00.")
		} else {
			setServerSchedule(message.GuildID, message.ChannelID, "daily", "quote", words[2], true)
			discord.ChannelMessageSend(message.ChannelID, "All times will be set based on EST. Message scheduling was succesful.")
		}
	// Delete schedule for daily joke
	case message.Content == "!j daily delete":
		key := message.ChannelID + "dailyjoke"
		deleteServerSchedule(key)
		discord.ChannelMessageSend(message.ChannelID, "Succesfully removed scheduled daily joke.")
	// Get joke daily
	case strings.Contains(message.Content, "!j daily"):
		words := strings.Fields(message.Content)
		if len(words) < 3 {
			discord.ChannelMessageSend(message.ChannelID, "Could not set daily message. Example usage: !j daily 5:00.")
		} else {
			setServerSchedule(message.GuildID, message.ChannelID, "daily", "joke", words[2], true)
			discord.ChannelMessageSend(message.ChannelID, "All times will be set based on EST. Message scheduling was succesful.")
		}
	// Delete schedule for daily advice
	case message.Content == "!a daily delete":
		key := message.ChannelID + "dailyadvice"
		deleteServerSchedule(key)
		discord.ChannelMessageSend(message.ChannelID, "Succesfully removed scheduled daily advice.")
	// Get advice daily
	case strings.Contains(message.Content, "!a daily"):
		words := strings.Fields(message.Content)
		if len(words) < 3 {
			discord.ChannelMessageSend(message.ChannelID, "Could not set daily advice. Example usage: !a daily 5:00.")
		} else {
			setServerSchedule(message.GuildID, message.ChannelID, "daily", "advice", words[2], true)
			discord.ChannelMessageSend(message.ChannelID, "All times will be set based on EST. Message scheduling was succesful.")
		}
	// Get Breaking Bad Quote
	case message.Content == "!q bb":
		quote := getBreakingBadQuote(message.Content)
		discord.ChannelMessageSendComplex(message.ChannelID, quote)
	// Get Game of Thrones quote
	case message.Content == "!q got":
		quote := getGameOfThronesQuote(message.Content)
		discord.ChannelMessageSendComplex(message.ChannelID, quote)
	// Get Lucifier quote
	case message.Content == "!q lucifier":
		quote := getLucifierQuote(message.Content)
		discord.ChannelMessageSendComplex(message.ChannelID, quote)
	// Get Stranger Things quote
	case message.Content == "!q stranger":
		quote := getStrangerThingsQuote(message.Content)
		discord.ChannelMessageSendComplex(message.ChannelID, quote)
	// Get South Park quote
	case message.Content == "!q south":
		quote := getSouthParkQuote(message.Content)
		discord.ChannelMessageSendComplex(message.ChannelID, quote)
	// Get quote
	case message.Content == "!q":
		quote := getQuote()
		discord.ChannelMessageSendComplex(message.ChannelID, quote)
	// Get basic hi
	case message.Content == "!help" || message.Content == "!h":
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
					{
						Name:   "!q daily hh:mm",
						Value:  "To get daily quote.",
						Inline: true,
					},
					{
						Name:   "!a daily hh:mm",
						Value:  "To get daily advice.",
						Inline: true,
					},
					{
						Name:   "!j daily hh:mm",
						Value:  "To get daily joke.",
						Inline: true,
					},
					{
						Name:   "!q daily delete",
						Value:  "To remove daily quote schedule.",
						Inline: true,
					},
					{
						Name:   "!a daily delete",
						Value:  "To remove daily advice schedule.",
						Inline: true,
					},
					{
						Name:   "!j daily delete",
						Value:  "To remove daily joke schedule.",
						Inline: true,
					},
				}}}}
		discord.ChannelMessageSendComplex(message.ChannelID, embed)
	}
}
