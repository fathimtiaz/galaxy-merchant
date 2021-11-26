package galaxymerchant

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/shopspring/decimal"
)

type GalaxyMerchant struct {
	numbers                 map[string]int
	translation             map[string]string
	translationReversed     map[string]string
	allowedRepition         map[string]int
	allowedSmallerPrecedent map[string]string
	prices                  map[string]decimal.Decimal
	queries                 []string
	Results                 []string
}

func NewGalaxyMerchant() *GalaxyMerchant {
	return &GalaxyMerchant{
		numbers:                 make(map[string]int),
		translation:             make(map[string]string),
		translationReversed:     make(map[string]string),
		allowedRepition:         make(map[string]int),
		allowedSmallerPrecedent: make(map[string]string),
		prices:                  make(map[string]decimal.Decimal),
		Results:                 make([]string, 0),
	}
}

func (gm *GalaxyMerchant) ParseInput(input string) (err error) {
	for i, line := range strings.Split(input, "\n") {
		line = normalizeInputLine(line)

		if line == "" {
			continue
		}

		if isDictionaryLineType(line) {
			gm.translate(line)
			continue
		}

		if isPriceLineType(line) {
			if err = gm.setPrices(line); err != nil {
				return err
			}
			continue
		}

		if isQueryLineType(line) {
			gm.queries = append(gm.queries, line)
			continue
		}

		return fmt.Errorf("invalid input on line %d", i+1)
	}

	return nil
}

// translate must be called after isDictionaryLineType
func (gm *GalaxyMerchant) translate(line string) {
	lineComps := strings.Split(line, " ")

	newWord := lineComps[0]
	dictWord := lineComps[2]

	number := originalNumbers[dictWord]

	gm.translation[newWord] = dictWord
	gm.translationReversed[dictWord] = newWord
	gm.numbers[newWord] = number
	gm.allowedRepition[newWord] = originalAllowedRepition[dictWord]
	gm.allowedSmallerPrecedent[newWord] = gm.translationReversed[originalAllowedSmallerPrecedent[dictWord]]
}

// setPrices must be called after isPriceLineType
func (gm *GalaxyMerchant) setPrices(line string) (err error) {
	var words = []string{}
	var mineral = ""
	var totalPrice, amount = decimal.Decimal{}, 0

	var lineComps = strings.Split(line, " is ")
	var beforeIsComps = strings.Split(lineComps[0], " ")
	var afterIsComps = strings.Split(lineComps[1], " ")

	for _, comp := range beforeIsComps {
		if comp == "" {
			continue
		}
		if _, ok := gm.numbers[comp]; !ok {
			mineral = comp
			break
		} else {
			words = append(words, comp)
		}
	}

	if amount, err = gm.calculateAmmount(words); err != nil {
		return err
	}

	totalPriceString, err := strconv.Atoi(afterIsComps[0])
	if err != nil {
		return err
	}

	totalPrice = decimal.NewFromInt(int64(totalPriceString))

	gm.prices[mineral] = totalPrice.Div(decimal.NewFromInt(int64(amount)))

	return nil
}

func (gm *GalaxyMerchant) calculateAmmount(words []string) (result int, err error) {
	var repetitionCount = 0

	for i, word := range words {
		result += gm.numbers[word]

		if i == 0 {
			repetitionCount = 1
			continue
		}

		if words[i-1] == word {
			repetitionCount++
		} else {
			repetitionCount = 1
		}

		if repetitionCount > gm.allowedRepition[word] {
			return 0, fmt.Errorf("invalid number of repitition for word %s", word)
		}

		if gm.numbers[words[i-1]] < gm.numbers[word] {
			if words[i-1] != gm.allowedSmallerPrecedent[word] {
				return 0, fmt.Errorf("invalid precedent %s for word %s", words[i-1], word)
			} else {
				result -= gm.numbers[words[i-1]] * 2
			}
		}
	}

	return result, nil
}

func (gm *GalaxyMerchant) SetResults() (err error) {
	for _, query := range gm.queries {
		var result string
		var queryComps = strings.Split(query, " is ")

		if len(queryComps) <= 1 {
			gm.Results = append(gm.Results, confusedResult)
			continue
		}

		var afterIsComps = strings.Split(queryComps[1], " ")

		if query[:8] == "how much" {
			if result, err = gm.getAmountResult(queryComps[1], afterIsComps); err != nil {
				return err
			}

			gm.Results = append(gm.Results, result)
		}

		if query[:8] == "how many" {
			if result, err = gm.getPriceResult(queryComps[1], afterIsComps); err != nil {
				return err
			}

			gm.Results = append(gm.Results, result)
		}
	}

	return nil
}

func (gm *GalaxyMerchant) getAmountResult(query string, afterIsComps []string) (result string, err error) {
	var words = []string{}
	var amount int

	for _, comp := range afterIsComps {
		if comp == "" || comp == "?" {
			continue
		}

		if _, ok := gm.numbers[comp]; !ok {
			return confusedResult, nil
		}

		words = append(words, comp)
	}

	if amount, err = gm.calculateAmmount(words); err != nil {
		return "", err
	}

	return fmt.Sprintf("%s is %d", strings.Trim(query, " ?"), amount), nil
}

func (gm *GalaxyMerchant) getPriceResult(query string, afterIsComps []string) (result string, err error) {
	var words = []string{}
	var unitPrice decimal.Decimal
	var amount int

	for _, comp := range afterIsComps {
		var ok bool

		if comp == "" || comp == "?" {
			continue
		}

		if _, ok = gm.numbers[comp]; ok {
			words = append(words, comp)
		} else if unitPrice, ok = gm.prices[comp]; !ok {
			return confusedResult, nil
		}
	}

	if amount, err = gm.calculateAmmount(words); err != nil {
		return "", err
	}

	return fmt.Sprintf("%s is %+v Credits", strings.Trim(query, " ?"), unitPrice.Mul(decimal.NewFromInt(int64(amount)))), nil
}
