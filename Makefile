.PHONY: go-fix-lint
go-fix-lint:
	find . -print | grep --regex '.*\.go$$' | xargs goimports -w -local "github.com/nnaakkaaii/daily-algorithm"

.PHONY: go-test
go-test:
	go test -cover -race -count=1 ./...
