all:
	go get github.com/sorcix/irc
	go get github.com/smotti/ircx
	go build -o tad tad.go
clean:
	rm -f tad
