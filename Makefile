
.PHONY: build inlama run


run: ./cmd/inlama/main.go inlama
	go run ./cmd/inlama

inlama: ./streams.go ./requests.go ./go.mod

