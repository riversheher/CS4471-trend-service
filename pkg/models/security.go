package models

import (
	"time"
)

type Security struct {
	Ticker        string
	Price         float64
	ChangeAmount  float64
	ChangePercent float64
	Volume        float64
	Timestamp     time.Time
}
