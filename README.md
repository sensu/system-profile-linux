# sensu-extensions-system-profile

TravisCI: [![TravisCI Build Status](https://travis-ci.org/sreejita-biswas/sensu-extensions-system-profile.svg?branch=master)](https://travis-ci.org/sreejita-biswas/sensu-extensions-system-profile)

## Functionality

 - sensu-extensions-system-profile
    This check extension collects Linux system metrics. This extension reads and parses /proc/stat, /proc/loadavg,   /proc/net/dev, and /proc/meminfo to produce Linux system metrics in the Graphite plaintext format

## Files

* /bin/sensu-extensions-system-profile

## Usage

**sensu-extensions-system-profile**
```
./sensu-extensions-system-profile --host="127.0.0.1" --port=8080 --interval=10 --prefix="com.sensuapp.demo"

Sample Output : 
cpu.user 2255 2019-03-17 17:07:10
cpu.nice 34 2019-03-17 17:07:10
cpu.system 2290 2019-03-17 17:07:10
cpu.idle 22625563 2019-03-17 17:07:10
cpu.iowait 6290 2019-03-17 17:07:10
cpu.irq 127 2019-03-17 17:07:10
cpu.softirq 456 2019-03-17 17:07:10
cpu0.user 1132 2019-03-17 17:07:10
cpu0.nice 34 2019-03-17 17:07:10
cpu0.system 1441 2019-03-17 17:07:10
cpu0.idle 11311718 2019-03-17 17:07:10
cpu0.iowait 3675 2019-03-17 17:07:10
cpu0.irq 127 2019-03-17 17:07:10
cpu0.softirq 438 2019-03-17 17:07:10
cpu1.user 1123 2019-03-17 17:07:10
cpu1.nice 0 2019-03-17 17:07:10
cpu1.system 849 2019-03-17 17:07:10
cpu1.idle 11313845 2019-03-17 17:07:10
cpu1.iowait 2614 2019-03-17 17:07:10
cpu1.irq 0 2019-03-17 17:07:10
cpu1.softirq 18 2019-03-17 17:07:10
intr 4 2019-03-17 17:07:10
ctxt 1990473 2019-03-17 17:07:10
btime 1062191376 2019-03-17 17:07:10
processes 2915 2019-03-17 17:07:10
procs_running 1 2019-03-17 17:07:10
procs_blocked 0 2019-03-17 17:07:10
load_avg.1_min 0.75 2019-03-17 17:07:10
load_avg.5_min 0.35 2019-03-17 17:07:10
load_avg.15_min 0.25 2019-03-17 17:07:10
memory.memtotal 1921988 2019-03-17 17:07:10
memory.memfree 1374408 2019-03-17 17:07:10
memory.buffers 32688 2019-03-17 17:07:10
memory.cached 370540 2019-03-17 17:07:10
swap.swapcached 0 2019-03-17 17:07:10
memory.active 344604 2019-03-17 17:07:10
swap.swaptotal 1048572 2019-03-17 17:07:10
swap.swapfree 1048572 2019-03-17 17:07:10
memory.committed_as 134216 2019-03-17 17:07:10
```
## External Dependency

```
go get github.com/marpaia/graphite-golang

```
