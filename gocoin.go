package main

import (
	"log"
	"maunium.net/go/mautrix"
	"maunium.net/go/mautrix/event"
	"maunium.net/go/mautrix/id"
	"os"
)

const sdRoomID = "!VYobXAAxPBwDxtamiQ:matrix.org"

type Sender interface {
	SendText(id.RoomID, string) (*mautrix.RespSendEvent, error)
}

type GocoinBot struct {
}

func (bot *GocoinBot) balance(s Sender, matrixEvent event.Event) {
	s.SendText(matrixEvent.RoomID, "0")
}

func login(config *Config) (*mautrix.Client, error) {
	client, err := mautrix.NewClient(*config.homeserver, "", "")
	if err != nil {
		return nil, err
	}

	resp, err := client.Login(&mautrix.ReqLogin{
		Type: "m.login.password",
		Identifier: mautrix.UserIdentifier{
			Type: mautrix.IdentifierType("m.id.user"),
			User: *config.username,
		},
		Password: *config.password,
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

	bot := GocoinBot{}

	syncer := client.Syncer.(*mautrix.DefaultSyncer)
	syncer.OnEventType(event.EventMessage, func(source mautrix.EventSource, matrixEvent *event.Event) {
		bot.balance(client, *matrixEvent)
	})

	err = client.Sync()
	if err != nil {
		log.Println(err)
	}
}
