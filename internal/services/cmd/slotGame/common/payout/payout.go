package payout

import (
	"gamePractice/internal/pkg/entity/settings/payout"
	log "github.com/sirupsen/logrus"
)

type Payout struct {
	SingleBet   float64
	BetBase     float64
	PayoutTable *payout.Table
}

func (p *Payout) QueryLineScore(count int, symbol string) float64 {
	for _, payoutRow := range p.PayoutTable.Rows {
		if payoutRow.Symbol != symbol {
			continue
		}
		return p.SingleBet * (payoutRow.Pays[count-1] / p.BetBase)
	}
	log.Error("No payout for symbol !!!!!!!!!", symbol)
	panic("No payout for symbol !!!!!!!!!")
	return 0
}

func (p *Payout) QueryWayScore(count int, symbol string, way float64) float64 {
	for _, payoutRow := range p.PayoutTable.Rows {
		if payoutRow.Symbol != symbol {
			continue
		}
		return p.SingleBet * (payoutRow.Pays[count-1] / p.BetBase) * way
	}
	log.Error("No payout for symbol !!!!!!!!!", symbol)
	panic("No payout for symbol !!!!!!!!!")
	return 0
}

//log.Info("找到同樣")
//log.Info(p.SingleBet)
//log.Info(payoutRow.Pays)
//log.Info(payoutRow.Pays[count-1])
//log.Info(p.BetBase)
//log.Info(way)
