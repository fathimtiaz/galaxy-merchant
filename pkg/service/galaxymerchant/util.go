package galaxymerchant

import (
	"strings"
)

func isDictionaryLineType(line string) bool {
	var lineComps []string = strings.Split(line, " ")

	if len(lineComps) != 3 {
		return false
	}

	if lineComps[1] != "is" {
		return false
	}

	if _, ok := originalNumbers[lineComps[2]]; !ok {
		return false
	}

	return true
}

func isPriceLineType(line string) bool {
	if len(line) < 16 {
		return false
	}

	if !strings.Contains(line, "is") {
		return false
	}

	if line[len(line)-7:] != "Credits" {
		return false
	}

	return true
}

func isQueryLineType(line string) bool {
	if len(line) < 15 {
		return false
	}

	if line[len(line)-2:] != " ?" {
		return false
	}

	if line[:8] != "how much" && line[:8] != "how many" {
		return false
	}

	return true
}
