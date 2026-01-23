package schwab

import (
	"context"
	"fmt"
	"net/url"
	"strconv"
	"strings"

	"github.com/Chalupa-Tech/go-schwab-api-individual/schwab/models"
)

// GetQuotes returns quotes for a list of symbols.
func (c *Client) GetQuotes(ctx context.Context, symbols []string, fields string, indicative bool) (models.QuoteResponse, error) {
	u, _ := url.Parse(MarketDataURL + "/quotes")
	q := u.Query()
	q.Set("symbols", strings.Join(symbols, ","))
	if fields != "" {
		q.Set("fields", fields)
	}
	q.Set("indicative", strconv.FormatBool(indicative))
	u.RawQuery = q.Encode()

	var result models.QuoteResponse
	err := c.do(ctx, "GET", u.String(), nil, &result)
	return result, err
}

// GetOptionChainParams represents parameters for GetOptionChain.
type GetOptionChainParams struct {
	Symbol                 string
	ContractType           string // CALL, PUT, ALL
	StrikeCount            int
	IncludeUnderlyingQuote bool
	Strategy               string // SINGLE, ANALYTICAL, etc.
	Interval               float64
	Strike                 float64
	Range                  string // ITM, NTM, OTM
	FromDate               string // yyyy-MM-dd
	ToDate                 string // yyyy-MM-dd
	Volatility             float64
	UnderlyingPrice        float64
	InterestRate           float64
	DaysToExpiration       int
	ExpMonth               string // JAN, FEB, etc.
	OptionType             string
}

// GetOptionChain returns option chain for a symbol.
func (c *Client) GetOptionChain(ctx context.Context, params GetOptionChainParams) (*models.OptionChain, error) {
	u, _ := url.Parse(MarketDataURL + "/chains")
	q := u.Query()
	q.Set("symbol", params.Symbol)
	if params.ContractType != "" {
		q.Set("contractType", params.ContractType)
	}
	if params.StrikeCount > 0 {
		q.Set("strikeCount", strconv.Itoa(params.StrikeCount))
	}
	q.Set("includeUnderlyingQuote", strconv.FormatBool(params.IncludeUnderlyingQuote))
	if params.Strategy != "" {
		q.Set("strategy", params.Strategy)
	}
	if params.Interval > 0 {
		q.Set("interval", fmt.Sprintf("%f", params.Interval))
	}
	if params.Strike > 0 {
		q.Set("strike", fmt.Sprintf("%f", params.Strike))
	}
	if params.Range != "" {
		q.Set("range", params.Range)
	}
	if params.FromDate != "" {
		q.Set("fromDate", params.FromDate)
	}
	if params.ToDate != "" {
		q.Set("toDate", params.ToDate)
	}
	if params.Volatility > 0 {
		q.Set("volatility", fmt.Sprintf("%f", params.Volatility))
	}
	if params.UnderlyingPrice > 0 {
		q.Set("underlyingPrice", fmt.Sprintf("%f", params.UnderlyingPrice))
	}
	if params.InterestRate > 0 {
		q.Set("interestRate", fmt.Sprintf("%f", params.InterestRate))
	}
	if params.DaysToExpiration > 0 {
		q.Set("daysToExpiration", strconv.Itoa(params.DaysToExpiration))
	}
	if params.ExpMonth != "" {
		q.Set("expMonth", params.ExpMonth)
	}
	if params.OptionType != "" {
		q.Set("optionType", params.OptionType)
	}
	u.RawQuery = q.Encode()

	var result models.OptionChain
	err := c.do(ctx, "GET", u.String(), nil, &result)
	return &result, err
}

// GetPriceHistoryParams represents parameters for GetPriceHistory.
type GetPriceHistoryParams struct {
	Symbol                string
	PeriodType            string // day, month, year, ytd
	Period                int
	FrequencyType         string // minute, daily, weekly, monthly
	Frequency             int
	StartDate             int64 // Epoch ms
	EndDate               int64 // Epoch ms
	NeedExtendedHoursData bool
	NeedPreviousClose     bool
}

// GetPriceHistory returns price history for a symbol.
func (c *Client) GetPriceHistory(ctx context.Context, params GetPriceHistoryParams) (*models.CandleList, error) {
	u, _ := url.Parse(MarketDataURL + "/pricehistory")
	q := u.Query()
	q.Set("symbol", params.Symbol)
	if params.PeriodType != "" {
		q.Set("periodType", params.PeriodType)
	}
	if params.Period > 0 {
		q.Set("period", strconv.Itoa(params.Period))
	}
	if params.FrequencyType != "" {
		q.Set("frequencyType", params.FrequencyType)
	}
	if params.Frequency > 0 {
		q.Set("frequency", strconv.Itoa(params.Frequency))
	}
	if params.StartDate > 0 {
		q.Set("startDate", strconv.FormatInt(params.StartDate, 10))
	}
	if params.EndDate > 0 {
		q.Set("endDate", strconv.FormatInt(params.EndDate, 10))
	}
	q.Set("needExtendedHoursData", strconv.FormatBool(params.NeedExtendedHoursData))
	q.Set("needPreviousClose", strconv.FormatBool(params.NeedPreviousClose))
	u.RawQuery = q.Encode()

	var result models.CandleList
	err := c.do(ctx, "GET", u.String(), nil, &result)
	return &result, err
}

// GetMovers returns top 10 securities movement for a specific index.
// symbol can be "$DJI", "$COMPX", "$SPX".
// sort can be "VOLUME", "TRADES", "PERCENT_CHANGE_UP", "PERCENT_CHANGE_DOWN".
// frequency can be 0, 1, 5, 10, 30, 60.
func (c *Client) GetMovers(ctx context.Context, symbol string, sort string, frequency int) (*models.MoverResponse, error) {
	u, _ := url.Parse(MarketDataURL + "/movers/" + symbol)
	q := u.Query()
	if sort != "" {
		q.Set("sort", sort)
	}
	if frequency > 0 {
		q.Set("frequency", strconv.Itoa(frequency))
	}
	u.RawQuery = q.Encode()

	var result models.MoverResponse
	err := c.do(ctx, "GET", u.String(), nil, &result)
	return &result, err
}

// GetMarketHours returns market hours for different markets.
// markets: comma separated list of markets (EQUITY, OPTION, BOND, FOREX, FUTURE).
// date: yyyy-MM-dd.
func (c *Client) GetMarketHours(ctx context.Context, markets string, date string) (models.MarketHoursResponse, error) {
	u, _ := url.Parse(MarketDataURL + "/markets")
	q := u.Query()
	if markets != "" {
		q.Set("markets", markets)
	}
	if date != "" {
		q.Set("date", date)
	}
	u.RawQuery = q.Encode()

	var result models.MarketHoursResponse
	err := c.do(ctx, "GET", u.String(), nil, &result)
	return result, err
}

// GetMarketHour returns market hours for a specific market.
func (c *Client) GetMarketHour(ctx context.Context, market_id string, date string) (models.MarketHoursResponse, error) {
	u, _ := url.Parse(MarketDataURL + "/markets/" + market_id)
	q := u.Query()
	if date != "" {
		q.Set("date", date)
	}
	u.RawQuery = q.Encode()

	var result models.MarketHoursResponse
	err := c.do(ctx, "GET", u.String(), nil, &result)
	return result, err
}

// GetInstruments returns instrument details by symbols and projections.
// projection: "symbol-search", "symbol-regex", "desc-search", "desc-regex", "fundamental".
func (c *Client) GetInstruments(ctx context.Context, symbol string, projection string) (*models.InstrumentListResponse, error) {
	u, _ := url.Parse(MarketDataURL + "/instruments")
	q := u.Query()
	q.Set("symbol", symbol)
	if projection != "" {
		q.Set("projection", projection)
	}
	u.RawQuery = q.Encode()

	var result models.InstrumentListResponse
	err := c.do(ctx, "GET", u.String(), nil, &result)
	return &result, err
}

