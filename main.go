package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os/exec"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	// Открываем или создаём базу данных
	db, err := sql.Open("sqlite3", "./users.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Создаём таблицу для пользователей, если она не существует
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS users (id INTEGER PRIMARY KEY, name TEXT)`)
	if err != nil {
		log.Fatal(err)
	}

	// Обработчик для главной страницы
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Получаем имена всех пользователей из базы
		rows, err := db.Query("SELECT name FROM users")
		if err != nil {
			http.Error(w, "Ошибка при чтении из базы данных", http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		var names []string
		for rows.Next() {
			var name string
			if err := rows.Scan(&name); err != nil {
				http.Error(w, "Ошибка при обработке данных", http.StatusInternalServerError)
				return
			}
			names = append(names, name)
		}

		// Выводим имена пользователей на сайт
		fmt.Fprintf(w, "Пользователи:\n")
		for _, name := range names {
			fmt.Fprintf(w, "%s\n", name)
		}
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

	// Раздача HLS видео-файлов
	http.Handle("/hls/", http.StripPrefix("/hls/", http.FileServer(http.Dir("C:/hls"))))

	// Добавляем пользователя из консоли в базу данных
	fmt.Print("Введите имя пользователя: ")
	var name string
	fmt.Scanln(&name)

	// Вставляем имя в базу данных
	_, err = db.Exec("INSERT INTO users (name) VALUES (?)", name)
	if err != nil {
		log.Fatal(err)
	}

	// Запуск сервера
	fmt.Println("Сервер запущен на порту 80")
	http.ListenAndServe(":80", nil)
}
