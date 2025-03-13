package leetcode

func robotSim(commands []int, obstacles [][]int) int {
	key := func(x, y int) [2]int {
		return [2]int{x, y}
	}

	mObstacles := map[[2]int]bool{}
	for _, o := range obstacles {
		x, y := o[0], o[1]
		mObstacles[key(x, y)] = true
	}

	var x, y, result, dir int
	var dx = []int{0, 1, 0, -1}
	var dy = []int{1, 0, -1, 0}

	for _, cmd := range commands {
		if cmd == -1 {
			dir = (dir + 1) % 4
			continue
		}
		if cmd == -2 {
			dir = (dir + 3) % 4
			continue
		}

		// step on
		for i := 0; i < cmd; i++ {
			nx, ny := x+dx[dir], y+dy[dir]
			if mObstacles[key(nx, ny)] {
				break
			}
			x, y = nx, ny
		}

		dis := x*x + y*y
		if result < dis {
			result = dis
		}
	}

	return result
}
