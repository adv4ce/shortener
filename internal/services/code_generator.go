package services

func CreateShortCode(id int) string {
	const chars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	if id == 0 {
		return string(chars[0])
	}

	var code []byte
	for id > 0 {
		code = append(code, chars[id%62])
		id /= 62
	}

	for i, j := 0, len(code) - 1; i < j; i, j = i + 1, j - 1 {
		code[i], code[j] = code[j], code[i]
	}
	return string(code)
}
