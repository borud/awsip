package main

import (
	"fmt"
	"sort"

	"github.com/borud/aws-ip-ranges/pkg/util"
	"github.com/jedib0t/go-pretty/v6/table"
)

//lint:file-ignore SA5008 Ignore duplicate struct tags
type regionsCmd struct {
	Format   string `short:"f" long:"format" default:"table" description:"output format" choice:"table" choice:"html" choice:"markdown" choice:"csv"`
	NoColors bool   `short:"c" long:"no-color" description:"do not colorize output"`
}

func (re *regionsCmd) Execute([]string) error {
	ipRanges, err := getRanges(opt)
	if err != nil {
		return err
	}

	m := map[string]struct{}{}

	for _, v := range ipRanges.Prefixes {
		m[v.Region] = struct{}{}
	}
	for _, v := range ipRanges.Ipv6Prefixes {
		m[v.Region] = struct{}{}
	}
	keys := []string{}
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	tab := util.NewTableOutput(!re.NoColors)
	tab.AppendHeader(table.Row{
		"regions",
	})

	for _, region := range keys {
		fmt.Println(region)
		tab.AppendRow(table.Row{region})
	}

	util.RenderTable(tab, re.Format)
	return nil
}
