.PHONY: run clean inputs answer install-hooks

run:
	@cd day$(DAY) && go run $(DAY).go

clean:
	@rm -rf bin/

inputs:
	@./scripts/fetch_inputs.sh

answer:
	@./scripts/post_answer.sh $(DAY) $(PART) $(ANSWER)

install-hooks:
	@cp hooks/pre-push .git/hooks/pre-push
	@chmod +x .git/hooks/pre-push
	@echo "✓ Installed git hooks"
