appName = status-checker
required = server/main.go server/getStatus.go server/openBrowser.go

help:
	@echo
	@echo list of make commands:
	@echo
	@echo make run - runs the app without building
	@echo make build - packages the app into a single executable called \"$(appName)\"
	@echo make clean - deletes the \"$(appName)\" executable
	@echo

run: $(required)
	@go run $(required)

build: 
	@go build -o $(appName) $(required)
	@echo built $(appName)

clean: $(appName)
	@rm $(appName)
	@echo removed $(appName)
