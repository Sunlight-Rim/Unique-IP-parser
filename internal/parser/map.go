package parser

import (
	"bufio"
	"os"

	"github.com/Sunlight-Rim/unique-ip-parser/internal/tools"
	"github.com/sirupsen/logrus"
)

func RunMap() {
	uniqueIPs := make(map[uint32]struct{}, 1024)

	filePath, err := tools.GetFilePath()
	if err != nil {
		logrus.Fatalf("Get file path: %v", err)
	}

	file, err := os.Open(filePath)
	if err != nil {
		logrus.Fatalf("Open file: %v", err)
	}

	defer func() {
		if err := file.Close(); err != nil {
			logrus.Errorf("Close file: %v", err)
		}
	}()

	scanner := bufio.NewScanner(file)

	// Scan every line
	for scanner.Scan() {
		// Parse IPv4 to number
		ipNumber := tools.ParseIPToNumber(scanner.Bytes())

		if _, ok := uniqueIPs[ipNumber]; !ok {
			uniqueIPs[ipNumber] = struct{}{}
		}
	}

	if err := scanner.Err(); err != nil {
		logrus.Fatalf("Read file: %v", err)
	}

	// Save profile
	if err := tools.SaveProfile("profile.prof"); err != nil {
		logrus.Errorf("save pprof profile: %v", err)
	}

	logrus.Printf("Unique IPs: %d", len(uniqueIPs))
}
