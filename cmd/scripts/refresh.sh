#!/bin/sh
Lib=/mnt/onboard/kobofileserver
SD=/mnt/sd/kobofileserver

[ ! -e "$Lib" ] && mkdir -p "$Lib" >/dev/null 2>&1
[ ! -e "$SD" ] && mkdir -p "$SD" >/dev/null 2>&1

mountpoint -q "$SD"
mount --bind "$Lib" "$SD"

echo sd add /dev/mmcblk1p1 >> /tmp/nickel-hardware-status