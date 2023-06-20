package dto

type TradeInput struct {
	OrderId       string
	InvestorId    string
	AssetId       string
	CurrentShares int
	Shares        int
	Price         int
	OrderType     string
}
