package pkg

const INFTY = 1000000

type Change struct {
	coins   []int
	changes map[int]int
}

func NewChange(c []int) *Change {
	return &Change{
		coins:   c,
		changes: map[int]int{},
	}
}

func (c *Change) Solve(total int) int {
	return c.solve(total, len(c.coins)-1)
}

func (c *Change) solve(total int, coins int) int {
	if total == 0 {
		return 0
	} else if total < 0 {
		return INFTY
	}
	if coins <= -1 {
		return INFTY
	}
	// c.coins[coins]を使わない場合
	i := c.solveWithCache(total, coins-1)
	// c.coins[coins]を使う場合
	j := c.solveWithCache(total-c.coins[coins], coins) + 1
	if i < j {
		return i
	} else {
		return j
	}
}

func (c *Change) solveWithCache(total int, coins int) int {
	x, ok := c.changes[total*len(c.coins)+coins]
	if !ok {
		x = c.solve(total, coins)
		c.changes[total*len(c.coins)+coins] = x
	}
	return x
}
