# A default Makefile for Create Go App project.
# Author: Vic Shóstak <vic@shostak.dev> (https://shostak.dev)
# For more information, please visit https://create-go.app/

.PHONY: test run build

FRONTEND_PATH = $(PWD)/frontend

test:
	@if [ -d "$(FRONTEND_PATH)" ]; then cd $(FRONTEND_PATH) && npm run test; fi

run: test
	@if [ -d "$(FRONTEND_PATH)" ]; then cd $(FRONTEND_PATH) && npm run dev; fi

build: test
	@if [ -d "$(FRONTEND_PATH)" ]; then cd $(FRONTEND_PATH) && npm run build; fi
