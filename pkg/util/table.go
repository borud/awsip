package util

import (
	"fmt"
	"os"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
)

var awsipStyle = table.Style{
	Name: "awsipStyle",
	Box:  table.StyleBoxLight,
	Color: table.ColorOptions{
		IndexColumn:  text.Colors{text.BgHiBlue, text.FgBlack},
		Footer:       text.Colors{text.BgBlue, text.FgBlack},
		Header:       text.Colors{text.BgBlue, text.FgWhite},
		Row:          text.Colors{text.BgHiWhite, text.FgBlack},
		RowAlternate: text.Colors{text.BgWhite, text.FgBlack},
	},
	Format: table.FormatOptions{
		Footer: text.FormatDefault,
		Header: text.FormatDefault,
		Row:    text.FormatDefault,
	},

	HTML:    table.DefaultHTMLOptions,
	Options: table.OptionsNoBordersAndSeparators,
	Title: table.TitleOptions{
		Align:  text.AlignLeft,
		Colors: text.Colors{text.BgYellow, text.FgBlack},
		Format: text.FormatDefault,
	},
}

// NewTableOutput creates a new table writer with the specified settings
func NewTableOutput(useColors bool) table.Writer {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)

	if useColors {
		t.SetStyle(awsipStyle)
	}
	return t
}

// RenderTable renders the table according to settings
func RenderTable(t table.Writer, format string) {
	switch format {
	case "csv":
		t.SetTitle("")
		t.RenderCSV()
	case "html":
		t.RenderHTML()
	case "markdown":
		t.RenderMarkdown()
	default:
		fmt.Println()
		t.Render()
		fmt.Println()
	}
}
