package handler

import (
	"database/sql"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"social-media-feed/internal/fake"
	"social-media-feed/internal/feed_data"
	"strconv"

	//"strconv"
	"time"

	"github.com/google/uuid"
)

func (h *Handler) createPost(w http.ResponseWriter, r *http.Request) {
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
       h.logger.Error(err.Error())
	   return
    }
	textarea := r.FormValue("textarea")
	file, fileHeader, err := r.FormFile("fileInput")
    if err != nil {
        h.logger.Error(err.Error())
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
        h.logger.Error(err.Error())
        http.Error(w, "Ошибка при добавлении поста", http.StatusInternalServerError)
        return
    }

}

func (h *Handler) mainPage(w http.ResponseWriter, r *http.Request) {
	// cookie, err := r.Cookie("user_id")
    // if err != nil {
    //     http.Error(w, "cookie 'user_id' не найден", http.StatusUnauthorized)
    //     return
    // }
    // // Преобразуем значение cookie в целое число
    // userId, err := strconv.Atoi(cookie.Value)
    // if err != nil {
    //     http.Error(w, "некорректный user_id в cookie", http.StatusBadRequest)
    //     return
    // }
	postWithLike, err := h.repo.PostsWithLikeGateway.JoinPostWithLike(fake.AdminId)
	if err != nil {
        http.Error(w, fmt.Sprintf("ошибка при чтении постов с лайками: %s", err), http.StatusBadRequest)
        return
    }

	main_html, err := h.getFilledMainTemplate(postWithLike)
	if err != nil {
		h.newErrorResponse(w, http.StatusInternalServerError, fmt.Sprintf("Не смогли получить сформированный шаблон: %s", err))
		return
	}

	w.Write([]byte(main_html))
}

func (h *Handler) deletePost(w http.ResponseWriter, r *http.Request) {
	userId, err := getUserId(r)
	if err != nil {
		h.newErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	if !fake.IsAdmin(userId) {
		h.newErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	params, err := url.ParseQuery(r.URL.RawQuery)
	if err != nil {
		err_str := fmt.Sprintf("[delete post] Not get url params: %s", err)
		h.newErrorResponse(w, http.StatusInternalServerError, err_str)
		return
	}
	
	idPost := params.Get("idPost")

	h.logger.Debug("data:", "idPost", idPost)

	if idPost == ""  {
		err_str := fmt.Sprintf("[delete post] Not parse url params: %s", err)
		h.newErrorResponse(w, http.StatusInternalServerError, err_str)
		return
	}

	fmtIdPost,err := strconv.Atoi(idPost) 
	if err != nil {
        err_str := fmt.Sprintf("[delete post] Not parse postId: %s", err)
		h.newErrorResponse(w, http.StatusInternalServerError, err_str)
        return
    }

	post, err := h.repo.PostGateway.GetById(fmtIdPost)
	if err != nil {
        err_str := fmt.Sprintf("[delete post] error of getting id of post in db: %s", err)
		h.newErrorResponse(w, http.StatusInternalServerError, err_str)
        return
    }
	err = os.Remove("./"+post.ImagePath)
	if err != nil {
		err_str := fmt.Sprintf("[delete post] error of delete file : %s", err)
		h.newErrorResponse(w, http.StatusInternalServerError, err_str)
		return
	}
	err = h.repo.PostGateway.Delete(fmtIdPost)
	if err != nil {
        err_str := fmt.Sprintf("[delete post] error of delete in db: %s", err)
		h.newErrorResponse(w, http.StatusInternalServerError, err_str)
        return
    }

	w.WriteHeader(http.StatusOK)
}