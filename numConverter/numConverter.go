package numConverter

import "strconv"

func ConvertFromBinary(number string) string { // convert binary string into decimal string
	// numberLength := len(number)
	// result := 0
	// for i := 0; i < len(number); i++ {
	// 	if number[i:i+1] == "1" {
	// 		result += power(2, numberLength - i - 1)
	// 	}
	// }
	// return strconv.Itoa(result)
	num, err := strconv.ParseInt(number, 2, 64)
	if err != nil {
		panic(err)
	}
	return strconv.FormatInt(num, 10)
}

func ConvertFromHex(number string) string { // convert hex string into decimal string
	// format := "0123456789ABCDEF"
	// numberLength := len(number)
	// result := 0
	// for i := 0; i < len(number); i++ {
	// 	index := strings.Index(format, number[i:i+1])
	// 	result += power(16, numberLength - i - 1) * index
	// }
	// return strconv.Itoa(result)
	num, err := strconv.ParseInt(number, 16, 64)
	if err != nil {
		panic(err)
	}
	return strconv.FormatInt(num, 10)
}