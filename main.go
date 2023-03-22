package main

import (
	"flag"
	"log"
	event_consumer "telegram-url-saver/consumer/event-consumer"
	"telegram-url-saver/events/telegram"
	"telegram-url-saver/storage/files"

	tgClient "telegram-url-saver/clients/telegram"
)

const (
	tgBotHost   = "api.telegram.org"
	storagePath = "./storage_files"
	batchSize   = 100
)

func main() {
	eventsProcessor := telegram.New(
		tgClient.New(tgBotHost, mustToken()),
		files.New(storagePath),
	)

	log.Print("service started")
	consumer := event_consumer.New(eventsProcessor, eventsProcessor, batchSize)

	if err := consumer.Start(); err != nil {
		log.Fatal("service is stopped", err)
	}
}

func mustToken() string {
	token := flag.String(
		"tg-bot-token",
		"",
		"token for access to telegram bot",
	)

	flag.Parse()

	if *token == "" {
		log.Fatal("token is not specified")
	}

	return *token
}
