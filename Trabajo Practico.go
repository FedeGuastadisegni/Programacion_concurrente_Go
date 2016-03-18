package main

import (
	"fmt"
	"time"
	"runtime"
)

func ready(ch chan string, w string,tiempo int64) {
	time.Sleep(time.Duration((0.5)*1e9))
	type Duration int64
	fmt.Println("Tiempo para que se haga el",w,": ",tiempo,"Segundos");
	time.Sleep(time.Duration((tiempo) * 1e9))
	ch <- w
	fmt.Println("el",w,"esta listo!")

}

func main() {

	runtime.GOMAXPROCS(2)				//Limita la cantidad de goroutines que se pueden ejecutar simultaneamente en el sistema//i
	t0 := time.Now()					//Tiempo inicial de la ejecucion//
	ch := make(chan string)				//Creo un canal tipo string para las bebidas//

	go ready(ch, "Te",5)				//Ejecuto las goroutines//
	go ready(ch, "Cafe",10)

	fmt.Println("Bebidas:") 

	a := <-ch 							//Leo las bebidas del canal//
	b := <-ch
	time.Sleep(time.Duration((2)*1e9))
	fmt.Println(a,"y",b,"leidos del canal")

						

	t1 := time.Now()

	fmt.Printf("La ejecucion tomo %v \n", t1.Sub(t0))
}