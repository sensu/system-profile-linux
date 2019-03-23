package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"time"
)

var metrics = []string{}

func parseProcStat() {
	contents, err := ioutil.ReadFile("/proc/stat")
	if err != nil {
		return
	}
	cpuMetrics := []string{"user", "nice", "system", "idle", "iowait", "irq", "softirq", "steal", "guest"}
	lines := strings.Split(string(contents), "\n")
	for _, line := range lines {
		fields := strings.Fields(line)
		if len(fields) > 1 && strings.HasPrefix(fields[0], "cpu") {
			for i := 1; i < len(fields) && i < 9; i++ {
				addMetric([]string{fields[0], cpuMetrics[i-1]}, fields[i])
			}
		} else if len(fields) > 1 {
			addMetric([]string{fields[0]}, fields[len(fields)-1])
		}
	}
	return
}

func procLoadavgMetrics() {
	contents, err := ioutil.ReadFile("/proc/loadavg")
	if err != nil {
		return
	}
	metrcis := []string{"1_min", "5_min", "15_min"}
	lines := strings.Split(string(contents), "\n")
	for _, line := range lines {
		fields := strings.Fields(line)
		for index, field := range fields {
			if index <= 2 {
				addMetric([]string{"load_avg", metrcis[index]}, field)
			}
		}
	}
	return
}

func procNetDevMetrics() {
	var interface_name string
	contents, err := ioutil.ReadFile("/proc/net/dev")
	if err != nil {
		return
	}
	metrcis := []string{"rxbytes", "rxpackets", "rxerrors", "rxdrops", "rxfifo", "rxframe", "rxcompressed", "rxmulticast", "txbytes", "txpackets", "txerrors", "txdrops", "txfifo", "txcolls", "txcarrier", "txcompressed"}
	lines := strings.Split(string(contents), "\n")
	for lineNumber, line := range lines {
		if lineNumber == 0 {
			continue
		}
		fields := strings.Fields(line)
		for index, field := range fields {
			if index == 0 {
				interface_name = strings.Split(field, ":")[0]
				continue
			}
			addMetric([]string{"net", interface_name, metrcis[index-1]}, field)
		}
	}
	return
}

func procMeminfoMetrics() {
	contents, err := ioutil.ReadFile("/proc/meminfo")
	if err != nil {
		return
	}
	lines := strings.Split(string(contents), "\n")
	for _, line := range lines {
		fields := strings.Fields(line)
		if len(fields) <= 0 {
			continue
		}
		metric := strings.ToLower(strings.Split(fields[0], ":")[0])
		if strings.HasPrefix(fields[0], "Mem") || (metric == "buffers") || (metric == "cached") || (metric == "active") || (metric == "committed_as") {
			addMetric([]string{"memory", metric}, fields[1])
		} else if strings.HasPrefix(fields[0], "Swap") {
			addMetric([]string{"swap", metric}, fields[1])
		}
	}
	return
}

func addMetric(metricType []string, value string) {
	metricName := strings.Join(metricType, ".")
	timeNow := time.Now().Unix()
	outputs := []string{metricName, value, strconv.FormatInt(timeNow, 10)}
	metrics = append(metrics, strings.Join(outputs, " "))
}

func flushMetrics() {
	metrics = []string{}
}

func main() {
	parseProcStat()
	procLoadavgMetrics()
	procMeminfoMetrics()
	procNetDevMetrics()
	for _, metric := range metrics {
		fmt.Println(fmt.Sprintf("%s", metric))
	}
	flushMetrics()
}
