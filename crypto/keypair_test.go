package crypto

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGeneratePrivateKey(t *testing.T) {
	privKey := GeneratePrivateKey()
	assert.NotNil(t, privKey)

	pubKey := privKey.PublicKey()
	assert.NotNil(t, pubKey)

	address := pubKey.Address()
	assert.NotNil(t, address)
}

func TestKeypairVerifySucess(t *testing.T) {
	privKey := GeneratePrivateKey()
	pubKey := privKey.PublicKey()

	msg := []byte("hello world")
	sig, err := privKey.Sign(msg)
	assert.Nil(t, err)

	sigVerify := sig.Verify(pubKey, msg)
	assert.True(t, sigVerify)
}

func TestKeypairVerifyFail(t *testing.T) {
	privKey := GeneratePrivateKey()
	pubKey := privKey.PublicKey()

	msg := []byte("hello world")
	sig, err := privKey.Sign(msg)
	assert.Nil(t, err)

	// check with other pubKey
	otherPrivKey := GeneratePrivateKey()
	otherPubKey := otherPrivKey.PublicKey()

	sigVerify := sig.Verify(otherPubKey, msg)
	assert.False(t, sigVerify)

	//check with other msg
	otherMsg := []byte("hello world2")
	sigVerify2 := sig.Verify(pubKey, otherMsg)
	assert.False(t, sigVerify2)

	//check with other msg and other key
	sigVerify3 := sig.Verify(otherPubKey, otherMsg)
	assert.False(t, sigVerify3)

}
