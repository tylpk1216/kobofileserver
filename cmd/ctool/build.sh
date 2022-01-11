#!/bin/bash

# set variables
TOOL_CHAIN_PATH=/home/tylpk/x-tools/arm-kobo-linux-gnueabihf/bin
KOXTOOLCHAIN_PATH=/home/tylpk/koxtoolchain

# set up env
export PATH=$TOOL_CHAIN_PATH:$PATH
source $KOXTOOLCHAIN_PATH/refs/x-compile.sh kobo env

# build
arm-kobo-linux-gnueabihf-gcc -o refresh -ldl refresh.c

