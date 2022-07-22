package sysaccess

type sshMgrOpts struct {
	customSSHDPort     int
	customSSHDCfgFile  string
	customHostKeyFiles []string
	manageDropletKeys  bool
}

// SSHManagerOpt allows creating the SSHManager instance with designated options
type SSHManagerOpt func(opt *sshMgrOpts)

// WithCustomSSHDPort indicates the SSHD is running on a custom port which is specified via command line argument
func WithCustomSSHDPort(port int) SSHManagerOpt {
	return func(opt *sshMgrOpts) {
		opt.customSSHDPort = port
	}
}

// WithCustomSSHDCfg specifies the path the custom sshd_config file that the sshd instance uses
func WithCustomSSHDCfg(cfgFile string) SSHManagerOpt {
	return func(opt *sshMgrOpts) {
		opt.customSSHDCfgFile = cfgFile
	}
}

// WithCustomHostKeyFiles provides a way to specify the locations of host key files
// W/O specifying the host key files, `HostKey` entries presented in the sshd_config will be used
// if no `HostKey` entry can be found in sshd_config, files with `ssh_host_*_key.pub` under the ssh config
// directory (default to `/etc/ssh` for linux) will be loaded.
func WithCustomHostKeyFiles(hostKeyFiles []string) SSHManagerOpt {
	return func(opt *sshMgrOpts) {
		opt.customHostKeyFiles = append(hostKeyFiles[:0:0], hostKeyFiles...)
	}
}

// WithoutManagingDropletKeys tells the agent to not attempt to manage the ssh keys
func WithoutManagingDropletKeys() SSHManagerOpt {
	return func(opt *sshMgrOpts) {
		opt.manageDropletKeys = false
	}
}

func defaultMgrOpts() *sshMgrOpts {
	return &sshMgrOpts{
		customSSHDPort:    0,
		customSSHDCfgFile: "",
		manageDropletKeys: true,
	}
}
