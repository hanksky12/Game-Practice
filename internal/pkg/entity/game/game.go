package game

import (
	"fmt"
	"gamePractice/internal/pkg/entity/settings/reel"
)

type Board struct {
	Name  string
	Shape []int
	Reels []*reel.Reel
}

func (t Board) String() string {
	//data, _ := json.MarshalIndent(t, "", "  ")
	//return string(data)
	return fmt.Sprintf("Board{Name: %s,Shape: %v, Reels: %v}", t.Name, t.Shape, t.ReelsToString())
}

func (t Board) ReelsToString() string {
	var result string
	for _, reel := range t.Reels {
		result += fmt.Sprintf("{ Items: %v} ", reel.Items)
	}
	return result
}
