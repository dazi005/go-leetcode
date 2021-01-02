package main

func minArray(numbers []int) int {
	var min = numbers[0]
	for i:=0;i<len(numbers);i++{
		if numbers[i] < min{
			min = numbers[i]
		}
	}
	return min
}



func main(){

}

