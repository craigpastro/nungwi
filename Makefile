.PHONY: test
test:
	python -m pytest
	
.PHONY: run-development
run-development:
	uvicorn app.main:app --reload
