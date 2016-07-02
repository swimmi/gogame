package utils

import (
	"time"
)

func PsGender(gender int) string {
	genders := []string{"女", "男"}
	return genders[gender]
}

func PsSmNum(number int) string {
	str := "零一二三四五六七八九十"
	if number <= 10 {
		return string([]rune(str)[number])
	}
	if number < 100 {
		return string([]rune(str)[number/10]) + "十" + string([]rune(str)[number%10])
	}
	return string(number)
}

func PsTime(time time.Time) string {

}
