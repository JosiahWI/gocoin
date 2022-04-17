package main

import (
	"log"
	"maunium.net/go/mautrix"
	"maunium.net/go/mautrix/id"
	"os"
)

const sdRoomID = "!VYobXAAxPBwDxtamiQ:matrix.org"

type GocoinClient interface {
	SendText(id.RoomID, string) (*mautrix.RespSendEvent, error)
}

func RegisterBalanceCommand(client GocoinClient, roomID id.RoomID) {
	client.SendText(roomID, "0")
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

	RegisterBalanceCommand(client, sdRoomID)
}
