package handle

import (
	"time"

	"github.com/sorcix/irc"

	"github.com/smotti/ircx"
	"github.com/smotti/tad/report"
)

// CmdOs handles the CMD_OS bot command, by sending back data about
// the hosts operating system gathered by CFEngine.
func CmdOs(s ircx.Sender, m *irc.Message) {
	r := report.Reports["hostInfo"].(*report.HostInfo)
	msg := r.Os.ToString()

	s.Send(&irc.Message{
		Command:  irc.PRIVMSG,
		Params:   Params(m),
		Trailing: msg,
	})
}

// CmdId handles the CMD_ID bot command.
func CmdId(s ircx.Sender, m *irc.Message) {
	r := report.Reports["hostInfo"].(*report.HostInfo)
	msg := r.Identity.ToString()

	s.Send(&irc.Message{
		Command:  irc.PRIVMSG,
		Params:   Params(m),
		Trailing: msg,
	})
}

// CmdCfe handles the CMD_CFE bot command.
func CmdCfe(s ircx.Sender, m *irc.Message) {
	r := report.Reports["hostInfo"].(*report.HostInfo)
	msg := r.Cfengine.ToString()

	s.Send(&irc.Message{
		Command:  irc.PRIVMSG,
		Params:   Params(m),
		Trailing: msg,
	})
}

// CmdSw handles the CMD_SW bot command.
func CmdSw(s ircx.Sender, m *irc.Message) {
	r := report.Reports["hostInfo"].(*report.HostInfo)

	for _, v := range r.Software {
		msg := v.ToString()

		s.Send(&irc.Message{
			Command:  irc.PRIVMSG,
			Params:   Params(m),
			Trailing: msg,
		})

		// Need to wait before sending the next msg, or else we will get
		// blocked by the IRC server.
		time.Sleep(100 * time.Millisecond)
	}
}

// CmdNetIf handles the CMD_NET_IF bot command.
func CmdNetIf(s ircx.Sender, m *irc.Message) {
	r := report.Reports["hostInfo"].(*report.HostInfo)

	for _, v := range r.Interfaces {
		msg := v.ToString()

		s.Send(&irc.Message{
			Command:  irc.PRIVMSG,
			Params:   Params(m),
			Trailing: msg,
		})

		// Need to wait before sending the next msg, or else we will get
		// blocked by the IRC server.
		time.Sleep(500 * time.Millisecond)
	}
}

// CmdNetPorts handle the CMD_NET_PORTS bot command.
func CmdNetPorts(s ircx.Sender, m *irc.Message) {
	r := report.Reports["hostInfo"].(*report.HostInfo)

	for _, v := range r.Ports {
		msg := v.ToString()

		s.Send(&irc.Message{
			Command:  irc.PRIVMSG,
			Params:   Params(m),
			Trailing: msg,
		})

		// Need to wait before sending the next msg, or else we will get
		// blocked by the IRC server.
		time.Sleep(500 * time.Millisecond)
	}
}
