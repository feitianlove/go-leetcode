package main

func main() {

}

// leetcode1237_找出给定方程的正整数解
func findSolution(customFunction func(int, int) int, z int) [][]int {
	res := make([][]int, 0)
	i := 1
	j := 1000
	for i <= 1000 && j >= 1 {
		if customFunction(i, j) == z {
			res = append(res, []int{i, j})
			i++
			j--
		} else if customFunction(i, j) > z {
			j--
		} else {
			i++
		}
	}
	return res
}
