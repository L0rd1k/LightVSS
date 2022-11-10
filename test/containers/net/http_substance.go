package net

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type H map[string]interface{}

type Substance struct {
	Writer     http.ResponseWriter //> Формирует ответ http, записывая туда данные (посредством вызова Fprintf) мы возвращаем пользователю содержимое страницы.
	Request    *http.Request       //> Запрос пользователя.
	Path       string              //> Путь.
	Method     string              //> Http метод запроса (GET/POST).
	StatusCode int                 //> Http код статуса.
}

/** Создаем новый контекст. **/
func newSubstance(_w http.ResponseWriter, _r *http.Request) *Substance {
	return &Substance{
		Writer:  _w,
		Request: _r,
		Path:    _r.URL.Path,
		Method:  _r.Method,
	}
}

func (subst *Substance) PostForm(key string) string {
	return subst.Request.FormValue(key)
}

func (subst *Substance) Query(key string) string {
	return subst.Request.URL.Query().Get(key)
}

func (subst *Substance) Status(code int) {
	subst.StatusCode = code
	subst.Writer.WriteHeader(code)
}

func (subst *Substance) SetHeader(key string, value string) {
	subst.Writer.Header().Set(key, value)
}

func (subst *Substance) String(code int, format string, values ...interface{}) {
	subst.SetHeader("Content-Type", "text/plain")
	subst.Status(code)
	subst.Writer.Write([]byte(fmt.Sprintf(format, values...)))
}

func (subst *Substance) JSON(code int, obj interface{}) {
	subst.SetHeader("Content-Type", "application/json")
	subst.Status(code)
	encoder := json.NewEncoder(subst.Writer)
	if err := encoder.Encode(obj); err != nil {
		http.Error(subst.Writer, err.Error(), 500)
	}
}

func (subst *Substance) Data(code int, data []byte) {
	subst.Status(code)
	subst.Writer.Write(data)
}

func (subst *Substance) HTML(code int, html string) {
	subst.SetHeader("Content-Type", "text/html")
	subst.Status(code)
	subst.Writer.Write([]byte(html))
}
