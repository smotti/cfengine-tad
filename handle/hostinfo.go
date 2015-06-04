package handle

import (
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
