package main

import (
    // stdlib
	"flag"
	"log"

    // third party
	"github.com/nickvanw/ircx"
	"github.com/sorcix/irc"

    // own
    "handle"
)

var (
	name     = flag.String("name", "tad", "Nick to use in IRC")
	server   = flag.String("server", "irc.internetz.me:6667", "Host:Port to connect to")
	channels = flag.String("chan", "#tad", "Channels to join")
)

func init() {
	flag.Parse()
}

func main() {
	bot := ircx.Classic(*server, *name)
	if err := bot.Connect(); err != nil {
		log.Panicln("Unable to dial IRC server ", err)
	}

	RegisterHandlers(bot)
	bot.CallbackLoop()
	log.Println("Exiting..")
}

func RegisterHandlers(bot *ircx.Bot) {
	bot.AddCallback(irc.RPL_WELCOME, ircx.Callback{Handler: ircx.HandlerFunc(RegisterConnect)})
	bot.AddCallback(irc.PING, ircx.Callback{Handler: ircx.HandlerFunc(PingHandler)})
}

func RegisterConnect(s ircx.Sender, m *irc.Message) {
	s.Send(&irc.Message{
		Command: irc.JOIN,
		Params:  []string{*channels},
	})
}

func PingHandler(s ircx.Sender, m *irc.Message) {
	s.Send(&irc.Message{
		Command:  irc.PONG,
		Params:   m.Params,
		Trailing: m.Trailing,
	})
}
