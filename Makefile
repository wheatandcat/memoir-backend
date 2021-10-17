.PHONY: moqgen
moqgen: 
	@-rm -f ./repository/moqs/*.go
	go generate ./repository/...