package global

import (
	"os"

	"github.com/olekukonko/tablewriter"
)

func NewTable() *tablewriter.Table {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeaderLine(false)
	table.SetBorder(false)
	table.SetColumnSeparator(" ")
	return table
}
