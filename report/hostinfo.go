package report

import (
	"encoding/json"
	"os"
	"strings"

	"github.com/smotti/tad/config"
)

type (
	// meta defines a reports meta data.
	meta struct {
		Timestamp int `json:"timestamp"`
	}

	// cfengine defines data about a hosts cfengine installation.
	cfengine struct {
		BootstrappedTo    string `json:"bootstrapped_to"`
		LastAgentRun      string `json:"last_agent_run"`
		PolicyLastUpdated string `json:"policy_last_updated"`
		PolicyReleaseId   string `json:"policy_release_id"`
		Version           string `json:"version"`
	}

	// identity defines identifying data about a host.
	identity struct {
		Fqdn string `json:"fqdn"`
		Id   string `json:"id"`
		Uqdn string `json:"uqdn"`
	}

	// networkInterface defines data about a network interface of a host.
	networkInterface struct {
		Flags string `json:"flags"`
		Ipv4  string `json:"ipv4"`
		Mac   string `json:"mac"`
		Name  string `json:"name"`
	}

	// networkPort defines data about a port listening on a host.
	networkPort struct {
		Inet     string `json:"inet"`
		Port     string `json:"port"`
		Protocol string `json:"type"`
	}

	// _os defines data about a hosts operating system.
	_os struct {
		Arch    string `json:"arch"`
		Flavor  string `json:"flavor"`
		Os      string `json:"os"`
		Release string `json:"release"`
		Uptime  string `json:"uptime"`
		Version string `json:"version"`
	}

	// software defines data about installed software on a host.
	software struct {
		Arch    string `json:"arch"`
		Method  string `json:"method"`
		Name    string `json:"name"`
		Version string `json:"version"`
	}

	// HostInfo defines a set of data (os, software, ...) about a host.
	HostInfo struct {
		Filename   string
		Meta       meta               `json:"_meta"`
		Cfengine   cfengine           `json:"cfengine"`
		Identity   identity           `json:"identity"`
		Interfaces []networkInterface `json:"network_interfaces"`
		Ports      []networkPort      `json:"network_ports"`
		Os         _os                `json:"os"`
		Software   []software         `json:"software"`
	}
)

// init registers the report at application startup.
func init() {
	var r Report
	r = &HostInfo{
		Filename: config.HostInfoReport,
	}

	Register("hostInfo", r)
}

// Read implements the Report interface for HostInfo.
func (h *HostInfo) Read() error {
	file, err := os.Open(h.Filename)
	if err != nil {
		return err
	}
	defer file.Close()

	err = json.NewDecoder(file).Decode(h)

	return err
}

// ToString for type _os.
func (d *_os) ToString() string {
	fields := []string{
		"arch: " + d.Arch,
		"flavor: " + d.Flavor,
		"os: " + d.Os,
		"release: " + d.Release,
		"uptime: " + d.Uptime,
		"version: " + d.Version,
	}
	return strings.Join(fields, ", ")
}

// ToString for type identity.
func (d *identity) ToString() string {
	fields := []string{
		"fqdn: " + d.Fqdn,
		"id: " + d.Id,
		"uqdn: " + d.Uqdn,
	}
	return strings.Join(fields, ", ")
}

// ToString for type cfengine.
func (d *cfengine) ToString() string {
	fields := []string{
		"bootstrappedTo: " + d.BootstrappedTo,
		"lastAgentRun: " + d.LastAgentRun,
		"policyLastUpdated: " + d.PolicyLastUpdated,
		"policyReleaseId: " + d.PolicyReleaseId,
		"version: " + d.Version,
	}
	return strings.Join(fields, ", ")
}

// ToString for type software.
func (d *software) ToString() string {
	fields := []string{
		"name: " + d.Name,
		"arch: " + d.Arch,
		"method: " + d.Method,
		"version: " + d.Version,
	}
	return strings.Join(fields, ", ")
}

// ToString for type networkInterface.
func (d *networkInterface) ToString() string {
	fields := []string{
		"name: " + d.Name,
		"mac: " + d.Mac,
		"ipv4: " + d.Ipv4,
		"flags: " + d.Flags,
	}
	return strings.Join(fields, ", ")
}

// ToString for type networkPort.
func (d *networkPort) ToString() string {
	fields := []string{
		"port: " + d.Port,
		"protocol: " + d.Protocol,
		"inet: " + d.Inet,
	}
	return strings.Join(fields, ", ")
}
