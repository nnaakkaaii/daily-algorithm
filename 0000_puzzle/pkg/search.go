package pkg

type Search interface {
	Search(f Field, limit int) int
}
