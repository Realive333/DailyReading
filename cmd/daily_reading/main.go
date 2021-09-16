package main

import (
	"fmt"
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

	info, err := information.CreateInfoStruct(os.Getenv("FILEPATH"))
	if err != nil {
		log.Fatal(err)
	}

	message := linebot.NewTextMessage(info.GetMessageStr())
	if _, err := bot.BroadcastMessage(message).Do(); err != nil {
		log.Fatal(err)
	}

	_, err = fmt.Scanln()
	if err != nil {
		log.Fatal(err)
	}
}
