package leetcode

func corpFlightBookings(bookings [][]int, n int) []int {
	delta := make([]int, n+3)
	for i := range bookings {
		l, r, v := bookings[i][0], bookings[i][1], bookings[i][2]
		delta[l] += v
		delta[r+1] -= v
	}

	ans := make([]int, n+1)
	for i := 1; i <= n; i++ {
		ans[i] = ans[i-1] + delta[i]
	}

	return ans[1:]
}
