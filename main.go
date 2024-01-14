package bot

import (
	"flag"
	"fmt"
	"log"
)

func main() {
	t := mustToken()
	fmt.Println(t)
	//tgClient := telegram.New(token)

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
