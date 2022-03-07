package main

func swap2 (x, y *int) (x2, y2 *int){
	*y2, *x2 = *x, *y
	return
}
// still not resolved
func main () {
	n, m := 10, 20
	swap2(&n, &m)
	println(n,m)
}