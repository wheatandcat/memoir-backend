.PHONY: moqgen
moqgen: 
	@-rm -f ./repository/moqs/*.go
	go generate ./repository/...
	go generate ./usecase/...
.PHONY: ci_setup_server
ci_setup_server:
	FIRESTORE_EMULATOR_HOST=localhost:3600 APP_ENV=local go run main.go &