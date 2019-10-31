package data

import (
	"bufio"
	"errors"
	"io"
	"os"
	"strconv"
	"strings"
)

func ParseData(r io.Reader) (flows [][]byte, err error) {
	var currentFlow []byte

	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := scanner.Text()
		// If the line starts with ">>> " then it marks the beginning
		// of a new flow.
		if strings.HasPrefix(line, ">>> ") {
			if len(currentFlow) > 0 || len(flows) > 0 {
				flows = append(flows, currentFlow)
				currentFlow = nil
			}
			continue
		}

		// Otherwise the line is a line of hex dump that looks like:
		// 00000170  fc f5 06 bf (...)  |.....X{&?......!|
		// (Some bytes have been omitted from the middle section.)

		if i := strings.IndexByte(line, ' '); i >= 0 {
			line = line[i:]
		} else {
			return nil, errors.New("invalid test data")
		}

		if i := strings.IndexByte(line, '|'); i >= 0 {
			line = line[:i]
		} else {
			return nil, errors.New("invalid test data")
		}

		hexBytes := strings.Fields(line)
		for _, hexByte := range hexBytes {
			val, err := strconv.ParseUint(hexByte, 16, 8)
			if err != nil {
				return nil, errors.New("invalid hex byte in test data: " + err.Error())
			}
			currentFlow = append(currentFlow, byte(val))
		}
	}

	if len(currentFlow) > 0 {
		flows = append(flows, currentFlow)
	}

	return flows, nil
}

func LoadData(path string) (flows [][]byte, err error) {
	in, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer in.Close()
	return ParseData(in)
}
