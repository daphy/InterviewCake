package stockPrice

import (
  "testing"
)

func TestGetMaxProfitFunc1(tt *testing.T) {
  prices := []int{10, 7, 5, 8, 11, 9}
  profit, err := getMaxProfit(prices)
  if err != nil {
    tt.Error(err)
  }
  if profit != 6 { // Buy at 5; sell at 11.
    tt.Errorf("Profit calculated wrong. Expected: 6; actual: %d", profit)
  }
}

func TestGetMaxProfitFunc2(tt *testing.T) {
  prices := []int{10, 7, 5, 11, 4, 2, 5}
  profit, err := getMaxProfit(prices)
  if err != nil {
    tt.Error(err)
  }
  if profit != 6 { // Buy at 5, sell at 11.
    tt.Errorf("Profit calculated wrong. Expected: 6; actual: %d", profit)
  }
}

func TestGetMaxProfitFunc3(tt *testing.T) {
  prices := []int{1, 7, 5, 11, 4, 2, 5, 14}
  profit, err := getMaxProfit(prices)
  if err != nil {
    tt.Error(err)
  }
  if profit != 13 { // Buy at a, sell at 14.
    tt.Errorf("Profit calculated wrong. Expected: 13; actual: %d", profit)
  }
}

func TestGetMaxProfitNoProfit(tt *testing.T) {
  prices := []int{19, 18, 17, 16, 15, 14, 13, 12}
  profit, err := getMaxProfit(prices)
  if err != nil {
    tt.Error(err)
  }
  if profit != 0 {
    tt.Errorf("Profit calculated wrong. Expected: 0; actual: %d", profit)
  }
}

func TestGetMaxProfitOnePoint(tt *testing.T) {
  prices := []int{19, 18, 17, 16, 17, 14, 13, 12}
  profit, err := getMaxProfit(prices)
  if err != nil {
    tt.Error(err)
  }
  if profit != 1 { // Nothing to buy or sell.
    tt.Errorf("Profit calculated wrong. Expected: 0; actual: %d", profit)
  }
}

func TestGetMaxProfitEmptySlice(tt *testing.T) {
  prices := []int{}
  profit, err := getMaxProfit(prices)
  if err != nil {
    tt.Error(err)
  }
  if profit != 0 { // Nothing to buy or sell.
    tt.Errorf("Profit calculated wrong. Expected: 0; actual: %d", profit)
  }
}

func TestGetMaxProfitOnePricePoint(tt *testing.T) {
  prices := []int{1}
  profit, err := getMaxProfit(prices)
  if err != nil {
    tt.Error(err)
  }
  if profit != 0 { // Nothing to buy or sell.
    tt.Errorf("Profit calculated wrong. Expected: 0; actual: %d", profit)
  }
}

func TestGetMaxProfitStockPriceZero(tt *testing.T) {
  prices := []int{0, 0, 0}
  profit, err := getMaxProfit(prices)
  if err != nil {
    tt.Error(err)
  }
  if profit != 0 {
    tt.Errorf("Profit calculated wrong. Expected: 0; actual: %d", profit)
  }
}

func TestGetMaxProfitNegativePrice(tt *testing.T) {
  prices := []int{-1, -3, 0}
  _, err := getMaxProfit(prices)
  if err == nil {
    tt.Errorf("Negative stock price should err out.")
  }
}
