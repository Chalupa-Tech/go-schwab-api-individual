package models

// TokenResponse represents the OAuth2 token response.
type TokenResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	ExpiresIn    int    `json:"expires_in"`
	TokenType    string `json:"token_type"`
	Scope        string `json:"scope"`
	IDToken      string `json:"id_token,omitempty"`
}

// AccountNumberHash represents the encrypted account ID.
type AccountNumberHash struct {
	AccountNumber string `json:"accountNumber"`
	HashValue     string `json:"hashValue"`
}

// Account represents a user account.
type Account struct {
	SecuritiesAccount SecuritiesAccount `json:"securitiesAccount"`
}

type SecuritiesAccount struct {
	Type                    string           `json:"type"` // "CASH" or "MARGIN"
	AccountNumber           string           `json:"accountNumber"`
	RoundTrips              int              `json:"roundTrips"`
	IsDayTrader             bool             `json:"isDayTrader"`
	IsClosingOnlyRestricted bool             `json:"isClosingOnlyRestricted"`
	PfcbFlag                bool             `json:"pfcbFlag"`
	Positions               []Position       `json:"positions"`
	InitialBalances         *InitialBalance  `json:"initialBalances,omitempty"`
	CurrentBalances         *CurrentBalance  `json:"currentBalances,omitempty"`
	ProjectedBalances       *ProjectedBalance `json:"projectedBalances,omitempty"`
}

type InitialBalance struct {
	AccruedInterest                 float64 `json:"accruedInterest"`
	CashAvailableForTrading         float64 `json:"cashAvailableForTrading"`
	CashAvailableForWithdrawal      float64 `json:"cashAvailableForWithdrawal"`
	CashBalance                     float64 `json:"cashBalance"`
	BondValue                       float64 `json:"bondValue"`
	CashReceipts                    float64 `json:"cashReceipts"`
	LiquidationValue                float64 `json:"liquidationValue"`
	LongOptionMarketValue           float64 `json:"longOptionMarketValue"`
	LongStockValue                  float64 `json:"longStockValue"`
	MoneyMarketFund                 float64 `json:"moneyMarketFund"`
	MutualFundValue                 float64 `json:"mutualFundValue"`
	ShortOptionMarketValue          float64 `json:"shortOptionMarketValue"`
	ShortStockValue                 float64 `json:"shortStockValue"`
	IsInCall                        float64 `json:"isInCall"`
	UnsettledCash                   float64 `json:"unsettledCash"`
	CashDebitCallValue              float64 `json:"cashDebitCallValue"`
	PendingDeposits                 float64 `json:"pendingDeposits"`
	AccountValue                    float64 `json:"accountValue"`
	// Margin specific
	AvailableFundsNonMarginableTrade float64 `json:"availableFundsNonMarginableTrade"`
	BuyingPower                      float64 `json:"buyingPower"`
	DayTradingBuyingPower            float64 `json:"dayTradingBuyingPower"`
	DayTradingBuyingPowerCall        float64 `json:"dayTradingBuyingPowerCall"`
	DayTradingEquityCall             float64 `json:"dayTradingEquityCall"`
	Equity                           float64 `json:"equity"`
	EquityPercentage                 float64 `json:"equityPercentage"`
	LongMarginValue                  float64 `json:"longMarginValue"`
	MaintenanceCall                  float64 `json:"maintenanceCall"`
	MaintenanceRequirement           float64 `json:"maintenanceRequirement"`
	Margin                           float64 `json:"margin"`
	MarginEquity                     float64 `json:"marginEquity"`
	RegTCall                         float64 `json:"regTCall"`
	ShortMarginValue                 float64 `json:"shortMarginValue"`
	TotalCash                        float64 `json:"totalCash"`
	MarginBalance                    float64 `json:"marginBalance"`
	ShortBalance                     float64 `json:"shortBalance"`
}

type CurrentBalance struct {
	CashAvailableForTrading         float64 `json:"cashAvailableForTrading"`
	CashAvailableForWithdrawal      float64 `json:"cashAvailableForWithdrawal"`
	CashCall                        float64 `json:"cashCall"`
	LongNonMarginableMarketValue    float64 `json:"longNonMarginableMarketValue"`
	TotalCash                       float64 `json:"totalCash"`
	CashDebitCallValue              float64 `json:"cashDebitCallValue"`
	UnsettledCash                   float64 `json:"unsettledCash"`
	// Margin specific
	AvailableFunds                   float64 `json:"availableFunds"`
	AvailableFundsNonMarginableTrade float64 `json:"availableFundsNonMarginableTrade"`
	BuyingPower                      float64 `json:"buyingPower"`
	BuyingPowerNonMarginableTrade    float64 `json:"buyingPowerNonMarginableTrade"`
	DayTradingBuyingPower            float64 `json:"dayTradingBuyingPower"`
	DayTradingBuyingPowerCall        float64 `json:"dayTradingBuyingPowerCall"`
	Equity                           float64 `json:"equity"`
	EquityPercentage                 float64 `json:"equityPercentage"`
	LongMarginValue                  float64 `json:"longMarginValue"`
	MaintenanceCall                  float64 `json:"maintenanceCall"`
	MaintenanceRequirement           float64 `json:"maintenanceRequirement"`
	MarginBalance                    float64 `json:"marginBalance"`
	RegTCall                         float64 `json:"regTCall"`
	ShortBalance                     float64 `json:"shortBalance"`
	ShortMarginValue                 float64 `json:"shortMarginValue"`
	Sma                              float64 `json:"sma"`
	IsInCall                         float64 `json:"isInCall"`
	StockBuyingPower                 float64 `json:"stockBuyingPower"`
	OptionBuyingPower                float64 `json:"optionBuyingPower"`
}

type ProjectedBalance struct {
	// Same as CurrentBalance generally
	CurrentBalance
}

type Position struct {
	ShortQuantity                  float64 `json:"shortQuantity"`
	AveragePrice                   float64 `json:"averagePrice"`
	CurrentDayProfitLoss           float64 `json:"currentDayProfitLoss"`
	CurrentDayProfitLossPercentage float64 `json:"currentDayProfitLossPercentage"`
	LongQuantity                   float64 `json:"longQuantity"`
	SettledLongQuantity            float64 `json:"settledLongQuantity"`
	SettledShortQuantity           float64 `json:"settledShortQuantity"`
	AgedQuantity                   float64 `json:"agedQuantity"`
	Instrument                     Instrument `json:"instrument"`
	MarketValue                    float64 `json:"marketValue"`
	MaintenanceRequirement         float64 `json:"maintenanceRequirement"`
	AverageLongPrice               float64 `json:"averageLongPrice"`
	AverageShortPrice              float64 `json:"averageShortPrice"`
	TaxLotAverageLongPrice         float64 `json:"taxLotAverageLongPrice"`
	TaxLotAverageShortPrice        float64 `json:"taxLotAverageShortPrice"`
	LongOpenProfitLoss             float64 `json:"longOpenProfitLoss"`
	ShortOpenProfitLoss            float64 `json:"shortOpenProfitLoss"`
	PreviousSessionLongQuantity    float64 `json:"previousSessionLongQuantity"`
	PreviousSessionShortQuantity   float64 `json:"previousSessionShortQuantity"`
	CurrentDayCost                 float64 `json:"currentDayCost"`
}

type Instrument struct {
	AssetType   string `json:"assetType"`
	Cusip       string `json:"cusip"`
	Symbol      string `json:"symbol"`
	Description string `json:"description"`
	InstrumentId int64  `json:"instrumentId"`
	NetChange   float64 `json:"netChange"`
	Type        string  `json:"type"` // "SWEEP_VEHICLE", "SAVINGS", etc.
}

type Order struct {
	Session                  string           `json:"session"`
	Duration                 string           `json:"duration"`
	OrderType                string           `json:"orderType"`
	CancelTime               string           `json:"cancelTime"`
	ComplexOrderStrategyType string           `json:"complexOrderStrategyType"`
	Quantity                 float64          `json:"quantity"`
	FilledQuantity           float64          `json:"filledQuantity"`
	RemainingQuantity        float64          `json:"remainingQuantity"`
	RequestedDestination     string           `json:"requestedDestination"`
	DestinationLinkName      string           `json:"destinationLinkName"`
	ReleaseTime              string           `json:"releaseTime"`
	StopPrice                float64          `json:"stopPrice"`
	StopPriceLinkBasis       string           `json:"stopPriceLinkBasis"`
	StopPriceLinkType        string           `json:"stopPriceLinkType"`
	StopPriceOffset          float64          `json:"stopPriceOffset"`
	StopType                 string           `json:"stopType"`
	PriceLinkBasis           string           `json:"priceLinkBasis"`
	PriceLinkType            string           `json:"priceLinkType"`
	Price                    float64          `json:"price"`
	TaxLotMethod             string           `json:"taxLotMethod"`
	OrderLegCollection       []OrderLeg       `json:"orderLegCollection"`
	ActivationPrice          float64          `json:"activationPrice"`
	SpecialInstruction       string           `json:"specialInstruction"`
	OrderStrategyType        string           `json:"orderStrategyType"`
	OrderId                  int64            `json:"orderId"`
	Cancelable               bool             `json:"cancelable"`
	Editable                 bool             `json:"editable"`
	Status                   string           `json:"status"`
	EnteredTime              string           `json:"enteredTime"`
	CloseTime                string           `json:"closeTime"`
	Tag                      string           `json:"tag"`
	AccountNumber            int64            `json:"accountNumber"`
	OrderActivityCollection  []OrderActivity  `json:"orderActivityCollection"`
	ReplacingOrderCollection []Order          `json:"replacingOrderCollection"`
	ChildOrderStrategies     []Order          `json:"childOrderStrategies"`
	StatusDescription        string           `json:"statusDescription"`
}

type OrderLeg struct {
	OrderLegType   string     `json:"orderLegType"`
	LegId          int64      `json:"legId"`
	Instrument     Instrument `json:"instrument"`
	Instruction    string     `json:"instruction"`
	PositionEffect string     `json:"positionEffect"`
	Quantity       float64    `json:"quantity"`
	QuantityType   string     `json:"quantityType"`
	DivCapGains    string     `json:"divCapGains"`
	ToSymbol       string     `json:"toSymbol"`
}

type OrderActivity struct {
	ActivityType           string         `json:"activityType"`
	ExecutionType          string         `json:"executionType"`
	Quantity               float64        `json:"quantity"`
	OrderRemainingQuantity float64        `json:"orderRemainingQuantity"`
	ExecutionLegs          []ExecutionLeg `json:"executionLegs"`
}

type ExecutionLeg struct {
	LegId             int64   `json:"legId"`
	Price             float64 `json:"price"`
	Quantity          float64 `json:"quantity"`
	MismarkedQuantity float64 `json:"mismarkedQuantity"`
	InstrumentId      int64   `json:"instrumentId"`
	Time              string  `json:"time"`
}

type Transaction struct {
	ActivityId     int64          `json:"activityId"`
	Time           string         `json:"time"`
	User           UserDetails    `json:"user"`
	Description    string         `json:"description"`
	AccountNumber  string         `json:"accountNumber"`
	Type           string         `json:"type"` // TransactionType
	Status         string         `json:"status"`
	SubAccount     string         `json:"subAccount"`
	TradeDate      string         `json:"tradeDate"`
	SettlementDate string         `json:"settlementDate"`
	PositionId     int64          `json:"positionId"`
	OrderId        int64          `json:"orderId"`
	NetAmount      float64        `json:"netAmount"`
	ActivityType   string         `json:"activityType"`
	TransferItems  []TransferItem `json:"transferItems"`
}

type UserDetails struct {
	CdDomainId     string `json:"cdDomainId"`
	Login          string `json:"login"`
	Type           string `json:"type"`
	UserId         int64  `json:"userId"`
	SystemUserName string `json:"systemUserName"`
	FirstName      string `json:"firstName"`
	LastName       string `json:"lastName"`
	BrokerRepCode  string `json:"brokerRepCode"`
}

type TransferItem struct {
	Instrument     Instrument `json:"instrument"`
	Amount         float64    `json:"amount"`
	Cost           float64    `json:"cost"`
	Price          float64    `json:"price"`
	FeeType        string     `json:"feeType"`
	PositionEffect string     `json:"positionEffect"`
}

// Market Data

type QuoteResponse map[string]QuoteResponseObject

type QuoteResponseObject struct {
	AssetMainType string          `json:"assetMainType"`
	AssetSubType  string          `json:"assetSubType"`
	Ssid          int64           `json:"ssid"`
	Symbol        string          `json:"symbol"`
	Realtime      bool            `json:"realtime"`
	QuoteType     string          `json:"quoteType"`
	Extended      *ExtendedMarket `json:"extended,omitempty"`
	Fundamental   *Fundamental    `json:"fundamental,omitempty"`
	Quote         *QuoteUnion     `json:"quote,omitempty"`
	Reference     *ReferenceUnion `json:"reference,omitempty"`
	Regular       *RegularMarket  `json:"regular,omitempty"`
}

type QuoteUnion struct {
	// Common fields merged for simplicity
	AskPrice           float64 `json:"askPrice"`
	AskSize            int     `json:"askSize"`
	BidPrice           float64 `json:"bidPrice"`
	BidSize            int     `json:"bidSize"`
	LastPrice          float64 `json:"lastPrice"`
	LastSize           int     `json:"lastSize"`
	OpenPrice          float64 `json:"openPrice"`
	HighPrice          float64 `json:"highPrice"`
	LowPrice           float64 `json:"lowPrice"`
	ClosePrice         float64 `json:"closePrice"`
	NetChange          float64 `json:"netChange"`
	TotalVolume        int64   `json:"totalVolume"`
	QuoteTime          int64   `json:"quoteTime"`
	TradeTime          int64   `json:"tradeTime"`
	Mark               float64 `json:"mark"`
	MarkChange         float64 `json:"markChange"`
	MarkPercentChange  float64 `json:"markPercentChange"`
	NetPercentChange   float64 `json:"netPercentChange"`
	SecurityStatus     string  `json:"securityStatus"`
	Volatility         float64 `json:"volatility"`
	FiftyTwoWeekHigh   float64 `json:"52WeekHigh"`
	FiftyTwoWeekLow    float64 `json:"52WeekLow"`
	NAV                float64 `json:"nAV"`
	Delta              float64 `json:"delta"`
	Gamma              float64 `json:"gamma"`
	Theta              float64 `json:"theta"`
	Vega               float64 `json:"vega"`
	Rho                float64 `json:"rho"`
	OpenInterest       float64 `json:"openInterest"`
	MoneyIntrinsicValue float64 `json:"moneyIntrinsicValue"`
	UnderlyingPrice    float64 `json:"underlyingPrice"`
	StrikePrice        float64 `json:"strikePrice"` // FutureOption
	TimeValue          float64 `json:"timeValue"`
}

type ReferenceUnion struct {
	Cusip              string  `json:"cusip"`
	Description        string  `json:"description"`
	Exchange           string  `json:"exchange"`
	ExchangeName       string  `json:"exchangeName"`
	IsTradable         bool    `json:"isTradable"`
	IsShortable        bool    `json:"isShortable"`
	IsHardToBorrow     bool    `json:"isHardToBorrow"`
	HtbQuantity        int     `json:"htbQuantity"`
	HtbRate            float64 `json:"htbRate"`
	ContractType       string  `json:"contractType"` // P or C
	Underlying         string  `json:"underlying"`
	ExpirationDate     int64   `json:"expirationDate"`
	DaysToExpiration   int     `json:"daysToExpiration"`
	Multiplier         float64 `json:"multiplier"`
}

type ExtendedMarket struct {
	AskPrice    float64 `json:"askPrice"`
	AskSize     int     `json:"askSize"`
	BidPrice    float64 `json:"bidPrice"`
	BidSize     int     `json:"bidSize"`
	LastPrice   float64 `json:"lastPrice"`
	LastSize    int     `json:"lastSize"`
	Mark        float64 `json:"mark"`
	QuoteTime   int64   `json:"quoteTime"`
	TotalVolume int64   `json:"totalVolume"`
	TradeTime   int64   `json:"tradeTime"`
}

type Fundamental struct {
	Avg10DaysVolume     float64 `json:"avg10DaysVolume"`
	Avg1YearVolume      float64 `json:"avg1YearVolume"`
	DivAmount           float64 `json:"divAmount"`
	DivYield            float64 `json:"divYield"`
	Eps                 float64 `json:"eps"`
	FundLeverageFactor  float64 `json:"fundLeverageFactor"`
	PeRatio             float64 `json:"peRatio"`
	// Add more if needed
}

type RegularMarket struct {
	RegularMarketLastPrice     float64 `json:"regularMarketLastPrice"`
	RegularMarketLastSize      int     `json:"regularMarketLastSize"`
	RegularMarketNetChange     float64 `json:"regularMarketNetChange"`
	RegularMarketPercentChange float64 `json:"regularMarketPercentChange"`
	RegularMarketTradeTime     int64   `json:"regularMarketTradeTime"`
}

type OptionChain struct {
	Symbol           string                 `json:"symbol"`
	Status           string                 `json:"status"`
	Underlying       *Underlying            `json:"underlying"`
	Strategy         string                 `json:"strategy"`
	Interval         float64                `json:"interval"`
	IsDelayed        bool                   `json:"isDelayed"`
	IsIndex          bool                   `json:"isIndex"`
	DaysToExpiration float64                `json:"daysToExpiration"`
	InterestRate     float64                `json:"interestRate"`
	UnderlyingPrice  float64                `json:"underlyingPrice"`
	Volatility       float64                `json:"volatility"`
	CallExpDateMap   map[string]OptionMap   `json:"callExpDateMap"`
	PutExpDateMap    map[string]OptionMap   `json:"putExpDateMap"`
}

type OptionMap map[string][]OptionContract

type OptionContract struct {
	PutCall           string  `json:"putCall"`
	Symbol            string  `json:"symbol"`
	Description       string  `json:"description"`
	ExchangeName      string  `json:"exchangeName"`
	BidPrice          float64 `json:"bidPrice"`
	AskPrice          float64 `json:"askPrice"`
	LastPrice         float64 `json:"lastPrice"`
	MarkPrice         float64 `json:"markPrice"`
	BidSize           int     `json:"bidSize"`
	AskSize           int     `json:"askSize"`
	LastSize          int     `json:"lastSize"`
	HighPrice         float64 `json:"highPrice"`
	LowPrice          float64 `json:"lowPrice"`
	OpenPrice         float64 `json:"openPrice"`
	ClosePrice        float64 `json:"closePrice"`
	TotalVolume       int     `json:"totalVolume"`
	TradeDate         int64   `json:"tradeDate"` // Note: spec says int, sometimes long
	NetChange         float64 `json:"netChange"`
	Volatility        float64 `json:"volatility"`
	Delta             float64 `json:"delta"`
	Gamma             float64 `json:"gamma"`
	Theta             float64 `json:"theta"`
	Vega              float64 `json:"vega"`
	Rho               float64 `json:"rho"`
	OpenInterest      float64 `json:"openInterest"`
	TimeValue         float64 `json:"timeValue"`
	TheoreticalOptionValue float64 `json:"theoreticalOptionValue"`
	StrikePrice       float64 `json:"strikePrice"`
	ExpirationDate    string  `json:"expirationDate"`
	DaysToExpiration  int     `json:"daysToExpiration"`
	Multiplier        float64 `json:"multiplier"`
	SettlementType    string  `json:"settlementType"`
	DeliverableNote   string  `json:"deliverableNote"`
	IsIndexOption     bool    `json:"isIndexOption"`
	PercentChange     float64 `json:"percentChange"`
	MarkChange        float64 `json:"markChange"`
	MarkPercentChange float64 `json:"markPercentChange"`
	IntrinsicValue    float64 `json:"intrinsicValue"`
	OptionRoot        string  `json:"optionRoot"`
}

type Underlying struct {
	Symbol      string  `json:"symbol"`
	Description string  `json:"description"`
	Change      float64 `json:"change"`
	PercentChange float64 `json:"percentChange"`
	Close       float64 `json:"close"`
	QuoteTime   int64   `json:"quoteTime"`
	TradeTime   int64   `json:"tradeTime"`
	Bid         float64 `json:"bid"`
	Ask         float64 `json:"ask"`
	Last        float64 `json:"last"`
	Mark        float64 `json:"mark"`
	TotalVolume int64   `json:"totalVolume"`
	HighPrice   float64 `json:"highPrice"`
	LowPrice    float64 `json:"lowPrice"`
	OpenPrice   float64 `json:"openPrice"`
}

type CandleList struct {
	Candles []Candle `json:"candles"`
	Empty   bool     `json:"empty"`
	Symbol  string   `json:"symbol"`
}

type Candle struct {
	Close    float64 `json:"close"`
	Datetime int64   `json:"datetime"`
	High     float64 `json:"high"`
	Low      float64 `json:"low"`
	Open     float64 `json:"open"`
	Volume   int64   `json:"volume"`
}

type MoverResponse struct {
	Screeners []Screener `json:"screeners"`
}

type Screener struct {
	Change      float64 `json:"change"`
	Description string  `json:"description"`
	Direction   string  `json:"direction"`
	Last        float64 `json:"last"`
	Symbol      string  `json:"symbol"`
	TotalVolume int64   `json:"totalVolume"`
}

type MarketHoursResponse map[string]map[string]Hours

type Hours struct {
	Date         string                `json:"date"`
	MarketType   string                `json:"marketType"`
	Exchange     string                `json:"exchange"`
	Category     string                `json:"category"`
	Product      string                `json:"product"`
	ProductName  string                `json:"productName"`
	IsOpen       bool                  `json:"isOpen"`
	SessionHours map[string][]Interval `json:"sessionHours"`
}

type Interval struct {
	Start string `json:"start"`
	End   string `json:"end"`
}

type InstrumentListResponse struct {
	Instruments []InstrumentResponse `json:"instruments"`
}

type InstrumentResponse struct {
	Cusip       string `json:"cusip"`
	Symbol      string `json:"symbol"`
	Description string `json:"description"`
	Exchange    string `json:"exchange"`
	AssetType   string `json:"assetType"`
	Type        string `json:"type"`
}
