/*
Calculates the difference between two strings

For example:

GAGCCTACTAACGGGAT
CATCGTAATGACGGCCT
^ ^ ^  ^ ^    ^^

The Hamming distance between these two DNA strands is 7.
*/

package hamming

import "errors"

func Distance(a, b string) (int, error) {

  if len(a) != len(b) {
    return -1, errors.New("DNA Strands are not of equal length")
  }

  hammingDistance := 0

  for i := 0; i < len(a); i++ {
    if a[i] != b[i] {
      hammingDistance += 1
    }
  }
  return hammingDistance, nil
}
