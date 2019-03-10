/* 
	Go - Functions As Values
	
	Functions can be created on the fly and can be used as values.
*/

package main

import ("fmt"; "math")	// need ;

func main(){
   /* declare a function variable */
   getSquareRoot := func(x float64) float64 {
      return math.Sqrt(x)
   }

   /* use the function */
   fmt.Println(getSquareRoot(9))
}