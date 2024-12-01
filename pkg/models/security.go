package models

import (
	"errors"
	"strconv"
	"strings"
	"time"
)

type Security struct {
	Ticker        string  `json:"ticker"`
	Price         float64 `json:"price"`
	ChangeAmount  float64 `json:"change_amount"`
	ChangePercent float64 `json:"change_percentage"`
	Volume        int     `json:"volume"`
	Timestamp     time.Time
}

type SecurityResponse struct {
	Ticker        string `json:"ticker"`
	Price         string `json:"price"`
	ChangeAmount  string `json:"change_amount"`
	ChangePercent string `json:"change_percentage"`
	Volume        string `json:"volume"`
}

type GainersLosersResponse struct {
	Metadata     string             `json:"metadata"`
	Last_updated string             `json:"last_updated"`
	Gainers      []SecurityResponse `json:"top_gainers"`
	Losers       []SecurityResponse `json:"top_losers"`
	Active       []SecurityResponse `json:"most_actively_traded"`
}

func (s *SecurityResponse) ToSecurity(time time.Time) (*Security, error) {

	price, err := strconv.ParseFloat(s.Price, 64)
	if err != nil {
		return nil, errors.New("failed to parse price: " + err.Error())
	}

	changeAmount, err := strconv.ParseFloat(s.ChangeAmount, 64)
	if err != nil {
		return nil, errors.New("failed to parse change amount: " + err.Error())
	}

	cleanedPercent := strings.Replace(s.ChangePercent, "%", "", -1)
	changePercent, err := strconv.ParseFloat(cleanedPercent, 64)
	if err != nil {
		return nil, errors.New("failed to parse change percent: " + err.Error())
	}
	changePercent = changePercent / 100

	volume, err := strconv.Atoi(s.Volume)
	if err != nil {
		return nil, errors.New("failed to parse volume: " + err.Error())
	}

	return &Security{
		Ticker:        s.Ticker,
		Price:         price,
		ChangeAmount:  changeAmount,
		ChangePercent: changePercent,
		Volume:        volume,
		Timestamp:     time,
	}, nil
}
