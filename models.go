package main

type Account struct {
	AccountNumber string
	AccountType   string
	TotalValue    string
	Stocks        []*Holding
	Funds         []*Holding
}

type Holding struct {
	Name          string
	Amount        float64
	MarketPrice   float64
	PurchasePrice float64
	ISIN          string
	Currency      string
}

type AvanzaStock struct {
	DataPoints         [][]float64   `json:"dataPoints"`
	TrendSeries        [][]float64   `json:"trendSeries"`
	AllowedResolutions []string      `json:"allowedResolutions"`
	DefaultResolution  string        `json:"defaultResolution"`
	OwnersPoints       []interface{} `json:"ownersPoints"`
	ChangePercent      float64       `json:"changePercent"`
	High               float64       `json:"high"`
	LastClosingPrice   float64       `json:"lastClosingPrice"`
	LastPrice          float64       `json:"lastPrice"`
	Low                float64       `json:"low"`
}

type AvanzaFundSearch struct {
	FundListViews []struct {
		Isin        string `json:"isin"`
		Name        string `json:"name"`
		OrderbookID int    `json:"orderbookId"`
		CompanyName string `json:"companyName"`
	} `json:"fundListViews"`
}

type AvanzaFund struct {
	Isin                   string  `json:"isin"`
	Name                   string  `json:"name"`
	Description            string  `json:"description"`
	Nav                    float64 `json:"nav"`
	NavDate                string  `json:"navDate"`
	Currency               string  `json:"currency"`
	Rating                 int     `json:"rating"`
	ProductFee             float64 `json:"productFee"`
	ManagementFee          float64 `json:"managementFee"`
	Risk                   int     `json:"risk"`
	RiskText               string  `json:"riskText"`
	DevelopmentOneDay      float64 `json:"developmentOneDay"`
	DevelopmentOneMonth    float64 `json:"developmentOneMonth"`
	DevelopmentThreeMonths float64 `json:"developmentThreeMonths"`
	DevelopmentSixMonths   float64 `json:"developmentSixMonths"`
	DevelopmentOneYear     float64 `json:"developmentOneYear"`
	DevelopmentThisYear    float64 `json:"developmentThisYear"`
	DevelopmentThreeYears  float64 `json:"developmentThreeYears"`
	DevelopmentFiveYears   float64 `json:"developmentFiveYears"`
	CountryChartData       []struct {
		Name        string      `json:"name"`
		Y           float64     `json:"y"`
		Type        string      `json:"type"`
		Currency    string      `json:"currency"`
		CountryCode string      `json:"countryCode"`
		Isin        interface{} `json:"isin"`
		OrderbookID interface{} `json:"orderbookId"`
	} `json:"countryChartData"`
	HoldingChartData []struct {
		Name        string  `json:"name"`
		Y           float64 `json:"y"`
		Type        string  `json:"type"`
		Currency    string  `json:"currency"`
		CountryCode string  `json:"countryCode"`
		Isin        string  `json:"isin"`
		OrderbookID string  `json:"orderbookId"`
	} `json:"holdingChartData"`
	SectorChartData []struct {
		Name        string      `json:"name"`
		Y           float64     `json:"y"`
		Type        string      `json:"type"`
		Currency    string      `json:"currency"`
		CountryCode interface{} `json:"countryCode"`
		Isin        interface{} `json:"isin"`
		OrderbookID interface{} `json:"orderbookId"`
	} `json:"sectorChartData"`
	LowCarbon         bool    `json:"lowCarbon"`
	IndexFund         bool    `json:"indexFund"`
	SharpeRatio       float64 `json:"sharpeRatio"`
	StandardDeviation float64 `json:"standardDeviation"`
	Capital           int64   `json:"capital"`
	StartDate         string  `json:"startDate"`
	FundManagers      []struct {
		Name      string `json:"name"`
		StartDate string `json:"startDate"`
	} `json:"fundManagers"`
	AdminCompany struct {
		Name    string `json:"name"`
		Country string `json:"country"`
		URL     string `json:"url"`
	} `json:"adminCompany"`
	PricingFrequency           string      `json:"pricingFrequency"`
	ProspectusLink             string      `json:"prospectusLink"`
	AumCoveredCarbon           float64     `json:"aumCoveredCarbon"`
	FossilFuelInvolvement      float64     `json:"fossilFuelInvolvement"`
	CarbonRiskScore            float64     `json:"carbonRiskScore"`
	Categories                 []string    `json:"categories"`
	FundTypeName               string      `json:"fundTypeName"`
	FundType                   string      `json:"fundType"`
	PrimaryBenchmark           string      `json:"primaryBenchmark"`
	HedgeFund                  bool        `json:"hedgeFund"`
	UcitsFund                  bool        `json:"ucitsFund"`
	RecommendedHoldingPeriod   string      `json:"recommendedHoldingPeriod"`
	PortfolioDate              string      `json:"portfolioDate"`
	PpmCode                    string      `json:"ppmCode"`
	SuperloanOrderbook         bool        `json:"superloanOrderbook"`
	EsgScore                   float64     `json:"esgScore"`
	EnvironmentalScore         float64     `json:"environmentalScore"`
	SocialScore                float64     `json:"socialScore"`
	GovernanceScore            float64     `json:"governanceScore"`
	ControversyScore           interface{} `json:"controversyScore"`
	CarbonSolutionsInvolvement float64     `json:"carbonSolutionsInvolvement"`
	ProductInvolvements        []struct {
		Product            string  `json:"product"`
		ProductDescription string  `json:"productDescription"`
		Value              float64 `json:"value"`
	} `json:"productInvolvements"`
	SustainabilityRating             int    `json:"sustainabilityRating"`
	SustainabilityRatingCategoryName string `json:"sustainabilityRatingCategoryName"`
}
