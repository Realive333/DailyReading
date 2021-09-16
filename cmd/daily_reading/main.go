package main

import (
	"log"
	"os"

	"github.com/Realive333/DailyReading/internal/information"

	"github.com/line/line-bot-sdk-go/linebot"
)

func main() {
	bot, err := linebot.New(
		os.Getenv("CHANNEL_SECRET"),
		os.Getenv("CHANNEL_TOKEN"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Starting Line Bot, DailyReading v0.1a")

	info, err := information.CreateInfoStruct(os.Getenv("FILEPATH"))
	if err != nil {
		log.Fatal(err)
	}

	message := linebot.NewTextMessage(info.GetMessageStr())
	if _, err := bot.BroadcastMessage(message).Do(); err != nil {
		log.Fatal(err)
	}
}
