package main

type CQueue struct {
	Top int
	End int
	List []int
}


func Constructor() CQueue {
	return CQueue{
		Top: -1,
		End: -1,
		List: make([]int,10000),
	}
}


func (this *CQueue) AppendTail(value int)  {
	this.Top ++
	this.List[this.Top] = value
}


func (this *CQueue) DeleteHead() int {
	if  this.End + 1 > this.Top{
		return -1
	}
	this.End ++
	return this.List[this.End]
}


/**
 * Your CQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AppendTail(value);
 * param_2 := obj.DeleteHead();
 */
