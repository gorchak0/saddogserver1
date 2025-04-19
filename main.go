package main

import (
	"fmt"
	"net/http"
	"os/exec"
	"time"
)

func main() {
	// Обработчик для главной страницы
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		currentTime := time.Now().Format("2006-01-02 15:04:05")
		fmt.Fprintf(w, "Время сервера: %s", currentTime)
	})

	// Обработчик для вебхука
	http.HandleFunc("/update", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Получен запрос от GitHub на /update") // <-- добавили лог
		// Запуск скрипта для обновления
		cmd := exec.Command("cmd", "/C", "H:\\saddogserver1\\deploy.bat")
		err := cmd.Run()
		if err != nil {
			http.Error(w, "Ошибка при выполнении скрипта обновления", http.StatusInternalServerError)
			return
		}
		fmt.Fprintln(w, "Обновление выполнено успешно!")
	})

	// Запуск сервера
	http.ListenAndServe(":8080", nil)
}
