# Saludo en Go

Con este paquete es un forma simple de obtener saludos aleatorios entrandole como parametro el nombre que tu quieras

## Instalacion
Ejecuta el siguiente comando para instalar el paquete
```bash
go get -u github.com/axie10/greeting
```

## Uso
Aqui tienes un ejemplo de como utilizar este paquete en tu codigo

```go
import (
	"fmt"
	"log"

	"github.com/axie10/greeting"
)

func main() {

	log.SetPrefix("greeting: ")
	log.SetFlags(0)

	names := []string{"Alex", "Roel", "Juan"}
	messages, err := greeting.Hellos(names)
	if err != nil {
		log.Fatal()
	}

	// message, err := greeting.Hello("Manu")
	// if err != nil {
	// 	log.Fatal()
	// }

	fmt.Println(messages)

}

```


