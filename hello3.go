package main

import "fmt"

func main() {
	w := 3.12
	pointer()
	fmt.Printf("%f\n", w)
	w1 := int(w)
	fmt.Printf("%d\n", w1)
	w2 := float32(w1)
	fmt.Printf("%f",w2)
}
func pointer() {
	var a, b = "Slava", 24
	p := &a
	c := &b
	var r [2]string
	t := [...]string{"red", "blue", "black"}
	r[0] = "Hello"
	r[1] = "Bro"
	fmt.Printf("My name is %s and me %d years old\n", a, b)
	fmt.Printf("%s\n %d\n", *p, p)
	fmt.Printf("%d\n %d\n", *c, c)
	*p = "Vyacheslav"
	fmt.Println(*p, a)
	fmt.Println(r[0] + " " + r[1])
	fmt.Println(t)
	t[1] = "white"
	fmt.Println(t)
}
