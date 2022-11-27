package pkg

const INFTY = 1000000

type Change struct {
	num     int
	coins   []int
	changes map[int]int
}

func NewChange(c []int) *Change {
	return &Change{
		num:     len(c),
		coins:   c,
		changes: map[int]int{},
	}
}

func (c *Change) Solve(total int) int {
	return c.solve(total, c.num-1)
}

func (c *Change) solve(total, num int) int {
	if total == 0 {
		return 0
	} else if total < 0 {
		return INFTY
	}
	if num <= -1 {
		return INFTY
	}
	// c.coins[coins]を使わない場合
	i := c.solveWithCache(total, num-1)
	// c.coins[coins]を使う場合
	j := c.solveWithCache(total-c.coins[num], num) + 1
	if i < j {
		return i
	} else {
		return j
	}
}

func (c *Change) solveWithCache(total, coins int) int {
	x, ok := c.changes[total*c.num+coins]
	if !ok {
		x = c.solve(total, coins)
		c.changes[total*c.num+coins] = x
	}
	return x
}
