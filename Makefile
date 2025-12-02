.PHONY: run clean inputs install-hooks

run:
	@cd day$(DAY) && go run $(DAY).go

clean:
	@rm -rf bin/

inputs:
	@./fetch_inputs.sh

install-hooks:
	@cp hooks/pre-push .git/hooks/pre-push
	@chmod +x .git/hooks/pre-push
	@echo "✓ Installed git hooks"
