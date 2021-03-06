package main

import (
	"fmt"
	"io"
)

type UnionFind struct {
	elems []int
}

func NewUnionFind(n int) *UnionFind {
	groups := make([]int, n)
	for idx, _ := range groups {
		groups[idx] = idx
	}
	return &UnionFind{elems: groups}
}

func (ug *UnionFind) Union(a, b int) {
	first, second := ug.elems[a], ug.elems[b]
	// This will change n - 1 elements and take n^2 to run
	for idx, _ := range ug.elems {
		if ug.elems[idx] == first {
			ug.elems[idx] = second
		}
	}
}

func (ug *UnionFind) Connected(a, b int) bool {
	return ug.elems[a] == ug.elems[b]
}

type QuickUnion struct {
	elems []int
}

func NewQuickUnion(n int) *QuickUnion {
	elements := make([]int, n)
	for idx, _ := range elements {
		elements[idx] = idx
	}
	return &QuickUnion{elements}
}

func (qu *QuickUnion) root(idx int) int {
	for ; idx != qu.elems[idx]; idx = qu.elems[idx] {
	}
	return idx
}

func (qu *QuickUnion) Union(a, b int) {
	i, j := qu.root(a), qu.root(b)
	qu.elems[i] = j
}

func (qu *QuickUnion) Connected(a, b int) bool {
	return qu.root(a) == qu.root(b)
}

func main() {
	var n int
	if _, err := fmt.Scanf("%d\n", &n); err != nil {
		panic("Unable to read from standard in")
	}

	ug := NewUnionFind(n)

	for {
		var a, b int
		if _, err := fmt.Scanf("%d %d\n", &a, &b); err == io.EOF {
			break
		}
		ug.Union(a, b)
	}

	for _, e := range ug.elems {
		fmt.Printf("%v\n", e)
	}
}
