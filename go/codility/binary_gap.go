package main

func Solution(N int) int {
	isStart := false
	count := 0
	max := 0
	for N > 0 {
		if N%2 == 1 {
			if !isStart {
				isStart = true
			} else {
				if count > max {
					max = count
				}
				count = 0
			}
		} else {
			if isStart {
				count++
			}
		}
		N = N / 2
	}
	return max
}
