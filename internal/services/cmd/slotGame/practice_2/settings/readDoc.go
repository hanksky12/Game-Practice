package settings

import (
	"gamePractice/internal/pkg/entity/settings/payout"
	"gamePractice/internal/pkg/entity/settings/reel"
	"gamePractice/internal/services/cmd/slotGame/common/settings"
)

type Doc struct {
	BasePath string
	BaseDoc  *settings.Doc
}

func (d *Doc) getPath(fileName string) string {
	return d.BasePath + "/internal/services/cmd/slotGame/practice_2/settings/file/" + fileName
}

func (d *Doc) ReadWheelTable() (bool, *reel.Table, *reel.Table) {
	path := d.getPath("50.csv")
	return d.BaseDoc.ReadWheelTable(path)
}

func (d *Doc) ReadPayTable() (bool, *payout.Table) {
	path := d.getPath("table.csv")
	return d.BaseDoc.ReadPayTable(path)
}
