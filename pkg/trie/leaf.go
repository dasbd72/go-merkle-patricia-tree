package trie

import "fmt"

type (
	Leaf struct {
		Key []byte
		Val []byte
	}
)

// Hash returns the hash of the leaf node
func (n *Leaf) Hash() []byte {
	return Hash(n)
}

// Serialize serializes the leaf node
func (n *Leaf) Serialize() []byte {
	return Serialize(n)
}

// Raw returns the raw representation of the leaf node
func (n *Leaf) Raw() interface{} {
	return []interface{}{n.HexKey(), n.Val}
}

// HexKey returns the hex key of the leaf node
func (n *Leaf) HexKey() []byte {
	flag := 2 + len(n.Key)%2
	buf := make([]byte, 0, len(n.Key)/2+2)
	if len(n.Key)%2 == 1 {
		buf = append(buf, byte(flag))
		buf = append(buf, n.Key...)
	} else {
		buf = append(buf, byte(flag))
		buf = append(buf, 0)
		buf = append(buf, n.Key...)
	}
	return NibblesToHex(buf)
}

// String returns a string representation of the leaf node
func (n *Leaf) String(full bool) string {
	if full {
		return fmt.Sprintf("%x: %s", n.HexKey(), string(n.Val))
	}
	return string(n.Val)
}
