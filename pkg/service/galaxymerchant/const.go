package galaxymerchant

const confusedResult = "I have no idea what you are talking about"

var originalNumbers map[string]int = map[string]int{
	"I": 1,
	"V": 5,
	"X": 10,
	"L": 50,
	"C": 100,
	"D": 500,
	"M": 1000,
}

var originalAllowedRepition = map[string]int{
	"I": 3,
	"X": 3,
	"C": 3,
	"M": 3,
	"D": 1,
	"L": 1,
	"V": 1,
}

var originalAllowedSmallerPrecedent = map[string]string{
	"V": "I",
	"X": "I",
	"L": "X",
	"C": "X",
	"D": "C",
	"M": "C",
}
