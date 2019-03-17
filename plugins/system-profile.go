package plugins

import (
	"fmt"
	"io/ioutil"
	"strings"
	"time"

	graphite "github.com/marpaia/graphite-golang"
)

var metrics = []graphite.Metric{}

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
			for i := 1; i < len(fields); i++ {
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
	metric := graphite.NewMetric(metricName, value, time.Now().Unix())
	metrics = append(metrics, metric)
}

func flushMetrics() {
	metrics = []graphite.Metric{}
}

func SendMetrics(graphiteEnabled bool, graphiteHost string, graphitePort int, metricsPrefix string, interval time.Duration) {
	var graphiteInstance *graphite.Graphite
	var err error
	if graphiteEnabled {
		graphiteInstance, err = graphite.NewGraphiteWithMetricPrefix(graphiteHost, graphitePort, metricsPrefix)
	} else {
		graphiteInstance = graphite.NewGraphiteNop(graphiteHost, graphitePort)
	}

	fmt.Printf("Loaded Graphite connection: %#v", graphiteInstance)
	// if you couldn't connect to graphite, use a nop
	if err != nil {
		graphiteInstance = graphite.NewGraphiteNop(graphiteHost, graphitePort)
	}

	for {
		parseProcStat()
		procLoadavgMetrics()
		procMeminfoMetrics()
		graphiteInstance.SendMetrics(metrics)
		time.Sleep(interval * time.Minute)
		flushMetrics()
	}
}
