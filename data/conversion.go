package data

import "strconv"

func ConvertToFloat(rows [][]string, skiplabel bool) [][]float64 {
	conv := make([][]float64, len(rows))
	// for each row
	for i, r := range rows {
		row := make([]float64, len(rows[i]))
		// for each col
		for j, v := range r {
			if j == 0 && skiplabel {
				row[j] = 0.0
			} else {
				f, _ := strconv.ParseFloat(v, 64)
				row[j] = f
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
		if j == 0 && skiplabel {
			conv[j] = xlabel
		} else {
			conv[j] = strconv.FormatFloat(v, 'f', -1, 64)
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
			if j == 0 && skiplabel {
				row[j] = 0.0
			} else {
				s, _ := strconv.ParseInt(v, 10, 64)
				row[j] = int(s)
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
		if j == 0 && skiplabel {
			conv[j] = xlabel
		} else {
			conv[j] = strconv.FormatInt(int64(v), 10)
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
			if j == 0 && skiplabel {
				row[j] = 0.0
			} else {
				t, _ := strconv.ParseInt(v, 10, 8)
				row[j] = uint8(t)
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
		if j == 0 && skiplabel {
			conv[j] = xlabel
		} else {
			conv[j] = strconv.FormatInt(int64(v), 10)
		}
	}
	return conv
}
