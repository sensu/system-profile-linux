package main

/* collects system metrics, using the graphite plain-text format */

import (
	"fmt"
	"os"
	"time"

	"github.com/sreejita-biswas/sensu-extensions-system-profile/plugins"

	"github.com/spf13/cobra"
)

var (
	interval        int64
	prefix          string
	host            string
	port            int
	graphiteEnabled bool
)

func main() {
	rootCmd := configureRootCommand()
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}

func run(cmd *cobra.Command, args []string) error {
	if len(args) != 0 {
		_ = cmd.Help()
		return fmt.Errorf("invalid argument(s) received")
	}
	plugins.SendMetrics(graphiteEnabled, host, port, prefix, time.Duration(interval))
	return nil
}

func configureRootCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "check-cloudwatch-alarms",
		Short: "The Sensu Go Aws Cloudwatch handler for alarms management",
		RunE:  run,
	}
	cmd.Flags().Int64Var(&interval, "interval", 10, "interval between two metrics data sent to graphite")
	cmd.Flags().StringVar(&prefix, "prefix", "com.sensuapp.demo", "prefix to be added before metric")
	cmd.Flags().StringVar(&host, "host", "127.0.0.1", "graphite host where to send the metrics")
	cmd.Flags().IntVar(&port, "port", 8080, "graphite host port")
	cmd.Flags().BoolVar(&graphiteEnabled, "graphiteEnabled", true, "graphite enabled or not")

	cmd.MarkFlagRequired("interval")
	cmd.MarkFlagRequired("prefix")
	cmd.MarkFlagRequired("host")
	cmd.MarkFlagRequired("port")
	return cmd
}
