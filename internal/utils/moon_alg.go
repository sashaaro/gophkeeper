package utils

// CheckMoonAlgorithm - вернёт true, если цифры из которых состоит строка удовлетворяют Лунному алгоритму.
// Не цифровые символы игнорируются
func CheckMoonAlgorithm(b string) bool {
	if b == "" {
		return false
	}
	chars := []uint8(b)
	sum := 0
	size := len(chars)
	for i := range b {
		c := chars[size-i-1]
		if c < '0' || c > '9' {
			return false
		}
		if i%2 == 0 {
			sum += int(c) - 48
		} else {
			t := (int(c) - 48) * 2
			if t > 9 {
				t -= 9
			}
			sum += t
		}
	}
	return sum%10 == 0
}
