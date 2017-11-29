# zabroniryi.ru SDK
[![CircleCI](https://circleci.com/gh/tmconsulting/zabroniryiru-sdk/tree/develop.svg?style=shield)](https://circleci.com/gh/tmconsulting/zabroniryiru-sdk)
[![MIT Licensed](https://img.shields.io/badge/license-MIT-blue.svg)](https://raw.githubusercontent.com/circleci/cci-demo-react/master/LICENSE)

## Using
```golang
package main
import (
	"fmt"
	"github.com/tmconsulting/zabroniryiru-sdk"
	"strconv"
)
var (
	auth = zabroniryiru_sdk.Auth{
		BuyerId:  "BuyerId",
		UserId:   "BuyerId",
		Password: "Password",
		Language: "ru",
	}
	zApi = zabroniryiru_sdk.NewApi(auth)
)
func main() {
	searchReq := zabroniryiru.HotelSearchRequest{
		City:           "2",
		Lat:            "",
		Lng:            "",
		Radius:         "30",
		ArrivalDate:    "09.11.2017",
		DepartureDate:  "10.11.2017",
		PriceFrom:      "2000",
		PriceTo:        "3000",
		NumberOfGuests: "1",
	}
	data, err := zApi.HotelSearchRequest(searchReq)
	if err != nil {
		panic(err)
	} else {
		//Вывод результата
		for _, d := range data {
			fmt.Println("Город - " + d.CityName)
			for _, r := range d.Rooms {
				price := strconv.Itoa(r.Price)
				fmt.Println("Номер - " + r.RoomName + " - " + price + "руб.")
			}
			fmt.Println("")
		}
	}
}
 ```

## Docs

### HotelSearchRequest(searchReq)
Расширенный поиск по гостиницам, метод
возвращает стоимость проживания, наличие номеров с моментальным
подтверждением бронирования, рассчитывает сумму комиссии агентства.

Входные параметры (* помечены обязательные параметры)
| Параметр       | Описание параметра               |
|-----------------|----------------------------------|
| BuyerId*        | Код компании                     |
| UserId*         | Код пользователя                 |
| Password*       | пароль                           |
| Language        | Язык вывода                      |
| City*           | Код города из CountryListRequest |
| Lat             | Долгота                          |
| Lng             | Широта                           |
| Radius          | Радиус поиска                    |
| ArrivalDate*    | Дата заезда                      |
| DepartureDate*  | Дата выезда                      |
| PriceFrom*      | Минимальная цена поиска          |
| PriceTo*        | Максимальная цена поиска         |
| NumberOfGuests* | Число гостей                     |

Выходные параметры
|       | Параметр       | Описание параметра                | Варианты значений                     |
|-------|----------------|-----------------------------------|---------------------------------------|
|       | HotelCode      | Код отеля                         |                                       |
|       | HotelName      | Название отеля                    |                                       |
|       | Address        | Адрес отеля                       |                                       |
|       | ImageUrl       | Фото отеля                        |                                       |
|       | Vat            | НДС                               |                                       |
|       | Description    | Описание отеля                    |                                       |
|       | Amenities      | Удобства, пока пустое             |                                       |
|       | CheckInTime    | Время заезда                      |                                       |
|       | CheckOutTime   | Время выезда                      |                                       |
|       | Timezone       | Часовой пояс                      |                                       |
|       | CityCode       | Код города                        |                                       |
|       | CityName       | Название города                   |                                       |
|       | HotelLatitude  | Широта отеля                      |                                       |
|       | HotelLongitude | Долгота отеля                     |                                       |
|       | CountryCode    | Код страны                        |                                       |
|       | CountryName    | Название страны                   |                                       |
|       | RatingCode     | Код типа отеля                    | от 4 до 14, подробнее,см таблицу ниже |
|       | RatingName     | Название типа отеля               | см таблицу ниже                       |
|       | StarsCode      | Класс отеля                       | от 0 до 5,подробнее, см таблицу ниже  |
|       | StarsName      | Название класса отеля             | см таблицу ниже                       |
|       | CurrencyCode   | Код валюты Из метода CurrencyList |                                       |
|       | CurrencyName   | Название валюты                   |                                       |
| Rooms |                | Описание номеров                  |                                       |
| Rooms     | RoomCode           | Код отеля                                          |                             |
| Rooms     | RoomName           | Название отеля                                     |                             |
| Rooms     | NumberOfGuests     |                                                    |                             |
| Rooms     | Price              |                                                    |                             |
| Rooms     | Rackrate           |                                                    |                             |
| Rooms     | Comission          | Комиссия агента                                    |                             |
| Comission | Room               | Комиссия по проживанию                             |                             |
| Comission | Meal               | Комиссия по питанию                                |                             |
| Comission | Total              | Комиссия Итого                                     |                             |
| Comission | TotalMealfree      | Комиссия за вычетомпитания                         |                             |
| Rooms     | PenaltySize        | В-на штрафа за несвоев. аннуляцию                  |                             |
| Rooms     | DeadlineDate       | Дата наступления штрафных санкций                  |                             |
| Rooms     | DeadlineTime       | Время наступления штрафных санкций                 |                             |
| Rooms     | PenaltyInfo        | Доп. информация по штрафу                          |                             |
| Rooms     | MealTypeCode       | Код типа питания                                   |                             |
| Rooms     | MealTypeName       | Название типа питания                              |                             |
| Rooms     | MealCategoryCode   | Код категории питания                              |                             |
| Rooms     | MealCategoryName   | Название категории питания                         |                             |
| Rooms     | MealName           | Название питания                                   |                             |
| Rooms     | MealPrice          | Цена питания                                       |                             |
| Rooms     | MealIsIncludedCode | Признак включенного завтрака (включен не включен)  | 1 - включен, 2 - не включен |
| Rooms     | MealIsIncludedName | Признак включенного завтрака (включен не включен)  | Да, Нет                     |
| Rooms     | AvailabilityCode   | Возможность бронирования по стратегии Free-sale    | 2 или 4                     |
| Rooms   | AvailabilityName | Название наличия номера                                                       | 2 - Free-sale, 4 - Продажа по запросу                      |
| Rooms   | PaymentTermsCode | Условия оплаты                                                                | 3 (временно фиксированное значение)                        |
| Rooms   | PaymentTermsName | Условия оплаты                                                                | Оплата согласно договору (временно фиксированное значение) |
| Rooms   | Periods          |                                                                               |                                                            |
| Periods | PeriodStart      | Начало периода действия цен                                                   |                                                            |
| Periods | PeriodEnd        | Окончание периода действия цен                                                |                                                            |
| Periods | PeriodDays       | Кол-во дней впериоде                                                          |                                                            |
| Periods | PeriodSummRoom   | Сумма оплата за одни сутки в периоде                                          |                                                            |
| Periods | PeriodSummMeal   | Сумма оплата за одни сутки в периоде                                          |                                                            |
| Periods | PeriodSummTotal  | Сумма оплата за одни сутки в периоде с учетом питания за дополнительную плату |                                                            |

Возможные типы отелей
| RatingCode | RatingName (ru) | RatingName (en)   |
|------------|-----------------|-------------------|
| 4          | Хостел          | Hostel            |
| 5          | Отель           | Hotel             |
| 6          | Мини-отель      | Mini-hotel        |
| 7          | Апарт-отель     | Apart-hotel       |
| 8          | Апартаменты     | Apartments        |
| 9          | Санаторий       | Health resort     |
| 10         | Вилла           | Villa             |
| 11         | Гостевой дом    | Guest home        |
| 12         | Бутик-отель     | Boutique hotel    |
| 13         | База отдыха     | Recreation center |
| 14         | Пансионат       | Holiday hotel     |

Возможные классы отелей
| StarsCode | StarsName (ru) | StarsName (en)    |
|-----------|----------------|-------------------|
| 0         | Без звезд      | Hostel            |
| 1         | 5 звезд        | 5 stars           |
| 2         | 4 звезды       | 4 stars           |
| 3         | 3 звезды       | 3 stars           |
| 4         | 2 звезды       | 2 stars           |
| 5         | 1 звезда       | 1 star            |

