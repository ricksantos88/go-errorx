package errorx

import (
	"errors"
	"testing"
)

func TestCustomError(t *testing.T) {
	err := &customError{
		operation: "test",
		message:   "message",
		context:   map[string]interface{}{"key": "value"},
		wrapped:   errors.New("original"),
	}

	t.Run("Error()", func(t *testing.T) {
		got := err.Error()
		want := "test: message [key=value] -> original"
		if got != want {
			t.Errorf("Error() = %v, want %v", got, want)
		}
	})

	t.Run("Unwrap()", func(t *testing.T) {
		if err.Unwrap().Error() != "original" {
			t.Error("Unwrap() falhou")
		}
	})

	t.Run("Context()", func(t *testing.T) {
		if err.Context()["key"] != "value" {
			t.Error("Context() falhou")
		}
	})

	t.Run("Operation()", func(t *testing.T) {
		if err.Operation() != "test" {
			t.Error("Operation() falhou")
		}
	})
}

func TestHandler(t *testing.T) {
	t.Run("With/WithMap", func(t *testing.T) {
		h := New("test").
			With("key1", "value1").
			WithMap(map[string]interface{}{"key2": "value2"})
		
		if h.context["key1"] != "value1" || h.context["key2"] != "value2" {
			t.Error("With/WithMap falhou")
		}
	})

	t.Run("Check", func(t *testing.T) {
		h := New("test")
		if h.Check(nil) {
			t.Error("Check com nil error falhou")
		}
		if !h.Check(errors.New("test")) {
			t.Error("Check com error falhou")
		}
	})

	t.Run("Must", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Error("Must não causou panic")
			}
		}()
		New("test").Must(errors.New("error"))
	})

	t.Run("Wrap", func(t *testing.T) {
		h := New("test").With("key", "value")
		wrapped := h.Wrap(errors.New("original"), "wrapped")
		
		if wrapped.Error() != "test: wrapped [key=value] -> original" {
			t.Error("Wrap falhou")
		}
	})
}

func TestUtilityFunctions(t *testing.T) {
	baseErr := errors.New("base error")
	wrapped := New("op").With("key", "value").Wrap(baseErr, "wrapped")

	t.Run("Is", func(t *testing.T) {
		if !Is(wrapped, baseErr) {
			t.Error("Is falhou")
		}
	})

	t.Run("As", func(t *testing.T) {
		var target ErrorWrapper
		if !As(wrapped, &target) {
			t.Error("As falhou")
		}
	})

	t.Run("Unwrap", func(t *testing.T) {
		if Unwrap(wrapped) != baseErr {
			t.Error("Unwrap falhou")
		}
	})

	t.Run("GetContext", func(t *testing.T) {
		if GetContext(wrapped)["key"] != "value" {
			t.Error("GetContext falhou")
		}
	})

	t.Run("GetOperation", func(t *testing.T) {
		if GetOperation(wrapped) != "op" {
			t.Error("GetOperation falhou")
		}
	})
}