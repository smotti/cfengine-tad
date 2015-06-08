package handle

import (
	"time"

	"github.com/smotti/ircx"
	"github.com/smotti/tad/report"
	"github.com/sorcix/irc"
)

// CmdClList handles the CMD_CL_LIST command and sends a list of all classes
// found in the config.Classes file.
func CmdClList(s ircx.Sender, m *irc.Message) {
	r := report.Reports["classes"].(*report.Context)

	for _, v := range r.Classes {
		s.Send(&irc.Message{
			Command:  irc.PRIVMSG,
			Params:   Params(m),
			Trailing: v,
		})
	}

	time.Sleep(600 * time.Millisecond)
}
