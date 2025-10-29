package daysteps

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/spentcalories"
)

const (
	// Длина одного шага в метрах
	stepLength = 0.65
	// Количество метров в одном километре
	mInKm = 1000
)

// Функция возвращает количество шагов и продолжительность прогулки
func parsePackage(data string) (int, time.Duration, error) {
	// TODO: реализовать функцию
	dataSplit := strings.Split(data, ",")

	if len(dataSplit) != 2 {
		return 0, 0, errors.New("dataSplit != 2")
	}

	stepsCount, err := strconv.Atoi(dataSplit[0]) // Количество шагов

	if err != nil {
		return 0, 0, err
	}

	if stepsCount <= 0 {
		return 0, 0, errors.New("stepCount <= 0")
	}

	walkDuration, err := time.ParseDuration(dataSplit[1]) // Продолжительность прогулки

	if walkDuration.Minutes() <= 0 {
		return 0, 0, errors.New("walkDuration <= 0")
	}

	if err != nil {
		return 0, 0, err
	}

	return stepsCount, walkDuration, nil

}

func DayActionInfo(data string, weight, height float64) string {
	// TODO: реализовать функцию
	steps, walkDuration, err := parsePackage(data)

	if steps <= 0 { // Количество шагов
		log.Println("Количество шагов <= 0")
		return ""
	}

	if err != nil {
		log.Println(err)
		return ""
	}

	stepDistance := float64(steps) * stepLength / mInKm // Пройденная дистанция

	calories, err := spentcalories.WalkingSpentCalories(steps, weight, height, walkDuration) // Количество каллорий

	if err != nil {
		log.Println(err)
		return ""
	}

	return fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n", steps, stepDistance, calories)
}
