/*
Convert a number to a string, the contents of which depend on the number's factors.

- If the number has 3 as a factor, output 'Pling'.
- If the number has 5 as a factor, output 'Plang'.
- If the number has 7 as a factor, output 'Plong'.
- If the number does not have 3, 5, or 7 as a factor,
  just pass the number's digits straight through.
*/

package raindrops

// import "math"
import "strconv"

func Convert(num int) string {

  result := ""

  if num % 3 == 0 {
    result += "Pling"
  }
  if num % 5 == 0 {
    result += "Plang"
  }
  if num % 7 == 0  {
    result +=  "Plong"
  }
  if result == ""{
    result += strconv.Itoa(num)
  }

  return result
}
