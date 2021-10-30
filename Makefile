.PHONY: moqgen
moqgen: 
	@-rm -f ./repository/moqs/*.go
	go generate ./repository/...
.PHONY: ci_setup_server
ci_setup_server:
	FIRESTORE_EMULATOR_HOST=localhost:3600 go run app.go &