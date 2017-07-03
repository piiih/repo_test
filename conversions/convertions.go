package conversions

import "strconv"

func ToInt(text string) int {
    number, _ := strconv.Atoi(text)
    return number
}

func ToString(number int) string {
    return strconv.Itoa(number)
}

func HexToString(data interface{}) string {
    val, _ := data.([]byte)

    text := string(val)

    return text
}
