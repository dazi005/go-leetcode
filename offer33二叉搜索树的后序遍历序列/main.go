package main

func verifyPostorder(postorder []int) bool {
	if len(postorder) <= 2 {
		return true
	}

	return juage(postorder, 0, len(postorder) - 1)
}

func juage(postorder []int, start, end int) bool {
	if start >= end {
		return true
	}

	var i int

	for i = start; i < end; i ++ {
		if postorder[i] > postorder[end] {
			break
		}
	}

	for j := i; j < end; j ++ {
		if postorder[j] < postorder[end] {
			return false
		}
	}

	return juage(postorder, start, i - 1) && juage(postorder, i, end - 1)
}

func main() {

}
