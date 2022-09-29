package model

import (
	"time"

	"github.com/shopspring/decimal"
)

type Event struct {
	Outcome   []string
	Odds      []decimal.Decimal
	StartTime time.Time
}
