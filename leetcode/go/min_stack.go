package main

type MinStack struct {
	items    []int
	minItems []int
}

func Constructor() MinStack {
	data := MinStack{}
	return data
}

func (this *MinStack) Push(item int) {
	this.items = append(this.items, item)
	if len(this.minItems) == 0 || item < this.GetMin() {
		this.minItems = append(this.minItems, item)
	}
}

func (this *MinStack) Pop() int {
	if len(this.items) == 0 {
		return 0
	} else {
		item := this.items[len(this.items)-1]
		this.items = this.items[:len(this.items)-1]
		return item
	}
}

func (this *MinStack) GetMin() int {
	if len(this.minItems) == 0 {
		return -1
	} else {
		return this.minItems[len(this.minItems) - 1]
	}
}

func (this *MinStack) Top() int {
	if len(this.items) == 0 {
		return -1
	} else {
		return this.items[len(this.items)-1]
	}
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.GetMin();
 */

// func main() {
// 	data := Constructor()
// 	data.Push(-2)
// 	data.Push(0)
// 	data.Push(-3)
// 	fmt.Println(data.GetMin())
// 	data.Pop()
// 	fmt.Println(data.Top())
// 	fmt.Println(data.GetMin())

// }
