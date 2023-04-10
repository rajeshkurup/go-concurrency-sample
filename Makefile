build:
	go build -o bin/stats_generator main.go

test:
	go test ./appmain ./fileops ./appconfig ./appstats ./reports

run:
	bin/stats_generator

clean:
	go clean -modcache
	rm bin/stats_generator
