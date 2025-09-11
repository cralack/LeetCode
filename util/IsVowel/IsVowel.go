package IsVowel

// (1 << 1) | (1 << 5) | (1 << 9) | (1 << 15) | (1 << 21)
const vowelMask = 0x208222

func IsVowel(ch byte) bool {
	return vowelMask>>(ch&31)&1 > 0
}
