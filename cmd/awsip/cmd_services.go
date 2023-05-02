package main

import (
	"sort"

	"github.com/borud/aws-ip-ranges/pkg/util"
	"github.com/jedib0t/go-pretty/v6/table"
)

//lint:file-ignore SA5008 Ignore duplicate struct tags
type servicesCmd struct {
	Format   string `short:"f" long:"format" default:"table" description:"output format" choice:"table" choice:"html" choice:"markdown" choice:"csv"`
	NoColors bool   `short:"c" long:"no-color" description:"do not colorize output"`
}

func (se *servicesCmd) Execute([]string) error {
	ipRanges, err := getRanges(opt)
	if err != nil {
		return err
	}

	m := map[string]struct{}{}

	for _, v := range ipRanges.Prefixes {
		m[v.Service] = struct{}{}
	}
	for _, v := range ipRanges.Ipv6Prefixes {
		m[v.Service] = struct{}{}
	}
	keys := []string{}
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	tab := util.NewTableOutput(!se.NoColors)
	tab.AppendHeader(table.Row{
		"service",
	})

	for _, service := range keys {
		tab.AppendRow(table.Row{service})
	}

	util.RenderTable(tab, se.Format)
	return nil
}
