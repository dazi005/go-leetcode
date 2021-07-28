package main

func febi(n int) int {
	if n == 1 {
		return 1
	} else if n == 2 {
		return 2
	}
	return febi(n - 1) + febi(n - 2)
}

func climbStairs(n int) int {
	if n < 3 {
		return n
	}
	t1, t2 := 1, 2
	for i := 3; i < n; i ++ {
		t1, t2 = t2, t1 + t2
	}
	return t2
}

func main()  {
	
}
