package main

import (
	"awesomeProject/internal/app"
	"errors"
	"log"
)

func main() {
	a, err := app.New()
	if err != nil {
		log.Println(errors.New("Error New in main"))
	}
	err = a.Run()
	if err != nil {
		log.Println(errors.New("Error Start in main"))
	}

}

//func main() {
//
//	newMongo := storage.New()
//	res, err := newMongo.SelectInfoUser("64f6a9fcc2ae96c8244a81ce")
//	fmt.Println(res, " \n", err)
//
//	res1, err := newMongo.UpdateRefresh("64e1ade986a817525a13d9fd", "blablabla")
//	fmt.Println(res1, " \n", err)
//}
