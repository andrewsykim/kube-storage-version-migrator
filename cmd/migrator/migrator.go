package main

import (
	goflag "flag"
	"fmt"
	"os"

	"github.com/kubernetes-sigs/kube-storage-version-migrator/cmd/migrator/app"
	"github.com/spf13/pflag"
	"k8s.io/klog"
)

func main() {
	klog.InitFlags(nil)
	pflag.CommandLine.AddGoFlagSet(goflag.CommandLine)
	pflag.Parse()
	pflag.VisitAll(func(flag *pflag.Flag) {
		klog.V(2).Infof("FLAG: --%s=%q", flag.Name, flag.Value)
	})
	command := app.NewMigratorCommand()
	if err := command.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}
