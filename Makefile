build:
	go build -o bin/banana
	GOOS=windows go build -o bin/banana.exe
