package v1

type CSV struct {
	ID     int64   `json:"id"`
	UNIX   int64   `json:"unix"`
	SYMBOL string  `json:"symbol"`
	OPEN   float64 `json:"open"`
	HIGH   float64 `json:"high"`
	LOW    float64 `json:"low"`
	CLOSE  float64 `json:"close"`
}
