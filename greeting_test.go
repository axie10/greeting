package greeting

import (
	"regexp"
	"testing"
)

//comprobamos que cuando le metemos un nombre nos devuelve un mensaje con el nombre
func TestHelloname(t *testing.T) {

	name := "Juan"
	want := regexp.MustCompile(`\b` + name + `\b`)
	msg,err := Hello("Juan")
	if !want.MatchString(msg) || err != nil{
		t.Fatalf(`Hello("Juan") = %q, %v, quiere coincidencia para %#q, nil`, msg, err, want)
	}

}


//comprobamos el manejo de error de nuestra funcion
func TestHelloEmpty(t *testing.T){

	msg, err := Hello("")
	if msg != "" || err == nil{
		t.Fatalf(`Hello("") = %q, %v, quiere "", error`, msg, err)
	}

}