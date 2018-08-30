package main

import "fmt"

func zero(ptr *int) {
	*ptr = 0
}

type Cart struct {
	Name  string
	Price int
}

type person struct {
	name string
	age  int
}

func (c Cart) GetPrice() {
	fmt.Println(c.Price)
}

func (c *Cart) UpdatePricebyPoint(price int) { //
	c.Price = price
}

func (c Cart) UpdatePrice(price int) {
	c.Price = price
}

func main() {
	fmt.Println(person{"Bob", 20})
	fmt.Println(person{name: "Alice", age: 30})
	fmt.Println(person{name: "Fred"})

	fmt.Println(&person{name: "Ann", age: 40})
	d := person{"larry", 23}
	s := person{name: "Sean", age: 50}
	fmt.Println(s.name)
	fmt.Println(d.name)
	x := 5
	zero(&x)
	fmt.Println(x)
	c := &Cart{"bage", 100}

	c.GetPrice()
	c.UpdatePricebyPoint(300)
	c.GetPrice()
	c.UpdatePrice(100)
	c.GetPrice()

}
