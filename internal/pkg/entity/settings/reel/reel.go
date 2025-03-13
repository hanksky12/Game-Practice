package reel

import "fmt"

type Table struct {
	Name  string
	Reels []*Reel
}

type Reel struct {
	Items []string
}

func (t Table) String() string {
	return fmt.Sprintf("Table{Name: %s, Reels: %v}", t.Name, t.ReelsToString())
}

func (t Table) ReelsToString() string {
	var result string
	for _, reel := range t.Reels {
		result += fmt.Sprintf("{Items: %v} ", reel.Items)
	}
	return result
}
