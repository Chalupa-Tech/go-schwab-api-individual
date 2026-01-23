package schwab

import (
	"context"
	"fmt"
	"net/url"
	"strconv"
	"strings"

	"github.com/placeholder/schwab-go/schwab/models"
)

// --- Market Data Service ---

// GetQuotes returns quotes for a list of symbols.
func (c *Client) GetQuotes(ctx context.Context, symbols []string, fields string, indicative bool) (models.QuoteResponse, error) {
	u := fmt.Sprintf("%s/quotes", c.MarketURL)
	v := url.Values{}
	v.Set("symbols", strings.Join(symbols, ","))
	if fields != "" {
		v.Set("fields", fields)
	}
	v.Set("indicative", strconv.FormatBool(indicative))
	u += "?" + v.Encode()

	var result models.QuoteResponse
	err := c.doRequest(ctx, "GET", u, nil, &result)
	return result, err
}

// GetQuote returns a quote for a single symbol.
func (c *Client) GetQuote(ctx context.Context, symbol string, fields string) (models.QuoteResponse, error) {
	u := fmt.Sprintf("%s/%s/quotes", c.MarketURL, url.PathEscape(symbol))
	if fields != "" {
		u += "?fields=" + url.QueryEscape(fields)
	}
	var result models.QuoteResponse
	err := c.doRequest(ctx, "GET", u, nil, &result)
	return result, err
}

// OptionChainParams holds parameters for GetOptionChain.
type OptionChainParams struct {
	Symbol           string
	ContractType     string // CALL, PUT, ALL
	StrikeCount      int
	IncludeQuotes    bool
	Strategy         string // SINGLE, ANALYTICAL, etc.
	Interval         float64
	Strike           float64
	Range            string
	FromDate         string // yyyy-MM-dd
	ToDate           string // yyyy-MM-dd
	Volatility       float64
	UnderlyingPrice  float64
	InterestRate     float64
	DaysToExpiration int
	ExpMonth         string // JAN, FEB, ...
	OptionType       string
	Entitlement      string // PN, NP, PP
}

// GetOptionChain returns option chain for an optionable Symbol.
func (c *Client) GetOptionChain(ctx context.Context, params OptionChainParams) (*models.OptionChain, error) {
	u := fmt.Sprintf("%s/chains", c.MarketURL)
	v := url.Values{}
	v.Set("symbol", params.Symbol)
	if params.ContractType != "" {
		v.Set("contractType", params.ContractType)
	}
	if params.StrikeCount > 0 {
		v.Set("strikeCount", strconv.Itoa(params.StrikeCount))
	}
	v.Set("includeUnderlyingQuote", strconv.FormatBool(params.IncludeQuotes))
	if params.Strategy != "" {
		v.Set("strategy", params.Strategy)
	}
	if params.Interval > 0 {
		v.Set("interval", fmt.Sprintf("%f", params.Interval))
	}
	if params.Strike > 0 {
		v.Set("strike", fmt.Sprintf("%f", params.Strike))
	}
	if params.Range != "" {
		v.Set("range", params.Range)
	}
	if params.FromDate != "" {
		v.Set("fromDate", params.FromDate)
	}
	if params.ToDate != "" {
		v.Set("toDate", params.ToDate)
	}
	if params.Volatility > 0 {
		v.Set("volatility", fmt.Sprintf("%f", params.Volatility))
	}
	if params.UnderlyingPrice > 0 {
		v.Set("underlyingPrice", fmt.Sprintf("%f", params.UnderlyingPrice))
	}
	if params.InterestRate > 0 {
		v.Set("interestRate", fmt.Sprintf("%f", params.InterestRate))
	}
	if params.DaysToExpiration > 0 {
		v.Set("daysToExpiration", strconv.Itoa(params.DaysToExpiration))
	}
	if params.ExpMonth != "" {
		v.Set("expMonth", params.ExpMonth)
	}
	if params.OptionType != "" {
		v.Set("optionType", params.OptionType)
	}
	if params.Entitlement != "" {
		v.Set("entitlement", params.Entitlement)
	}
	u += "?" + v.Encode()

	var result models.OptionChain
	err := c.doRequest(ctx, "GET", u, nil, &result)
	return &result, err
}

// GetExpirationChain returns option expiration chain for an optionable symbol.
func (c *Client) GetExpirationChain(ctx context.Context, symbol string) (*models.ExpirationChain, error) {
	u := fmt.Sprintf("%s/expirationchain", c.MarketURL)
	u += "?symbol=" + url.QueryEscape(symbol)
	var result models.ExpirationChain
	err := c.doRequest(ctx, "GET", u, nil, &result)
	return &result, err
}

// PriceHistoryParams holds parameters for GetPriceHistory.
type PriceHistoryParams struct {
	Symbol                string
	PeriodType            string // day, month, year, ytd
	Period                int
	FrequencyType         string // minute, daily, weekly, monthly
	Frequency             int
	StartDate             int64 // milliseconds
	EndDate               int64 // milliseconds
	NeedExtendedHoursData bool
	NeedPreviousClose     bool
}

// GetPriceHistory returns PriceHistory for a single symbol and date ranges.
func (c *Client) GetPriceHistory(ctx context.Context, params PriceHistoryParams) (*models.CandleList, error) {
	u := fmt.Sprintf("%s/pricehistory", c.MarketURL)
	v := url.Values{}
	v.Set("symbol", params.Symbol)
	if params.PeriodType != "" {
		v.Set("periodType", params.PeriodType)
	}
	if params.Period > 0 {
		v.Set("period", strconv.Itoa(params.Period))
	}
	if params.FrequencyType != "" {
		v.Set("frequencyType", params.FrequencyType)
	}
	if params.Frequency > 0 {
		v.Set("frequency", strconv.Itoa(params.Frequency))
	}
	if params.StartDate > 0 {
		v.Set("startDate", strconv.FormatInt(params.StartDate, 10))
	}
	if params.EndDate > 0 {
		v.Set("endDate", strconv.FormatInt(params.EndDate, 10))
	}
	v.Set("needExtendedHoursData", strconv.FormatBool(params.NeedExtendedHoursData))
	v.Set("needPreviousClose", strconv.FormatBool(params.NeedPreviousClose))
	u += "?" + v.Encode()

	var result models.CandleList
	err := c.doRequest(ctx, "GET", u, nil, &result)
	return &result, err
}

// GetMovers returns Movers for a specific index.
func (c *Client) GetMovers(ctx context.Context, symbolId string, sort string, frequency int) (*models.Movers, error) {
	u := fmt.Sprintf("%s/movers/%s", c.MarketURL, url.PathEscape(symbolId))
	v := url.Values{}
	if sort != "" {
		v.Set("sort", sort)
	}
	if frequency > 0 {
		v.Set("frequency", strconv.Itoa(frequency))
	}
	if len(v) > 0 {
		u += "?" + v.Encode()
	}

	var result models.Movers
	err := c.doRequest(ctx, "GET", u, nil, &result)
	return &result, err
}

// GetMarketHours returns Market Hours for different markets.
func (c *Client) GetMarketHours(ctx context.Context, markets []string, date string) (models.MarketHours, error) {
	u := fmt.Sprintf("%s/markets", c.MarketURL)
	v := url.Values{}
	v.Set("markets", strings.Join(markets, ","))
	if date != "" {
		v.Set("date", date)
	}
	u += "?" + v.Encode()

	var result models.MarketHours
	err := c.doRequest(ctx, "GET", u, nil, &result)
	return result, err
}

// GetMarketHour returns Market Hours for a single market.
func (c *Client) GetMarketHour(ctx context.Context, marketId string, date string) (models.MarketHours, error) {
	u := fmt.Sprintf("%s/markets/%s", c.MarketURL, url.PathEscape(marketId))
	if date != "" {
		u += "?date=" + url.QueryEscape(date)
	}
	var result models.MarketHours
	err := c.doRequest(ctx, "GET", u, nil, &result)
	return result, err
}

// GetInstruments returns Instruments by symbols and projections.
func (c *Client) GetInstruments(ctx context.Context, symbol string, projection string) (*models.InstrumentResponse, error) {
	u := fmt.Sprintf("%s/instruments", c.MarketURL)
	v := url.Values{}
	v.Set("symbol", symbol)
	v.Set("projection", projection)
	u += "?" + v.Encode()

	var result models.InstrumentResponse
	err := c.doRequest(ctx, "GET", u, nil, &result)
	return &result, err
}

// GetInstrumentByCusip returns Instrument by specific cusip.
func (c *Client) GetInstrumentByCusip(ctx context.Context, cusipId string) (*models.InstrumentResponse, error) {
	u := fmt.Sprintf("%s/instruments/%s", c.MarketURL, url.PathEscape(cusipId))
	var result models.InstrumentResponse
	err := c.doRequest(ctx, "GET", u, nil, &result)
	return &result, err
}
