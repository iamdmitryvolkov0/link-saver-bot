package bot

import (
	"bot/clients/telegram"
	"flag"
	"log"
)

const tgBotHost = "api.telegram.org"

func main() {
	t := mustToken()
	tgClient := telegram.New(tgBotHost, t)

	//fetcher := fetcher.New(tgClient)

	//processor := processor.New(tgClient)

	//consumer.Start(fetcher, processor)
}

func mustToken() string {
	t := flag.String(
		"bot-token",
		"",
		"Token to access to telegram bot",
	)

	flag.Parse()

	if *t == "" {
		log.Fatal("token is empty")
	}

	return t
}
