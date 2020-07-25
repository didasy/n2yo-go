package n2yo

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

const (
	tleResponse = `{
		"info": {
		  "satid": 25544,
		  "satname": "SPACE STATION",
		  "transactionscount": 1
		},
		"tle": "1 25544U 98067A   20207.26175341 -.00002477  00000-0 -36149-4 0  9999\r\n2 25544  51.6430 160.0151 0001351 151.1358 178.4154 15.49502101237919"
	  }`
	positionsResponse = `{
		"info": {
		  "satid": 25544,
		  "satname": "SPACE STATION",
		  "transactionscount": 1
		},
		"positions": [
		  {
			"satlatitude": 23.76462945,
			"satlongitude": -113.2852807,
			"azimuth": 61.37,
			"elevation": -68.18,
			"ra": 339.23812471,
			"dec": 16.02308192,
			"timestamp": 1595666167
		  }
		]
	  }`
	visualPassesResponse = `{
		"info": {
		  "satid": 25544,
		  "satname": "SPACE STATION",
		  "transactionscount": 1
		},
		"passes_type": "visual"
	  }`
	radioPassesResponse = `{
		"info": {
		  "satid": 25544,
		  "satname": "SPACE STATION",
		  "transactionscount": 1
		},
		"passes": [
		  {
			"startAz": 318.19,
			"startAzCompass": "NW",
			"startEl": 0,
			"startUTC": 1595705475,
			"maxAz": 234.84,
			"maxAzCompass": "SW",
			"maxEl": 58.46,
			"maxUTC": 1595705795,
			"endAz": 149.69,
			"endAzCompass": "SE",
			"endEl": 0,
			"endUTC": 1595706115
		  },
		  {
			"startAz": 234.26,
			"startAzCompass": "SW",
			"startEl": 0,
			"startUTC": 1595747095,
			"maxAz": 304.47,
			"maxAzCompass": "NW",
			"maxEl": 24.54,
			"maxUTC": 1595747400,
			"endAz": 15.39,
			"endAzCompass": "NNE",
			"endEl": 0,
			"endUTC": 1595747700
		  }
		],
		"passes_type": "radio"
	  }`
)

func mockHTTPHandler(w http.ResponseWriter, r *http.Request) {
	if strings.Contains(r.URL.Path, "tle") {
		fmt.Fprintln(w, tleResponse)
		return
	}
	if strings.Contains(r.URL.Path, "positions") {
		fmt.Fprintln(w, positionsResponse)
		return
	}
	if strings.Contains(r.URL.Path, "visualpasses") {
		fmt.Fprintln(w, visualPassesResponse)
		return
	}
	if strings.Contains(r.URL.Path, "radiopasses") {
		fmt.Fprintln(w, radioPassesResponse)
		return
	}

	w.WriteHeader(http.StatusNotFound)
	fmt.Fprintf(w, "not found")
}

func TestClient(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(mockHTTPHandler))
	defer srv.Close()
	dummyHost := srv.URL

	apiKey := "myapikey"
	issNORADID := 25544
	lat := -6.200000
	long := 106.816666
	alt := 5.0
	sec := 1
	days := 1
	minVisibility := 3
	minElevation := 15

	client := New(apiKey)

	var wg sync.WaitGroup
	wg.Add(2)

	t.Run("Set custom HTTP client", func(t *testing.T) {
		defer wg.Done()

		cli := &http.Client{}
		cli.Timeout = time.Minute
		client.CustomHTTPClient(cli)
	})

	t.Run("Set BaseURL", func(t *testing.T) {
		defer wg.Done()

		client.SetBaseURL(dummyHost)
	})

	wg.Wait()

	t.Run("When GetTLE success", func(t *testing.T) {
		resp, err := client.GetTLE(issNORADID)

		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(t, issNORADID, resp.Info.SatelliteID)
		require.NotEmpty(t, resp.TLE)
	})

	t.Run("When GetPositions success", func(t *testing.T) {
		resp, err := client.GetPositions(issNORADID, lat, long, alt, sec)

		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(t, issNORADID, resp.Info.SatelliteID)
		require.NotEmpty(t, resp.Positions)
	})

	t.Run("When GetVisualPasses success", func(t *testing.T) {
		resp, err := client.GetVisualPasses(issNORADID, lat, long, alt, days, minVisibility)

		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(t, issNORADID, resp.Info.SatelliteID)
		require.Equal(t, PassTypeVisual, resp.PassesType)
	})

	t.Run("When GetRadioPasses success", func(t *testing.T) {
		resp, err := client.GetRadioPasses(issNORADID, lat, long, alt, days, minElevation)

		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(t, issNORADID, resp.Info.SatelliteID)
		require.Equal(t, PassTypeRadio, resp.PassesType)
	})
}
