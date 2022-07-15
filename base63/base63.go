package base63

const (
	alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789_"
	length   = int64(len(alphabet))
)

// Encode переводит id в base63
// Так как в условии сказано, что длина сокращённого URL должна быть
// равна 10 символам, всё оставшееся место заполняется символами a,
// если значение id слишком мало
func Encode(id int64) string {
	res := make([]byte, 10)
	for i := 0; i < 10; i++ {
		if id > 0 {
			res[i] = alphabet[id%length]
			id /= length
		} else {
			res[i] = alphabet[0]
		}
	}
	return string(res)
}
