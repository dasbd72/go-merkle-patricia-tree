package trie

import (
	"fmt"
)

type (
	Extension struct {
		Key   []byte
		Child Node
	}
)

func (n *Extension) Hash() []byte {
	return Hash(n)
}

func (n *Extension) Serialize() []byte {
	return Serialize(n)
}

func (n *Extension) Raw() interface{} {
	raw := make([]interface{}, 2)
	raw[0] = n.HexKey()
	if len(n.Child.Serialize()) >= 32 {
		raw[1] = n.Child.Hash()
	} else {
		raw[1] = n.Child.Raw()
	}
	return raw
}

func (n *Extension) HexKey() []byte {
	flag := len(n.Key) % 2
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

// String returns a string representation of the extension node
func (n *Extension) String(full bool) string {
	if full {
		return fmt.Sprintf("[%x: %s]", n.HexKey(), n.Child.String(full))
	}
	return fmt.Sprintf("[%s, %s]", nibblesToString(n.Key), n.Child.String(full))
}
