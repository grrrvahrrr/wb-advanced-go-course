package main

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

type node struct {
	key   int
	val   int
	left  *node
	right *node
}

func (n *node) insert(newNode *node) {
	switch {
	case newNode == nil:
		return
	case newNode.key < n.key:
		if n.left != nil {
			n.left.insert(newNode)
			return
		}
		n.left = newNode
		return
	case newNode.key > n.key:
		if n.right != nil {
			n.right.insert(newNode)
			return
		}
		n.right = newNode
		return
	}
}

func (n *node) findPrevious(key int) *node {
	if key < n.key {
		if n.left == nil || n.left.key == key {
			return n
		}
		return n.left.findPrevious(key)
	}

	if n.right == nil || n.right.key == key {
		return n
	}
	return n.right.findPrevious(key)
}

func (n *node) forEach(action func(int, int)) {
	if n == nil {
		return
	}
	n.left.forEach(action)
	action(n.key, n.val)
	n.right.forEach(action)
}

type OrderedMap struct {
	node *node
	size int
}

func NewOrderedMap() OrderedMap {
	return OrderedMap{}
}

func (m *OrderedMap) Insert(key, value int) {
	newNode := &node{
		key: key,
		val: value,
	}
	if m.node == nil {
		m.node = newNode
	} else {
		m.node.insert(newNode)
	}
	m.size++
}

func (m *OrderedMap) Erase(key int) {
	if m.node == nil {
		return
	}

	if m.node.key == key {
		node := m.node
		m.node = node.right
		m.node.insert(node.left)
		m.size--
		return
	}

	p := m.node.findPrevious(key)
	if key < p.key {
		if p.left == nil {
			return
		}
		left := p.left
		p.left = left.right
		p.left.insert(left.left)
		m.size--
		return
	}

	if p.right == nil {
		return
	}

	right := p.right
	p.right = right.left
	p.right.insert(right.right)
	m.size--
}

func (m *OrderedMap) Contains(key int) bool {
	for c := m.node; c != nil; {
		switch {
		case key < c.key:
			c = c.left
		case key > c.key:
			c = c.right
		default:
			return true
		}
	}
	return false
}

func (m *OrderedMap) Size() int {
	return m.size
}

func (m *OrderedMap) ForEach(action func(int, int)) {
	m.node.forEach(action)
}

func TestCircularQueue(t *testing.T) {
	data := NewOrderedMap()
	assert.Zero(t, data.Size())

	data.Insert(10, 10)
	data.Insert(5, 5)
	data.Insert(15, 15)
	data.Insert(2, 2)
	data.Insert(4, 4)
	data.Insert(12, 12)
	data.Insert(14, 14)

	assert.Equal(t, 7, data.Size())
	assert.True(t, data.Contains(4))
	assert.True(t, data.Contains(12))
	assert.False(t, data.Contains(3))
	assert.False(t, data.Contains(13))

	var keys []int
	expectedKeys := []int{2, 4, 5, 10, 12, 14, 15}
	data.ForEach(func(key, _ int) {
		keys = append(keys, key)
	})

	assert.True(t, reflect.DeepEqual(expectedKeys, keys))

	data.Erase(15)
	data.Erase(14)
	data.Erase(2)

	assert.Equal(t, 4, data.Size())
	assert.True(t, data.Contains(4))
	assert.True(t, data.Contains(12))
	assert.False(t, data.Contains(2))
	assert.False(t, data.Contains(14))

	keys = nil
	expectedKeys = []int{4, 5, 10, 12}
	data.ForEach(func(key, _ int) {
		keys = append(keys, key)
	})

	assert.True(t, reflect.DeepEqual(expectedKeys, keys))
}
