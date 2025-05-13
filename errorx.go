package errorx

import (
	// "errors"
	"fmt"
	"log"
	"runtime"
)

// Handler fornece métodos para tratamento de erros
type Handler struct {
	Op      string
	Context map[string]interface{}
}

// NewHandler cria um novo Handler
func NewHandler(op string) *Handler {
	return &Handler{
		Op:      op,
		Context: make(map[string]interface{}),
	}
}

// WithContext adiciona contexto ao handler
func (h *Handler) WithContext(key string, value interface{}) *Handler {
	h.Context[key] = value
	return h
}

// Check verifica o erro e loga se não for nil
func (h *Handler) Check(err error) bool {
	if err != nil {
		_, file, line, _ := runtime.Caller(1)
		log.Printf("%s:%d - op=%s, error=%v, context=%v", file, line, h.Op, err, h.Context)
		return true
	}
	return false
}

// Must panics se err não for nil
func (h *Handler) Must(err error) {
	if err != nil {
		_, file, line, _ := runtime.Caller(1)
		log.Panicf("%s:%d - op=%s, error=%v, context=%v", file, line, h.Op, err, h.Context)
	}
}

// Wrap envolve o erro com contexto adicional
func (h *Handler) Wrap(err error, message string) error {
	if err == nil {
		return nil
	}
	return fmt.Errorf("%s: %w (context: %v)", message, err, h.Context)
}

// Try executa a função e retorna o erro tratado
func Try(op string, fn func() error) error {
	h := NewHandler(op)
	err := fn()
	if err != nil {
		return h.Wrap(err, "operation failed")
	}
	return nil
}

// MustDo executa a função e faz panic se houver erro
func MustDo(op string, fn func() error) {
	h := NewHandler(op)
	h.Must(fn())
}