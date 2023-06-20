package entity

type Investor struct {
	ID             string
	Name           string
	AssetPossition []*InvertorAssetPosition
}

type InvertorAssetPosition struct {
	AssetId string
	Shares  int
}
