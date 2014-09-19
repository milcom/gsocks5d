#!/bin/sh

rm  gsocks5d*.64
GOOS=linux go build -o gsocks5d.linux_64
go build -o gsocks5d.darwin_64

