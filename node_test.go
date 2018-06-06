package avl

import (
	"fmt"
	"testing"
)

var n *node
var comparer = func(f, s interface{}) Comparison {
	intF, ok := f.(string)
	if !ok {
		return IsLesser
	}
	intS, ok := s.(string)
	if !ok {
		return IsLesser
	}
	switch {
	case intF < intS:
		return IsLesser
	case intF > intS:
		return IsGreater
	}
	return AreEqual
}

func TestInsert(t *testing.T) {

	fmt.Println("Insert A")
	n = n.insert("A", nil, comparer)
	t.Log(n)
	fmt.Println("Insert L")
	n = n.insert("L", nil, comparer)
	t.Log(n)
	fmt.Println("Insert G")
	n = n.insert("G", nil, comparer)
	t.Log(n)
	fmt.Println("Insert O")
	n = n.insert("O", nil, comparer)
	t.Log(n)
	fmt.Println("Insert R")
	n = n.insert("R", nil, comparer)
	t.Log(n)
	fmt.Println("Insert I")
	n = n.insert("I", nil, comparer)
	t.Log(n)
	fmt.Println("Insert T")
	n = n.insert("T", nil, comparer)
	t.Log(n)
	fmt.Println("Insert H")
	n = n.insert("H", nil, comparer)
	t.Log(n)
	fmt.Println("Insert M")
	n = n.insert("M", nil, comparer)
	t.Log(n)
}

func TestFind(t *testing.T) {
	t.Log("Find A")
	if _, found := n.find("A", comparer); !found {
		t.Fatal("expected found")
	}
	t.Log("Find X")
	if _, found := n.find("X", comparer); found {
		t.Fatal("expected not found")
	}
	t.Log("Find L")
	if _, found := n.find("L", comparer); !found {
		t.Fatal("expected found")
	}
	t.Log("Find M")
	if _, found := n.find("M", comparer); !found {
		t.Fatal("expected found")
	}
	t.Log("Find W")
	if _, found := n.find("W", comparer); found {
		t.Fatal("expected not found")
	}

}

func TestDelete(t *testing.T) {
	t.Log("Delete A")
	fmt.Println("Delete A")
	n = n.delete("A", comparer)
	t.Log(n)
	t.Log("Delete L")
	fmt.Println("Delete L")
	n = n.delete("L", comparer)
	t.Log(n)
	t.Log("Delete G")
	fmt.Println("Delete G")
	n = n.delete("G", comparer)
	t.Log(n)
	t.Log("Delete O")
	fmt.Println("Delete O")
	n = n.delete("O", comparer)
	t.Log(n)
	t.Log("Delete R")
	fmt.Println("Delete R")
	n = n.delete("R", comparer)
	t.Log(n)
	t.Log("Delete I")
	fmt.Println("Delete I")
	n = n.delete("I", comparer)
	t.Log(n)
	t.Log("Delete T")
	fmt.Println("Delete T")
	n = n.delete("T", comparer)
	t.Log(n)
	t.Log("Delete H")
	fmt.Println("Delete H")
	n = n.delete("H", comparer)
	t.Log(n)
	t.Log("Delete M")
	fmt.Println("Delete M")
	n = n.delete("M", comparer)
	t.Log(n)
}
