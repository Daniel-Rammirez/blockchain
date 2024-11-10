package core

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"io"

	"github.com/Daniel-Rammirez/blockchain/types"
)

type Header struct {
	Version   uint32
	PrevBlock types.Hash
	Timestamp int64
	Height    uint32
	Nonce     uint64
}

func writeBinary(w io.Writer, data interface{}) error {
	return binary.Write(w, binary.LittleEndian, data)
}

func readBinary(r io.Reader, data interface{}) error {
	return binary.Read(r, binary.LittleEndian, data)
}

func (h *Header) EncodeBinary(w io.Writer) error {
	fields := []interface{}{h.Version, h.PrevBlock, h.Timestamp, h.Height, h.Nonce}

	for _, field := range fields {
		if err := writeBinary(w, field); err != nil {
			return err
		}
	}

	return nil
}

func (h *Header) DecodeBinary(r io.Reader) error {
	fields := []interface{}{&h.Version, &h.PrevBlock, &h.Timestamp, &h.Height, &h.Nonce}

	for _, field := range fields {
		if err := readBinary(r, field); err != nil {
			return err
		}
	}

	return nil
}

type Block struct {
	Header
	Transaction []Transaction

	// Cached version of the header hash
	hash types.Hash
}

func (b *Block) Hash() types.Hash {
	buf := &bytes.Buffer{}
	b.Header.EncodeBinary(buf)

	if b.hash.IsZero() {
		b.hash = types.Hash(sha256.Sum256(buf.Bytes()))
	}

	return b.hash
}

func (b *Block) EncodeBinary(w io.Writer) error {

	if err := b.Header.EncodeBinary(w); err != nil {
		return err
	}
	for _, tx := range b.Transaction {
		if err := tx.EncodeBinary(w); err != nil {
			return err
		}
	}

	return nil
}

func (b *Block) DecodeBinary(r io.Reader) error {
	if err := b.Header.DecodeBinary(r); err != nil {
		return err
	}

	for _, tx := range b.Transaction {
		if err := tx.DecodeBinary(r); err != nil {
			return err
		}
	}

	return nil
}
