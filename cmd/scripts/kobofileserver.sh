#!/bin/sh

WORK_DIR="${0%/*}"

FBINK=/usr/bin/fbink

$WORK_DIR/kobofileserver& >>$WORK_DIR/log.txt 2>&1

sleep 2s

[ -e "$FBINK" ] && $FBINK -g file=$WORK_DIR/qrcode.png,halign=CENTER,valign=CENTER -W A2
