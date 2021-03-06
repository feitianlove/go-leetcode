package main

import "fmt"

func main() {
	fmt.Println(reverseStr("abcdefg", 2))
}

func reverseStr(s string, k int) string {
	arr := []byte(s)
	for i := 0; i < len(s); i = i + k {
		if i%(2*k) == 0 {
			j := i + k
			if len(arr) < j {
				j = len(arr)
			}
			reverse(arr[i:j])
		}
	}
	return string(arr)
}

func reverse(arr []byte) {
	i, j := 0, len(arr)-1
	for i < j {
		arr[i], arr[j] = arr[j], arr[i]
		i++
		j--
	}
}
