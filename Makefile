.PHONY: run clean inputs

run:
	@cd day$(DAY) && go run $(DAY).go

clean:
	@rm -rf bin/

inputs:
	@./fetch_inputs.sh
