package pkg

const INFTY = 1000000

type Knapsack struct {
	num     int
	values  []int
	weights []int
	burdens map[int]int
}

func NewKnapsack(vs []int, ws []int) *Knapsack {
	return &Knapsack{
		num:     len(vs),
		values:  vs,
		weights: ws,
		burdens: map[int]int{},
	}
}

func (k *Knapsack) Solve(w int) int {
	return k.solve(w, k.num-1)
}

func (k *Knapsack) solve(w, num int) int {
	if w == 0 {
		return 0
	} else if w < 0 {
		return -INFTY
	}
	if num <= -1 {
		return -INFTY
	}
	i := k.solveWithCache(w, num-1)
	j := k.solveWithCache(w-k.weights[num], num-1) + k.values[num]
	if i < j {
		return j
	} else {
		return i
	}
}

func (k *Knapsack) solveWithCache(w, num int) int {
	x, ok := k.burdens[w*k.num+num]
	if !ok {
		x = k.solve(w, num)
		k.burdens[w*k.num+num] = x
	}
	return x
}
