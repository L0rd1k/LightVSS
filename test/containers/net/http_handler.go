package net

import (
	"log"
	"net/http"
)

/** Функция обработчик запросов. **/
// type HandlerFunc func(http.ResponseWriter, *http.Request)
type HandlerFunc func(*Substance)

/** Интерфейс http обработчика. Хранит функции обработика в ассоц. контейнере. **/
type HttpHandler struct {
	// router map[string]HandlerFunc
	router *router
}

/** Конструктор обработчика запросов. **/
func New() *HttpHandler {
	// return &HttpHandler{router: make(map[string]HandlerFunc)}
	return &HttpHandler{router: newRouter()}
}

/** Добавление маршрута для определенного типа запроса. **/
func (httphandler *HttpHandler) addRoute(method string, pattern string, handler HandlerFunc) {
	// key := method + "-" + pattern
	log.Printf("Route %4s - %s", method, pattern)
	// //> Добавляем ключ-путь: тип запроса-путь, и его обработчик.
	// httphandler.router[key] = handler
	httphandler.router.addRoute(method, pattern, handler)
}

/** GET http запрос. **/
func (httphandler *HttpHandler) Get(pattern string, handler HandlerFunc) {
	httphandler.addRoute("GET", pattern, handler)
}

/** POST http запрос. **/
func (httphandler *HttpHandler) Post(pattern string, handler HandlerFunc) {
	httphandler.addRoute("POST", pattern, handler)
}

func (httphandler *HttpHandler) ServeHTTP(_writer http.ResponseWriter, _request *http.Request) {
	// key := _request.Method + "-" + _request.URL.Path
	// if handler, ok := httphandler.router[key]; ok {
	// 	//> Если ключ с таким типом запросом и путем существует, вызваем обработчик.
	// 	handler(_writer, _request)
	// } else {
	// 	//> Если ключа нет, кидаем 404 код ответа HTTP
	// 	fmt.Fprintf(_writer, "404 NOT FOUND: %s\n", _request.URL)
	// }
	s := newSubstance(_writer, _request)
	httphandler.router.handle(s)
}

/** Запуск http сервера, с кастомным http обработчиком **/
func (httphandler *HttpHandler) Run(addr string) (err error) {
	return http.ListenAndServe(addr, httphandler)
}
