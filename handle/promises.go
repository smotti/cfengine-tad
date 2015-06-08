package handle

import (
	"time"

	"github.com/smotti/ircx"
	"github.com/smotti/tad/report"
	"github.com/sorcix/irc"
)

// CmdPList handles the command CMD_P_LIST and will send each promise's fields
// line by line, due to ircs message limitation of 512 bytes.
func CmdPList(s ircx.Sender, m *irc.Message) {
	r := report.Reports["promises"].(*report.Promises)

	for _, v := range r.List {
		s.Send(&irc.Message{
			Command:  irc.PRIVMSG,
			Params:   Params(m),
			Trailing: "class: " + v.Class + ", handler: " + v.Handler + ", promiser: " + v.Promiser + ", outcome: " + v.Outcome,
		})

		time.Sleep(600 * time.Millisecond)
	}
}
