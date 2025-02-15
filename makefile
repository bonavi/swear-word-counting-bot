deploy-check-all: go-mod-tidy-all
	cd pkg; make deploy-check
	cd swear; make deploy-check

go-mod-tidy-all:
	cd pkg; go mod tidy
	cd swear; go mod tidy

update-pkg:
	cd pkg; git pull origin main

update-linter:
	brew upgrade golangci-lint
