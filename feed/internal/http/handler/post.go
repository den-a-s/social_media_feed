package handler

import (
	"database/sql"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"social-media-feed/internal/fake"
	"social-media-feed/internal/feed_data"
	"time"

	"github.com/google/uuid"
)

func (h *Handler) createItem(w http.ResponseWriter, r *http.Request) {
	userId, err := getUserId(r)
	if err != nil {
		h.newErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	if !fake.IsAdmin(userId) {
		h.newErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	http.ServeFile(w, r, "web/templates/formPost.html")
	
}

func (h *Handler) postFormCreateItem(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(10 << 20) // Ограничение на 10 МБ
	if err != nil {
       h.logger.Info(err.Error())
    }
	textarea := r.FormValue("textarea")
	file, fileHeader, err := r.FormFile("fileInput")
    if err != nil {
        h.logger.Info(err.Error())
        http.Error(w, "Ошибка при получении файла", http.StatusBadRequest)
        return
    }
    defer file.Close() // Закрываем файл после обработки
	fmt.Println("Полученные данные формы:")
    fmt.Println(textarea+" , "+fileHeader.Filename)

	dir := "./resources/postImg"
	// Генерация уникального имени файла
    uniqueID := uuid.New().String() // Генерируем уникальный идентификатор
    ext := filepath.Ext(fileHeader.Filename) // Получаем расширение файла
    uniqueFileName := fmt.Sprintf("%s_%s%s", uniqueID, time.Now().Format("20060102150405"), ext) // Формируем уникальное имя
	filePath := filepath.Join(dir, uniqueFileName) 

	// Создаем файл в указанной директории
	outFile, err := os.Create(filePath)
	if err != nil {
		http.Error(w, "Ошибка при создании файла", http.StatusInternalServerError)
		return
	}
	defer outFile.Close()
	// Копируем содержимое загруженного файла в созданный файл
	_, err = io.Copy(outFile, file)
	if err != nil {
		http.Error(w, "Ошибка при записи файла", http.StatusInternalServerError)
		return
	}
    // Создаем экземпляр структуры Post
    post := feed_data.Post{
        ImagePath: filePath, // Путь к загруженному изображению
        Content:   sql.NullString{String: textarea, Valid: true}, // Содержимое текстовой области
    }
	
    // Вызов метода Create для добавления поста в базу данных
    newID, err := h.repo.PostGateway.Create(post)
	fmt.Println("new id is = ")
	fmt.Println(newID)
    if err != nil {
        h.logger.Info(err.Error())
        http.Error(w, "Ошибка при добавлении поста", http.StatusInternalServerError)
        return
    }

}

func (h *Handler) mainPage(w http.ResponseWriter, r *http.Request) {
	posts, err := h.repo.PostGateway.GetAll()
	if err != nil {
		h.newErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	main_html, err := h.getFilledMainTemplate(posts)
	if err != nil {
		h.newErrorResponse(w, http.StatusInternalServerError, fmt.Sprintf("Не смогли получить сформированный шаблон: %s", err))
		return
	}

	w.Write([]byte(main_html))
}


func (h *Handler) deleteItem(w http.ResponseWriter, r *http.Request) {
	userId, err := getUserId(r)
	if err != nil {
		h.newErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	if !fake.IsAdmin(userId) {
		h.newErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	// Доделать
}