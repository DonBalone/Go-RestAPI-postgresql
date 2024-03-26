package apiserver

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAPIServer_HandleHello(t *testing.T) {
	// создание сервера
	s := New(NewConfig())
	// создаем новый экземпляр
	// httptest.ResponseRecorder
	// нужен для записи
	// ответов http запроса
	rec := httptest.NewRecorder()
	// создаем ноый http.Request
	// с get методом запроса
	// url - соответсвующий endpoint
	// _ стоит, чтобы не принимать ошибку,
	// т.к. нам похуй
	// тела запроса нет
	req, _ := http.NewRequest(http.MethodGet, "/hello", nil)
	// принимает rec и req чтобы
	// записать и передать http ответ
	// на наш сервер s
	s.handleHello().ServeHTTP(rec, req)
	// проверяем ожидаемый ответ от действительного
	// если содержимое rec.Body.String() != "Hello!"
	// то мы выведем ошибку
	// а t - это структура для тестов
	// имееющая методы: error, fail, fatal, log
	// и т.д.
	assert.Equal(t, rec.Body.String(), "Hello!")
}
