.PHONY: install
install:
	pip install -r requirements.txt

.PHONY: test
test:
	python3 -m pytest
	
.PHONY: run-development
run-development:
	uvicorn app.main:app --reload

.PHONY: run
run:
	uvicorn app.main:app
