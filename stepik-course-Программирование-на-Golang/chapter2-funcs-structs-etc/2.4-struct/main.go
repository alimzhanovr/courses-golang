package main

import "fmt"

type GtaPerson struct {
	On          bool
	Ammo, Power int
}

func NewGtaPerson(on bool, ammo, power int) *GtaPerson {
	return &GtaPerson{on, ammo, power}
}

func (p *GtaPerson) RideBike() bool {
	if !p.On {
		return false
	}
	if p.Power <= 0 {
		return false
	}
	p.Power--
	return true
}
func (p *GtaPerson) Shoot() bool {
	if !p.On {
		return false
	}
	if p.Ammo <= 0 {
		return false
	}
	p.Ammo--
	return true
}
func main() {
	//1
	//В рамках этого урока мы постарались представить себе уже привычные нам переменные и функции,
	//как объекты из реальной жизни. Чтобы закрепить результат мы предлагаем вам небольшую творческую задачу.
	//Вам необходимо реализовать структуру со свойствами-полями On, Ammo и Power, с типами bool, int, int соответственно.
	//У этой структуры должны быть методы: Shoot и RideBike, которые не принимают аргументов, но возвращают значение bool.
	//Если значение On == false, то оба метода вернут false.
	//Делать Shoot можно только при наличии Ammo (тогда Ammo уменьшается на единицу, а метод возвращает true),
	//если его нет, то метод вернет false. Метод RideBike работает также, но только зависит от свойства Power.
	//Чтобы проверить, что вы все сделали правильно, вы должны создать указатель на экземпляр этой структуры
	//с именем testStruct в функции main, в дальнейшем программа проверит результат.
	//Закрывающая фигурная скобка в конце main() вам не видна, но она есть.
	//Пакет main объявлять не нужно!
	//Удачи!
	p1 := NewGtaPerson(false, 100, 100)
	p2 := NewGtaPerson(true, 0, 100)
	p3 := NewGtaPerson(true, 100, 0)
	p4 := NewGtaPerson(true, 100, 100)
	fmt.Println(p1.Shoot(), p1.RideBike())
	fmt.Println(p2.Shoot(), p2.RideBike())
	fmt.Println(p3.Shoot(), p3.RideBike())
	fmt.Println(p4.Shoot(), p4.RideBike())
}
