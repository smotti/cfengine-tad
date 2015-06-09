package handle

import (
	"log"
	"regexp"
	"strings"
	"time"

	"github.com/smotti/ircx"
	"github.com/smotti/tad/report"
	"github.com/sorcix/irc"
)

// CmdPList handles the command CMD_P_LIST and will send each report entry
// line by line.
// NOTE: This might mean problems if any of the lines is larger than
// 512 bytes because the max size of an irc msg is 512 bytes.
func CmdPList(s ircx.Sender, m *irc.Message) {
	r := report.Reports["promises"].(*report.Promises)

	for _, v := range r.List {
		s.Send(&irc.Message{
			Command:  irc.PRIVMSG,
			Params:   Params(m),
			Trailing: v.ToString(),
		})

		time.Sleep(600 * time.Millisecond)
	}
}

// CmdPSearch handles the comman CMD_P_SEARCH and will sean each match line
// by line.
func CmdPSearch(s ircx.Sender, m *irc.Message) {
	r := report.Reports["promises"].(*report.Promises)

	search := strings.Split(m.Trailing, " ")[1]

	for _, v := range r.List {
		mclass, err := regexp.MatchString(search, v.Class)
		if err != nil {
			log.Println("Error:", err)
			continue
		}

		mhandler, err := regexp.MatchString(search, v.Handler)
		if err != nil {
			log.Println("Error:", err)
			continue
		}

		mpromiser, err := regexp.MatchString(search, v.Promiser)
		if err != nil {
			log.Println("Error:", err)
			continue
		}

		moutcome, err := regexp.MatchString(search, v.Outcome)
		if err != nil {
			log.Println("Error:", err)
			continue
		}

		if mclass || mhandler || mpromiser || moutcome {
			s.Send(&irc.Message{
				Command:  irc.PRIVMSG,
				Params:   Params(m),
				Trailing: v.ToString(),
			})
		}
	}
}
