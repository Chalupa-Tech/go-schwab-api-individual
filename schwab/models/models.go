package models

// TokenResponse represents the OAuth2 token response.
type TokenResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	TokenType    string `json:"token_type"`
	ExpiresIn    int    `json:"expires_in"`
	Scope        string `json:"scope"`
	IdToken      string `json:"id_token,omitempty"`
}

// --- Account Access Models ---

type AccountNumberHash struct {
	AccountNumber string `json:"accountNumber"`
	HashValue     string `json:"hashValue"`
}

type Account struct {
	SecuritiesAccount SecuritiesAccount `json:"securitiesAccount"`
}

type SecuritiesAccount struct {
	Type                    string            `json:"type"` // "CASH" or "MARGIN"
	AccountNumber           string            `json:"accountNumber"`
	RoundTrips              int               `json:"roundTrips"`
	IsDayTrader             bool              `json:"isDayTrader"`
	IsClosingOnlyRestricted bool              `json:"isClosingOnlyRestricted"`
	PfcbFlag                bool              `json:"pfcbFlag"`
	Positions               []Position        `json:"positions,omitempty"`
	InitialBalances         *InitialBalance   `json:"initialBalances,omitempty"`
	CurrentBalances         *CurrentBalance   `json:"currentBalances,omitempty"`
	ProjectedBalances       *ProjectedBalance `json:"projectedBalances,omitempty"`
}

// Combined InitialBalance for Cash and Margin to simplify
type InitialBalance struct {
	AccruedInterest                  float64 `json:"accruedInterest,omitempty"`
	AvailableFundsNonMarginableTrade float64 `json:"availableFundsNonMarginableTrade,omitempty"`
	BondValue                        float64 `json:"bondValue,omitempty"`
	BuyingPower                      float64 `json:"buyingPower,omitempty"`
	CashBalance                      float64 `json:"cashBalance,omitempty"`
	CashAvailableForTrading          float64 `json:"cashAvailableForTrading,omitempty"`
	CashReceipts                     float64 `json:"cashReceipts,omitempty"`
	DayTradingBuyingPower            float64 `json:"dayTradingBuyingPower,omitempty"`
	DayTradingBuyingPowerCall        float64 `json:"dayTradingBuyingPowerCall,omitempty"`
	DayTradingEquityCall             float64 `json:"dayTradingEquityCall,omitempty"`
	Equity                           float64 `json:"equity,omitempty"`
	EquityPercentage                 float64 `json:"equityPercentage,omitempty"`
	LiquidationValue                 float64 `json:"liquidationValue,omitempty"`
	LongMarginValue                  float64 `json:"longMarginValue,omitempty"`
	LongOptionMarketValue            float64 `json:"longOptionMarketValue,omitempty"`
	LongStockValue                   float64 `json:"longStockValue,omitempty"`
	MaintenanceCall                  float64 `json:"maintenanceCall,omitempty"`
	MaintenanceRequirement           float64 `json:"maintenanceRequirement,omitempty"`
	Margin                           float64 `json:"margin,omitempty"`
	MarginEquity                     float64 `json:"marginEquity,omitempty"`
	MoneyMarketFund                  float64 `json:"moneyMarketFund,omitempty"`
	MutualFundValue                  float64 `json:"mutualFundValue,omitempty"`
	RegTCall                         float64 `json:"regTCall,omitempty"`
	ShortMarginValue                 float64 `json:"shortMarginValue,omitempty"`
	ShortOptionMarketValue           float64 `json:"shortOptionMarketValue,omitempty"`
	ShortStockValue                  float64 `json:"shortStockValue,omitempty"`
	TotalCash                        float64 `json:"totalCash,omitempty"`
	IsInCall                         float64 `json:"isInCall,omitempty"`
	UnsettledCash                    float64 `json:"unsettledCash,omitempty"`
	PendingDeposits                  float64 `json:"pendingDeposits,omitempty"`
	MarginBalance                    float64 `json:"marginBalance,omitempty"`
	ShortBalance                     float64 `json:"shortBalance,omitempty"`
	AccountValue                     float64 `json:"accountValue,omitempty"`
	CashAvailableForWithdrawal       float64 `json:"cashAvailableForWithdrawal,omitempty"`
	CashDebitCallValue               float64 `json:"cashDebitCallValue,omitempty"`
}

// Combined CurrentBalance and ProjectedBalance for Cash and Margin
type CurrentBalance struct {
	AvailableFunds                   float64 `json:"availableFunds,omitempty"`
	AvailableFundsNonMarginableTrade float64 `json:"availableFundsNonMarginableTrade,omitempty"`
	BuyingPower                      float64 `json:"buyingPower,omitempty"`
	BuyingPowerNonMarginableTrade    float64 `json:"buyingPowerNonMarginableTrade,omitempty"`
	DayTradingBuyingPower            float64 `json:"dayTradingBuyingPower,omitempty"`
	DayTradingBuyingPowerCall        float64 `json:"dayTradingBuyingPowerCall,omitempty"`
	Equity                           float64 `json:"equity,omitempty"`
	EquityPercentage                 float64 `json:"equityPercentage,omitempty"`
	LongMarginValue                  float64 `json:"longMarginValue,omitempty"`
	MaintenanceCall                  float64 `json:"maintenanceCall,omitempty"`
	MaintenanceRequirement           float64 `json:"maintenanceRequirement,omitempty"`
	MarginBalance                    float64 `json:"marginBalance,omitempty"`
	RegTCall                         float64 `json:"regTCall,omitempty"`
	ShortBalance                     float64 `json:"shortBalance,omitempty"`
	ShortMarginValue                 float64 `json:"shortMarginValue,omitempty"`
	Sma                              float64 `json:"sma,omitempty"`
	IsInCall                         float64 `json:"isInCall,omitempty"`
	StockBuyingPower                 float64 `json:"stockBuyingPower,omitempty"`
	OptionBuyingPower                float64 `json:"optionBuyingPower,omitempty"`
	CashAvailableForTrading          float64 `json:"cashAvailableForTrading,omitempty"`
	CashAvailableForWithdrawal       float64 `json:"cashAvailableForWithdrawal,omitempty"`
	CashCall                         float64 `json:"cashCall,omitempty"`
	LongNonMarginableMarketValue     float64 `json:"longNonMarginableMarketValue,omitempty"`
	TotalCash                        float64 `json:"totalCash,omitempty"`
	CashDebitCallValue               float64 `json:"cashDebitCallValue,omitempty"`
	UnsettledCash                    float64 `json:"unsettledCash,omitempty"`
}

type ProjectedBalance CurrentBalance

type Position struct {
	ShortQuantity                  float64             `json:"shortQuantity"`
	AveragePrice                   float64             `json:"averagePrice"`
	CurrentDayProfitLoss           float64             `json:"currentDayProfitLoss"`
	CurrentDayProfitLossPercentage float64             `json:"currentDayProfitLossPercentage"`
	LongQuantity                   float64             `json:"longQuantity"`
	SettledLongQuantity            float64             `json:"settledLongQuantity"`
	SettledShortQuantity           float64             `json:"settledShortQuantity"`
	AgedQuantity                   float64             `json:"agedQuantity"`
	Instrument                     *AccountsInstrument `json:"instrument"`
	MarketValue                    float64             `json:"marketValue"`
	MaintenanceRequirement         float64             `json:"maintenanceRequirement"`
	AverageLongPrice               float64             `json:"averageLongPrice"`
	AverageShortPrice              float64             `json:"averageShortPrice"`
	TaxLotAverageLongPrice         float64             `json:"taxLotAverageLongPrice"`
	TaxLotAverageShortPrice        float64             `json:"taxLotAverageShortPrice"`
	LongOpenProfitLoss             float64             `json:"longOpenProfitLoss"`
	ShortOpenProfitLoss            float64             `json:"shortOpenProfitLoss"`
	PreviousSessionLongQuantity    float64             `json:"previousSessionLongQuantity"`
	PreviousSessionShortQuantity   float64             `json:"previousSessionShortQuantity"`
	CurrentDayCost                 float64             `json:"currentDayCost"`
}

type AccountsInstrument struct {
	AssetType        string  `json:"assetType"`
	Cusip            string  `json:"cusip"`
	Symbol           string  `json:"symbol"`
	Description      string  `json:"description"`
	InstrumentId     int64   `json:"instrumentId"`
	NetChange        float64 `json:"netChange"`
	Type             string  `json:"type,omitempty"`             // For specific types like CASH_EQUIVALENT, MUTUAL_FUND etc.
	MaturityDate     string  `json:"maturityDate,omitempty"`     // Fixed Income
	Factor           float64 `json:"factor,omitempty"`           // Fixed Income
	VariableRate     float64 `json:"variableRate,omitempty"`     // Fixed Income
	PutCall          string  `json:"putCall,omitempty"`          // Option
	OptionMultiplier int32   `json:"optionMultiplier,omitempty"` // Option
	UnderlyingSymbol string  `json:"underlyingSymbol,omitempty"` // Option
}

type Order struct {
	Session                  string          `json:"session"`
	Duration                 string          `json:"duration"`
	OrderType                string          `json:"orderType"`
	CancelTime               string          `json:"cancelTime,omitempty"`
	ComplexOrderStrategyType string          `json:"complexOrderStrategyType,omitempty"`
	Quantity                 float64         `json:"quantity"`
	FilledQuantity           float64         `json:"filledQuantity"`
	RemainingQuantity        float64         `json:"remainingQuantity"`
	RequestedDestination     string          `json:"requestedDestination,omitempty"`
	DestinationLinkName      string          `json:"destinationLinkName,omitempty"`
	ReleaseTime              string          `json:"releaseTime,omitempty"`
	StopPrice                float64         `json:"stopPrice,omitempty"`
	StopPriceLinkBasis       string          `json:"stopPriceLinkBasis,omitempty"`
	StopPriceLinkType        string          `json:"stopPriceLinkType,omitempty"`
	StopPriceOffset          float64         `json:"stopPriceOffset,omitempty"`
	StopType                 string          `json:"stopType,omitempty"`
	PriceLinkBasis           string          `json:"priceLinkBasis,omitempty"`
	PriceLinkType            string          `json:"priceLinkType,omitempty"`
	Price                    float64         `json:"price,omitempty"`
	TaxLotMethod             string          `json:"taxLotMethod,omitempty"`
	OrderLegCollection       []OrderLeg      `json:"orderLegCollection,omitempty"`
	ActivationPrice          float64         `json:"activationPrice,omitempty"`
	SpecialInstruction       string          `json:"specialInstruction,omitempty"`
	OrderStrategyType        string          `json:"orderStrategyType"`
	OrderId                  int64           `json:"orderId,omitempty"`
	Cancelable               bool            `json:"cancelable"`
	Editable                 bool            `json:"editable"`
	Status                   string          `json:"status,omitempty"`
	EnteredTime              string          `json:"enteredTime,omitempty"`
	CloseTime                string          `json:"closeTime,omitempty"`
	Tag                      string          `json:"tag,omitempty"`
	AccountNumber            string          `json:"accountNumber,omitempty"`
	OrderActivityCollection  []OrderActivity `json:"orderActivityCollection,omitempty"`
	ReplacingOrderCollection []Order         `json:"replacingOrderCollection,omitempty"`
	ChildOrderStrategies     []Order         `json:"childOrderStrategies,omitempty"`
	StatusDescription        string          `json:"statusDescription,omitempty"`
}

type OrderRequest struct {
	Session                  string          `json:"session"`
	Duration                 string          `json:"duration"`
	OrderType                string          `json:"orderType"`
	CancelTime               string          `json:"cancelTime,omitempty"`
	ComplexOrderStrategyType string          `json:"complexOrderStrategyType,omitempty"`
	Quantity                 float64         `json:"quantity,omitempty"`
	FilledQuantity           float64         `json:"filledQuantity,omitempty"`
	RemainingQuantity        float64         `json:"remainingQuantity,omitempty"`
	DestinationLinkName      string          `json:"destinationLinkName,omitempty"`
	ReleaseTime              string          `json:"releaseTime,omitempty"`
	StopPrice                float64         `json:"stopPrice,omitempty"`
	StopPriceLinkBasis       string          `json:"stopPriceLinkBasis,omitempty"`
	StopPriceLinkType        string          `json:"stopPriceLinkType,omitempty"`
	StopPriceOffset          float64         `json:"stopPriceOffset,omitempty"`
	StopType                 string          `json:"stopType,omitempty"`
	PriceLinkBasis           string          `json:"priceLinkBasis,omitempty"`
	PriceLinkType            string          `json:"priceLinkType,omitempty"`
	Price                    float64         `json:"price,omitempty"`
	TaxLotMethod             string          `json:"taxLotMethod,omitempty"`
	OrderLegCollection       []OrderLeg      `json:"orderLegCollection,omitempty"`
	ActivationPrice          float64         `json:"activationPrice,omitempty"`
	SpecialInstruction       string          `json:"specialInstruction,omitempty"`
	OrderStrategyType        string          `json:"orderStrategyType"`
	OrderId                  int64           `json:"orderId,omitempty"`
	Cancelable               bool            `json:"cancelable,omitempty"`
	Editable                 bool            `json:"editable,omitempty"`
	Status                   string          `json:"status,omitempty"`
	EnteredTime              string          `json:"enteredTime,omitempty"`
	CloseTime                string          `json:"closeTime,omitempty"`
	AccountNumber            string          `json:"accountNumber,omitempty"`
	OrderActivityCollection  []OrderActivity `json:"orderActivityCollection,omitempty"`
	ReplacingOrderCollection []OrderRequest  `json:"replacingOrderCollection,omitempty"`
	ChildOrderStrategies     []OrderRequest  `json:"childOrderStrategies,omitempty"`
	StatusDescription        string          `json:"statusDescription,omitempty"`
}

type OrderLeg struct {
	OrderLegType   string              `json:"orderLegType,omitempty"`
	LegId          int64               `json:"legId,omitempty"`
	Instrument     *AccountsInstrument `json:"instrument,omitempty"`
	Instruction    string              `json:"instruction"`
	PositionEffect string              `json:"positionEffect,omitempty"`
	Quantity       float64             `json:"quantity"`
	QuantityType   string              `json:"quantityType,omitempty"`
	DivCapGains    string              `json:"divCapGains,omitempty"`
	ToSymbol       string              `json:"toSymbol,omitempty"`
}

type OrderActivity struct {
	ActivityType           string         `json:"activityType"` // EXECUTION, ORDER_ACTION
	ExecutionType          string         `json:"executionType,omitempty"`
	Quantity               float64        `json:"quantity,omitempty"`
	OrderRemainingQuantity float64        `json:"orderRemainingQuantity,omitempty"`
	ExecutionLegs          []ExecutionLeg `json:"executionLegs,omitempty"`
}

type ExecutionLeg struct {
	LegId             int64   `json:"legId"`
	Price             float64 `json:"price"`
	Quantity          float64 `json:"quantity"`
	MismarkedQuantity float64 `json:"mismarkedQuantity,omitempty"`
	InstrumentId      int64   `json:"instrumentId"`
}

type PreviewOrder struct {
	OrderId               int64                  `json:"orderId,omitempty"`
	OrderStrategy         Order                  `json:"orderStrategy"`
	OrderValidationResult *OrderValidationResult `json:"orderValidationResult,omitempty"`
	CommissionAndFee      *CommissionAndFee      `json:"commissionAndFee,omitempty"`
}

type OrderValidationResult struct {
	Alerts  []OrderValidationDetail `json:"alerts,omitempty"`
	Accepts []OrderValidationDetail `json:"accepts,omitempty"`
	Rejects []OrderValidationDetail `json:"rejects,omitempty"`
	Reviews []OrderValidationDetail `json:"reviews,omitempty"`
	Warns   []OrderValidationDetail `json:"warns,omitempty"`
}

type OrderValidationDetail struct {
	ValidationRuleName string `json:"validationRuleName,omitempty"`
	Message            string `json:"message,omitempty"`
	ActivityMessage    string `json:"activityMessage,omitempty"`
	OriginalSeverity   string `json:"originalSeverity,omitempty"`
	OverrideName       string `json:"overrideName,omitempty"`
	OverrideSeverity   string `json:"overrideSeverity,omitempty"`
}

type CommissionAndFee struct {
	Commission     *Commission `json:"commission,omitempty"`
	Fee            *Fees       `json:"fee,omitempty"`
	TrueCommission *Commission `json:"trueCommission,omitempty"`
}

type Commission struct {
	CommissionLegs []CommissionLeg `json:"commissionLegs,omitempty"`
}

type CommissionLeg struct {
	CommissionValues []CommissionValue `json:"commissionValues,omitempty"`
}

type CommissionValue struct {
	Value float64 `json:"value"`
	Type  string  `json:"type"`
}

type Fees struct {
	FeeLegs []FeeLeg `json:"feeLegs,omitempty"`
}

type FeeLeg struct {
	FeeValues []FeeValue `json:"feeValues,omitempty"`
}

type FeeValue struct {
	Value float64 `json:"value"`
	Type  string  `json:"type"`
}

type Transaction struct {
	ActivityId     int64          `json:"activityId"`
	User           *UserDetails   `json:"user,omitempty"`
	Description    string         `json:"description,omitempty"`
	AccountNumber  string         `json:"accountNumber,omitempty"`
	Type           string         `json:"type,omitempty"`
	Status         string         `json:"status,omitempty"`
	SubAccount     string         `json:"subAccount,omitempty"`
	TradeDate      string         `json:"tradeDate,omitempty"`
	SettlementDate string         `json:"settlementDate,omitempty"`
	PositionId     int64          `json:"positionId,omitempty"`
	OrderId        int64          `json:"orderId,omitempty"`
	NetAmount      float64        `json:"netAmount,omitempty"`
	ActivityType   string         `json:"activityType,omitempty"`
	TransferItems  []TransferItem `json:"transferItems,omitempty"`
}

type UserDetails struct {
	CdDomainId     string `json:"cdDomainId,omitempty"`
	Login          string `json:"login,omitempty"`
	Type           string `json:"type,omitempty"`
	UserId         int64  `json:"userId,omitempty"`
	SystemUserName string `json:"systemUserName,omitempty"`
	FirstName      string `json:"firstName,omitempty"`
	LastName       string `json:"lastName,omitempty"`
	BrokerRepCode  string `json:"brokerRepCode,omitempty"`
}

type TransferItem struct {
	Instrument     *TransactionInstrument `json:"instrument,omitempty"`
	Amount         float64                `json:"amount,omitempty"`
	Cost           float64                `json:"cost,omitempty"`
	Price          float64                `json:"price,omitempty"`
	FeeType        string                 `json:"feeType,omitempty"`
	PositionEffect string                 `json:"positionEffect,omitempty"`
}

type TransactionInstrument struct {
	AssetType    string  `json:"assetType"`
	Cusip        string  `json:"cusip,omitempty"`
	Symbol       string  `json:"symbol,omitempty"`
	Description  string  `json:"description,omitempty"`
	InstrumentId int64   `json:"instrumentId,omitempty"`
	NetChange    float64 `json:"netChange,omitempty"`
	// Additional fields for specific types can be added here as needed
}

type UserPreference struct {
	Accounts     []UserPreferenceAccount `json:"accounts,omitempty"`
	StreamerInfo []StreamerInfo          `json:"streamerInfo,omitempty"`
	Offers       []Offer                 `json:"offers,omitempty"`
}

type UserPreferenceAccount struct {
	AccountNumber      string `json:"accountNumber"`
	PrimaryAccount     bool   `json:"primaryAccount"`
	Type               string `json:"type"`
	NickName           string `json:"nickName,omitempty"`
	AccountColor       string `json:"accountColor,omitempty"`
	DisplayAcctId      string `json:"displayAcctId,omitempty"`
	AutoPositionEffect bool   `json:"autoPositionEffect"`
}

type StreamerInfo struct {
	StreamerSocketUrl      string `json:"streamerSocketUrl"`
	SchwabClientCustomerId string `json:"schwabClientCustomerId"`
	SchwabClientCorrelId   string `json:"schwabClientCorrelId"`
	SchwabClientChannel    string `json:"schwabClientChannel"`
	SchwabClientFunctionId string `json:"schwabClientFunctionId"`
}

type Offer struct {
	Level2Permissions bool   `json:"level2Permissions"`
	MktDataPermission string `json:"mktDataPermission"`
}

// --- Market Data Models ---

type QuoteResponse map[string]QuoteResponseObject

type QuoteResponseObject struct {
	AssetMainType string          `json:"assetMainType,omitempty"`
	AssetSubType  string          `json:"assetSubType,omitempty"`
	SSID          int64           `json:"ssid,omitempty"`
	Symbol        string          `json:"symbol,omitempty"`
	Realtime      bool            `json:"realtime,omitempty"`
	QuoteType     string          `json:"quoteType,omitempty"`
	Extended      *ExtendedMarket `json:"extended,omitempty"`
	Fundamental   *Fundamental    `json:"fundamental,omitempty"`
	Quote         *QuoteData      `json:"quote,omitempty"` // Generic quote data, covers Equity, Option, etc. fields
	Reference     *ReferenceData  `json:"reference,omitempty"`
	Regular       *RegularMarket  `json:"regular,omitempty"`
}

// Merging QuoteEquity, QuoteOption, QuoteFuture, etc. into a single struct for simplicity
// or we can use map[string]interface{} but struct is safer.
// I will include fields from all of them.
type QuoteData struct {
	WeekHigh52             float64 `json:"52WeekHigh,omitempty"`
	WeekLow52              float64 `json:"52WeekLow,omitempty"`
	AskMICId               string  `json:"askMICId,omitempty"`
	AskPrice               float64 `json:"askPrice,omitempty"`
	AskSize                int32   `json:"askSize,omitempty"`
	AskTime                int64   `json:"askTime,omitempty"`
	BidMICId               string  `json:"bidMICId,omitempty"`
	BidPrice               float64 `json:"bidPrice,omitempty"`
	BidSize                int32   `json:"bidSize,omitempty"`
	BidTime                int64   `json:"bidTime,omitempty"`
	ClosePrice             float64 `json:"closePrice,omitempty"`
	HighPrice              float64 `json:"highPrice,omitempty"`
	LastMICId              string  `json:"lastMICId,omitempty"`
	LastPrice              float64 `json:"lastPrice,omitempty"`
	LastSize               int32   `json:"lastSize,omitempty"`
	LowPrice               float64 `json:"lowPrice,omitempty"`
	Mark                   float64 `json:"mark,omitempty"`
	MarkChange             float64 `json:"markChange,omitempty"`
	MarkPercentChange      float64 `json:"markPercentChange,omitempty"`
	NetChange              float64 `json:"netChange,omitempty"`
	NetPercentChange       float64 `json:"netPercentChange,omitempty"`
	OpenPrice              float64 `json:"openPrice,omitempty"`
	QuoteTime              int64   `json:"quoteTime,omitempty"`
	SecurityStatus         string  `json:"securityStatus,omitempty"`
	TotalVolume            int64   `json:"totalVolume,omitempty"`
	TradeTime              int64   `json:"tradeTime,omitempty"`
	Volatility             float64 `json:"volatility,omitempty"`
	Delta                  float64 `json:"delta,omitempty"`
	Gamma                  float64 `json:"gamma,omitempty"`
	Theta                  float64 `json:"theta,omitempty"`
	Vega                   float64 `json:"vega,omitempty"`
	Rho                    float64 `json:"rho,omitempty"`
	TimeValue              float64 `json:"timeValue,omitempty"`
	OpenInterest           float64 `json:"openInterest,omitempty"` // spec says float or int, use float
	MoneyIntrinsicValue    float64 `json:"moneyIntrinsicValue,omitempty"`
	TheoreticalOptionValue float64 `json:"theoreticalOptionValue,omitempty"`
	UnderlyingPrice        float64 `json:"underlyingPrice,omitempty"`
	ImpliedYield           float64 `json:"impliedYield,omitempty"`
	FuturePercentChange    float64 `json:"futurePercentChange,omitempty"`
	QuotedInSession        bool    `json:"quotedInSession,omitempty"`
	SettleTime             int64   `json:"settleTime,omitempty"`
	Tick                   float64 `json:"tick,omitempty"`
	TickAmount             float64 `json:"tickAmount,omitempty"`
	NAV                    float64 `json:"nAV,omitempty"`
}

type ExtendedMarket struct {
	AskPrice    float64 `json:"askPrice,omitempty"`
	AskSize     int32   `json:"askSize,omitempty"`
	BidPrice    float64 `json:"bidPrice,omitempty"`
	BidSize     int32   `json:"bidSize,omitempty"`
	LastPrice   float64 `json:"lastPrice,omitempty"`
	LastSize    int32   `json:"lastSize,omitempty"`
	Mark        float64 `json:"mark,omitempty"`
	QuoteTime   int64   `json:"quoteTime,omitempty"`
	TotalVolume int64   `json:"totalVolume,omitempty"`
	TradeTime   int64   `json:"tradeTime,omitempty"`
}

type Fundamental struct {
	Avg10DaysVolume    float64 `json:"avg10DaysVolume,omitempty"`
	Avg1YearVolume     float64 `json:"avg1YearVolume,omitempty"`
	DeclarationDate    string  `json:"declarationDate,omitempty"`
	DivAmount          float64 `json:"divAmount,omitempty"`
	DivExDate          string  `json:"divExDate,omitempty"`
	DivFreq            int     `json:"divFreq,omitempty"`
	DivPayAmount       float64 `json:"divPayAmount,omitempty"`
	DivPayDate         string  `json:"divPayDate,omitempty"`
	DivYield           float64 `json:"divYield,omitempty"`
	EPS                float64 `json:"eps,omitempty"`
	FundLeverageFactor float64 `json:"fundLeverageFactor,omitempty"`
	FundStrategy       string  `json:"fundStrategy,omitempty"`
	NextDivExDate      string  `json:"nextDivExDate,omitempty"`
	NextDivPayDate     string  `json:"nextDivPayDate,omitempty"`
	PERatio            float64 `json:"peRatio,omitempty"`
}

type ReferenceData struct {
	Cusip                 string  `json:"cusip,omitempty"`
	Description           string  `json:"description,omitempty"`
	Exchange              string  `json:"exchange,omitempty"`
	ExchangeName          string  `json:"exchangeName,omitempty"`
	FsiDesc               string  `json:"fsiDesc,omitempty"`
	HtbQuantity           int32   `json:"htbQuantity,omitempty"`
	HtbRate               float64 `json:"htbRate,omitempty"`
	IsHardToBorrow        bool    `json:"isHardToBorrow,omitempty"`
	IsShortable           bool    `json:"isShortable,omitempty"`
	OtcMarketTier         string  `json:"otcMarketTier,omitempty"`
	ContractType          string  `json:"contractType,omitempty"`
	DaysToExpiration      int32   `json:"daysToExpiration,omitempty"`
	Deliverables          string  `json:"deliverables,omitempty"`
	ExerciseType          string  `json:"exerciseType,omitempty"`
	ExpirationDay         int32   `json:"expirationDay,omitempty"`
	ExpirationMonth       int32   `json:"expirationMonth,omitempty"`
	ExpirationType        string  `json:"expirationType,omitempty"`
	ExpirationYear        int32   `json:"expirationYear,omitempty"`
	IsPennyPilot          bool    `json:"isPennyPilot,omitempty"`
	LastTradingDay        int64   `json:"lastTradingDay,omitempty"`
	Multiplier            float64 `json:"multiplier,omitempty"`
	SettlementType        string  `json:"settlementType,omitempty"`
	StrikePrice           float64 `json:"strikePrice,omitempty"`
	Underlying            string  `json:"underlying,omitempty"`
	FutureActiveSymbol    string  `json:"futureActiveSymbol,omitempty"`
	FutureExpirationDate  int64   `json:"futureExpirationDate,omitempty"`
	FutureIsActive        bool    `json:"futureIsActive,omitempty"`
	FutureMultiplier      float64 `json:"futureMultiplier,omitempty"`
	FuturePriceFormat     string  `json:"futurePriceFormat,omitempty"`
	FutureSettlementPrice float64 `json:"futureSettlementPrice,omitempty"`
	FutureTradingHours    string  `json:"futureTradingHours,omitempty"`
	Product               string  `json:"product,omitempty"`
}

type RegularMarket struct {
	RegularMarketLastPrice     float64 `json:"regularMarketLastPrice,omitempty"`
	RegularMarketLastSize      int32   `json:"regularMarketLastSize,omitempty"`
	RegularMarketNetChange     float64 `json:"regularMarketNetChange,omitempty"`
	RegularMarketPercentChange float64 `json:"regularMarketPercentChange,omitempty"`
	RegularMarketTradeTime     int64   `json:"regularMarketTradeTime,omitempty"`
}

type OptionChain struct {
	Symbol           string                                 `json:"symbol"`
	Status           string                                 `json:"status"`
	Underlying       *Underlying                            `json:"underlying,omitempty"`
	Strategy         string                                 `json:"strategy"`
	Interval         float64                                `json:"interval"`
	IsDelayed        bool                                   `json:"isDelayed"`
	IsIndex          bool                                   `json:"isIndex"`
	DaysToExpiration float64                                `json:"daysToExpiration"`
	InterestRate     float64                                `json:"interestRate"`
	UnderlyingPrice  float64                                `json:"underlyingPrice"`
	Volatility       float64                                `json:"volatility"`
	CallExpDateMap   map[string]map[string][]OptionContract `json:"callExpDateMap"` // Map string -> Map string -> List
	PutExpDateMap    map[string]map[string][]OptionContract `json:"putExpDateMap"`
}

type Underlying struct {
	Symbol            string  `json:"symbol"`
	Description       string  `json:"description,omitempty"`
	Change            float64 `json:"change,omitempty"`
	PercentChange     float64 `json:"percentChange,omitempty"`
	Close             float64 `json:"close,omitempty"`
	Bid               float64 `json:"bid,omitempty"`
	Ask               float64 `json:"ask,omitempty"`
	Last              float64 `json:"last,omitempty"`
	Mark              float64 `json:"mark,omitempty"`
	MarkChange        float64 `json:"markChange,omitempty"`
	MarkPercentChange float64 `json:"markPercentChange,omitempty"`
	TotalVolume       int64   `json:"totalVolume,omitempty"`
	TradeTime         int64   `json:"tradeTime,omitempty"`
	QuoteTime         int64   `json:"quoteTime,omitempty"`
	HighPrice         float64 `json:"highPrice,omitempty"`
	LowPrice          float64 `json:"lowPrice,omitempty"`
	OpenPrice         float64 `json:"openPrice,omitempty"`
}

type OptionContract struct {
	PutCall                string  `json:"putCall"`
	Symbol                 string  `json:"symbol"`
	Description            string  `json:"description"`
	ExchangeName           string  `json:"exchangeName"`
	Bid                    float64 `json:"bidPrice"`
	Ask                    float64 `json:"askPrice"`
	Last                   float64 `json:"lastPrice"`
	Mark                   float64 `json:"markPrice"`
	BidSize                int32   `json:"bidSize"`
	AskSize                int32   `json:"askSize"`
	LastSize               int32   `json:"lastSize"`
	HighPrice              float64 `json:"highPrice"`
	LowPrice               float64 `json:"lowPrice"`
	OpenPrice              float64 `json:"openPrice"`
	ClosePrice             float64 `json:"closePrice"`
	TotalVolume            int32   `json:"totalVolume"`
	TradeDate              int64   `json:"tradeDate"` // spec says number format integer, usually long?
	QuoteTime              int32   `json:"quoteTimeInLong"`
	TradeTime              int32   `json:"tradeTimeInLong"`
	NetChange              float64 `json:"netChange"`
	Volatility             float64 `json:"volatility"`
	Delta                  float64 `json:"delta"`
	Gamma                  float64 `json:"gamma"`
	Theta                  float64 `json:"theta"`
	Vega                   float64 `json:"vega"`
	Rho                    float64 `json:"rho"`
	TimeValue              float64 `json:"timeValue"`
	OpenInterest           float64 `json:"openInterest"`
	IsInTheMoney           bool    `json:"isInTheMoney"`
	TheoreticalOptionValue float64 `json:"theoreticalOptionValue"`
	TheoreticalVolatility  float64 `json:"theoreticalVolatility"`
	IsMini                 bool    `json:"isMini"`
	IsNonStandard          bool    `json:"isNonStandard"`
	StrikePrice            float64 `json:"strikePrice"`
	ExpirationDate         string  `json:"expirationDate"`
	DaysToExpiration       int     `json:"daysToExpiration"`
	ExpirationType         string  `json:"expirationType"`
	LastTradingDay         int64   `json:"lastTradingDay"`
	Multiplier             float64 `json:"multiplier"`
	SettlementType         string  `json:"settlementType"`
	DeliverableNote        string  `json:"deliverableNote"`
	IsIndexOption          bool    `json:"isIndexOption"`
	PercentChange          float64 `json:"percentChange"`
	MarkChange             float64 `json:"markChange"`
	MarkPercentChange      float64 `json:"markPercentChange"`
	IsPennyPilot           bool    `json:"isPennyPilot"`
	IntrinsicValue         float64 `json:"intrinsicValue"`
	OptionRoot             string  `json:"optionRoot"`
}

type ExpirationChain struct {
	Status         string       `json:"status"`
	ExpirationList []Expiration `json:"expirationList"`
}

type Expiration struct {
	DaysToExpiration int    `json:"daysToExpiration"`
	Expiration       string `json:"expiration"`
	ExpirationType   string `json:"expirationType"`
	Standard         bool   `json:"standard"`
	SettlementType   string `json:"settlementType"`
	OptionRoots      string `json:"optionRoots"`
}

type CandleList struct {
	Candles                  []Candle `json:"candles"`
	Empty                    bool     `json:"empty"`
	PreviousClose            float64  `json:"previousClose"`
	PreviousCloseDate        int64    `json:"previousCloseDate"`
	PreviousCloseDateISO8601 string   `json:"previousCloseDateISO8601"`
	Symbol                   string   `json:"symbol"`
}

type Candle struct {
	Close           float64 `json:"close"`
	Datetime        int64   `json:"datetime"`
	DatetimeISO8601 string  `json:"datetimeISO8601"`
	High            float64 `json:"high"`
	Low             float64 `json:"low"`
	Open            float64 `json:"open"`
	Volume          int64   `json:"volume"`
}

type InstrumentResponse struct {
	Instruments []Instrument `json:"instruments"`
}

type Instrument struct {
	Cusip          string           `json:"cusip"`
	Symbol         string           `json:"symbol"`
	Description    string           `json:"description"`
	Exchange       string           `json:"exchange"`
	AssetType      string           `json:"assetType"`
	BondFactor     string           `json:"bondFactor,omitempty"`
	BondMultiplier string           `json:"bondMultiplier,omitempty"`
	BondPrice      float64          `json:"bondPrice,omitempty"`
	Fundamental    *FundamentalInst `json:"fundamental,omitempty"`
}

type FundamentalInst struct {
	Symbol              string  `json:"symbol"`
	High52              float64 `json:"high52"`
	Low52               float64 `json:"low52"`
	DividendAmount      float64 `json:"dividendAmount"`
	DividendYield       float64 `json:"dividendYield"`
	DividendDate        string  `json:"dividendDate"`
	PERatio             float64 `json:"peRatio"`
	PEGRatio            float64 `json:"pegRatio"`
	PBRatio             float64 `json:"pbRatio"`
	PRRatio             float64 `json:"prRatio"`
	PCFRatio            float64 `json:"pcfRatio"`
	GrossMarginTTM      float64 `json:"grossMarginTTM"`
	NetProfitMarginTTM  float64 `json:"netProfitMarginTTM"`
	OperatingMarginTTM  float64 `json:"operatingMarginTTM"`
	ReturnOnEquity      float64 `json:"returnOnEquity"`
	ReturnOnAssets      float64 `json:"returnOnAssets"`
	ReturnOnInvestment  float64 `json:"returnOnInvestment"`
	QuickRatio          float64 `json:"quickRatio"`
	CurrentRatio        float64 `json:"currentRatio"`
	InterestCoverage    float64 `json:"interestCoverage"`
	TotalDebtToCapital  float64 `json:"totalDebtToCapital"`
	LTDebtToEquity      float64 `json:"ltDebtToEquity"`
	TotalDebtToEquity   float64 `json:"totalDebtToEquity"`
	EPSTTM              float64 `json:"epsTTM"`
	EPSChangePercentTTM float64 `json:"epsChangePercentTTM"`
	EPSChangeYear       float64 `json:"epsChangeYear"`
	RevChangeYear       float64 `json:"revChangeYear"`
	RevChangeTTM        float64 `json:"revChangeTTM"`
	SharesOutstanding   float64 `json:"sharesOutstanding"`
	MarketCapFloat      float64 `json:"marketCapFloat"`
	MarketCap           float64 `json:"marketCap"`
	BookValuePerShare   float64 `json:"bookValuePerShare"`
	ShortIntToFloat     float64 `json:"shortIntToFloat"`
	ShortIntDayToCover  float64 `json:"shortIntDayToCover"`
	DivGrowthRate3Year  float64 `json:"divGrowthRate3Year"`
	DividendPayAmount   float64 `json:"dividendPayAmount"`
	DividendPayDate     string  `json:"dividendPayDate"`
	Beta                float64 `json:"beta"`
	Vol1DayAvg          float64 `json:"vol1DayAvg"`
	Vol10DayAvg         float64 `json:"vol10DayAvg"`
	Vol3MonthAvg        float64 `json:"vol3MonthAvg"`
	Avg10DaysVolume     int64   `json:"avg10DaysVolume"`
	Avg1DayVolume       int64   `json:"avg1DayVolume"`
	Avg3MonthVolume     int64   `json:"avg3MonthVolume"`
	DeclarationDate     string  `json:"declarationDate"`
	DividendFreq        int32   `json:"dividendFreq"`
	EPS                 float64 `json:"eps"`
	CorpActionDate      string  `json:"corpactionDate"`
	DtnVolume           int64   `json:"dtnVolume"`
	NextDividendPayDate string  `json:"nextDividendPayDate"`
	NextDividendDate    string  `json:"nextDividendDate"`
	FundLeverageFactor  float64 `json:"fundLeverageFactor"`
	FundStrategy        string  `json:"fundStrategy"`
}

type Movers struct {
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

type MarketHours map[string]map[string]Hours

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
