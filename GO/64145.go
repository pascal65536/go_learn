package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
	"time"
	"errors"
)

// Заглушка для функции TimeNow
func TimeNow() time.Time {
	return time.Now()
}

func currentDayOfTheWeek() string {
	t := TimeNow()
	switch t.Weekday() {
	case time.Monday:
		return "Понедельник"
	case time.Tuesday:
		return "Вторник"
	case time.Wednesday:
		return "Среда"
	case time.Thursday:
		return "Четверг"
	case time.Friday:
		return "Пятница"
	case time.Saturday:
		return "Суббота"
	case time.Sunday:
		return "Воскресенье"
	default:
		return ""
	}
}

func dayOrNight() string {
	t := TimeNow()
	hour := t.Hour()
	if hour >= 10 && hour <= 22 {
		return "День"
	}
	return "Ночь"
}

func nextFriday() int {
	t := TimeNow()
	currentWeekday := t.Weekday()

	var daysUntilFriday int
	if currentWeekday == time.Friday {
		daysUntilFriday = 0
	} else {
		daysUntilFriday = (int(time.Friday) - int(currentWeekday) + 7) % 7
	}
	return daysUntilFriday
}


func CheckCurrentDayOfTheWeek(answer string) bool {
	currentDay := strings.ToLower(currentDayOfTheWeek())
	answerLower := strings.ToLower(strings.TrimSpace(answer))
	return currentDay == answerLower
}

func CheckNowDayOrNight(answer string) (bool, error) {
	if utf8.RuneCountInString(answer) != 4 {
		return false, errors.New("исправь свой ответ, а лучше ложись поспать")
	}
	currentPeriod := dayOrNight()
	answerLower := strings.ToLower(strings.TrimSpace(answer))
	currentPeriodLower := strings.ToLower(currentPeriod)
	return answerLower == currentPeriodLower, nil
}


func main() {
	fmt.Println(currentDayOfTheWeek())
	fmt.Println(dayOrNight())
	fmt.Println(nextFriday())
}

