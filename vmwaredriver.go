package main

/*
 * Executable entry point used by docker-machine and minikube
 *
 * Created by Chris Baumbauer <cab@cabnetworks.net> 2016
 * Licensed under the Apache v2 License.
 */

import (
	"github.com/cab105/machine-vmware/vmwaredriver"
	"github.com/docker/machine/libmachine/drivers/plugin"
)

func main() {
	plugin.RegisterDriver(vmwaredriver.NewDriver("", ""))
}
