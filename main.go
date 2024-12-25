package main

import (
	"Did-You-Know-Bot/bot"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	envFile, _ := godotenv.Read(".env")
	// Load environment variables
	botToken, ok := envFile["BOT_TOKEN"]
	if !ok {
		log.Fatal("Must set Discord token as env variable: BOT_TOKEN")
	}

	//Start the bot
	bot.BotToken = botToken
	bot.Run()
}
