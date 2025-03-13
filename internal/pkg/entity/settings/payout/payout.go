package payout

import "fmt"

type Table struct {
	Name string
	Rows []*Row
}

type Row struct {
	Symbol string
	Pays   []float64
}

func (t Table) String() string {
	return fmt.Sprintf("Table{Name: %s, Rows: %v}", t.Name, t.RowsToString())
}

func (t Table) RowsToString() string {
	var result string
	for _, row := range t.Rows {
		result += fmt.Sprintf("{Symbol: %s, P1: %f, P2: %f, P3: %f, P4: %f, P5: %f} ", row.Symbol, row.Pays[0], row.Pays[1], row.Pays[2], row.Pays[3], row.Pays[4])
	}
	return result
}
