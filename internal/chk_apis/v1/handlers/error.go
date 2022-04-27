package handlers

import "fmt"

var (
	errCSVHeader        error = fmt.Errorf("csv_header_error")
	errCSVWrongDataType error = fmt.Errorf("csv_wrong_data_type")
)
