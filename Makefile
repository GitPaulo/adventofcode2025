.PHONY: run clean inputs

run:
	@cd day$(DAY) && go run 01.go

clean:
	@rm -rf bin/

inputs:
	@./fetch_inputs.sh
