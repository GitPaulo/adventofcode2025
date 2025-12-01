.PHONY: run clean

run:
	@cd day$(DAY) && go run 01.go

clean:
	@rm -rf bin/
