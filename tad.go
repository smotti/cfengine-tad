package main

import (
	// stdlib
	"crypto/tls"
	"flag"
	"log"

	// third party
	"github.com/sorcix/irc"

	// own
	"github.com/smotti/ircx"
	"github.com/smotti/tad/handle"
)

var (
	name     = flag.String("name", "tad", "Nick to use in IRC")
	server   = flag.String("server", "irc.internetz.me:6697", "Host:Port to connect to")
	channels = flag.String("chan", "#tad", "Channels to join")
	ssl      = flag.Bool("ssl", true, "Use SSL/TLS")

	cmd = map[string]string{
		"CMD_OS":        "^!os$",
		"CMD_CFE":       "^!cfe$",
		"CMD_ID":        "^!id$",
		"CMD_NET_IF":    "^!net if$",
		"CMD_NET_PORTS": "^!net ports$",
		"CMD_SW":        "^!sw$",
	}

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

	bot.Commands = cmd
	RegisterHandlers(bot)
	bot.CallbackLoop()
	log.Println("Exiting..")
}

// RegisterHandlers registers the bots handler functions.
func RegisterHandlers(bot *ircx.Bot) {
	bot.AddCallback(irc.RPL_WELCOME, ircx.Callback{
		Handler: ircx.HandlerFunc(RegisterConnect),
	})
	bot.AddCallback(irc.PING, ircx.Callback{
		Handler: ircx.HandlerFunc(PingHandler),
	})
	bot.AddCallback(irc.PRIVMSG, ircx.Callback{
		Handler: ircx.HandlerFunc(PrivMsgHandler),
	})
	bot.AddCallback("CMD_OS", ircx.Callback{
		Handler: ircx.HandlerFunc(handle.CmdOs),
	})
	bot.AddCallback("CMD_CFE", ircx.Callback{
		Handler: ircx.HandlerFunc(handle.CmdCfe),
	})
	bot.AddCallback("CMD_ID", ircx.Callback{
		Handler: ircx.HandlerFunc(handle.CmdId),
	})
	bot.AddCallback("CMD_SW", ircx.Callback{
		Handler: ircx.HandlerFunc(handle.CmdSw),
	})
	bot.AddCallback("CMD_NET_IF", ircx.Callback{
		Handler: ircx.HandlerFunc(handle.CmdNetIf),
	})
	bot.AddCallback("CMD_NET_PORTS", ircx.Callback{
		Handler: ircx.HandlerFunc(handle.CmdNetPorts),
	})
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

// PrivMsgHandler logs incoming PRIVMSGs.
func PrivMsgHandler(s ircx.Sender, m *irc.Message) {
	log.Println(m.String())
}
