package main

import (
	"errors"
	"fmt"
)

type BookingError struct {
	Code    int
	Message string
	Flight  string
}

type TestCase struct {
	Flights  map[string]int
	Flight   string
	Seat     int
	Passport string
}

func (b *BookingError) Error() string {
	return fmt.Sprintf("Code %d | Message: %s | Flight: %s", b.Code, b.Message, b.Flight)
}

func CheckingFlightExist(flights map[string]int, flight string) error {
	if _, ok := flights[flight]; !ok {
		return &BookingError{
			Code:    10,
			Message: "такого рейса нет",
			Flight:  flight,
		}
	}
	return nil
}

func CheckSeatsAvailable(flights map[string]int, flight string, seat int) error {
	kolMest := flights[flight]
	if kolMest < seat {
		return &BookingError{
			Code:    11,
			Message: "недостаточно мест",
			Flight:  flight,
		}
	}
	return nil
}

func CheckPassportValid(passport string) error {
	if len(passport) == 0 {
		return &BookingError{
			Code:    12,
			Message: "паспорт не валиден",
			Flight:  "",
		}
	}
	return nil
}

func BookTicket(flights map[string]int, flight string, seat int, passport string) error {
	a := CheckingFlightExist(flights, flight)
	b := CheckSeatsAvailable(flights, flight, seat)
	c := CheckPassportValid(passport)

	return errors.Join(a, b, c)
}

func main() {
	flights := map[string]int{"SU123": 10, "LH404": 0, "FR999": 3}

	testCases := []TestCase{
		{flights, "SU123", 2, "21"},
		{flights, "qwe", 2, "123"},
		{flights, "LH404", 1, "2"},
		{flights, "SU123", 3, ""},
		{flights, "qwe", 3, ""},
	}

	for _, t := range testCases {
		fmt.Printf("\nПроверка рейса %s (место: %d, паспорт: %q)\n", t.Flight, t.Seat, t.Passport)
		err := BookTicket(t.Flights, t.Flight, t.Seat, t.Passport)

		if err != nil {
			fmt.Println("Не удалось проверить рейс.")

			if je, ok := err.(interface{ Unwrap() []error }); ok {
				for _, e := range je.Unwrap() {
					fmt.Println("  ошибка:", e)
				}
			} else {
				fmt.Println("  ошибка:", err)
			}
		} else {
			fmt.Println("Рейс успешно прошёл проверку ✅✅✅✅✅")
		}
	}
} // я функции делал 10-20 минут,а main часа на 2, скорее всего, даже больше, мне если чото не понятно было
// я просил чатжпт объяснить и я что-нибудь разбирал и потом переделывал как я понял, и каждый раз там
// что-нибудь полностью не правильно было и надо было абсолютно в другом разбираться, и из-за этого main
// затянулся очень и я уже жду когда ии заменит программистов наконец-то
