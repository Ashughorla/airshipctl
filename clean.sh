#!/bin/sh
echo "Hello,This is clean up script"
echo "#############################Get VM's############################"
sudo virsh list
echo "#########################Destroy air-ephemeral###################"
sudo virsh destroy air-ephemeral
sudo virsh undefine air-ephemeral
echo "#########################Destroy air-target-1####################"
sudo virsh destroy air-target-1
sudo virsh undefine air-target-1
sudo virsh list
echo "###################Remove config file################s###########"
sudo rm -rf $HOME/.airship/
export GOROOT=/usr/local/go/
export GOPATH=/home/node-53/go/
export PATH=$PATH:/usr/local/go/bin
sudo rm -rf /etc/libvirt/qemu/*
sudo virsh net-list
sudo virsh net-destroy air_nat
sudo virsh net-destroy air_prov
sudo virsh net-list
