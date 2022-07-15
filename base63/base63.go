package base63

const (
	ALPHABET = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789_"
	length   = int64(len(ALPHABET))
)

// Encode переводит id в base63
// Так как в условии сказано, что длина сокращённого URL должна быть
// равна 10 символам, всё оставшееся место заполняется символами a,
// если значение id слишком мало
func Encode(id int64) string {
	res := make([]byte, 10)
	for i := 0; i < 10; i++ {
		if id > 0 {
			res[i] = ALPHABET[id%length]
			id /= length
		} else {
			res[i] = ALPHABET[0]
		}
	}
	return string(res)
}
