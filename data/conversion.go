package data

import "strconv"

func ConvertToFloat(rows [][]string, skiplabel bool) [][]float64 {
	conv := make([][]float64, len(rows))
	// for each row
	for i, r := range rows {
		row := make([]float64, len(rows[i]))
		// for each col
		for j, v := range r {
			if j == 1 && skiplabel {
				row[j] = 0.0
			} else {
				row[j] = strconv.ParseFloat(v, 64)
			}
		}
		conv[i] = row
	}
	return conv
}

func ConvertFromFloat(row []float64, skiplabel bool, xlabel string) []string {
	conv := make([]string, len(row))
	// for each row
	for j, v := range row {
		if j == 1 && skiplabel {
			row[j] = xlabel
		} else {
			row[j] = strconv.FormatFloat(v, "f", -1, 64)
		}
	}
	return conv
}

func ConvertToInt(rows [][]string, skiplabel bool) [][]int {
	conv := make([][]int, len(rows))
	// for each row
	for i, r := range rows {
		row := make([]int, len(rows[i]))
		// for each col
		for j, v := range r {
			if j == 1 && skiplabel {
				row[j] = 0.0
			} else {
				row[j] = strconv.ParseInt(v, 10, 64)
			}
		}
		conv[i] = row
	}
	return conv
}

func ConvertFromInt(row []int, skiplabel bool, xlabel string) []string {
	conv := make([]string, len(row))
	// for each row
	for j, v := range row {
		if j == 1 && skiplabel {
			row[j] = xlabel
		} else {
			row[j] = strconv.FormatInt(v, 10)
		}
	}
	return conv
}

func ConvertToByte(rows [][]string, skiplabel bool) [][]uint8 {
	conv := make([][]uint8, len(rows))
	// for each row
	for i, r := range rows {
		row := make([]uint8, len(rows[i]))
		// for each col
		for j, v := range r {
			if j == 1 && skiplabel {
				row[j] = 0.0
			} else {
				row[j] = strconv.ParseInt(v, 10, 8)
			}
		}
		conv[i] = row
	}
	return conv
}

func ConvertFromByte(row []uint8, skiplabel bool, xlabel string) []string {
	conv := make([]string, len(row))
	// for each row
	for j, v := range row {
		if j == 1 && skiplabel {
			row[j] = xlabel
		} else {
			row[j] = strconv.FormatInt(v, 10)
		}
	}
	return conv
}
