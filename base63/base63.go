package base63

const (
	alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789_"
	length   = int64(len(alphabet))
)

func Encode(n int64) string {
	res := make([]byte, 10)
	for i := 0; i < 10; i++ {
		if n > 0 {
			res[i] = alphabet[n%length]
			n /= length
		} else {
			res[i] = alphabet[0]
		}
	}
	return string(res)
}
