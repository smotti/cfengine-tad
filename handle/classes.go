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

// CmdClSearch handle the CMD_CL_SEARCH command and sends classes matching
// the fiven search term line by line.
func CmdClSearch(s ircx.Sender, m *irc.Message) {
	r := report.Reports["classes"].(*report.Context)

	pattern := strings.Split(m.Trailing, " ")[1]

	for _, v := range r.Classes {
		matched, err := regexp.MatchString(pattern, v)
		if err != nil {
			log.Println("Error:", err)
		}
		if matched {
			s.Send(&irc.Message{
				Command:  irc.PRIVMSG,
				Params:   Params(m),
				Trailing: v,
			})

			time.Sleep(600 * time.Millisecond)
		}
	}
}
