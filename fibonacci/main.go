package main

import "fmt"

func fibonacci(n int) []int {
	if n < 2 {
		return nil
	}
	ans := []int{1, 1}
	for i := 2; i < n; i++ {
		ans = append(ans, ans[len(ans)-1]+ans[len(ans)-2])
	}
	return ans
}

func answer(n int) []int {
	if n < 2 {
		return make([]int, 0)
	}
	ans := make([]int, n)
	ans[0] = 1
	ans[1] = 1
	for i := 2; i < n; i++ {
		ans[i] = ans[i-1] + ans[i-2]
	}
	return ans
}

func main() {
	var n int

	fmt.Scanln(&n)
	fmt.Printf("%v\n", fibonacci(n))

	fmt.Scanln(&n)
	fmt.Printf("%v\n", answer(n))
}
