package main

import (
	"fmt"
	"net/http"
	"os/exec"
	"time"
)

// 123
func main() {
	// Обработчик для главной страницы
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		currentTime := time.Now().Format("2006-01-02 15:04:05")
		fmt.Fprintf(w, "Время сервера: %s", currentTime)
	})

	// Обработчик для вебхука
	http.HandleFunc("/update", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Получен запрос от GitHub на /update")

		cmd := exec.Command("cmd", "/C", "H:\\deploy.bat")
		output, err := cmd.CombinedOutput()

		if err != nil {
			fmt.Printf("Ошибка при выполнении: %s\nВывод скрипта:\n%s\n", err.Error(), string(output))
			http.Error(w, "Ошибка при выполнении скрипта обновления", http.StatusInternalServerError)
			return
		}

		fmt.Fprintln(w, "Обновление выполнено успешно!")
	})

	// Запуск сервера
	http.ListenAndServe(":80", nil)
}
