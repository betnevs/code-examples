package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(math.MaxInt64)
}

func strToInt(str string) int {
	if len(str) == 0 {
		return 0
	}
	var nums []byte
	var isNegative bool
	index := 1
	// 丢弃开头空格
	if str[0] == byte(' ') {
		for index < len(str) {
			if str[index] != byte(' ') {
				break
			}
			index++
		}
		str = str[index:]
	}
	if len(str) == 0 {
		return 0
	}
	// 获取开头的符号
	var isSymbol bool
	if str[0] == byte('-') {
		isNegative = true
		isSymbol = true
		str = str[1:]
		if len(str) == 0 {
			return 0
		}
	}
	if str[0] == byte('+') {
		if isSymbol {
			return 0
		}
		str = str[1:]
		if len(str) == 0 {
			return 0
		}
	}
	if !isDigit(str[0]) {
		return 0
	}
	for i := 0; i < len(str); i++ {
		if !isDigit(str[i]) {
			break
		}
		nums = append(nums, str[i]-48)
	}
	fmt.Println(nums)
	return calculate(nums, isNegative)
}

func isDigit(a byte) bool {
	if a >= 48 && a <= 57 {
		return true
	}
	return false
}

func calculate(nums []byte, isNegative bool) int {
	if len(nums) == 0 {
		return 0
	}
	// 去除前面的 0
	index := 1
	if nums[0] == 0 {
		for index < len(nums) {
			if nums[index] != 0 {
				break
			}
			index++
		}
		nums = nums[index:]
	}
	r := 1
	sum := 0
	for i := len(nums) - 1; i >= 0; i-- {
		sum += int(nums[i]) * r
		if !isNegative && (sum > math.MaxInt32 || r > math.MaxInt32) {
			return math.MaxInt32
		}
		if isNegative && (-sum < math.MinInt32 || -r < math.MinInt32) {
			return math.MinInt32
		}
		r *= 10
	}
	if isNegative {
		sum = -sum
	}
	return sum
}
