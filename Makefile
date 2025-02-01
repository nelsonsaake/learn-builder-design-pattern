.PHONY: run commit
.SILENT:
.ONESHELL:

run:
	go run .

commit:
	git add .
	git commit -m "chore: commit everything"
	git push origin main