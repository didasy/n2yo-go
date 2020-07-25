package n2yo

const (
	PassTypeVisual = "visual"
	PassTypeRadio  = "radio"
)

type Response struct {
	Error       string     `json:"error,omitempty"`
	Info        Info       `json:"info,omitempty"`
	TLE         string     `json:"tle,omitempty"`
	Positions   []Position `json:"positions,omitempty"`
	Passes      []Pass     `json:"passes,omitempty"`
	PassesCount int        `json:"passescount,omitempty"`
	PassesType  string     `json:"passes_type,omitempty"`
}

type Info struct {
	SatelliteID       int    `json:"satid"`
	SatelliteName     string `json:"satname"`
	TransactionsCount int    `json:"transactionscount"`
}

type Position struct {
	Latitude            float64 `json:"satlatitude"`
	Longitude           float64 `json:"satlongitude"`
	Azimuth             float64 `json:"azimuth"`
	Elevation           float64 `json:"elevation"`
	RightAscensionAngle float64 `json:"ra"`
	DeclinationAngle    float64 `json:"dec"`
	UnixTimestamp       int     `json:"timestamp"`
}

type Pass struct {
	StartAzimuth           float64 `json:"startAz"`
	StartAzimuthCompass    string  `json:"startAzCompass"`
	StartElevation         float64 `json:"startEl"`
	StartUnixTimestamp     int     `json:"startUTC"`
	MaxAzimuth             float64 `json:"maxAz"`
	MaxAzimuthCompass      string  `json:"maxAzCompass"`
	MaxElevation           float64 `json:"maxEl"`
	MaxUnixTimestamp       int     `json:"maxUTC"`
	EndAzimuth             float64 `json:"endAz"`
	EndAzimuthCompass      string  `json:"endAzCompass"`
	EndElevation           float64 `json:"endEl"`
	EndUnixTimestamp       int     `json:"endUTC"`
	VisualMagnitude        float64 `json:",omitempty"`
	VisibleDurationSeconds int     `json:",omitempty"`
}
