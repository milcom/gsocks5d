package main

import (
	"fmt"
	"os"

	"github.com/armon/go-socks5"
	"github.com/jessevdk/go-flags"
)

const version = "0.0.1"

type Options struct {
	BindAddr string `short:"b" long:"bindaddr" env:"ADDR" default:"" description:"Adress to listen on."`
	Port     int    `short:"p" long:"port" env:"PORT" default:"8.8.8.8" description:"Port to listen on"`
	Version  bool   `short:"v" long:"version" default:"false" description:"Show version and exit"`
}

func main() {
	opts := Options{}
	if _, err := flags.Parse(&opts); err != nil {
		if err.(*flags.Error).Type == flags.ErrHelp {
			os.Exit(0)
		} else {
			os.Exit(1)
		}
	}

	if opts.Version {
		fmt.Println(version)
		os.Exit(0)
	}

	// see https://godoc.org/github.com/armon/go-socks5
	conf := &socks5.Config{}
	server, err := socks5.New(conf)
	if err != nil {
		panic(err)
	}

	if err := server.ListenAndServe("tcp", fmt.Sprintf("%s:%d", opts.BindAddr, opts.Port)); err != nil {
		panic(err)
	}
}
