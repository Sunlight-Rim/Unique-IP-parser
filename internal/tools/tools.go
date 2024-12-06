package tools

import (
	"flag"
	"os"
	"runtime/pprof"
	"strconv"

	"github.com/pkg/errors"
)

// GetFilePath returns file path from startup argument.
func GetFilePath() (string, error) {
	flag.Parse()

	filePath := flag.Arg(0)

	if filePath == "" {
		return "", errors.New("no file were provided")
	}

	return filePath, nil
}

// ParseIPToNumber parses IPv4 to uint32 number.
func ParseIPToNumber(ip []byte) uint32 {
	var (
		octets        [4]int64
		octetStartIdx int
		octetNumber   int
	)

	for i := range ip {
		if ip[i] == byte('.') {
			octet, _ := strconv.ParseInt(string(ip[octetStartIdx:i]), 0, 64)
			octets[octetNumber] = octet

			octetNumber++
			octetStartIdx = i + 1
		}
	}
	// Add last octet
	octet, _ := strconv.ParseInt(string(ip[octetStartIdx:]), 0, 64)
	octets[3] = octet

	// Convert IPv4 octets to number
	return uint32(octets[0]*16777216 + octets[1]*65536 + octets[2]*256 + octets[3])
}

// SaveProfile saves pprof profile.
func SaveProfile(filename string) error {
	f, err := os.Create(filename)
	if err != nil {
		return errors.Wrap(err, "create profile file")
	}

	defer f.Close()

	if err := pprof.Lookup("heap").WriteTo(f, 2); err != nil {
		return errors.Wrap(err, "write heap profile")
	}

	return nil
}
