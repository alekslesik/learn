package main

import (
	"fmt"
	"net/http"
	"net/http/httptrace"
)

func main() {
	// Создание нового HTTP-запроса
	req, _ := http.NewRequest("GET", "http://example.com", nil)

	// Инициализация трассировщика
	trace := &httptrace.ClientTrace{
		GotConn: func(info httptrace.GotConnInfo) {
			fmt.Println("Подключение установлено")
		},
		GotFirstResponseByte: func() {
			fmt.Println("Получен первый байт ответа")
		},
	}

	// Присоединение трассировщика к контексту
	ctx := httptrace.WithClientTrace(req.Context(), trace)
	req = req.WithContext(ctx)

	// Выполнение запроса
	_, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("Ошибка запроса:", err)
		return
	}

	fmt.Println("Запрос выполнен успешно")
}
