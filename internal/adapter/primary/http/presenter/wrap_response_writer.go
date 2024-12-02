package presenter

import "net/http"

type WrapResponseWriter struct {
	http.ResponseWriter
	StatusCode int
	Length     int64
	Err        error
}

func NewWrapResponseWriter(w http.ResponseWriter) *WrapResponseWriter {
	return &WrapResponseWriter{ResponseWriter: w, StatusCode: http.StatusOK}
}

func GetWrapResponseWriter(w http.ResponseWriter) *WrapResponseWriter {
	if rw, ok := w.(*WrapResponseWriter); ok {
		return rw
	}
	return NewWrapResponseWriter(w)
}

// Write インターフェースのオーバーライド。
func (rw *WrapResponseWriter) Write(b []byte) (int, error) {
	n, err := rw.ResponseWriter.Write(b)
	rw.Length += int64(n)
	return n, err
}

// WriteHeader インターフェースのオーバーライド。status code を更新する時のみ呼び出される
func (rw *WrapResponseWriter) WriteHeader(statusCode int) {
	rw.StatusCode = statusCode
	rw.ResponseWriter.WriteHeader(statusCode)
}

// WriteError Middleware または、JSONWriter からのエラーをセットする。
func (rw *WrapResponseWriter) WriteError(err error) {
	rw.Err = err
}
