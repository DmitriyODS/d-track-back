all: t_assistant

t_assistant:
	go build -mod vendor -v -o ./bin/dtrack_back ./cmd
