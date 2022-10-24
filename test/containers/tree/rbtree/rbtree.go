package rbtree

import (
	"fmt"

	"github.com/L0rd1k/LightVSS/test/containers/tree"
	"github.com/L0rd1k/LightVSS/test/containers/utils"
)

var _ tree.Tree = (*Tree)(nil)

type color bool

const black color = true
const red color = false

/** Implementation of Red-Black Tree. **/
type Tree struct {
	Root       *Node            /** Mother Node from where tree starts. **/
	Comparator utils.Comparator /** Caompare tree keys **/
	size       int              /** Size of tree. **/
}

/** Leaf node structure. **/
type Node struct {
	Key    interface{} /** Node's key. **/
	Value  interface{} /** Node's value. **/
	Left   *Node       /** Left inheritor node **/
	Right  *Node       /** Right inheritor node **/
	Parent *Node       /** Root node for current node **/
	color  color       /** Color of the node  **/
}

/*============================================================*/

/** Create Tree with custom comparator. */
func NewTree(comparator utils.Comparator) *Tree {
	return &Tree{Comparator: comparator}
}

/** Put Node value to tree. **/
func (tree *Tree) Put(key interface{}, value interface{}) {
	var tempNode *Node
	//< If the root node is empty, create root node. **/
	if tree.Root == nil {
		tree.Comparator(key, key)
		tree.Root = &Node{Key: key, Value: value, color: red} //< Create root node value.
		tempNode = tree.Root
	} else {
		node := tree.Root    //< Select root node for start.
		isInsideLoop := true //< Running state for while loop.
		for isInsideLoop {
			compare := tree.Comparator(key, node.Key)
			switch {
			/** If keys are equal, replace existed key value. **/
			case compare == 0:
				{
					node.Key = key
					node.Value = value
					return
				}
			/** If key to insert less then root node. **/
			case compare < 0:
				{
					/** If Left node is empty, paste key there as red, else make left node root **/
					if node.Left == nil {
						node.Left = &Node{Key: key, Value: value, color: red}
						tempNode = node.Left
						isInsideLoop = false
					} else {
						node = node.Left
					}
				}
			/** If key to insert more then root node. **/
			case compare > 0:
				{
					/** If Right node is empty, paste key there as red, else make right node root **/
					if node.Right == nil {
						node.Right = &Node{Key: key, Value: value, color: red}
						tempNode = node.Right
						isInsideLoop = false
					} else {
						node = node.Right
					}
				}
			}
		}
		tempNode.Parent = node /** Current node. **/
	}
	tree.insertInitialState(tempNode)
	tree.size++ /** Increase tree size **/
}

/** Manage color node and where to paste it. **/
func (tree *Tree) insertInitialState(node *Node) {
	//< If parent node doesn't exist, mark current node as black.
	if node.Parent == nil {
		node.color = black
	} else {
		tree.insertParentExist(node)
	}
}

/** Insert node if parent node exist. **/
func (tree *Tree) insertParentExist(node *Node) {
	//< If parent node is black
	if nodeColor(node.Parent) == black {
		return
	}
	uncle := node.uncle()
	if nodeColor(uncle) == red {
		node.Parent.color = black
		uncle.color = black
		node.grandparent().color = red
		tree.insertInitialState(node.grandparent())
	} else {
		grandparent := node.grandparent()
		if node == node.Parent.Right && node.Parent == grandparent.Left {
			tree.rotateLeft(node.Parent)
			node = node.Left
		} else if node == node.Parent.Left && node.Parent == grandparent.Right {
			tree.rotateRight(node.Parent)
			node = node.Right
		}
		node.Parent.color = black
		grandparent = node.grandparent()
		grandparent.color = red
		if node == node.Parent.Left && node.Parent == grandparent.Left {
			tree.rotateRight(grandparent)
		} else if node == node.Parent.Right && node.Parent == grandparent.Right {
			tree.rotateLeft(grandparent)
		}
	}
}

func (tree *Tree) rotateLeft(node *Node) {
	right := node.Right
	tree.replaceNode(node, right)
	node.Right = right.Left
	if right.Left != nil {
		right.Left.Parent = node
	}
	right.Left = node
	node.Parent = right
}

func (tree *Tree) rotateRight(node *Node) {
	left := node.Left
	tree.replaceNode(node, left)
	node.Left = left.Right
	if left.Right != nil {
		left.Right.Parent = node
	}
	left.Right = node
	node.Parent = left
}

func (tree *Tree) replaceNode(old *Node, new *Node) {
	if old.Parent == nil {
		tree.Root = new
	} else {
		if old == old.Parent.Left {
			old.Parent.Left = new
		} else {
			old.Parent.Right = new
		}
	}
	if new != nil {
		new.Parent = old.Parent
	}
}

/** Select node to put data. Check if parents nodes exist. **/
func (node *Node) uncle() *Node {
	if node == nil || node.Parent == nil || node.Parent.Parent == nil {
		return nil
	}
	return node.Parent.sibling()
}

/** Select which sibling node use to paste data (right ot left). **/
func (node *Node) sibling() *Node {
	if node == nil || node.Parent == nil {
		return nil
	}
	if node == node.Parent.Left {
		return node.Parent.Right
	}
	return node.Parent.Left
}

/** Return grandparent node **/
func (node *Node) grandparent() *Node {
	if node != nil && node.Parent != nil {
		return node.Parent.Parent
	}
	return nil
}

/* Check color node */
func nodeColor(node *Node) color {
	if node == nil {
		return black
	}
	return node.color
}

/*============================================================*/

func (tree *Tree) ToString() string {
	return fmt.Sprintln("String output!")
}

func (tree *Tree) GetValues() []interface{} {
	return make([]interface{}, tree.size)
}

func (tree *Tree) Clear() {
}

func (tree *Tree) Empty() bool {
	return false
}

func (tree *Tree) Size() int {
	return tree.size
}
