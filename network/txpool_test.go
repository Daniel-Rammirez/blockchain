package network

import (
	"testing"

	"github.com/Daniel-Rammirez/blockchain/core"
	"github.com/stretchr/testify/assert"
)

func TextTxPool(t *testing.T) {
	p := NewTxPool()

	assert.Equal(t, p.Len(), 0)
}

func TestTxPoolAddTx(t *testing.T) {
	p := NewTxPool()
	tx := core.NewTransaction([]byte("hello world"))

	assert.Nil(t, p.Add(tx))
	assert.Equal(t, p.Len(), 1)

	_ = core.NewTransaction([]byte("hello world"))
	assert.Equal(t, p.Len(), 1)

	p.Flush()
	assert.Equal(t, p.Len(), 0)
}
