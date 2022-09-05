package main

import (
	"fmt"
	"l1/01/human"
)


func main() {
	fmt.Println("[01]: Human Structure\n")

	// Создание объекта
	var hmn = human.Human{}
	// Установим значения в поля структуры. Поряд полей не важен
	hmn = human.Human{
		LastName: "Shamilovich",
		Email: "onemgvv@gmail.com",
		FamilyName: "Gasanov",
		GivenName: "Magomed",
	}
  // Вызываем сеттер для установление значения в приватное поле
	hmn.SetAge(19)

	// Второй способ установления значений.
	// В таком случае порядок полей важен
	action := human.Action{hmn}
	// У action можно неявно вызвать метод струтктуры Human
	fmt.Printf("Fullname is: %s\n", action.FullName())
	// Методы Human можно вызвать явно
	fmt.Printf("Age is: %d\n", action.Human.Age())
	// Метод SayHello струтуры Human
	hmn.SayHello()
	// Метод, который переопределили в Action
	action.SayHello("Tommy")
	// action.SayHello() - не будет работать тк нет аргумента
	action.Human.SayHello() // Рабтает, тк это метод струтуры Human
	// Метод стуртуры action
	action.SendMail()
}
