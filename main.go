package main

import (
	"log"
	"maunium.net/go/mautrix"
	"maunium.net/go/mautrix/event"
	"os"
)

func login(config Config) (*mautrix.Client, error) {
	client, err := mautrix.NewClient(config.Homeserver, "", "")
	if err != nil {
		return nil, err
	}

	resp, err := client.Login(&mautrix.ReqLogin{
		Type: "m.login.password",
		Identifier: mautrix.UserIdentifier{
			Type: mautrix.IdentifierType("m.id.user"),
			User: config.Username,
		},
		Password: config.Password,
	})
	if err != nil {
		return nil, err
	}

	client.SetCredentials(resp.UserID, resp.AccessToken)
	log.Println("logged in to matrix.org")

	return client, err
}

func main() {
	config, err := parseArgs()
	if err != nil {
		log.Println(err)
		os.Exit(0)
	}

	client, err := login(config)
	if err != nil {
		log.Println(err)
		os.Exit(0)
	}

	bot := GocoinBot{ID: client.UserID}

	syncer := client.Syncer.(*mautrix.DefaultSyncer)
	syncer.OnEventType(event.EventMessage, func(source mautrix.EventSource, matrixEvent *event.Event) {
		bot.HandleMessage(client, *matrixEvent)
	})

	err = client.Sync()
	if err != nil {
		log.Println(err)
	}
}
