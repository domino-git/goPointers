package main

import "fmt"

func print_1() {

	var val int = 5
	var p_val *int
	var pp_val **int

	p_val = &val
	pp_val = &p_val

	tl := "|------------------------------------------------------------------------------------------------------|"
	fmt.Println(tl)
	fmt.Println("|   name |  type  |\t     address\t  |\t    value\t |\tvalue\t\t |\tvalue  |")
	fmt.Println(tl)
	fmt.Printf("|    val |   %T  |    &val: %p |    val: %v            |\n", val, &val, val)
	fmt.Printf("|  p_val |  %T  |  &p_val: %p |  p_val: %p |  *p_val: %v\n", p_val, &p_val, p_val, *p_val)
	fmt.Printf("| pp_val | %T  | &pp_val: %p | pp_val: %p | *pp_val: %p | **pp_val: %v |\n", pp_val, &pp_val, pp_val, *pp_val, **pp_val)
	fmt.Println(tl)
}

func main() {

	print_1()
	/*
		|------------------------------------------------------------------------------------------------------|
		|   name |  type  |	     address	  |	    value	 |	value		 |	value  |
		|------------------------------------------------------------------------------------------------------|
		|    val |   int  |    &val: 0xc42000a418 |    val: 5            |
		|  p_val |  *int  |  &p_val: 0xc420034020 |  p_val: 0xc42000a418 |  *p_val: 5
		| pp_val | **int  | &pp_val: 0xc420034028 | pp_val: 0xc420034020 | *pp_val: 0xc42000a418 | **pp_val: 5 |
		|------------------------------------------------------------------------------------------------------|
	*/

}
