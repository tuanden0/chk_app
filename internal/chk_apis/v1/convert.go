package v1

import "chk/internal/chk_apis/v1/handlers"

func convertRowsToCSV(in []*handlers.Rows) (out []*CSV) {

	for _, row := range in {
		out = append(out, &CSV{
			UNIX:   row.UNIX,
			SYMBOL: row.SYMBOL,
			OPEN:   row.OPEN,
			HIGH:   row.HIGH,
			LOW:    row.LOW,
			CLOSE:  row.CLOSE,
		})
	}

	return
}
