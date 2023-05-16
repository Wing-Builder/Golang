package main

func main() {
	w := func() {
		println("Je suis une fonction anonyme sans return")
	}

	w()

	z := func() string {
		println("Je suis une fonction anonyme.")
		return "Alex"
	}()

	println(z)

	x, y := func() (int, int) {
		println("Aucun paramètre. Deux returns")
		return 5, 7
	}()

	println(x)
	println(y)

	func(a, b int) {
		println("A * A + B * B = ", a*a+b*b)
	}(x, y)
}
