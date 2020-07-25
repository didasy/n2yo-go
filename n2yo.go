package n2yo

import (
	"net/http"
	"time"

	"github.com/pkg/errors"
)

const (
	DefaultClientTimeout       = 30 * time.Second
	DefaultMaxTransactionCount = 1000

	APIKeyQuery = "apiKey"
)

var (
	ErrMaxTransactionCountReached = errors.New("maximum transactions count reached")
)

type N2YOer interface {
	GetTLE(id int) (Response, error)
	GetPositions(id int, obsLat, obsLang, obsAlt float64, seconds int) (Response, error)
	GetVisualPasses(id int, obsLat, obsLang, obsAlt float64, days, minVisibility int) (Response, error)
	GetRadioPasses(id int, obsLat, obsLang, obsAlt float64, days, minElevation int) (Response, error)
	CustomHTTPClient(cl *http.Client)
	SetBaseURL(url string)
}
