package human

import "fmt"

type Human struct {
	FamilyName string
	GivenName string
	LastName  string
	age       uint			// приватное* поле, недоступно вне пакета human
	Email     string
}

// * - В Go используется правило регистра первой буквы названия.
//  Если оно начинается в Заглавной буквы, доступ публичный. Если со строчной доступ приватный

// Встраиваем структуру Human в Action
type Action struct {
	Human
}

func (h *Human) FullName() string {
	return fmt.Sprintf("%s %s %s", h.FamilyName, h.GivenName, h.LastName)
}

func (h *Human) SayHello() {
	fmt.Printf("[%s]: Hello, wildberries!\n", h.FullName())
}

// getter приватного поля age
func (h *Human) Age() int {
	return int(h.age)
}

// setter приватного поля age
func (h *Human) SetAge(age int) {
	h.age = uint(age)
}

func (h *Action) SendMail() {
	fmt.Printf("From [%s] send mail\n", h.Email)
}

// Переопределение метода SayHello в Action
func (h *Action) SayHello(target string) {
	fmt.Printf("[%s]: Hello, %s\n", h.FullName(), target)
}