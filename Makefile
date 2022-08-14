build-tasks:
	go build -o ./bin/ ./cmd/tasks/
run-tasks:
	./bin/tasks
deploy:
	sh ./deploy/deploy.sh
