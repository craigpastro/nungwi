.PHONY: run-development
run-development:
	uvicorn main:app --reload
