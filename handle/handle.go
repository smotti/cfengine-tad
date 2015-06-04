package handle

import (
	"github.com/smotti/ircx"
	"github.com/sorcix/irc"
)

// Params returns the params for sending an irc message by checking if the
// message received was a private query or a message from a channel.
func Params(m *irc.Message) []string {
	var params []string

	if ircx.IsQuery(m) {
		params = append(params, m.Prefix.Name)
	} else {
		params = m.Params
	}

	return params
}
