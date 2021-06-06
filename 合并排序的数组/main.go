package main

func main() {

}

func merge(a []int, m int, b []int, n int) {
	sorted := make([]int, 0, m + n)
	p1, p2 := 0, 0
	for {
		// 边界条件
		if p1 == m {
			sorted = append(sorted, b[p2:]...)
			break
		}

		if p2 == n {
			sorted = append(sorted, a[p1:]...)
			break
		}

		if a[p1] < b[p2] {
			sorted = append(sorted, a[p1])
			p1 ++
		} else {
			sorted = append(sorted, b[p2])
			p2 ++
		}
	}
	copy(a, sorted)
}