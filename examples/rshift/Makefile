all: simp comp

simp:
	@echo "simple"
	@echo "# 'go run' for release build" 
	go run ./simple
	@echo "# 'go run -tags debug' for testing build, causing an panic" 
	- go run -tags debug ./simple

comp:
	@echo "composite"
	@echo "# 'go run' for release build" 
	go run ./composite
	@echo "# 'go run -tags debug' for testing build, causing an panic" 
	- go run -tags debug ./composite
