PKG := github.com/w32blaster/bot-price-monitoring
PKG_LIST=$(shell go list ${PKG}/... | grep -v /vendor/)
FILES := $$(find . -name '*.go' | grep -vE 'vendor')

test:
	@go test -v -race ${PKG_LIST}

build:
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -a -ldflags "-s -w" -o bot-price-monitoring  github.com/w32blaster/bot-price-monitoring

regenerate:
	@echo "Regenerate the code using bot-scaffolding"
	@bot-scaffolding -o generate

up:
	@echo "Run bot"
	@go run `ls *.go` -t="-----" | grep -v _test.go

.POHNY: test build generateTest up