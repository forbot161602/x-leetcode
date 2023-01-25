package main

func climbStairs(n int) int {
	if n <= 2 {
		return n
	}
	stepWays := make([]int, n, n)
	stepWays[0] = 1
	stepWays[1] = 2
	for steps := 2; steps < n; steps++ {
		stepWays[steps] = stepWays[steps-1] + stepWays[steps-2]
	}
	return stepWays[n-1]
}
