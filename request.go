package n2yo

const (
	BaseURL = "https://www.n2yo.com"

	TLEPathFormat          = "/rest/v1/satellite/tle/%d"
	PositionsPathFormat    = "/rest/v1/satellite/positions/%d/%f/%f/%f/%d"
	VisualPassesPathFormat = "/rest/v1/satellite/visualpasses/%d/%f/%f/%f/%d/%d"
	RadioPassesPathFormat  = "/rest/v1/satellite/radiopasses/%d/%f/%f/%f/%d/%d"
)
