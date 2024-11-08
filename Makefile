dev: wire
	go run main.go -conf=dev.yaml

.PHONY: wire
wire:
	(cd wire && wire)

