# nbuload
Пакет реализует функции загрузки курсоа валют с сайтп НБУ (https:\\bank.gov.ua)
Функции:
1) func LoadRates() []NBURates - получение курса на текщую дату
2) func LoadRatesPeriod(from, to time.Time) []NBURates - получение курсов за период дат
3) func PrintData(data []NBURates) - вспомогательная функция, печатает загруженные курсы в виде таблицы

Сделано в рамках изучения языка Go
