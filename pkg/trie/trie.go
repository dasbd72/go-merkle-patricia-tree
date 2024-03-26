package trie

import (
	"bytes"
	"fmt"
)

// Trie represents a Merkle Patricia trie
type Trie struct {
	root Node
}

// New creates a new Merkle Patricia trie
func New() *Trie {
	return &Trie{root: nil}
}

// Hash returns the hash of the root node
func (t *Trie) Hash() []byte {
	if t.root == nil {
		return []byte{}
	}
	return t.root.Hash()
}

// Get returns the value associated with the key.
func (t *Trie) Get(key []byte) ([]byte, bool) {
	return t.get(t.root, keyToNibbles(key))
}

func (t *Trie) get(curr Node, key []byte) ([]byte, bool) {
	switch n := curr.(type) {
	case nil:
		return nil, false
	case *Branch:
		if len(key) != 0 {
			return t.get(n.Children[key[0]], key[1:])
		}
		return n.Val, true
	case *Extension:
		if bytes.HasPrefix(key, n.Key) {
			return t.get(n.Child, key[len(n.Key):])
		}
		return nil, false
	case *Leaf:
		if bytes.Equal(n.Key, key) {
			return n.Val, true
		}
		return nil, false
	default:
		panic(fmt.Sprintf("invalid node type: %T, value: %v", curr, curr))
	}
}

// Put inserts the key-value pair into the trie.
func (t *Trie) Put(key, value []byte) error {
	n, err := t.put(t.root, keyToNibbles(key), value)
	if err != nil {
		return err
	}
	if n != nil {
		t.root = n
	}
	return nil
}

func (t *Trie) put(curr Node, key, value []byte) (Node, error) {
	switch n := curr.(type) {
	case nil:
		// Only when the root is nil
		return &Leaf{Key: key, Val: value}, nil
	case *Branch:
		if len(key) != 0 {
			newChild, err := t.put(n.Children[key[0]], key[1:], value)
			if err != nil {
				return nil, err
			}
			n.Children[key[0]] = newChild
			return n, nil
		}
		n.Val = value
		return n, nil
	case *Extension:
		i := sharedPrefixLen(n.Key, key)
		if i == 0 {
			// No shared prefix
			// Replace the current extension node with a branch node
			b := &Branch{}
			// Insert the current extension node
			b.Children[n.Key[0]] = n
			n.Key = n.Key[1:]
			// Insert the new leaf node
			b.Children[key[0]] = &Leaf{Key: key[1:], Val: value}
			return b, nil
		} else if i == len(n.Key) {
			newChild, err := t.put(n.Child, key[len(n.Key):], value)
			if err != nil {
				return nil, err
			}
			n.Child = newChild
			return n, nil
		} else {
			// Shared prefix
			// Replace the current extension node with an extension node
			e := &Extension{}
			e.Key = n.Key[:i]
			n.Key = n.Key[i:]
			key = key[i:]
			// Append a branch node to the extension node
			b := &Branch{}
			// Insert the current extension node
			b.Children[n.Key[0]] = n
			n.Key = n.Key[1:]
			// Insert the new leaf node
			b.Children[key[0]] = &Leaf{Key: key[1:], Val: value}
			e.Child = b
			return e, nil
		}
	case *Leaf:
		if bytes.Equal(n.Key, key) {
			n.Val = value
			return n, nil
		}
		i := sharedPrefixLen(n.Key, key)
		if i == 0 {
			// No shared prefix
			// Replace the current leaf node with a branch node
			b := &Branch{}
			// Insert the current leaf node
			b.Children[n.Key[0]] = &Leaf{Key: n.Key[1:], Val: n.Val}
			// Insert the new leaf node
			b.Children[key[0]] = &Leaf{Key: key[1:], Val: value}
			return b, nil
		} else {
			// Shared prefix
			// Replace the current leaf node with an extension node
			e := &Extension{}
			e.Key = n.Key[:i]
			n.Key = n.Key[i:]
			key = key[i:]
			// Append a branch node to the extension node
			b := &Branch{}
			// Insert the current leaf node
			if len(n.Key) == 0 {
				b.Val = n.Val
			} else {
				b.Children[n.Key[0]] = &Leaf{Key: n.Key[1:], Val: n.Val}
			}
			// Insert the new leaf node
			if len(key) == 0 {
				b.Val = value
			} else {
				b.Children[key[0]] = &Leaf{Key: key[1:], Val: value}
			}
			e.Child = b
			return e, nil
		}
	default:
		panic(fmt.Sprintf("invalid node type: %T, value: %v", curr, curr))
	}
}

func sharedPrefixLen(a, b []byte) int {
	i := 0
	for i < len(a) && i < len(b) && a[i] == b[i] {
		i++
	}
	return i
}

// String returns a string representation of the trie.
func (t *Trie) String(full bool) string {
	if t.root == nil {
		return ""
	}
	return t.root.String(full)
}
