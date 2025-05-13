package errorx

import (
	"errors"
	"testing"
)

func TestHandler_Check(t *testing.T) {
	h := NewHandler("test operation")
	
	t.Run("sem erro", func(t *testing.T) {
		if h.Check(nil) {
			t.Error("Check deveria retornar false para nil error")
		}
	})
	
	t.Run("com erro", func(t *testing.T) {
		if !h.Check(errors.New("erro de teste")) {
			t.Error("Check deveria retornar true para error não-nil")
		}
	})
}

func TestHandler_Must(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("Must deveria causar panic com erro")
		}
	}()
	
	h := NewHandler("test must")
	h.Must(errors.New("erro forçado"))
}

func TestHandler_Wrap(t *testing.T) {
	h := NewHandler("wrap test").WithContext("key", "value")
	
	err := h.Wrap(errors.New("original"), "wrapper message")
	if err == nil {
		t.Fatal("Wrap não deveria retornar nil")
	}
	
	unwrapped := errors.Unwrap(err)
	if unwrapped == nil || unwrapped.Error() != "original" {
		t.Error("Unwrap falhou")
	}
}

func TestTry(t *testing.T) {
	t.Run("sem erro", func(t *testing.T) {
		err := Try("success op", func() error {
			return nil
		})
		if err != nil {
			t.Errorf("Try deveria retornar nil, retornou %v", err)
		}
	})
	
	t.Run("com erro", func(t *testing.T) {
		err := Try("failing op", func() error {
			return errors.New("operation failed")
		})
		if err == nil {
			t.Error("Try deveria retornar erro")
		}
	})
}

func TestMustDo(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Log("MustDo panicou como esperado:", r)
		}
	}()
	
	MustDo("failing must", func() error {
		return errors.New("must do error")
	})
	
	t.Error("MustDo deveria ter panicado")
}