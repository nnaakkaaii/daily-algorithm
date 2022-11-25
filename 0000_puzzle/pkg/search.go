package pkg

type SearchAlgorithm interface {
	Search(f Field, limit int) int
}
