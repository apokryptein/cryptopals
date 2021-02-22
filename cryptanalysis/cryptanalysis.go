package cryptanalysis

import "unicode"

// takes xor'd byte array and tests and scores text against English letter frequency
// returns phraseScore
func ScoreResult(inBytes []byte) float64 {
	// [rune]flat64 map
	// Statistical letter frequencies taken from https://cs.wellesley.edu/~fturbak/codman/letterfreq.html
	// Value for space was estimated to 0.20
	var letterFreq = make(map[rune]float64)
	letterFreq['E'] = 0.124167
	letterFreq['T'] = 0.0969225
	letterFreq['A'] = 0.0820011
	letterFreq['O'] = 0.0714095
	letterFreq['I'] = 0.0768052
	letterFreq['N'] = 0.0764055
	letterFreq['S'] = 0.0706768
	letterFreq['R'] = 0.0668132
	letterFreq['H'] = 0.0350386
	letterFreq['D'] = 0.0363709
	letterFreq['L'] = 0.0448308
	letterFreq['U'] = 0.028777
	letterFreq['C'] = 0.0344391
	letterFreq['M'] = 0.0281775
	letterFreq['F'] = 0.0235145
	letterFreq['Y'] = 0.0189182
	letterFreq['W'] = 0.0135225
	letterFreq['G'] = 0.0181188
	letterFreq['P'] = 0.0203171
	letterFreq['B'] = 0.0106581
	letterFreq['V'] = 0.0124567
	letterFreq['K'] = 0.00393019
	letterFreq['X'] = 0.00219824
	letterFreq['Q'] = 0.0009325
	letterFreq['J'] = 0.0019984
	letterFreq['Z'] = 0.000599
	letterFreq[' '] = 0.20

	// variable containing score
	var phraseScore float64 = 0

	// iterates through byte string, looks for rune value in frequency table.
	// If in table, it adds the associated value to the phraseScore variable
	for _, let := range inBytes {
		val, ok := letterFreq[unicode.ToUpper(rune(let))]

		if ok {
			phraseScore += val
		}

	}
	return phraseScore
}

// Function taken from https://stackoverflow.com/questions/53069040/checking-a-string-contains-only-ascii-characters
// Modified
func IsASCII(s []byte) bool {
	for i := 0; i < len(s); i++ {
		if s[i] > unicode.MaxASCII {
			return false
		}
	}
	return true
}
