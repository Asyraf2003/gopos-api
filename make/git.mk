##@ Git

.PHONY: git-status push

git-status: ## Show short git status
	git status --short

push: ## Run the scripted add, commit, and push flow using MSG
	bash scripts/git_push.sh
