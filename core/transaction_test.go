package core

import (
	"testing"

	"github.com/Daniel-Rammirez/blockchain/crypto"
	"github.com/stretchr/testify/assert"
)

func TestSignTransaction(t *testing.T) {
	privKey := crypto.GeneratePrivateKey()
	tx := &Transaction{
		Data: []byte("foo"),
	}

	assert.Nil(t, tx.Sign(privKey))
	assert.NotNil(t, tx.Signature)
}

func TestVerifyTransaction(t *testing.T) {
	privKey := crypto.GeneratePrivateKey()
	// pubKey := privKey.PublicKey()

	tx := &Transaction{
		Data: []byte("foo"),
	}

	assert.Nil(t, tx.Signature)
	assert.Nil(t, tx.Sign(privKey))
	assert.NotNil(t, tx.From)
	assert.NotNil(t, tx.Signature)
	assert.Nil(t, tx.Verify())

	// we create another priv and pub key
	otherPrivKey := crypto.GeneratePrivateKey()
	// and then we tamper the existing tx
	// trying to hack the system
	tx.From = otherPrivKey.PublicKey()

	// this must be not nil, bc should throw an error
	// the reason s we have a sig with anohter pubKey
	assert.NotNil(t, tx.Verify())

}

func randomTxWithSignature(t *testing.T) *Transaction {
	privKey := crypto.GeneratePrivateKey()
	tx := &Transaction{
		Data: []byte("foo"),
	}

	assert.Nil(t, tx.Sign(privKey))

	return tx
}
