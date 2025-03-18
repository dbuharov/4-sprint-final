package daysteps

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/dbuharov/4-sprint-final/internal/spentcalories"
)

var (
	StepLength = 0.65 // длина шага в метрах
)

func parsePackage(data string) (int, time.Duration, error) {
	// ваш код ниже
	var duration time.Duration
	vals := strings.Split(data, ",") //1.разделили строку на две подстроки
	if len(vals) != 2 {              //2.проверили на валидность входящих данных
		return 0, duration, errors.New("not enough agruments ")
	}
	steps, err := strconv.Atoi(vals[0]) //3. конвертируем подстроку для шагов и обработали ошибку
	if err != nil {
		return 0, duration, errors.New("incorect steps data ")
	}
	if steps <= 0 { //4. Проверили количество шагов
		return 0, duration, errors.New("numbers of steps is incorrect ")
	}
	duration, err = time.ParseDuration(vals[1]) //5. конвертируем 2-ю подстроку во время
	if err != nil {
		return 0, duration, errors.New("incorect time data ")
	}
	return steps, duration, nil //6. возвращаем результат всех преобразований
}

// DayActionInfo обрабатывает входящий пакет, который передаётся в
// виде строки в параметре data. Параметр storage содержит пакеты за текущий день.
// Если время пакета относится к новым суткам, storage предварительно
// очищается.
// Если пакет валидный, он добавляется в слайс storage, который возвращает
// функция. Если пакет невалидный, storage возвращается без изменений.
func DayActionInfo(data string, weight, height float64) string {
	// ваш код ниже
	steps, duration, err := parsePackage(data) //1. получили данные
	if err != nil {
		fmt.Printf("Error receiving data %v\n", err)
		return err.Error()
	}
	if steps <= 0 { //2. проверили кол-во шагов
		return " "
	}
	distance := float64(steps) * StepLength                                                                                        //3. дистанция в метрах
	kmdistance := distance / 1000                                                                                                  //4. перевели дистанцию в километры
	calories := spentcalories.WalkingSpentCalories(steps, weight, height, duration)                                                //5.получили количество калорий
	return fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n", steps, kmdistance, calories) //6.
}
