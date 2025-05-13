package errorx

import (
	"errors"
	"fmt"
	"log"
	"runtime"
	"strings"
)

// ErrorWrapper é a interface que estende a funcionalidade de erro padrão
type ErrorWrapper interface {
	error
	Unwrap() error
	Context() map[string]interface{}
	Operation() string
}

// customError implementa ErrorWrapper
type customError struct {
	operation string
	message  string
	context  map[string]interface{}
	wrapped  error
}

func (e *customError) Error() string {
	var sb strings.Builder
	sb.WriteString(e.operation + ": " + e.message)
	
	if len(e.context) > 0 {
		sb.WriteString(" [")
		first := true
		for k, v := range e.context {
			if !first {
				sb.WriteString(", ")
			}
			fmt.Fprintf(&sb, "%s=%v", k, v)
			first = false
		}
		sb.WriteString("]")
	}
	
	if e.wrapped != nil {
		sb.WriteString(" -> " + e.wrapped.Error())
	}
	
	return sb.String()
}

func (e *customError) Unwrap() error {
	return e.wrapped
}

func (e *customError) Context() map[string]interface{} {
	return e.context
}

func (e *customError) Operation() string {
	return e.operation
}

// Handler fornece métodos para tratamento genérico de erros
type Handler struct {
	operation string
	context   map[string]interface{}
}

// New cria um novo Handler
func New(op string) *Handler {
	return &Handler{
		operation: op,
		context:   make(map[string]interface{}),
	}
}

// With adiciona contexto ao handler
func (h *Handler) With(key string, value interface{}) *Handler {
	h.context[key] = value
	return h
}

// WithMap adiciona múltiplos pares chave-valor ao contexto
func (h *Handler) WithMap(context map[string]interface{}) *Handler {
	for k, v := range context {
		h.context[k] = v
	}
	return h
}

// Check verifica o erro e loga se não for nil
func (h *Handler) Check(err error) bool {
	if err != nil {
		_, file, line, _ := runtime.Caller(1)
		log.Printf("%s:%d - op=%s, error=%v, context=%v", file, line, h.operation, err, h.context)
		return true
	}
	return false
}

// Must panics se err não for nil
func (h *Handler) Must(err error) {
	if err != nil {
		_, file, line, _ := runtime.Caller(1)
		err = &customError{
			operation: h.operation,
			message:   "critical error",
			context:   h.context,
			wrapped:   err,
		}
		log.Panicf("%s:%d - %v", file, line, err)
	}
}

// Wrap envolve o erro com contexto adicional
func (h *Handler) Wrap(err error, message string) error {
	if err == nil {
		return nil
	}
	return &customError{
		operation: h.operation,
		message:   message,
		context:   h.context,
		wrapped:   err,
	}
}

// Try executa a função e retorna o erro tratado
func Try(op string, fn func() error) error {
	h := New(op)
	err := fn()
	if err != nil {
		return h.Wrap(err, "operation failed")
	}
	return nil
}

// MustDo executa a função e faz panic se houver erro
func MustDo(op string, fn func() error) {
	h := New(op)
	h.Must(fn())
}

// Is verifica se o erro é do tipo especificado
func Is(err error, target error) bool {
	return errors.Is(err, target)
}

// As verifica se o erro pode ser convertido para o tipo especificado
func As(err error, target interface{}) bool {
	return errors.As(err, target)
}

// Unwrap desencapsula o erro
func Unwrap(err error) error {
	return errors.Unwrap(err)
}

// GetContext retorna o contexto de um erro encapsulado
func GetContext(err error) map[string]interface{} {
	var wrapper ErrorWrapper
	if errors.As(err, &wrapper) {
		return wrapper.Context()
	}
	return nil
}

// GetOperation retorna a operação de um erro encapsulado
func GetOperation(err error) string {
	var wrapper ErrorWrapper
	if errors.As(err, &wrapper) {
		return wrapper.Operation()
	}
	return ""
}