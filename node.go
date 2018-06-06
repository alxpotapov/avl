package avl

import (
	"fmt"
)

type node struct {
	height   int
	key      interface{}
	value    interface{}
	children [2]*node
}

func newNode(key, value interface{}) *node {
	return &node{key: key, value: value, height: 1}
}

func (n *node) insert(key, value interface{}, comparer Comparer) *node {
	// если не найден узел, то возвращаем новый
	if n == nil {
		return newNode(key, value)
	}
	var offset offset
	switch comparer(n.key, key) {
	case IsGreater:
		offset = left
	case IsLesser:
		offset = right
	case AreEqual:
		// либо заменяем value в случае set(map),
		// либо добавляем в список в случае multiset(multimap)
		return n
	}
	n.children[offset] = n.children[offset].insert(key, value, comparer)
	return n.fixBalanceViolation()
}

func (n *node) find(key interface{}, comparer Comparer) (interface{}, bool) {
	// если не найден узел, то возвращаем новый
	if n == nil {
		return nil, false
	}
	var offset offset
	switch comparer(n.key, key) {
	case IsGreater:
		offset = left
	case IsLesser:
		offset = right
	case AreEqual:
		return n.value, true
	}
	return n.children[offset].find(key, comparer)
}

func (n *node) delete(key interface{}, comparer Comparer) *node {
	if n == nil {
		return nil
	}
	var offset offset
	switch comparer(n.key, key) {
	case IsGreater:
		offset = left
	case IsLesser:
		offset = right
	case AreEqual:
		return n.splice(comparer)
	}
	n.children[offset] = n.children[offset].delete(key, comparer)
	return n.fixBalanceViolation()
}

// String - вывод на экран
func (n *node) String() string {
	if n == nil {
		return ""
	}
	s := ""
	s += n.children[left].String()
	s += fmt.Sprintf("%v:%d", n.key, n.height)
	s += n.children[right].String()
	return "(" + s + ")"
}

func (n *node) splice(comparer Comparer) *node {
	//Удалить узел и вернуть nil
	if n.children[left] == nil && n.children[right] == nil {
		return nil
	}
	//Удалить узел и вернуть его левую подветвь
	if n.children[right] == nil {
		return n.children[left]
	}
	//Удалить узел и вернуть его правую подветвь
	if n.children[left] == nil {
		return n.children[right]
	}
	tempNode := n.children[left].findMax()
	n.key = tempNode.key
	n.value = tempNode.value
	n.children[left] = n.children[left].delete(n.key, comparer)
	return n.fixBalanceViolation()
}

// findMax - вернуть узел с максимальным значением из левой подветви
func (n *node) findMax() *node {
	if n.children[right] != nil {
		n = n.children[right].findMax()
	}
	return n
}

// getBalance - оберткой для поля balance
func (n *node) getHeight() int {
	if n == nil {
		return 0
	}
	return n.height
}

// balanceFactor ...
func (n *node) balanceFactor() int {
	return n.children[right].getHeight() - n.children[left].getHeight()
}

// fixBalance - восстанавливает корректное значение поля balance заданного узла
func (n *node) fixHeight() {
	hl := n.children[left].getHeight()
	hr := n.children[right].getHeight()
	if hl > hr {
		n.height = hl
	} else {
		n.height = hr
	}
	n.height++
}

// rotate ...
func (n *node) rotate(offset offset) *node {
	//	fmt.Println("Rotate", offset)
	//	fmt.Println(" -src", n)
	root := n.children[offset.other()]
	n.children[offset.other()] = root.children[offset]
	root.children[offset] = n
	//	fmt.Println(" -dst", n)
	//	fmt.Println("Fix balance")
	//	fmt.Println(" -src", n)
	n.fixHeight()
	root.fixHeight()
	//	fmt.Println(" -dst", n)
	return root
}

func (n *node) fixBalanceViolation() *node {
	n.fixHeight()
	switch n.balanceFactor() {
	case 2:
		if n.children[right].balanceFactor() < 0 {
			n.children[right] = n.children[right].rotate(right)
		}
		return n.rotate(left)
	case -2:
		if n.children[left].balanceFactor() > 0 {
			n.children[left] = n.children[left].rotate(left)
		}
		return n.rotate(right)
	}
	return n
}
