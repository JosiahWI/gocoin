package main

import (
	"github.com/JosiahWI/gocoin/mocks"
	"github.com/golang/mock/gomock"
	"maunium.net/go/mautrix/event"
	"maunium.net/go/mautrix/id"
	"testing"
)

const exampleRoomID = id.RoomID("!qporfwt:matrix.org")

func TestGocoin(t *testing.T) {
	t.Run("sends to correct room", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		mockSender := mocks.NewMockSender(mockCtrl)

		mockSender.EXPECT().SendText(exampleRoomID, gomock.Any())

		bot := GocoinBot{}
		bot.HandleMessage(mockSender, event.Event{Sender: "tester", RoomID: exampleRoomID})
	})

	t.Run("bot does not reply to own message", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		mockSender := mocks.NewMockSender(mockCtrl)

		mockSender.EXPECT().SendText(gomock.Any(), gomock.Any()).Times(0)

		bot := GocoinBot{ID: id.UserID("gocoin")}
		bot.HandleMessage(mockSender, event.Event{Sender: "gocoin"})
	})

	t.Run("sends correct balance", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		mockSender := mocks.NewMockSender(mockCtrl)

		mockSender.EXPECT().SendText(gomock.Any(), "0")

		bot := GocoinBot{}
		bot.HandleMessage(mockSender, event.Event{Sender: "tester"})
	})
}
