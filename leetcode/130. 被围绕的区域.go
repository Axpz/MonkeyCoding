package leetcode

import "fmt"

// 给你一个 m x n 的矩阵 board ，由若干字符 'X' 和 'O' 组成，捕获 所有 被围绕的区域
//示例 1：
// 输入：board = [["X","X","X","X"],["X","O","O","X"],["X","X","O","X"],["X","O","X","X"]]

// 输出：[["X","X","X","X"],["X","X","X","X"],["X","X","X","X"],["X","O","X","X"]]
// 即：
// X X X X
// X O O X
// X X O X
// X O X X
// ->
// X X X X
// X X X X
// X X X X
// X O X X

// 示例 2：

// 输入：board = [["X"]]

// 输出：[["X"]]

// 分析，替换所有被X包围的O为X

func solve_dfs(board [][]byte) {

	// 方法1. deep first search, 沿着所有边缘开始上下所有展开深度搜索，如果根当前O相连，则标记， 最后处理未被标记的O

	//时间复杂度：O(n×m)，其中 n 和 m 分别为矩阵的行数和列数。深度优先搜索过程中，每一个点至多只会被标记一次。
	//空间复杂度：O(n×m)，其中 n 和 m 分别为矩阵的行数和列数。主要为深度优先搜索的栈的开销。

	m, n := len(board), len(board[0])
	tag := map[int]bool{}

	var dfs func(x, y int)
	dfs = func(x, y int) {
		// if tagged or is X, return
		if x < 0 || x >= m || y < 0 || y >= n || tag[x*n+y] || board[x][y] == 'X' {
			return
		}

		// tag it and go for directions

		tag[x*n+y] = true
		//上下左右
		dfs(x+1, y)
		dfs(x-1, y)
		dfs(x, y+1)
		dfs(x, y-1)
	}

	// 从i边缘开始展开搜索
	for i := 0; i < m; i++ {
		dfs(i, 0)
		dfs(i, n-1)
	}

	// 沿j边缘开始搜索
	for j := 1; j < n-1; j++ {
		dfs(0, j)
		dfs(m-1, j)
	}

	// 找到所有未标记的O替换为X
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if !tag[i*n+j] && board[i][j] == 'O' {
				board[i][j] = 'X'
			}
		}
	}
}

func solve_bfs(board [][]byte) {
	// 方法2. 广度优先搜索
	//时间复杂度：O(n×m)，其中 n 和 m 分别为矩阵的行数和列数。广度优先搜索过程中，每一个点至多只会被标记一次。
	//空间复杂度：O(n×m)，其中 n 和 m 分别为矩阵的行数和列数。主要为广度优先搜索的队列的开销。

	// 我们这次不从边缘开始
	// 直接遍历每个点，找到为O的点，上下左右展开并找出为相邻的值为O的点,
	// 如果所有展开的点中有边缘的点(x = 0/n-1, y = 0/m-1)，则说明这个O不被X包围

	m, n := len(board), len(board[0])

	// tag 的值有如下意义
	// 1 表示已经访问过
	// 2 表示被包围，需要替换为X
	// 隐含条件为 > 0 的时候都被访问过，即==0 表示没有被访问过
	visit := map[int]bool{}

	// 需要替换的片区
	coordinates := [][]int{}

	var bfs = func(x, y int) {
		// fast path
		if visit[x*n+y] {
			return
		}

		// 标记已经访问过
		visit[x*n+y] = true

		// 基于当前点开始展开的满足条件的一组结果
		var result [][]int = [][]int{{x, y}}

		var q [][]int = [][]int{{x, y}}

		for len(q) > 0 {
			// pop
			i, j := q[0][0], q[0][1]
			q = q[1:]

			// 上下左右
			directions := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
			for _, d := range directions {
				// calculate next i, next j
				ni, nj := i+d[0], j+d[1]
				// board 范围内并且是O，且没有被访问过
				if ni >= 0 && ni < m && nj >= 0 && nj < n && board[ni][nj] == 'O' && !visit[ni*n+nj] {
					visit[ni*n+nj] = true
					result = append(result, []int{ni, nj})
					q = append(q, []int{ni, nj})
				}
			}
		}

		// 判断当前点展开的，搜索的结果，如果这组结果，即片区，中有一个点是边缘点则说明这整个片区都与边缘相连，即大陆相连，无效的结果
		for _, r := range result {
			if r[0] == 0 || r[0] == m-1 || r[1] == 0 || r[1] == n-1 {
				return
			}
		}

		for _, r := range result {
			coordinates = append(coordinates, r)
		}
	}

	// 开始遍历寻找为O的点，作bfs搜索
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == 'O' {
				bfs(i, j)
			}
		}
	}

	for _, c := range coordinates {
		x, y := c[0], c[1]
		board[x][y] = 'X'
	}
}

func solve_union_find(board [][]byte) {
	// 方法3. 并查集 "Union-Find" 或 "Disjoint Set Union"（DSU）
	//时间复杂度：O(n×m)，其中 n 和 m 分别为矩阵的行数和列数。广度优先搜索过程中，每一个点至多只会被标记一次。
	//空间复杂度：O(n×m)，其中 n 和 m 分别为矩阵的行数和列数。主要为广度优先搜索的队列的开销。

	// 我们把所有的O定义为两个区间，一个是外部联通区间，其他的事非外部联通区间，即孤岛区间
	// 1. 定义一个虚拟节点作为所有的外部联通区间，定义为 m*n, 并连通所有
	// 2. 非上边的虚拟节点的都为孤岛区间，需要使用X重写

	m, n := len(board), len(board[0])

	// 定义并查集
	var parent []int = make([]int, m*n+2)
	var find func(int) int
	var union func(x, y int)

	for i := range parent {
		parent[i] = i
	}

	find = func(x int) int {
		if x != parent[x] {
			//如果根节点不是自己，则查找并压缩路径
			parent[x] = find(parent[x])
		}

		return parent[x]
	}

	union = func(x, y int) {
		// 把Y合并到X
		rootX, rootY := find(x), find(y)
		if rootX != rootY {
			//不相等， 则合并y到x到根节点
			parent[rootY] = rootX
		}
	}

	outside := m * n

	//遍历所有的节点，并计算所有O的并查集，注意 parent[m*n+1] 为虚拟节点
	dx := []int{-1, 1, 0, 0}
	dy := []int{0, 0, -1, 1}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == 'X' {
				continue
			}

			// 计算前后左右节点
			for k := range dx {
				x, y := i+dx[k], j+dy[k]
				if x < 0 || x >= m || y < 0 || y >= n {
					union(outside, i*n+j)
					continue
				}
				if board[x][y] == 'O' {
					union(i*n+j, x*n+y)
				}
			}
		}
	}

	// union(outside, 3)

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			// 如果不是虚拟连通区域
			if board[i][j] == 'O' && find(i*n+j) != find(outside) {
				board[i][j] = 'X'
			}
			fmt.Printf("\t%d", find(i*n+j))
		}
		fmt.Printf("\n")
	}
}
