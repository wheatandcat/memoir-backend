.PHONY: create_loggin_yaml
create_login_yaml: 
	go run script/main.go
.PHONY: scenarigo
scenarigo:
	HOST=localhost:8080 scenarigo run
.PHONY: local_scenarigo
local_scenarigo:
	npx firebase --config ./simulator/firebase.json emulators:exec 'HOST=localhost:8080 scenarigo run'