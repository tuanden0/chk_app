package v1

import "time"

const (
	WAIT_FOR_SHUTDOWN time.Duration = 30 * time.Second
	SERVER_ADDR       string        = "0.0.0.0:80"
	UPLOAD_DIR        string        = "upload"
)
