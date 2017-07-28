# SYS-STATS [![Build Status](https://travis-ci.org/bicofino/sys-status.svg?branch=master)](https://travis-ci.org/bicofino/sys-stats)
Sys-stats is a system that provides machine information via HTTP.

Below what information is available:

* CPU usage
* Disk usage
* Host Information
* Memory usage
* Network usage

## Installation

If you have Go tools installed to your system, enter the command below into your terminal:
(I tested against OSX and Linux)
```bash
$ go get github.com/bicofino/sys-stats
```

## Build from clone

```bash
go get -d
go build
./sys-stats
```

## Usage

You can access the information via your web browser or other HTTP clients.

List of endpoints:

1. `http://localhost:8888/memory`
1. `http://localhost:8888/cpu`
1. `http://localhost:8888/disks`
1. `http://localhost:8888/hostinfo`
1. `http://localhost:8888/load`
1. `http://localhost:8888/network`

## Examples

### 1. CPU

Provides CPU usage information

*Request:*

`GET http://localhost:8888/cpu`

*Response:*

```json
{
  "idle": 91.29885282332786,
  "iowait": 0,
  "irq": 0,
  "nice": 0,
  "stolen": 0,
  "system": 2.9273780185822904,
  "used": 8.701147176672137,
  "user": 5.773769158089837
}
```

### 2. Disks

Provides disk usage information

*Request:*

`GET http://localhost:8888/disks
