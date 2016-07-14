package data

import (
	"encoding/csv"
	"fmt"
	"os"
	"io"
)

func Resample(src *csv.Reader, dest *csv.Writer, mode *Mode, samples int, ylabels, xlabels bool ) {
	var err error
	// write header if specified
	if ylabels {
		header, err := src.Read()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to read header of SRC, reason: %s\n", err.Error())
		}
		err = dest.Write(header)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to write header to DEST, reason: %s\n",err.Error())
		}
	}
	row, err := src.Read()
	count := 1;
	ss := make([][]string,0)
	for err == nil {
		if count % samples == 0 {
			err = dest.Write(mode(ss,xlabels))
			ss = make([][]string,0)
		}
		ss = append(ss, row)
		row, err = src.Read()
		count++
	}
	if err != nil && err != io.EOF {
		fmt.Fprintf(os.Stderr, "Failed to read line of SRC, reason: %s\n", err.Error())
	}
}
