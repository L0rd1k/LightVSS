package tree

type Tree interface {
	GetValues() []interface{}
	ToString() string
	Clear()
	Empty() bool
	Size() int /* Ouput the size of the tree. */
}
