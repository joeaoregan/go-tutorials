/* Go - Returning Multiple Values From A Function */

package main

import "fmt"

func swap(x, y string) (string, string) {
   return y, x
}
func main() {
   a, b := swap("O'Regan", "Joe")
   fmt.Println(a, b)
}