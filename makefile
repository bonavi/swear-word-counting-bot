deploy-check-all: go-mod-tidy-all
	cd pkg; make deploy-check
	cd server; make deploy-check

go-mod-tidy-all:
	cd pkg; go mod tidy
	cd server; go mod tidy

update-pkg:
	cd pkg; git pull origin main

update-linter:
	brew upgrade golangci-lint
