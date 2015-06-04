package handle

import (
    "github.com/sorcix/irc"
    "github.com/smotti/ircx"
    "github.com/smotti/tad/report"
    "github.com/smotti/tad/config"
)

// CmdHostOs handles the CMD_HOST_OS bot command, by sending back data about
// the hosts operating system gathered by CFEngine.
func CmdHostOs(s ircx.Sender, m *irc.Message) {
    report := report.HostInfo{
        Filename: config.HostInfoReport,
    }
    var msg string
    if err := report.Read(); err != nil {
        msg = "Failed to read report file"
    } else {
        msg = report.Os.ToString()
    }

    var params []string
    if ircx.IsQuery(m) {
        params = append(params, m.Prefix.Name)
    } else {
        params = m.Params
    }

    s.Send(&irc.Message{
        Command: irc.PRIVMSG,
        Params: params,
        Trailing: msg,
    })
}

// CmdHostId handle the CMD_HOST_ID bot command.
func CmdHostId(s ircx.Sender, m *irc.Message) {
}
