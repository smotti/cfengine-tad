# Description

An IRC bot that provides information about an autonomous CFEngine node.

Features:

* Reads reports created by the vacana lib report bundles
    * Host info report
    * Classes
    * Promises
* Search classes and promises
* Watch report file for changes and get notified about it
* Notify about repaired and failed promises at a given interval

# Installation

## Binary

TODO

## Source

```
git clone https://github.com/smotti/cfengine-tad.git
cd cfengine-tad
make
```

The above will compile the bot and build a binary named **tad**, which you
then can drop where you like.

# Usage

```
Usage of ./tad:
  -chan="#tad": Channels to join
  -classes="./data/classes.log": Path to classes report
  -config="": Path to ini config for using in go flags. May be relative to the current executable path.
  -configUpdateInterval=0: Update interval for re-reading config file set via -config flag. Zero disables config file re-reading.
  -dumpflags=false: Dumps values for all flags defined in the app into stdout in ini-compatible syntax and terminates the app.
  -hostInfo="./data/va_host_info_report.json": Path to host info report
  -listenChannel=false: Listen for command on public channels
  -name="tad": Nick to use in IRC
  -notifyFailed=true: Notify about failed promises
  -notifyInterval=5m0s: Interval on when to notify about repaired and/or failed promises
  -notifyRepaired=true: Notify about repaired promises
  -promises="./data/promises_outcome.log": Path to promises report
  -server="127.0.0.1:6668": Host:Port to connect to
  -ssl=false: Use SSL/TLS
  -watchInterval=1s: Interval when to check files in watch dir (in seconds)
```

All flags can also be specified in a .ini file and provided to the bot via the
-config flag.

# Commands

* !os
  * Display the following data about the hosts OS: arch, flavor, os, release,
    uptime, version
