# graphite-memcached

## What is this?

Memcached statistics to graphite gateway.

## Building it

1. Install [go](http://golang.org/doc/install)

2. Install "graphite-golang" go get -u github.com/marpaia/graphite-golang

2. Install "go-logging" go get -u github.com/op/go-logging

4. Compile graphite-memcached

        git clone git://github.com/akashihi/graphite-memcached.git
        cd graphite-memcached
        go build .

## Running it

Generally:

    graphite-memcached -host 127.0.0.1 -port 11211 -metrics-host 192.168.1.1 -metrics-port 2003 -metrics-prefix test -period 60

All parameters could be omited. Run with --help to het parameters description

## License 

See LICENSE file.

Copyright 2015 Denis V Chapligin <akashihi@gmail.com>
