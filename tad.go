package main

import (
	// stdlib
	"crypto/tls"
	"log"

	// third party
	"github.com/sorcix/irc"

	// own
	"github.com/smotti/ircx"
	"github.com/smotti/tad/config"
	"github.com/smotti/tad/handle"
	"github.com/smotti/tad/report"
)

var (
	cmd = map[string]string{
		"CMD_OS":        "^!os$",
		"CMD_CFE":       "^!cfe$",
		"CMD_ID":        "^!id$",
		"CMD_NET_IF":    "^!net if$",
		"CMD_NET_PORTS": "^!net ports$",
		"CMD_SW":        "^!sw$",
		"CMD_SW_SEARCH": "^!sw ([a-z0-9\\-_]+)$",
		"CMD_CL_LIST":   "^!cl$",
		"CMD_CL_SEARCH": "^!cl ([a-z0-9_]+)$",
		"CMD_P_LIST":    "^!p$",
		"CMD_P_SEARCH":  "^!p ([a-z0-9\\-_/\\*\\s&\\|]+)$",
	}

	bot *ircx.Bot
)

func main() {
	// Establish the connection to the given irc server.
	if *config.Ssl {
		tlsConfig := &tls.Config{InsecureSkipVerify: true}
		bot = ircx.WithTLS(*config.Server, *config.Name, tlsConfig)
	} else {
		bot = ircx.Classic(*config.Server, *config.Name)
	}
	if err := bot.Connect(); err != nil {
		log.Panicln("Unable to dial IRC server ", err)
	}

	// Assign the boot commands.
	bot.Commands = cmd

	// Set the listenChannel option based on the value of the flag.
	bot.Options = map[string]bool{
		"rejoin":        bot.Options["rejoin"],
		"connectd":      bot.Options["connected"],
		"listenChannel": *config.Listen,
	}

	// Register command handlers.
	RegisterHandlers(bot)

	// Watch report for changes.
	for _, v := range report.Reports {
		go v.Watch(bot.Data)
	}

	// Start callback loop.
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
	bot.AddCallback("CMD_SW_SEARCH", ircx.Callback{
		Handler: ircx.HandlerFunc(handle.CmdSwSearch),
	})
	bot.AddCallback("CMD_CL_LIST", ircx.Callback{
		Handler: ircx.HandlerFunc(handle.CmdClList),
	})
	bot.AddCallback("CMD_CL_SEARCH", ircx.Callback{
		Handler: ircx.HandlerFunc(handle.CmdClSearch),
	})
	bot.AddCallback("CMD_P_LIST", ircx.Callback{
		Handler: ircx.HandlerFunc(handle.CmdPList),
	})
	bot.AddCallback("CMD_P_SEARCH", ircx.Callback{
		Handler: ircx.HandlerFunc(handle.CmdPSearch),
	})
}

// RegisterConnect takes care of joining provided channels.
func RegisterConnect(s ircx.Sender, m *irc.Message) {
	s.Send(&irc.Message{
		Command: irc.JOIN,
		Params:  []string{*config.Channels},
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
	if m.Prefix == nil {
		s.Send(&irc.Message{
			Command:  m.Command,
			Params:   m.Params,
			Trailing: m.Trailing,
		})
	}
}
