package trie

type (
	Branch struct {
		Children [16]Node
		Val      []byte
	}
)

func (n *Branch) Hash() []byte {
	return Hash(n)
}

func (n *Branch) Serialize() []byte {
	return Serialize(n)
}

func (n *Branch) Raw() interface{} {
	raw := make([]interface{}, 17)
	for i, child := range &n.Children {
		if child == nil {
			raw[i] = []byte{}
		} else {
			if len(child.Serialize()) >= 32 {
				raw[i] = child.Hash()
			} else {
				raw[i] = child.Raw()
			}
		}
	}
	raw[16] = n.Val
	return raw
}

// String returns a string representation of the branch node
func (n *Branch) String(full bool) string {
	resp := "[ "
	for i, node := range &n.Children {
		if i > 0 {
			resp += ", "
		}
		if node == nil {
			resp += "<>"
		} else {
			resp += node.String(full)
		}
	}
	resp += ", "
	if n.Val == nil {
		resp += "<>"
	}
	if n.Val != nil {
		resp += string(n.Val)
	}
	return resp + " ]"
}
