.PHONY: run clean inputs answer install-hooks

run:
	@cd day$(DAY) && go run $(DAY).go

clean:
	@rm -rf bin/

inputs:
	@./scripts/fetch_inputs.sh

answer:
	@DAY=$(DAY) PART=$(PART) ANSWER=$(ANSWER) ./scripts/post_answer.sh

install-hooks:
	@cp hooks/pre-push .git/hooks/pre-push
	@chmod +x .git/hooks/pre-push
	@echo "✓ Installed git hooks"
