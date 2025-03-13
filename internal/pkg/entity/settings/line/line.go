package line

import "fmt"

type Table struct {
	Name string
	Rows []*Row
}

type Row struct {
	Case []int
}

func (t Table) String() string {
	return fmt.Sprintf("Table{Name: %s, Rows: %v}", t.Name, t.RowsToString())
}

func (t Table) RowsToString() string {
	var result string
	for _, row := range t.Rows {
		result += fmt.Sprintf("{Case: %v} ", row.Case)
	}
	return result
}
