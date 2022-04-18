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
	t.Run("send balance to right room", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)

		mockSender := mocks.NewMockSender(mockCtrl)

		mockSender.EXPECT().SendText(exampleRoomID, "0").Times(1)

		bot := GocoinBot{}
		bot.balance(mockSender, event.Event{RoomID: exampleRoomID})
	})

}
