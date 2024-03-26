package main

import (
	"fmt"
	"math/big"

	"github.com/dasbd72/go-merkle-patricia-trie/pkg/trie"
)

func main() {
	testCase1()
	testCase2()
	testCase3()
}

func testCase1() {
	fmt.Println("Test Case 1")

	sequence := []struct {
		key []byte
		val []byte
	}{
		{[]byte("a711355"), []byte("45")},
		{[]byte("a77d337"), []byte("1")},
		{[]byte("a7f9365"), []byte("2")},
		{[]byte("a77d397"), []byte("12")},
	}
	want := "5838ad5578f346f40d3e6b71f9a82ae6e5198dd39c52e18deec63734da512055"

	trie := trie.New()
	for _, s := range sequence {
		trie.Put(s.key, s.val)
	}

	got := fmt.Sprintf("%x", trie.Hash())
	fmt.Printf("got: %s\n", got)
	fmt.Println(trie.String(false))
	if got == want {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
	fmt.Println()
}

func testCase2() {
	fmt.Println("Test Case 2")

	sequence := []struct {
		key []byte
		val []byte
	}{
		{[]byte("7c3002ad756d76a643cb09cd45409608abb642d9"), []byte("10")},
		{[]byte("7c303333756d555643cb09cd45409608abb642d9"), []byte("20")},
		{[]byte("7c303333756d777643cb09c999409608abb642d9"), []byte("30")},
		{[]byte("7c303333756d777643cb09caaa409608abb642d9"), []byte("40")},
		{[]byte("111102ad756d76a643cb09cd45409608abb642d9"), []byte("50")},
	}
	want := "b3506d16d769a8aaf5e2fe2f4449a673b408472c04ba0e0837aba0bc9d5364cd"

	t := trie.New()
	for _, s := range sequence {
		fmt.Printf("key: %s, val: %s\n", s.key, s.val)

		t.Put(s.key, s.val)
		got, ok := t.Get(s.key)
		if !ok {
			panic(fmt.Sprintf("key not found: %s", s.key))
		}
		if string(got) != string(s.val) {
			panic(fmt.Sprintf("value mismatch: %s != %s", got, s.val))
		}

		fmt.Println(t.String(true))
		fmt.Println()
	}

	got := fmt.Sprintf("%x", t.Hash())
	fmt.Printf("got: %s\n", got)
	fmt.Println(t.String(false))
	if got == want {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
	fmt.Println()
}

func testCase3() {
	fmt.Println("Test Case 3")

	sequence := []struct {
		key []byte
		val []byte
	}{
		{[]byte("7c3002ad756d76a643cb09cd45409608abb642d9"), []byte("10")},
		{[]byte("7c303333756d555643cb09cd45409608abb642d9"), []byte("20")},
		{[]byte("7c303333756d777643cb09c999409608abb642d9"), []byte("30")},
		{[]byte("7c303333756d777643cb09caaa409608abb642d9"), []byte("40")},
		{[]byte("111102ad756d76a643cb09cd45409608abb642d9"), []byte("50")},
	}
	transactions := []struct {
		from []byte
		to   []byte
		val  []byte
	}{
		{[]byte("7c3002ad756d76a643cb09cd45409608abb642d9"), []byte("7c303333756d777643cb09caaa409608abb642d9"), []byte("2")},
		{[]byte("7c303333756d777643cb09c999409608abb642d9"), []byte("11113333756d76a643cb09cd45409608abb642d9"), []byte("6")},
	}
	want := "eff402b46c2b81e230797cf224c5440aefde9335594271e19da8c75ecc476d08"
	results := []struct {
		key []byte
		val []byte
	}{
		{[]byte("7c3002ad756d76a643cb09cd45409608abb642d9"), []byte("8")},
		{[]byte("7c303333756d555643cb09cd45409608abb642d9"), []byte("20")},
		{[]byte("7c303333756d777643cb09c999409608abb642d9"), []byte("24")},
		{[]byte("7c303333756d777643cb09caaa409608abb642d9"), []byte("42")},
		{[]byte("111102ad756d76a643cb09cd45409608abb642d9"), []byte("50")},
		{[]byte("11113333756d76a643cb09cd45409608abb642d9"), []byte("6")},
	}

	t := trie.New()
	for _, s := range sequence {
		t.Put(s.key, s.val)
	}

	for _, trx := range transactions {
		val := new(big.Int)
		val.SetString(string(trx.val), 10)

		src, ok := t.Get(trx.from)
		if !ok {
			panic(fmt.Sprintf("key not found: %s", trx.from))
		}
		srcVal := new(big.Int)
		srcVal.SetString(string(src), 10)
		srcVal.Sub(srcVal, val)
		t.Put(trx.from, []byte(srcVal.Text(10)))

		dst, ok := t.Get(trx.to)
		dstVal := big.NewInt(0)
		if ok {
			dstVal.SetString(string(dst), 10)
		}
		dstVal.Add(dstVal, val)
		t.Put(trx.to, []byte(dstVal.Text(10)))
	}

	for _, r := range results {
		got, ok := t.Get(r.key)
		if !ok {
			panic(fmt.Sprintf("key not found: %s", r.key))
		}
		if string(got) != string(r.val) {
			panic(fmt.Sprintf("value mismatch: %s != %s", got, r.val))
		}
	}

	got := fmt.Sprintf("%x", t.Hash())
	fmt.Printf("got: %s\n", got)
	fmt.Println(t.String(false))
	if got == want {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
	fmt.Println()
}
