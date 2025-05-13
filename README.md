# errorx - Biblioteca de Tratamento de Erros para Go

Biblioteca para simplificar o tratamento de erros em Go, reduzindo boilerplate code.

## Instalação

```bash
go github.com/ricksantos88/go-errorx

```

## USO BASICO

```
import "github.com/seuusuario/errorx"

func main() {
    h := errorx.NewHandler("file operation")
    
    // Verificação simples
    if h.Check(someFunction()) {
        return
    }
    
    // Com contexto
    h.WithContext("attempt", 3).Check(otherFunction())
    
    // Try pattern
    errorx.Try("complex operation", func() error {
        // várias operações
        return nil
    })
}
```

Métodos Disponíveis
NewHandler(op string) *Handler

Check(err error) bool

Must(err error)

Wrap(err error, message string) error

Try(op string, fn func() error) error

MustDo(op string, fn func() error)