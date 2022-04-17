package main

import (
	"maunium.net/go/mautrix"
	"maunium.net/go/mautrix/id"
	"testing"
)

const exampleRoomID = id.RoomID("!qporfwt:matrix.org")

type MockMautrixClient struct {
	SentValue string
	SentTo    id.RoomID
}

func (client *MockMautrixClient) SendText(roomID id.RoomID, message string) (*mautrix.RespSendEvent, error) {
	client.SentValue = message
	client.SentTo = roomID
	return nil, nil
}

func AssertEqual(t *testing.T, got interface{}, want interface{}) {
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestGocoin(t *testing.T) {
	t.Run("send balance to right room", func(t *testing.T) {
		mockClient := &MockMautrixClient{}

		RegisterBalanceCommand(mockClient, exampleRoomID)

		got := mockClient.SentTo
		want := exampleRoomID
		AssertEqual(t, got, want)
	})

	t.Run("balance is right amount", func(t *testing.T) {
		mockClient := &MockMautrixClient{}

		RegisterBalanceCommand(mockClient, exampleRoomID)

		got := mockClient.SentValue
		want := "0"
		AssertEqual(t, got, want)
	})

}
