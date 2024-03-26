package trie

import (
	"fmt"
	"testing"
)

func TestTrie(t *testing.T) {
	type sequence []struct {
		key, value, hash string
	}
	testCases := []struct {
		name string
		seq  sequence
	}{
		{
			name: "Balance",
			seq: sequence{
				{"a711355", "45", "a9116924943abeddebf1c0da975ebef7b2006ede340b0f9e18504b65b52948ed"},
				{"a7f9365", "2", "39067a59d2192dbde0af0968ba50ac88d02a41e3a9e06834e6f3490edec03cb5"},
				{"a77d337", "1", "608b7c482ee39d36c1aadbbf38d8d4d7a557dbe5d0484c02a44a8bdb3f87f1e6"},
				{"a77d397", "12", "5838ad5578f346f40d3e6b71f9a82ae6e5198dd39c52e18deec63734da512055"},
				{"a711356", "46", "65e702c5a1fa7d34b1ccfbd130c9a1dc97bff6d83f3237ed06da1a7f801a0754"},
				{"a711357", "47", "0214f87faeb8417f4e5a73df8ee4aaaf904571fb9f859e2e8aa64f6f003ba3bf"},
			},
		},
		{
			name: "Words",
			seq: sequence{
				{"646f", "verb", "014f07ed95e2e028804d915e0dbd4ed451e394e1acfd29e463c11a060b2ddef7"},
				{"646f67", "puppy", "779db3986dd4f38416bfde49750ef7b13c6ecb3e2221620bcad9267e94604d36"},
				{"646f6765", "coin", "ef7b2fe20f5d2c30c46ad4d83c39811bcbf1721aef2e805c0e107947320888b6"},
				{"686f727365", "stallion", "5991bb8c6514148a29db676a14ac506cd2cd5775ace63c30a4fe457715e9ac84"},
			},
		},
		{
			name: "Balance2",
			seq: sequence{
				{"7c3002ad756d76a643cb09cd45409608abb642d9", "10", "b2c77bfa815fbf9f806588c0cc1f1632902ec85b2b8bb5e1241de87f121a8815"},
				{"7c303333756d555643cb09cd45409608abb642d9", "20", "8ce417bd2ffeceb05c76e1cacdd1b0c59918327f8489900022d721d59f6c0efb"},
				{"7c303333756d777643cb09c999409608abb642d9", "30", "9562db763cd5d915cc7eed8935e9aeedd33752e2b207c38cecc6cea527bdaee7"},
				{"7c303333756d777643cb09caaa409608abb642d9", "40", "6aff53fadf211e9b4a452414b7ec0c464b96a23365df6419b8fde90e8a3742fa"},
				{"111102ad756d76a643cb09cd45409608abb642d9", "50", "b3506d16d769a8aaf5e2fe2f4449a673b408472c04ba0e0837aba0bc9d5364cd"},
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			trie := New()
			for _, s := range tc.seq {
				trie.Put([]byte(s.key), []byte(s.value))
				if got := fmt.Sprintf("%x", trie.Hash()); got != s.hash {
					t.Errorf("Hash() = %q, want %q", got, s.hash)
				}
				if got, ok := trie.Get([]byte(s.key)); !ok || string(got) != s.value {
					t.Errorf("Get(%q) = %q, %t, want %q, true", s.key, got, ok, s.value)
				}
			}
		})
	}
}
