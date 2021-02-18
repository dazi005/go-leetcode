package main

func jumpFloor( number int ) int {
	switch number {
	case 1:
		return 1
	case 2:
		return 2
	}
	n := jumpFloor(number - 1) + jumpFloor(number - 2)
	return n
}

func main() {

}
