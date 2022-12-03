.PHONY: run-development
run-development:
	uvicorn app.main:app --reload
