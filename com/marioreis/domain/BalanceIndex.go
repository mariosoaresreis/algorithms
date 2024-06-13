package domain

func BalanceIndex(arr []int) int {
	n := len(arr)
	for i := range arr {
		sumForward := 0

		for f := i + 1; f < n; f++ {
			sumForward += arr[f]
		}

		sumBackward := 0

		for b := i - 1; b >= 0; b-- {
			sumBackward += arr[b]
		}

		if sumForward == sumBackward {
			return i
		}

	}

	return -1
}

func Equilibrium2(a []int, n int) int {
	if n == 1 {
		return 0
	}
	front := make([]int, n)
	back := make([]int, n)

	for i := 0; i < n; i++ {
		if i != 0 {
			front[i] = front[i-1] + a[i]
		} else {
			front[i] = a[i]
		}
	}

	for i := n - 1; i > 0; i-- {
		if i <= n-2 {
			back[i] = back[i+1] + a[i]
		} else {
			back[i] = a[i]
		}
	}

	for i := 0; i < n; i++ {
		if front[i] == back[i] {
			return i
		}
	}

	return -1
}

func Equilibrium(arr []int, n int) int {
	var i, j int
	var leftsum, rightsum int

	for i = 0; i < n; i++ {

		leftsum = 0
		for j = 0; j < i; j++ {
			leftsum += arr[j]
		}

		rightsum = 0
		for j = i + 1; j < n; j++ {
			rightsum += arr[j]
		}

		if leftsum == rightsum {
			return i
		}
	}

	return -1
}
