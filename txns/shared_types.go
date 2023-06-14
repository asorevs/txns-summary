package txns

import "time"

type Transaction struct {
	ID          int
	Date        time.Time
	Transaction float64
}
