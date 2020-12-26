TARGET=energomera-cli

all: energomera-cli.go
	go build -o $(TARGET)

clean:
	go clean
	rm -f $(TARGET)

