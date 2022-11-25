package pkg

type IterativeDeepening struct {
	IterSearch Search
}

func NewIterativeDeepening(l List) Search {
	return &IterativeDeepening{
		IterSearch: NewSimpleSearch(l),
	}
}

func (d *IterativeDeepening) Search(f Field, limit int) int {
	for i := 1; i <= limit; i++ {
		r := d.IterSearch.Search(f, i)
		if r != -1 {
			return r
		}
	}
	return -1
}
