/*
Matrix package for the matrix challenge on the exercism GO track
*/

package matrix

import (
    "strings"
    "strconv"
    "errors"
  )

type matrix [][]int

/*
Used to create a matrix ([][]int) from a string
parameters
  strMartix string:
returns
  martix type
*/
func New(strMatrix string) (matrix, error){

  var err error

  rows := strings.Split(strMatrix, "\n")

  matrix :=make([][]int, len(rows))

  for i , row := range rows {
    row = strings.TrimSpace(row)
    numbers := strings.Split(row, " ")
    matrix[i] = make([]int, len(numbers))

    if 0 < i {
      if len(matrix[i]) != len(matrix[i-1]){
        return matrix, errors.New("Matrix rows are uneven")
      }
    }

    for j, num := range numbers {
      var integer int
      integer, err = strconv.Atoi(num)
      if err != nil {
        return matrix, err
      } else {
        matrix[i][j] = integer
      }
    }
  }
  return matrix, err
}

/*
Returns a list of the Martix's rows ([][]int)
reading each row left-to-right top-to bottom
*/
func (m matrix) Rows() ([][]int) {
  duplicate := make([][]int, len(m))

  for i := range m {
    duplicate[i] = make([]int, len(m[i]))
    copy(duplicate[i], m[i])
  }
  return duplicate
}

/*
Returns a list ([][]int) of the martix's columns
reading each column from top-to-bottom while moving left-to-right
*/
func (m matrix) Cols() ([][]int) {

  result := make([][]int, len(m[0]))

  for i := range result {
    result[i] = make([]int, len(m))
    for j := range result[i]{
      result[i][j] = m[j][i]
    }
  }
  return result
}

/*
Sets a martix element
*/
func (m matrix) Set(row int, col int, value int) bool{
  ok := true

  //Validate row and col values
  if len(m) > row && row >= 0 {
    if len(m[row]) > col && col >= 0{
      m[row][col] = value
    } else {
      return !ok
    }
  } else {
    return !ok
  }
  return ok
}
