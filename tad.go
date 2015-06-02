package main

import (
	// stdlib
	"crypto/tls"
	"flag"
	"log"

	// third party
	"github.com/nickvanw/ircx"
	"github.com/sorcix/irc"

	// own
	"github.com/smotti/tad/report"
)

var (
	name     = flag.String("name", "tad", "Nick to use in IRC")
	server   = flag.String("server", "irc.internetz.me:6697", "Host:Port to connect to")
	channels = flag.String("chan", "#tad", "Channels to join")
	ssl      = flag.Bool("ssl", true, "Use SSL/TLS")

	bot *ircx.Bot
)

func init() {
	flag.Parse()
}

func main() {
	if *ssl {
		tlsConfig := &tls.Config{InsecureSkipVerify: true}
		bot = ircx.WithTLS(*server, *name, tlsConfig)
	} else {
		bot = ircx.Classic(*server, *name)
	}
	if err := bot.Connect(); err != nil {
		log.Panicln("Unable to dial IRC server ", err)
	}

	RegisterHandlers(bot)
	bot.CallbackLoop()
	log.Println("Exiting..")
}

// RegisterHandlers registers the bots handler functions.
func RegisterHandlers(bot *ircx.Bot) {
	bot.AddCallback(irc.RPL_WELCOME, ircx.Callback{Handler: ircx.HandlerFunc(RegisterConnect)})
	bot.AddCallback(irc.PING, ircx.Callback{Handler: ircx.HandlerFunc(PingHandler)})
}

// RegisterConnect takes care of joining provided channels.
func RegisterConnect(s ircx.Sender, m *irc.Message) {
	s.Send(&irc.Message{
		Command: irc.JOIN,
		Params:  []string{*channels},
	})
}

// PingHandler responds to PING commands from the server. On some servers the
// connection might be killed if the client doesn't respond to them.
func PingHandler(s ircx.Sender, m *irc.Message) {
	s.Send(&irc.Message{
		Command:  irc.PONG,
		Params:   m.Params,
		Trailing: m.Trailing,
	})
}
