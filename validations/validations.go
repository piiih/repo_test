package validations

import "strconv"

func IsNumeric(text string) bool {
    _, err := strconv.ParseFloat(text, 64)
    return err == nil
}
