.PHONY: create_login_yaml
create_login_yaml: 
	go run script/main.go
.PHONY: scenarigo
scenarigo:
	HOST=localhost:8080 scenarigo run --verbose --parallel 1
.PHONY: local_scenarigo
local_scenarigo:
	npx firebase --config ./simulator/firebase.json emulators:exec 'HOST=localhost:8080 scenarigo run --verbose --parallel 1'