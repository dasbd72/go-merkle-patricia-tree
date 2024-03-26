package trie

import (
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/rlp"
)

type (
	Node interface {
		Hash() []byte
		Serialize() []byte
		Raw() interface{}
		String(bool) string
	}
)

func NibblesToHex(b []byte) []byte {
	buf := make([]byte, 0, len(b)/2)
	for i := 0; i < len(b); i += 2 {
		buf = append(buf, b[i]*16+b[i+1])
	}
	return buf
}

// Hash returns the hash of a node
func Hash(n Node) []byte {
	return crypto.Keccak256(n.Serialize())
}

// Serialize serializes a node
func Serialize(n Node) []byte {
	raw := n.Raw()
	rlp, err := rlp.EncodeToBytes(raw)
	if err != nil {
		panic(err)
	}
	return rlp
}
