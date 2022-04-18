package main

import (
	"github.com/JosiahWI/gocoin/mocks"
	"github.com/golang/mock/gomock"
	"maunium.net/go/mautrix/event"
	"maunium.net/go/mautrix/id"
	"testing"
)

const exampleRoomID = id.RoomID("!qporfwt:matrix.org")
const exampleSenderID = id.UserID("@tester:matrix.org")

var validBalanceCheckEvent event.Event = event.Event{
	Sender: exampleSenderID,
	RoomID: exampleRoomID,
	Content: event.Content{
		Parsed: &event.MessageEventContent{
			Body: "g!bal",
		},
	},
}

func TestGocoin(t *testing.T) {
	t.Run("sends to correct room", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		mockSender := mocks.NewMockSender(mockCtrl)

		mockSender.EXPECT().SendText(exampleRoomID, gomock.Any())

		bot := GocoinBot{}
		bot.HandleMessage(mockSender, validBalanceCheckEvent)
	})

	t.Run("bot does not reply to own message", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		mockSender := mocks.NewMockSender(mockCtrl)

		mockSender.EXPECT().SendText(gomock.Any(), gomock.Any()).Times(0)

		bot := GocoinBot{ID: exampleSenderID}
		bot.HandleMessage(mockSender, validBalanceCheckEvent)
	})

	t.Run("sends correct balance", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		mockSender := mocks.NewMockSender(mockCtrl)

		mockSender.EXPECT().SendText(gomock.Any(), "0")

		bot := GocoinBot{}
		bot.HandleMessage(mockSender, validBalanceCheckEvent)
	})

	t.Run("does not reply to message not starting with g!bal", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		mockSender := mocks.NewMockSender(mockCtrl)

		mockSender.EXPECT().SendText(gomock.Any(), gomock.Any()).Times(0)

		bot := GocoinBot{}
		content := event.Content{Parsed: &event.MessageEventContent{Body: "hello world"}}

		bot.HandleMessage(mockSender, event.Event{Sender: "tester", Content: content})
	})
}
