// +build windows

/*
 * Copyright 2016 Chris Baumbauer <cab@cabnetworks.net>  Licensed under the Apache v2 License.
 */

package vmwaredriver

import "github.com/docker/machine/libmachine/drivers"

func NewDriver(hostName, storePath string) drivers.Driver {
	return drivers.NewDriverNotSupported("vmwaredriver", hostName, storePath)
}
