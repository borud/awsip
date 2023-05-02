package main

import (
	"github.com/borud/awsip/pkg/util"
	"github.com/jedib0t/go-pretty/v6/table"
)

//lint:file-ignore SA5008 Ignore duplicate struct tags
type rangeCmd struct {
	Service  string `short:"s" long:"service" default:"AMAZON" description:"service, use AMAZON for superset of all services"`
	Region   string `short:"r" long:"region" default:"GLOBAL" description:"region, use GLOBAL for superset of all regions"`
	IPV      int    `short:"i" long:"ip" default:"0" description:"IP version, use 0 for both IPv4 and IPv6" choice:"4" choice:"6" choice:"0"`
	Format   string `short:"f" long:"format" default:"table" description:"output format" choice:"table" choice:"html" choice:"markdown" choice:"csv"`
	NoColors bool   `short:"c" long:"no-color" description:"do not colorize output"`
}

func (r *rangeCmd) Execute([]string) error {
	ipRanges, err := getRanges(opt)
	if err != nil {
		return err
	}

	tab := util.NewTableOutput(!r.NoColors)
	tab.SetTitle("IP ranges for service=%s, region=%s", r.Service, r.Region)
	tab.AppendHeader(table.Row{
		"region",
		"service",
		"border group",
		"prefix",
	})

	if r.IPV == 0 || r.IPV == 4 {
		for _, v := range ipRanges.Prefixes {
			if v.Region == r.Region && v.Service == r.Service {
				tab.AppendRow(table.Row{
					v.Region,
					v.Service,
					v.NetworkBorderGroup,
					*v.IPPrefix,
				})
			}
		}
	}
	if r.IPV == 0 || r.IPV == 6 {
		for _, v := range ipRanges.Ipv6Prefixes {
			if v.Region == r.Region && v.Service == r.Service {
				tab.AppendRow(table.Row{
					v.Region,
					v.Service,
					v.NetworkBorderGroup,
					*v.Ipv6Prefix,
				})
			}
		}
	}

	util.RenderTable(tab, r.Format)

	return nil
}
