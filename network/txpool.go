package network

import (
	"github.com/Daniel-Rammirez/blockchain/core"
	"github.com/Daniel-Rammirez/blockchain/types"
)

type TxPool struct {
	transactions map[types.Hash]*core.Transaction
}

func NewTxPool() *TxPool {
	return &TxPool{
		transactions: make(map[types.Hash]*core.Transaction),
	}
}

// Add adds a transaction to the pool, the caller is responsible if the tx already exist.
func (p *TxPool) Add(tx *core.Transaction) error {
	hash := tx.Hash(core.TxHasher{})

	p.transactions[hash] = tx

	return nil
}

func (p *TxPool) Has(hash types.Hash) bool {
	_, ok := p.transactions[hash]

	return ok
}

func (p *TxPool) Len() int {
	return len(p.transactions)
}

func (p *TxPool) Flush() {
	p.transactions = make(map[types.Hash]*core.Transaction)
}
