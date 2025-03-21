package settings

import (
	"gamePractice/internal/pkg/const/game"
	"gamePractice/internal/pkg/entity/settings/line"
	"gamePractice/internal/pkg/entity/settings/payout"
	"gamePractice/internal/pkg/entity/settings/reel"
	"gamePractice/internal/pkg/util/file"
	log "github.com/sirupsen/logrus"
	"strconv"
	"strings"
)

type Doc struct {
	BasePath string
}

func (d *Doc) ReadWheelTable(path string) (bool, *reel.Table, *reel.Table) {
	records, ok := d.readCsv(path)
	//records, ok := d.readCsv(d.getPath("Fake.csv"))
	if !ok {
		return false, &reel.Table{}, &reel.Table{}
	}
	records = d.keepFirst10Row(records)
	d.clearEmptyAndUpper(records)
	mgReelRecords := records[:5]
	fgReelRecords := records[5:]
	//log.Info("mgReelRecords: ", mgReelRecords)
	//log.Info("fgReelRecords: ", fgReelRecords)
	mgTable := d.toReel(game.Main, mgReelRecords)
	fgTable := d.toReel(game.Free, fgReelRecords)
	//log.Info("MgReel: ", mgTable)
	//log.Info("FgReel: ", fgTable)
	return true, mgTable, fgTable
}

func (d *Doc) ReadPayTable(path string) (bool, *payout.Table) {
	records, ok := d.readCsv(path)
	if !ok {
		return false, &payout.Table{}
	}
	table := d.toPayOut(records)
	//log.Info("PayTable: ", table)
	return true, table
}

func (d *Doc) ReadLineTable(path string) (bool, *line.Table) {
	records, ok := d.readCsv(path)
	if !ok {
		return false, &line.Table{}
	}
	//log.Info("原始 Table: ", records)
	table := d.toLine(records)
	//log.Info("LineTable: ", *table)
	//log.Info("LineTable: ", len(table.Rows))
	return true, table
}

func (d *Doc) toLine(records [][]string) *line.Table {
	var lines []*line.Row
	for i := 1; i < len(records); i++ { // 跳過第一行表頭
		rowData := records[i]
		indexes := d.getLineIndexes(rowData)
		lines = append(lines, &line.Row{Case: indexes})
	}
	return &line.Table{Name: "LineTable", Rows: lines}
}

func (d *Doc) getLineIndexes(rowData []string) []int {
	// 動態解析 Case 欄位
	var indexes []int
	for j := 1; j < len(rowData); j++ { // 從第 2 欄開始 (跳過 Num)
		index, _ := strconv.Atoi(rowData[j])
		indexes = append(indexes, index)
	}
	return indexes
}

func (d *Doc) toPayOut(records [][]string) *payout.Table {
	var payouts []*payout.Row
	for _, row := range records {
		if len(row) < 6 {
			continue
		}
		p1, _ := strconv.ParseFloat(row[1], 64)
		p2, _ := strconv.ParseFloat(row[2], 64)
		p3, _ := strconv.ParseFloat(row[3], 64)
		p4, _ := strconv.ParseFloat(row[4], 64)
		p5, _ := strconv.ParseFloat(row[5], 64)

		payouts = append(payouts, &payout.Row{
			Symbol: strings.TrimSpace(row[0]),
			Pays:   []float64{p1, p2, p3, p4, p5},
		})
	}
	return &payout.Table{Name: "PayTable", Rows: payouts}
}

func (d *Doc) toReel(name string, records [][]string) *reel.Table {
	var reels []*reel.Reel

	for _, row := range records {
		var newRecords []string
		for _, item := range row {
			newRecords = append(newRecords, strings.TrimSpace(item))
		}
		//log.Info("row: ", row)
		reels = append(reels, &reel.Reel{
			Items: newRecords,
		})
	}
	return &reel.Table{Name: name, Reels: reels}
}

func (d *Doc) readCsv(path string) ([][]string, bool) {
	c := file.CSV{}
	records, err := c.Read(path)
	if err != nil {
		log.Info("ReadTable: ", err)
		return nil, false
	}
	//log.Info("原始 Table: ", records)
	return records, true
}

func (d *Doc) clearEmptyAndUpper(records [][]string) {
	for i := 0; i < len(records); i++ {
		var newRow []string
		for _, value := range records[i] {
			if value != "" { // 非空保留
				newRow = append(newRow, strings.ToUpper(value))
			}
		}
		records[i] = newRow // 更新原本
	}
}

func (d *Doc) keepFirst10Row(records [][]string) [][]string {
	limit := 10
	var filteredRecords [][]string
	for i, row := range records {
		if i+1 > limit {
			continue
		}
		filteredRecords = append(filteredRecords, row[1:])
	}
	return filteredRecords
}
