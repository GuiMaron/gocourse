clean:
	clear && rm -rf ./pb/*

gen:
	make clean && clear && protoc --go_out=pb --go_opt=paths=source_relative --go-grpc_out=pb --go-grpc_opt=paths=source_relative ./proto/*.proto

init:
	clear && rm -f go.mod && rm -f go.sum && go mod init github.com/GuiMaron/gocourse/pcbook

run:
	clear && go run *.go

test:
	clear && go test -v ./serializer

tidy:
	clear && go mod tidy
