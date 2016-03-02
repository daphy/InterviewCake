/*******************************************************************************
https://www.interviewcake.com/question/python/stock-price

Writing programming interview questions hasn't made me rich. Maybe trading
Apple stocks will.
Suppose we could access yesterday's stock prices as a list, where:

The indices are the time in minutes past trade opening time, which was 9:30am
local time.
The values are the price in dollars of Apple stock at that time.
So if the stock cost $500 at 10:30am, stock_prices_yesterday[60] = 500.

Write an efficient function that takes stock_prices_yesterday and returns the
best profit I could have made from 1 purchase and 1 sale of 1 Apple stock
yesterday.

For example:

  stock_prices_yesterday = [10, 7, 5, 8, 11, 9]

get_max_profit(stock_prices_yesterday)
# returns 6 (buying for $5 and selling for $11)

*******************************************************************************/
package stockPrice

import (
  "fmt"
)

func getMaxProfit(stockPrices []int) (int, error) {
  currentMin := 99999 // TODO: use max_int
  currentMaxProfit := 0
  for _, price := range stockPrices {
    if price < 0 {
      return currentMaxProfit, fmt.Errorf("Stock price can not be negative.")
    }
    if price <= currentMin {
      currentMin = price
      continue
    }
    // price > currentMin. There is a chance for profit!
    tempProfit := price - currentMin
    if tempProfit > currentMaxProfit {
      currentMaxProfit = tempProfit
    }
  }
  return currentMaxProfit, nil
}
