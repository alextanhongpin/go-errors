gen-slide:
	go-present -in main.slide -out ./docs/index.html

present:
	go run golang.org/x/tools/cmd/present
