package avl

// Tree ...
type Tree struct {
	root     *node    // Корневой узел
	comparer Comparer // Функция сравнения 2-х ключей
}

// NewTree - создает новое дерево. Аргумент - фнкция сравнения двух ключей
func NewTree(comparer Comparer) *Tree {
	return &Tree{
		comparer: comparer,
	}
}

// Insert ...
func (t *Tree) Insert(key, value interface{}) {
	t.root = t.root.insert(key, value, t.comparer)
}

// Find ...
func (t *Tree) Find(key interface{}) (interface{}, bool) {
	return t.root.find(key, t.comparer)
}

// Delete ...
func (t *Tree) Delete(key interface{}) {
	t.root = t.root.delete(key, t.comparer)
}

// Clear ...
func (t *Tree) Clear() {
	t.root = nil
}

// Empty ...
func (t *Tree) Empty() bool {
	return t.root == nil
}

// String ...
func (t *Tree) String() string {
	return t.root.String()
}
