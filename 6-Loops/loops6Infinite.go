/*
	Go - Infinite Loop

	A loop becomes an infinite loop if its condition never becomes false.
*/

package main

import "fmt"

func main() {
   for true  {
       fmt.Printf("This loop will run forever.\n");
   }
}