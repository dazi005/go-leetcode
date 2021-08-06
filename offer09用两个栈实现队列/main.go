package main

type CQueue struct {
	Top  int
	End  int
	List []int
}

func Constructor() CQueue {
	return CQueue{
		Top:  -1,
		End:  -1,
		List: make([]int, 10000),
	}
}

func (c *CQueue) AppendTail(value int) {
	c.End++
	c.List[c.End] = value
}

func (c *CQueue) DeleteHead() int {
	if c.Top+1 > c.End {
		return -1
	}
	c.Top++
	return c.List[c.Top]
}

func main() {

}
