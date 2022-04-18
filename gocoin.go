package main

import (
	"maunium.net/go/mautrix"
	"maunium.net/go/mautrix/event"
	"maunium.net/go/mautrix/id"
	"strings"
)

type Sender interface {
	SendText(id.RoomID, string) (*mautrix.RespSendEvent, error)
}

type GocoinBot struct {
	ID id.UserID
}

func (bot *GocoinBot) GetID() id.UserID {
	return bot.ID
}

func (bot *GocoinBot) HandleMessage(s Sender, matrixEvent event.Event) {
	content := matrixEvent.Content.Parsed.(*event.MessageEventContent).Body
	if bot.GetID() != matrixEvent.Sender && strings.HasPrefix(content, "g!bal") {
		s.SendText(matrixEvent.RoomID, "0")
	}
}
