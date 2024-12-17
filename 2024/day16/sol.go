package day16

import (
	"aoc/stl"
	"fmt"
	"math"

	pq "github.com/emirpasic/gods/queues/priorityqueue"
)

func Solution() {
	ip := stl.ReadFile("extc.txt")

	score := bfs(ip)
	sc, spots := dijkstra(ip)
	fmt.Println(score)
	fmt.Println(sc, spots)
}

func bfs(ip []string) int {
	dirs := [4][2]int{{0, 1}, {-1, 0}, {0, -1}, {1, 0}}

	queue := [][4]int{} // row, col, direction, currScore
	for r, line := range ip {
		for c, ch := range line {
			if ch == 'S' {
				queue = append(queue, [4]int{r, c, 0, 0})
			}
		}
	}

	visited := make([][]int, len(ip))
	for i := range visited {
		visited[i] = make([]int, len(ip[0]))
		for j := range visited[i] {
			visited[i][j] = math.MaxInt32
		}
	}

	score := math.MaxInt32

	for len(queue) != 0 {
		top := queue[0]
		queue = queue[1:]

		row, col, d, cs := top[0], top[1], top[2], top[3]
		visited[row][col] = cs

		if ip[row][col] == 'E' {
			score = min(score, cs)
		}

		for dd := -1; dd <= 1; dd++ {
			nd := (d + dd + 4) % 4
			dx, dy := dirs[nd][0], dirs[nd][1]
			nr, nc := row+dx, col+dy

			if ip[nr][nc] != '#' {
				if dd == 0 && visited[nr][nc] > cs+1 {
					queue = append(queue, [4]int{nr, nc, nd, cs + 1})
				} else if visited[nr][nc] > cs+1001 {
					queue = append(queue, [4]int{nr, nc, nd, cs + 1001})
				}
			}
		}
	}
	return score
}

// type tile [4]int
type tile struct {
	x, y, dir int
	score     int
	path      map[pt]struct{}
}

type pt struct {
	x, y int
}

func eww(ip []string) (int, int) {
	dirs := [4][2]int{{0, 1}, {-1, 0}, {0, -1}, {1, 0}}

	queue := pq.NewWith(func(a, b interface{}) int {
		return a.(tile).score - b.(tile).score
	})

	for r, line := range ip {
		for c, ch := range line {
			if ch == 'S' {
				newTile := tile{x: r, y: c, dir: 0, score: 0, path: make(map[pt]struct{})}
				newTile.path[pt{x: r, y: c}] = struct{}{}
				// fmt.Println("start point tile", newTile)
				queue.Enqueue(newTile)
				break
			}
		}
	}

	visited := make([][]int, len(ip)) // min score to reach a tile
	for i := range visited {
		visited[i] = make([]int, len(ip[0]))
		for j := range visited[i] {
			visited[i][j] = math.MaxInt32
		}
	}

	for !queue.Empty() {
		ele, _ := queue.Dequeue()
		top := ele.(tile)

		row, col, d, cs := top.x, top.y, top.dir, top.score
		pathTillNow := top.path

		visited[row][col] = cs
		pathTillNow[pt{x: row, y: col}] = struct{}{}

		if ip[row][col] == 'E' {
			return cs, len(pathTillNow)
		}

		for dd := -1; dd <= 1; dd++ {
			nd := (d + dd + 4) % 4
			dx, dy := dirs[nd][0], dirs[nd][1]
			nr, nc := row+dx, col+dy

			if ip[nr][nc] != '#' {
				if dd == 0 && visited[nr][nc] > cs+1 {
					pathTillNow[pt{x: nr, y: nc}] = struct{}{}
					newTile := tile{x: nr, y: nc, dir: 0, score: cs + 1, path: pathTillNow}
					queue.Enqueue(newTile)
				} else if visited[nr][nc] > cs+1001 {
					pathTillNow[pt{x: nr, y: nc}] = struct{}{}
					newTile := tile{x: nr, y: nc, dir: 0, score: cs + 1001, path: pathTillNow}
					queue.Enqueue(newTile)
				}
			}
		}
	}
	return 0, 0
}
