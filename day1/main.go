package main

import "fmt"


func find2020(values []int) (int, int, int, error) {
	for _, num := range values {
		target := 2020 - num
		for _, num2 := range values {
			target2 := target - num2
			for _, num3 :=range values {
				if num3 == target2 {
					return num, num2, num3, nil
				}
			}

		}
	}
	return 0, 0, 0, fmt.Errorf("no matches found")
}


func MultiplyFind2020(values []int) int {
	a, b, c, err := find2020(values)
	if err !=nil {
		return 0
	}
	return a*b*c
}

func main(){
	ans:= MultiplyFind2020(input)
	fmt.Println(ans)
}
