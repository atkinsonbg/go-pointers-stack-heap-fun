bench:
	go test -bench . -benchmem

escape:
	go build -gcflags '-m -l'
