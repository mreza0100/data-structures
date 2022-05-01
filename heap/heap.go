package heap

import "fmt"

type Tree struct {
	value int

	left  *Tree
	right *Tree
}

func newHeapTree(n int) *Tree {
	return &Tree{value: n}
}

func NewHeapTree(root int) *Tree {
	return newHeapTree(root)
}

var (
	truthy = true
	falsy  = false
)

func (t *Tree) Print() {
	t.print(0, nil)
}

func (t *Tree) print(deep int, isRight *bool) {
	if t.right != nil {
		t.right.print(deep+1, &truthy)
	}
	if t.left != nil {
		t.left.print(deep+1, &falsy)
	}

	for i := 0; i < deep; i++ {
		fmt.Print("\t")
	}
	if isRight != nil {
		if *isRight {
			fmt.Print("/")
		} else {
			fmt.Print("\\")
		}
	}
	fmt.Printf("%d\n", t.value)
}

func (t *Tree) push(n int) {
	if t.value == 0 {
		t.value = n
		return
	}

	if n > t.value {
		if t.right == nil {
			t.right = newHeapTree(n)
		} else {
			t.right.push(n)
		}
	} else {
		if t.left == nil {
			t.left = newHeapTree(n)
		} else {
			t.left.push(n)
		}
	}

	t.Balance()
}

func (t *Tree) Push(n int) {
	t.push(n)
}

func (t *Tree) Get(value int) (int, bool) {
	p := t.get(value)

	if p != nil {
		return *p, true
	}
	return 0, false
}

func (t *Tree) get(value int) *int {
	if t.value == value {
		return &value
	}

	if t.right != nil {
		if temp := t.right.get(value); temp != nil {
			return temp
		}
	}
	if t.left != nil {
		if temp := t.left.get(value); temp != nil {
			return temp
		}
	}

	return nil
}

func (t *Tree) Balance() {
	for t.balance() {
	}
}

func (t *Tree) balance() bool {
	if t.right != nil {
		if t.right.value > t.value {
			t.right.value, t.value = t.value, t.right.value
		}
		t.right.balance()
		return false
	}
	if t.left != nil {
		if t.left.value > t.value {
			t.left.value, t.value = t.value, t.left.value
		}
		t.left.balance()
		return false
	}

	return true
}
