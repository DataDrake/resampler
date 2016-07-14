package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"github.com/DataDrake/resampler/data"
	"os"
	"strconv"
)

func Usage() {
	fmt.Println("USAGE: resampler [OPTION]... MODE TYPE SAMPLES SRC DEST \n")
	fmt.Println("Generate request statistics for an ApacheLog2DB import\n")
	fmt.Println("\tMODE:\thow to combine the samples\n")
	fmt.Println("\t\t AVG, MAX, MIN, SUM\n")
	fmt.Println("\tTYPE:\ttype of data to resample\n")
	fmt.Println("\t\tFLOAT, INT, BYTE\n")
	fmt.Println("\tSAMPLES:\tnumber of rows to combine\n")
	fmt.Println("\tSRC:\tpath to a CSV file to resample\n")
	fmt.Println("\tDEST:\tpath to a CSV file for output\n")
}

func main() {
	flag.Usage = func() { Usage() }
	ylabels := flag.Bool("ylabels", false, "Preserve first row as column labels")
	xlabels := flag.Bool("xlabels", false, "Preserve first column as row labels, by taking first value in each sample set")
	flag.Parse()

	args := flag.Args()

	// Check arg length
	if len(args) != 5 {
		Usage()
		os.Exit(1)
	}

	// get the sampling window size
	samples, err := strconv.Atoi(args[2])
	if err != nil {
		fmt.Fprintf(os.Stderr, "SAMPLES not specified as an integer: %s\n", args[2])
		Usage()
		os.Exit(1)
	}
	if samples <= 0 {
		fmt.Fprintf(os.Stderr, "SAMPLES must be a non-zero, positive number: %s\n", args[2])
		Usage()
		os.Exit(1)
	}

	// open the source CSV
	src_file, err := os.Open(args[3])
	if err != nil {
		fmt.Fprintf(os.Stderr, "SRC could not be opened, reason: %s\n", err.Error())
		Usage()
		os.Exit(1)
	}
	defer src_file.Close()

	src_csv := csv.NewReader(src_file)

	// open the dest CSV
	dest_file, err := os.Create(args[4])
	if err != nil {
		fmt.Fprintf(os.Stderr, "DEST could not be created, reason: %s\n", err.Error())
		Usage()
		os.Exit(1)
	}
	defer dest_file.Close()

	dest_csv := csv.NewWriter(dest_file)

	// Check for valid data type and set type functions
	var mode *data.Mode
	switch args[1] {
	case "FLOAT":
		mode = data.GetFloatMode(args[2])
	case "BYTE":
		mode = data.GetByteMode(args[2])
	case "INT":
		mode = data.GetIntMode(args[2])
	default:
		fmt.Fprintf(os.Stderr, "Invalid data TYPE specified: %s\n", args[1])
		Usage()
		os.Exit(1)
	}

	data.Resample(src_csv, dest_csv, mode, samples, *ylabels, *xlabels)

	dest_csv.Flush()

	os.Exit(0)
}
