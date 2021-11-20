package galaxymerchant

import (
	"testing"
)

func TestGalaxyMerchant_calculateAmmount(t *testing.T) {
	type fields struct {
		numbers                 map[string]int
		translation             map[string]string
		allowedRepition         map[string]int
		allowedSmallerPrecedent map[string]string
		prices                  map[string]int
		queries                 []string
		Results                 []string
	}
	type args struct {
		words []string
	}

	f := fields{
		translation: map[string]string{
			"z": "M",
			"y": "D",
			"x": "C",
		},
		numbers: map[string]int{
			"z": 1000,
			"y": 500,
			"x": 100,
		},
		allowedRepition: map[string]int{
			"z": 3,
			"y": 1,
			"x": 3,
		},
		allowedSmallerPrecedent: map[string]string{
			"z": "x",
			"y": "x",
		},
	}

	tests := []struct {
		name       string
		fields     fields
		args       args
		wantResult int
		wantErr    bool
	}{
		{
			fields:     f,
			args:       args{words: []string{"z", "z", "y"}},
			wantResult: 2500,
		},
		{
			fields:     f,
			args:       args{words: []string{"z", "z", "z"}},
			wantResult: 3000,
		},
		{
			fields:     f,
			args:       args{words: []string{"z", "z", "z", "x", "z"}},
			wantResult: 3900,
		},
		{
			fields:  f,
			args:    args{words: []string{"z", "z", "z", "z", "x", "z"}},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gm := &GalaxyMerchant{
				numbers:                 tt.fields.numbers,
				translation:             tt.fields.translation,
				allowedRepition:         tt.fields.allowedRepition,
				allowedSmallerPrecedent: tt.fields.allowedSmallerPrecedent,
				prices:                  tt.fields.prices,
				queries:                 tt.fields.queries,
				Results:                 tt.fields.Results,
			}
			gotResult, err := gm.calculateAmmount(tt.args.words)
			if (err != nil) != tt.wantErr {
				t.Errorf("GalaxyMerchant.calculateAmmount() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotResult != tt.wantResult {
				t.Errorf("GalaxyMerchant.calculateAmmount() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}

func TestGalaxyMerchant_setPrices(t *testing.T) {
	type fields struct {
		numbers                 map[string]int
		translation             map[string]string
		translationReversed     map[string]string
		allowedRepition         map[string]int
		allowedSmallerPrecedent map[string]string
		prices                  map[string]int
		queries                 []string
		Results                 []string
	}
	type args struct {
		line string
	}

	f := fields{
		translation: map[string]string{
			"glob": "I",
			"prok": "V",
			"pish": "X",
			"tegj": "L",
		},
		translationReversed: map[string]string{
			"I": "glob",
			"V": "prok",
			"X": "pish",
			"L": "tegj",
		},
		numbers: map[string]int{
			"glob": 1,
			"prok": 5,
			"pish": 10,
			"tegj": 50,
		},
		allowedRepition: map[string]int{
			"glob": 3,
			"prok": 1,
			"pish": 3,
			"tegj": 1,
		},
		allowedSmallerPrecedent: map[string]string{
			"prok": "glob",
			"pish": "glob",
			"tegj": "pish",
		},
		prices: map[string]int{},
	}

	tests := []struct {
		name       string
		fields     fields
		args       args
		wantPrices map[string]int
		wantErr    bool
	}{
		{
			fields:     f,
			args:       args{line: "glob glob silver is 2000 Credits"},
			wantPrices: map[string]int{"silver": 1000},
			wantErr:    false,
		},
		{
			fields:     f,
			args:       args{line: "pish pish Iron is 3910 Credits"},
			wantPrices: map[string]int{"Iron": 195},
			wantErr:    false,
		},
		{
			fields:     f,
			args:       args{line: "glob prok Gold is 57800 Credits"},
			wantPrices: map[string]int{"Gold": 14450},
			wantErr:    false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gm := &GalaxyMerchant{
				numbers:                 tt.fields.numbers,
				translation:             tt.fields.translation,
				allowedRepition:         tt.fields.allowedRepition,
				allowedSmallerPrecedent: tt.fields.allowedSmallerPrecedent,
				prices:                  tt.fields.prices,
				queries:                 tt.fields.queries,
				Results:                 tt.fields.Results,
			}
			if err := gm.setPrices(tt.args.line); (err != nil) != tt.wantErr {
				t.Errorf("GalaxyMerchant.setPrices() error = %v, wantErr %v", err, tt.wantErr)
			}
			for mineral, price := range tt.wantPrices {
				if price != tt.fields.prices[mineral] {
					t.Errorf("GalaxyMerchant.setPrices() field[%s].Price = %d, wantPrice[%s] %d", mineral, tt.fields.prices[mineral], mineral, price)
				}
			}
		})
	}
}

func TestGalaxyMerchant_getAmountResult(t *testing.T) {
	type fields struct {
		numbers                 map[string]int
		translation             map[string]string
		translationReversed     map[string]string
		allowedRepition         map[string]int
		allowedSmallerPrecedent map[string]string
		prices                  map[string]int
		queries                 []string
		Results                 []string
	}

	f := fields{
		translation: map[string]string{
			"glob": "I",
			"prok": "V",
			"pish": "X",
			"tegj": "L",
		},
		translationReversed: map[string]string{
			"I": "glob",
			"V": "prok",
			"X": "pish",
			"L": "tegj",
		},
		numbers: map[string]int{
			"glob": 1,
			"prok": 5,
			"pish": 10,
			"tegj": 50,
		},
		allowedRepition: map[string]int{
			"glob": 3,
			"prok": 1,
			"pish": 3,
			"tegj": 1,
		},
		allowedSmallerPrecedent: map[string]string{
			"prok": "glob",
			"pish": "glob",
			"tegj": "pish",
		},
		prices: map[string]int{},
	}

	type args struct {
		query        string
		afterIsComps []string
	}
	tests := []struct {
		name       string
		fields     fields
		args       args
		wantResult string
		wantErr    bool
	}{
		{
			fields:     f,
			args:       args{query: "pish tegj glob glob ?", afterIsComps: []string{"pish", "tegj", "glob", "glob", "?"}},
			wantResult: "pish tegj glob glob is 42",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gm := &GalaxyMerchant{
				numbers:                 tt.fields.numbers,
				translation:             tt.fields.translation,
				translationReversed:     tt.fields.translationReversed,
				allowedRepition:         tt.fields.allowedRepition,
				allowedSmallerPrecedent: tt.fields.allowedSmallerPrecedent,
				prices:                  tt.fields.prices,
				queries:                 tt.fields.queries,
				Results:                 tt.fields.Results,
			}
			gotResult, err := gm.getAmountResult(tt.args.query, tt.args.afterIsComps)
			if (err != nil) != tt.wantErr {
				t.Errorf("GalaxyMerchant.getAmountResult() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotResult != tt.wantResult {
				t.Errorf("GalaxyMerchant.getAmountResult() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}

func TestGalaxyMerchant_getPriceResult(t *testing.T) {
	type fields struct {
		numbers                 map[string]int
		translation             map[string]string
		translationReversed     map[string]string
		allowedRepition         map[string]int
		allowedSmallerPrecedent map[string]string
		prices                  map[string]int
		queries                 []string
		Results                 []string
	}
	type args struct {
		query        string
		afterIsComps []string
	}

	f := fields{
		translation: map[string]string{
			"glob": "I",
			"prok": "V",
			"pish": "X",
			"tegj": "L",
		},
		translationReversed: map[string]string{
			"I": "glob",
			"V": "prok",
			"X": "pish",
			"L": "tegj",
		},
		numbers: map[string]int{
			"glob": 1,
			"prok": 5,
			"pish": 10,
			"tegj": 50,
		},
		allowedRepition: map[string]int{
			"glob": 3,
			"prok": 1,
			"pish": 3,
			"tegj": 1,
		},
		allowedSmallerPrecedent: map[string]string{
			"prok": "glob",
			"pish": "glob",
			"tegj": "pish",
		},
		prices: map[string]int{"Silver": 8},
	}

	tests := []struct {
		name       string
		fields     fields
		args       args
		wantResult string
		wantErr    bool
	}{
		{
			fields: f,
			args: args{query: "glob prok Silver ?", afterIsComps: []string{"glob", "prok", "Silver", "?"}},
			wantResult: "glob prok Silver is 32 Credits",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gm := &GalaxyMerchant{
				numbers:                 tt.fields.numbers,
				translation:             tt.fields.translation,
				translationReversed:     tt.fields.translationReversed,
				allowedRepition:         tt.fields.allowedRepition,
				allowedSmallerPrecedent: tt.fields.allowedSmallerPrecedent,
				prices:                  tt.fields.prices,
				queries:                 tt.fields.queries,
				Results:                 tt.fields.Results,
			}
			gotResult, err := gm.getPriceResult(tt.args.query, tt.args.afterIsComps)
			if (err != nil) != tt.wantErr {
				t.Errorf("GalaxyMerchant.getPriceResult() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotResult != tt.wantResult {
				t.Errorf("GalaxyMerchant.getPriceResult() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}
