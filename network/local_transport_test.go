package network

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func InitialSetUp() (*LocalTransport, *LocalTransport) {
	tra := NewLocalTransport("A")
	trb := NewLocalTransport("B")

	tra.Connect(trb)
	trb.Connect(tra)

	// Convert trb and tra to *LocalTransport
	localTra := tra
	localTrb := trb

	return localTra, localTrb
}

func TestConnect(t *testing.T) {
	tra := NewLocalTransport("A")
	trb := NewLocalTransport("B")

	tra.Connect(trb)
	trb.Connect(tra)

	localTra := tra
	localTrb := trb

	assert.Equal(t, localTrb.peers[localTra.addr], localTra)
	assert.Equal(t, localTra.peers[localTrb.addr], localTrb)
}

func TestSendMessage(t *testing.T) {
	localTra, localTrb := InitialSetUp()

	localTra.SendMessage(localTrb.addr, []byte("Buy Bitcoin"))

	// consume the message in the localTrb chanel
	receivedRPC := <-localTrb.Consume()

	assert.Equal(t, receivedRPC.Payload, []byte("Buy Bitcoin"))
	assert.Equal(t, receivedRPC.From, localTra.addr)
}
