[![Build Status](https://travis-ci.com/sensu-skunkworks/system-profile-linux.svg?branch=master)](https://travis-ci.com/sensu-skunkworks/system-profile-linux)

# System Profile Linux

A Sensu plugin for collecting system resource metrics from the procfile system,
with the aim to provide a Sensu-native alternative to Collectd's built-in
system resource telemetry.

- [Overview](#overview)
- [Usage Example(s)](#usage-examples)
- [Configuration](#configuration)
  - [Asset Manifest](#asset-manifest)
  - [Check Manifest(s)](#check-manifests)
- [Metrics](#metrics)
  - [`/proc/stat` metrics](#proc-stat-metrics)
  - [`/proc/loadavg` metrics](#proc-loadavg-metrics)
  - [`/proc/net/dev` metrics](#proc-net-dev-metrics)
  - [`/proc/meminfo` metrics](#proc-meminfo-metrics)
- [Service Checks](#service-checks)

## Overview

This plugin collects system metrics from Linux and various other unix-like
systems. The plugin reads and parses the following procfiles, and generates
output in [Graphite Plaintext Protocol][graphite-plaintext] format:

- `/proc/stat`
- `/proc/loadavg`
- `/proc/net/dev`
- `/proc/meminfo`

[graphite-plaintext]: https://graphite.readthedocs.io/en/latest/feeding-carbon.html#the-plaintext-protocol

## Usage Example(s)

The `system-profile-linux` check plugin can be run as a command with no options:

```shell
$ ./system-profile-linux
cpu.user 24642 1553373819
cpu.nice 9 1553373819
cpu.system 11256 1553373819
cpu.idle 14813304 1553373819
cpu.iowait 2333 1553373819
cpu.irq 0 1553373819
cpu.softirq 263 1553373819
cpu.steal 8174 1553373819
cpu0.user 6292 1553373819
cpu0.nice 2 1553373819
cpu0.system 2736 1553373819
cpu0.idle 405122 1553373819
cpu0.iowait 541 1553373819
cpu0.irq 0 1553373819
cpu0.softirq 40 1553373819
cpu0.steal 2011 1553373819
cpu1.user 6250 1553373819
cpu1.nice 2 1553373819
cpu1.system 2736 1553373819
cpu1.idle 4803200 1553373819
cpu1.iowait 492 1553373819
cpu1.irq 0 1553373819
cpu1.softirq 35 1553373819
cpu1.steal 1996 1553373819
cpu2.user 5585 1553373819
cpu2.nice 1 1553373819
cpu2.system 2981 1553373819
cpu2.idle 4802700 1553373819
cpu2.iowait 880 1553373819
cpu2.irq 0 1553373819
cpu2.softirq 40 1553373819
cpu2.steal 2160 1553373819
cpu3.user 6513 1553373819
cpu3.nice 2 1553373819
cpu3.system 2801 1553373819
cpu3.idle 4802281 1553373819
cpu3.iowait 418 1553373819
cpu3.irq 0 1553373819
cpu3.softirq 148 1553373819
cpu3.steal 2005 1553373819
intr 0 1553373819
ctxt 2654426 1553373819
btime 1553325668 1553373819
processes 4651 1553373819
procs_running 1 1553373819
procs_blocked 0 1553373819
softirq 377285 1553373819
load_avg.1_min 0.27 1553373819
load_avg.5_min 0.26 1553373819
load_avg.15_min 0.17 1553373819
memory.memtotal 11149428 1553373819
memory.memfree 9236996 1553373819
memory.memavailable 9986604 1553373819
memory.buffers 568 1553373819
memory.cached 1230644 1553373819
swap.swapcached 0 1553373819
memory.active 1107960 1553373819
swap.swaptotal 0 1553373819
swap.swapfree 0 1553373819
memory.committed_as 3278596 1553373819
net.face.rxbytes |bytes 1553373819
net.face.rxpackets packets 1553373819
net.face.rxerrors errs 1553373819
net.face.rxdrops drop 1553373819
net.face.rxfifo fifo 1553373819
net.face.rxframe frame 1553373819
net.face.rxcompressed compressed 1553373819
net.face.rxmulticast multicast|bytes 1553373819
net.face.txbytes packets 1553373819
net.face.txpackets errs 1553373819
net.face.txerrors drop 1553373819
net.face.txdrops fifo 1553373819
net.face.txfifo colls 1553373819
net.face.txcolls carrier 1553373819
net.face.txcarrier compressed 1553373819
net.eth0.rxbytes 3004129 1553373819
net.eth0.rxpackets 25490 1553373819
net.eth0.rxerrors 0 1553373819
net.eth0.rxdrops 0 1553373819
net.eth0.rxfifo 0 1553373819
net.eth0.rxframe 0 1553373819
net.eth0.rxcompressed 0 1553373819
net.eth0.rxmulticast 0 1553373819
net.eth0.txbytes 48605827 1553373819
net.eth0.txpackets 22666 1553373819
net.eth0.txerrors 0 1553373819
net.eth0.txdrops 0 1553373819
net.eth0.txfifo 0 1553373819
net.eth0.txcolls 0 1553373819
net.eth0.txcarrier 0 1553373819
net.eth0.txcompressed 0 1553373819
net.lo.rxbytes 1227 1553373819
net.lo.rxpackets 15 1553373819
net.lo.rxerrors 0 1553373819
net.lo.rxdrops 0 1553373819
net.lo.rxfifo 0 1553373819
net.lo.rxframe 0 1553373819
net.lo.rxcompressed 0 1553373819
net.lo.rxmulticast 0 1553373819
net.lo.txbytes 1227 1553373819
net.lo.txpackets 15 1553373819
net.lo.txerrors 0 1553373819
net.lo.txdrops 0 1553373819
net.lo.txfifo 0 1553373819
net.lo.txcolls 0 1553373819
net.lo.txcarrier 0 1553373819
net.lo.txcompressed 0 1553373819
net.br-5cb21acb080f.rxbytes 0 1553373819
net.br-5cb21acb080f.rxpackets 0 1553373819
net.br-5cb21acb080f.rxerrors 0 1553373819
net.br-5cb21acb080f.rxdrops 0 1553373819
net.br-5cb21acb080f.rxfifo 0 1553373819
net.br-5cb21acb080f.rxframe 0 1553373819
net.br-5cb21acb080f.rxcompressed 0 1553373819
net.br-5cb21acb080f.rxmulticast 0 1553373819
net.br-5cb21acb080f.txbytes 0 1553373819
net.br-5cb21acb080f.txpackets 0 1553373819
net.br-5cb21acb080f.txerrors 0 1553373819
net.br-5cb21acb080f.txdrops 0 1553373819
net.br-5cb21acb080f.txfifo 0 1553373819
net.br-5cb21acb080f.txcolls 0 1553373819
net.br-5cb21acb080f.txcarrier 0 1553373819
net.br-5cb21acb080f.txcompressed 0 1553373819
net.docker0.rxbytes 0 1553373819
net.docker0.rxpackets 0 1553373819
net.docker0.rxerrors 0 1553373819
net.docker0.rxdrops 0 1553373819
net.docker0.rxfifo 0 1553373819
net.docker0.rxframe 0 1553373819
net.docker0.rxcompressed 0 1553373819
net.docker0.rxmulticast 0 1553373819
net.docker0.txbytes 0 1553373819
net.docker0.txpackets 0 1553373819
net.docker0.txerrors 0 1553373819
net.docker0.txdrops 0 1553373819
net.docker0.txfifo 0 1553373819
net.docker0.txcolls 0 1553373819
net.docker0.txcarrier 0 1553373819
net.docker0.txcompressed 0 1553373819
```

## Configuration

### Asset Manifest

```yaml
---
type: Asset
api_version: core/v2
metadata:
  name: system-profile-linux-amd64
  namespace: default
spec:
  url: https://github.com/sensu-skunkworks/system-profile-linux/releases/download/0.4/system-profile-linux_0.4_linux_amd64.tar.gz
  sha512: b1eaeb7e9b3aadb3b184cc49995a0cb411c5962cdb4a921b8c295b03ad9c551c
```

### Check Manifest(s)

```yaml
---
type: CheckConfig
api_version: core/v2
metadata:
  name: system-profile-linux
  namespace: default
spec:
  command: system-profile-linux
  runtime_assets: system-profile-linux-amd64
  subscriptions:
  - linux
  interval: 10
  timeout: 5
  output_metric_format: graphite_plaintext
  output_metric_handlers:
  - influxdb
```

## Metrics

### `/proc/stat` metrics

- **name:** `cpu.user` (or `cpuN.user`)  
  **description:** Normal processes executing in user mode; `cpu.user` is the
  aggregate of the `cpuN.user` metrics (one per physical or virtual CPU).  
  **unit:** [jiffies][jiffies]  

- **name:** `cpu.nice`  
  **description:** Niced processes executing in user mode; `cpu.nice` is the
  aggregate of the `cpuN.nice` metrics (one per physical or virtual CPU).  
  **unit:** [jiffies][jiffies]  

- **name:** `cpu.system`  
  **description:** Processes executing in kernel mode; `cpu.system` is the
  aggregate of the `cpuN.system` metrics (one per physical or virtual CPU).  
  **unit:** [jiffies][jiffies]  

- **name:** `cpu.idle`  
  **description:** Twiddling thumbs; `cpu.idle` is the aggregate of the
  `cpuN.idle` metrics (one per physical or virtual CPU).  
  **unit:** [jiffies][jiffies]  

- **name:** `cpu.iowait`  
  **description:** Waiting for I/O to complete; `cpu.iowait` is the aggregate of
  the `cpuN.iowait` metrics (one per physical or virtual CPU).  
  **unit:** [jiffies][jiffies]  

- **name:** `cpu.irq`  
  **description:** Servicing interrupt requests; `cpu.irq` is the aggregate of
  the `cpuN.irq` metrics (one per physical or virtual CPU).
  **unit:** [jiffies][jiffies]

- **name:** `cpu.softirq`  
  **description:** Servicing soft interrupt requests; `cpu.softirq` is the
  aggregate of the `cpuN.softirq` metrics (one per physical or virtual CPU).  
  **unit:** [jiffies][jiffies]  

- **name:** `cpu.steal`  
  **description:**  
  **unit:**  

- **name:** `intr`  
  **description:** Interrupts serviced since boot time.  
  **unit:** counter  

- **name:** `ctxt`  
  **description:** Total number of context switches across all CPUs.  
  **unit:** counter  

- **name:** `btime`  
  **description:** The time the system booted.  
  **unit:** seconds since the Unix epoch  

- **name:** `processes`  
  **description:** The number of processes and threads created, which includes
  (but is not limited to) those created by calls to the fork() and clone()
  system calls.  
  **unit:** counter  

- **name:** `procs_running`  
  **description:** The total number of processes running on all CPUs.  
  **unit:** counter  

- **name:** `procs_blocked`  
  **description:** The number of processes currently blocked, waiting for I/O to
  complete.  
  **unit:** counter  

- **name:** `softirq`  
  **description:**  
  **unit:**  

References:
- [LinuxHowTos: `/proc/stat` explained](http://www.linuxhowtos.org/System/procstat.htm)  
- [Wikipedia: Interrupt request (PC architecture)](https://en.wikipedia.org/wiki/Interrupt_request_(PC_architecture))  

[jiffies]: http://man7.org/linux/man-pages/man7/time.7.html

### `/proc/loadavg` metrics

Coming soon.

### `/proc/net/dev` metrics

Coming soon.

### `/proc/meminfo` metrics

Coming soon.

## Service Checks

This plugin does not provide any service health checking capabilities.
