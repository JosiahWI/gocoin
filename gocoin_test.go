package main

import (
	"github.com/JosiahWI/gocoin/mocks"
	"github.com/golang/mock/gomock"
	"maunium.net/go/mautrix/id"
	"testing"
)

func TestGocoin(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	mockClient := mocks.NewMockMautrixClient(mockCtrl)

	mockClient.EXPECT().SendText(id.RoomID("!qporfwt:matrix.org"), "0")
	RegisterBalanceCommand(mockClient)

}
