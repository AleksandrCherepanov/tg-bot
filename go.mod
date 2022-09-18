module github.com/AleksandrCherepanov/tg-bot

go 1.19

require (
	github.com/AleksandrCherepanov/go_telegram v0.0.4
	github.com/gorilla/mux v1.8.0
)

replace github.com/gorilla/mux => ../../go/pkg/mod/github.com/gorilla/mux@v1.8.0

replace github.com/AleksandrCherepanov/go_telegram => ../../go/pkg/mod/github.com/!aleksandr!cherepanov/go_telegram@v0.0.4
