appName = status-checker
required = server/main.go server/getStatus.go server/openBrowser.go

run: $(required)
	@go run $(required)

build: 
	@go build -o $(appName) $(required)
	@echo built $(appName)

clean: $(appName)
	@rm $(appName)
	@echo removed $(appName)