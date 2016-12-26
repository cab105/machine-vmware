This is a fork of the vmware-fusion driver provided to the docker-machine
community in order to support the Linux based VMWare Workstation installation
in addition to VMWare Fusion.

To build:
    mkdir -p $SRC/machine-vmware/src/github.com/cab105
    export GOPATH=$SRC/machine-vmware
    cd $SRC/machine-vmware-unix/src/github.com/cab105
    git clone https://github.com/cab105/machine-vmware.git
    cd machine-vmware
    go get
    go build
    mv machine-vmware $DOCKER_MACHINE_PATH/docker-machine-driver-vmwaredriver

To use it with docker-machine:
    docker-machine cmd -d vmwaredriver ...
