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
	case strings.Contains(message.Content, "!advice"):
		advice := getAdvice(message.Content)
		discord.ChannelMessageSendComplex(message.ChannelID, advice)
	// Get cat fact
	case strings.Contains(message.Content, "!cat"):
		catFact := getCatFact(message.Content)
		discord.ChannelMessageSendComplex(message.ChannelID, catFact)
	// Get dog fact
	case strings.Contains(message.Content, "!dog"):
		dogFact := getDogFact(message.Content)
		discord.ChannelMessageSendComplex(message.ChannelID, dogFact)
	// Get math fact
	case strings.Contains(message.Content, "!math"):
		mathFact := getMathFact(message.Content)
		discord.ChannelMessageSendComplex(message.ChannelID, mathFact)
	// Get Breaking Bad Quote
	case strings.Contains(message.Content, "!quote bb"):
		quote := getBreakingBadQuote(message.Content)
		discord.ChannelMessageSendComplex(message.ChannelID, quote)
	// Get quote
	case strings.Contains(message.Content, "!quote"):
		quote := getQuote(message.Content)
		discord.ChannelMessageSendComplex(message.ChannelID, quote)
	// Get dog picture
	case strings.Contains(message.Content, "!dog pic"):
		dogImg := getDogImg(message.Content)
		discord.ChannelMessageSendComplex(message.ChannelID, dogImg)
	// Get basic hi
	case strings.Contains(message.Content, "!dyk"):
		discord.ChannelMessageSend(message.ChannelID, "Hi there!")
	}
}
