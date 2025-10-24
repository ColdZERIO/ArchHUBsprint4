package spentcalories

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"
)

// Основные константы, необходимые для расчетов.
const (
	lenStep                    = 0.65 // средняя длина шага.
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе
)

func parseTraining(data string) (int, string, time.Duration, error) {
	// TODO: реализовать функцию
	dataSplit := strings.Split(data, ",")

	if len(dataSplit) != 3 {
		return 0, "", 0, errors.New("dataSplit != 3")
	}

	stepsCount, stepErr := strconv.Atoi(dataSplit[0]) // Количество шагов

	if stepErr != nil {
		return 0, "", 0, stepErr
	}

	if stepsCount <= 0 {
		return 0, "", 0, errors.New("stepCount <= 0")
	}

	walkDuration, walkErr := time.ParseDuration(dataSplit[2]) // Продолжительность прогулки

	if walkDuration.Minutes() <= 0 {
		return 0, "", 0, errors.New("walkDuration <= 0")
	}

	if walkErr != nil {
		return 0, "", 0, walkErr
	}

	return stepsCount, dataSplit[1], walkDuration, nil
}

func distance(steps int, height float64) float64 {
	// TODO: реализовать функцию
	stepLength := height * stepLengthCoefficient
	walkDistance := stepLength * float64(steps)

	return walkDistance / mInKm
}

func meanSpeed(steps int, height float64, duration time.Duration) float64 {
	// TODO: реализовать функцию
	if duration <= 0 {
		return 0
	}
	walkDistance := distance(steps, height)

	return walkDistance / duration.Hours()
}

func TrainingInfo(data string, weight, height float64) (string, error) {
	// TODO: реализовать функцию

	steps, typeOfTraning, duration, parseErr := parseTraining(data)

	if parseErr != nil {
		log.Println(parseErr)
		return "", parseErr
	}

	var traningInfo float64
	var err error

	switch typeOfTraning {
	case "Ходьба":
		traningInfo, err = WalkingSpentCalories(steps, weight, height, duration)

		if err != nil {
			return "", err
		}
	case "Бег":
		traningInfo, err = RunningSpentCalories(steps, weight, height, duration)

		if err != nil {
			return "", err
		}
	default:
		return "", errors.New("неизвестный тип тренировки")
	}

	SetInfo := fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n", typeOfTraning, duration.Hours(), distance(steps, height), meanSpeed(steps, height, duration), traningInfo)

	return SetInfo, nil
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию
	if steps <= 0 {
		return 0, errors.New("stepCount <= 0")
	}

	if weight <= 0 {
		return 0, errors.New("weight <= 0")
	}

	if height <= 0 {
		return 0, errors.New("height <= 0")
	}

	if duration <= 0 {
		return 0, errors.New("duratino <= 0")
	}

	walkSpeed := meanSpeed(steps, height, duration)

	CaloriesCount := weight * walkSpeed * duration.Minutes() / minInH

	return CaloriesCount, nil
}

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию
	if steps <= 0 {
		return 0, errors.New("stepCount <= 0")
	}

	if weight <= 0 {
		return 0, errors.New("weight <= 0")
	}

	if height <= 0 {
		return 0, errors.New("height <= 0")
	}

	if duration <= 0 {
		return 0, errors.New("duratino <= 0")
	}

	walkSpeed := meanSpeed(steps, height, duration)

	CaloriesCount := weight * walkSpeed * duration.Minutes() / minInH

	return CaloriesCount * walkingCaloriesCoefficient, nil
}
