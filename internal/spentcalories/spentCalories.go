package spentcalories

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
)

// Основные константы, необходимые для расчетов.
const (
	lenStep   = 0.65  // средняя длина шага.
	mInKm     = 1000  // количество метров в километре.
	minInH    = 60    // количество минут в часе.
	kmhInMsec = 0.278 // коэффициент для преобразования км/ч в м/с.
	cmInM     = 100   // количество сантиметров в метре.
)

func parseTraining(data string) (int, string, time.Duration, error) { //"3456,Ходьба,3h00m" формат на входе
	// ваш код
	var activity string //создали переменную для удобства чтения и понимания кода
	var actTime time.Duration

	vals := strings.Split(data, ",") //1.разделили строку на две подстроки
	if len(vals) != 3 {              //2.проверили на валидность входящих данных
		return 0, activity, actTime, errors.New("not enough agruments ")
	}
	steps, err := strconv.Atoi(vals[0]) //3. конвертируем подстроку для шагов
	if err != nil {
		return 0, activity, actTime, errors.New("incorect steps data ")
	}
	activity = vals[1]                         // задали значение переменной
	actTime, err = time.ParseDuration(vals[2]) //4. конвертируем 3-ю подстроку во время
	if err != nil {
		return 0, activity, actTime, errors.New("incorect time data ")
	}
	return steps, activity, actTime, nil
}

// distance возвращает дистанцию(в километрах), которую преодолел пользователь за время тренировки.
//
// Параметры:
//
// steps int — количество совершенных действий (число шагов при ходьбе и беге).
func distance(steps int) float64 {
	// ваш код ниже
	return lenStep * float64(steps) / mInKm

}

// meanSpeed возвращает значение средней скорости движения во время тренировки.
//
// Параметры:
//
// steps int — количество совершенных действий(число шагов при ходьбе и беге).
// duration time.Duration — длительность тренировки.
func meanSpeed(steps int, duration time.Duration) float64 {
	// ваш код ниже
	if duration <= 0 { //1. проверка на ноль
		return 0
	}
	distance := distance(steps)        //2. получили дистанцию
	return distance / duration.Hours() //3. вычислили среднюю скорость

}

// ShowTrainingInfo возвращает строку с информацией о тренировке.
//
// Параметры:
//
// data string - строка с данными.
// weight, height float64 — вес и рост пользователя.
func TrainingInfo(data string, weight, height float64) string { //"3456,Ходьба,3h00m" на входе
	// ваш код ниже
	var dist float64
	var speed float64
	var calories float64
	steps, activity, actTime, err := parseTraining(data) // 1. получаем данные
	if err != nil {
		fmt.Printf("%v", err)
		return ""
	}
	switch activity { //2. получаем тип тренировки
	case "Бег": //3. Бег
		calories = RunningSpentCalories(steps, weight, actTime)
		speed = meanSpeed(steps, actTime)
		dist = distance(steps)

	case "Ходьба": //3. ходьба
		// = WalkingSpentCalories(steps, weight, height, actTime)
		// = meanSpeed(steps, actTime)
		//distance = distance(steps)

	default:
		return "неизвестный тип тренировки" // если что-то пошло не так
	}
	return fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n", activity, actTime.Hours(), dist, speed, calories)
}

// Константы для расчета калорий, расходуемых при беге.
const (
	runningCaloriesMeanSpeedMultiplier = 18.0 // множитель средней скорости.
	runningCaloriesMeanSpeedShift      = 20.0 // среднее количество сжигаемых калорий при беге.
)

// RunningSpentCalories возвращает количество потраченных колорий при беге.
//
// Параметры:
//
// steps int - количество шагов.
// weight float64 — вес пользователя.
// duration time.Duration — длительность тренировки.
func RunningSpentCalories(steps int, weight float64, duration time.Duration) float64 {
	// ваш код здесь
	meanSpeed := meanSpeed(steps, duration)                                                            //1. получаем среднюю скорость
	return ((runningCaloriesMeanSpeedMultiplier * meanSpeed) - runningCaloriesMeanSpeedShift) * weight // получаем калории
}

// Константы для расчета калорий, расходуемых при ходьбе.
const (
	walkingCaloriesWeightMultiplier = 0.035 // множитель массы тела.
	walkingSpeedHeightMultiplier    = 0.029 // множитель роста.
)

// WalkingSpentCalories возвращает количество потраченных калорий при ходьбе.
//
// Параметры:
//
// steps int - количество шагов.
// duration time.Duration — длительность тренировки.
// weight float64 — вес пользователя.
// height float64 — рост пользователя.
func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) float64 {
	// ваш код здесь
	meanSpeed := meanSpeed(steps, duration)
	return ((walkingCaloriesWeightMultiplier * weight) + (meanSpeed*meanSpeed/height)*walkingSpeedHeightMultiplier) * duration.Hours() * minInH
}
