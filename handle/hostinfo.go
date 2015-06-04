package handle

import (
	"github.com/sorcix/irc"

	"github.com/smotti/ircx"
	"github.com/smotti/tad/report"
)

// CmdHostOs handles the CMD_HOST_OS bot command, by sending back data about
// the hosts operating system gathered by CFEngine.
func CmdHostOs(s ircx.Sender, m *irc.Message) {
	r := report.Reports["hostInfo"].(*report.HostInfo)
	msg := r.Os.ToString()

	s.Send(&irc.Message{
		Command:  irc.PRIVMSG,
		Params:   Params(m),
		Trailing: msg,
	})
}

// CmdHostId handles the CMD_HOST_ID bot command.
func CmdHostId(s ircx.Sender, m *irc.Message) {
	r := report.Reports["hostInfo"].(*report.HostInfo)
	msg := r.Identity.ToString()

	s.Send(&irc.Message{
		Command:  irc.PRIVMSG,
		Params:   Params(m),
		Trailing: msg,
	})
}

// CmdHostCfe handles the CMD_HOST_CFE bot command.
func CmdHostCfe(s ircx.Sender, m *irc.Message) {
	r := report.Reports["hostInfo"].(*report.HostInfo)
	msg := r.Cfengine.ToString()

	s.Send(&irc.Message{
		Command:  irc.PRIVMSG,
		Params:   Params(m),
		Trailing: msg,
	})
}
