package settings

import (
	"gamePractice/internal/pkg/entity/settings/line"
	"gamePractice/internal/pkg/entity/settings/payout"
	"gamePractice/internal/pkg/entity/settings/reel"
	"gamePractice/internal/services/cmd/slotGame/common/settings"
)

type Doc struct {
	BaseDoc  *settings.Doc
	BasePath string
}

func (d *Doc) getPath(fileName string) string {
	return d.BasePath + "/internal/services/cmd/slotGame/practice_1/settings/file/" + fileName
}

func (d *Doc) ReadWheelTable() (bool, *reel.Table, *reel.Table) {
	path := d.getPath("CFZS.csv")
	return d.BaseDoc.ReadWheelTable(path)
}

func (d *Doc) ReadPayTable() (bool, *payout.Table) {
	path := d.getPath("table.csv")
	return d.BaseDoc.ReadPayTable(path)
}

func (d *Doc) ReadLineTable() (bool, *line.Table) {
	path := d.getPath("lines.csv")
	return d.BaseDoc.ReadLineTable(path)
}
