package data

import (
	"math"
	"errors"
)

type Mode func([][]string,bool) []string

func AVGFloat(samples [][]string, skiplabel bool) []string {
	total := make([]float64,len(samples[0]))
	count := make([]float64,len(samples[0]))
	floats := ConvertToFloat(samples,skiplabel)
	// for each column
	for j,_ := range floats[0] {
		// for each row
		for _, row := range floats {
			total += row[j]
			count[j]++
		}
		total[j] /= count[j]
	}
	return ConvertFromFloat(total,skiplabel,samples[0][0])
}

func AVGInt(samples [][]string, skiplabel bool) []string {
	total := make([]int,len(samples[0]))
	count := make([]int,len(samples[0]))
	ints := ConvertToInt(samples,skiplabel)
	// for each column
	for j,_ := range ints[0] {
		// for each row
		for _, row := range ints {
			total += row[j]
			count[j]++
		}
		total[j] /= count[j]
	}
	return ConvertFromInt(total,skiplabel,samples[0][0])
}

func AVGByte(samples [][]string, skiplabel bool) []string {
	total := make([]uint8,len(samples[0]))
	count := make([]uint8,len(samples[0]))
	bytes := ConvertToByte(samples,skiplabel)
	// for each column
	for j,_ := range bytes[0] {
		// for each row
		for _, row := range bytes {
			total += row[j]
			count[j]++
		}
		total[j] /= count[j]
	}
	return ConvertFromByte(total,skiplabel,samples[0][0])
}

func MAXFloat(samples [][]string, skiplabel bool) []string {
	max := make([]float64,len(samples[0]))
	floats := ConvertToFloat(samples,skiplabel)
	// for each column
	for j,_ := range floats[0] {
		max[j] = math.Inf(-1)
		// for each row
		for _, row := range floats {
			if row[j] > max[j] {
				max = row[j]
			}
		}
	}
	return ConvertFromFloat(max,skiplabel,samples[0][0])
}

func MAXInt(samples [][]string, skiplabel bool) []string {
	max := make([]int,len(samples[0]))
	ints := ConvertToInt(samples,skiplabel)
	// for each column
	for j,_ := range ints[0] {
		max[j] = math.MinInt64
		// for each row
		for _, row := range ints {
			if row[j] > max[j] {
				max = row[j]
			}
		}
	}
	return ConvertFromInt(max,skiplabel,samples[0][0])
}

func MAXByte(samples [][]string, skiplabel bool) []string {
	max := make([]uint8,len(samples[0]))
	bytes := ConvertToByte(samples,skiplabel)
	// for each column
	for j,_ := range bytes[0] {
		max[j] = uint8(0)
		// for each row
		for _, row := range bytes {
			if row[j] > max[j] {
				max = row[j]
			}
		}
	}
	return ConvertFromByte(max,skiplabel,samples[0][0])
}

func MINFloat(samples [][]string, skiplabel bool) []string {
	min := make([]float64,len(samples[0]))
	floats := ConvertToFloat(samples,skiplabel)
	// for each column
	for j,_ := range floats[0] {
		min[j] = math.Inf(1)
		// for each row
		for _, row := range floats {
			if row[j] < min[j] {
				min = row[j]
			}
		}
	}
	return ConvertFromFloat(min,skiplabel,samples[0][0])
}

func MINInt(samples [][]string, skiplabel bool) []string {
	min := make([]int,len(samples[0]))
	ints := ConvertToInt(samples,skiplabel)
	// for each column
	for j,_ := range ints[0] {
		min[j] = math.MaxInt64
		// for each row
		for _, row := range ints {
			if row[j] < min[j] {
				min = row[j]
			}
		}
	}
	return ConvertFromInt(min,skiplabel,samples[0][0])
}

func MINByte(samples [][]string, skiplabel bool) []string {
	min := make([]uint8,len(samples[0]))
	bytes := ConvertToByte(samples,skiplabel)
	// for each column
	for j,_ := range bytes[0] {
		min[j] = uint8(255)
		// for each row
		for _, row := range bytes {
			if row[j] < min[j] {
				min = row[j]
			}
		}
	}
	return ConvertFromByte(min,skiplabel,samples[0][0])
}


func SUMFloat(samples [][]string, skiplabel bool) []string {
	sum := make([]float64,len(samples[0]))
	floats := ConvertToFloat(samples,skiplabel)
	// for each column
	for j,_ := range floats[0] {
		sum[j] = 0
		// for each row
		for _, row := range floats {
			sum += row[j]
		}
	}
	return ConvertFromFloat(sum,skiplabel,samples[0][0])
}

func SUMInt(samples [][]string, skiplabel bool) []string {
	sum := make([]int,len(samples[0]))
	ints := ConvertToInt(samples,skiplabel)
	// for each column
	for j,_ := range ints[0] {
		sum[j] = 0
		// for each row
		for _, row := range ints {
			sum += row[j]
		}
	}
	return ConvertFromInt(sum,skiplabel,samples[0][0])
}

func SUMByte(samples [][]string, skiplabel bool) []string {
	sum := make([]uint8,len(samples[0]))
	bytes := ConvertToByte(samples,skiplabel)
	// for each column
	for j,_ := range bytes[0] {
		sum[j] = 0
		// for each row
		for _, row := range bytes {
			sum += row[j]
		}
	}
	return ConvertFromByte(sum,skiplabel,samples[0][0])
}

func GetFloatMode(mode string) (*Mode, error) {
	// Check for valid mode and set mode function
	switch mode {
	case "AVG":
		return &AVGFloat,nil
	case "MAX":
		return &MAXFloat,nil
	case "MIN":
		return &MINFloat,nil
	case "SUM":
		return &SUMFloat,nil
	default:
		return nil, errors.New("Invalid MODE specified: " + mode + "\n")
	}
}

func GetIntMode(mode string) (*Mode, error) {
	// Check for valid mode and set mode function
	switch mode {
	case "AVG":
		return &AVGInt,nil
	case "MAX":
		return &MAXInt,nil
	case "MIN":
		return &MINInt,nil
	case "SUM":
		return &SUMInt,nil
	default:
		return nil, errors.New("Invalid MODE specified: " + mode + "\n")
	}
}

func GetByteMode(mode string) (*Mode, error) {
	// Check for valid mode and set mode function
	switch mode {
	case "AVG":
		return &AVGByte,nil
	case "MAX":
		return &MAXByte,nil
	case "MIN":
		return &MINByte,nil
	case "SUM":
		return &SUMByte,nil
	default:
		return nil, errors.New("Invalid MODE specified: " + mode + "\n")
	}
}