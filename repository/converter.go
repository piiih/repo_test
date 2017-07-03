package repository

import "strconv"

func toInt(word string) int {
    number, _ := strconv.Atoi(word)
    return number
}

func toString(number int) string {
    return strconv.Itoa(number)
}
