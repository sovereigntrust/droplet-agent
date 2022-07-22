// SPDX-License-Identifier: Apache-2.0

package config

import (
	"flag"
	"strings"
)

type cliArgs struct {
	debugMode    bool
	useSyslog    bool
	sshdPort     int
	sshdCfgFile  string
	hostKeyFiles []string
}

func parseCLIArgs() *cliArgs {
	ret := &cliArgs{}
	hostKeyFilesStr := ""
	flag.BoolVar(&ret.debugMode, "debug", false, "Turn on debug mode")
	flag.BoolVar(&ret.useSyslog, "syslog", false, "Use syslog service for logging")
	flag.IntVar(&ret.sshdPort, "sshd_port", 0, "The port sshd is binding to")
	flag.StringVar(&ret.sshdCfgFile, "sshd_config", "", "The location of sshd_config")
	flag.StringVar(&hostKeyFilesStr, "host_key_files", "", "Paths to Host Key file, comma separated")

	flag.Parse()
	ret.hostKeyFiles = strings.Split(hostKeyFilesStr, ",")
	for k, v := range ret.hostKeyFiles {
		ret.hostKeyFiles[k] = strings.Trim(v, " \t")
	}

	return ret
}
