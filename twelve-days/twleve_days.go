/*
Exercism twleve days track

Contains:

  Verse
  Song
*/

package twelve


func Verse(verse int) string{
  /*
  Verse
    returns a string for the corrosponding verse in 'twleve days of christmas'

    param: verse int
    return: s string
  */

  s := ""
  if verse == 1 {
    s = "On the first day of Christmas my true love gave to me: a Partridge in a Pear Tree."
  } else if verse == 2 {
    s = "On the second day of Christmas my true love gave to me: two Turtle Doves, and a Partridge in a Pear Tree."
  } else if verse == 3 {
    s = "On the third day of Christmas my true love gave to me: three French Hens, two Turtle Doves, and a Partridge in a Pear Tree."
  } else if verse == 4 {
    s = "On the fourth day of Christmas my true love gave to me: four Calling Birds, three French Hens, two Turtle Doves, and a Partridge in a Pear Tree."
  } else if verse == 5 {
    s = "On the fifth day of Christmas my true love gave to me: five Gold Rings, four Calling Birds, three French Hens, two Turtle Doves, and a Partridge in a Pear Tree."
  } else if verse == 6 {
    s = "On the sixth day of Christmas my true love gave to me: six Geese-a-Laying, five Gold Rings, four Calling Birds, three French Hens, two Turtle Doves, and a Partridge in a Pear Tree."
  } else if verse == 7 {
    s = "On the seventh day of Christmas my true love gave to me: seven Swans-a-Swimming, six Geese-a-Laying, five Gold Rings, four Calling Birds, three French Hens, two Turtle Doves, and a Partridge in a Pear Tree."
  } else if verse == 8 {
    s = "On the eighth day of Christmas my true love gave to me: eight Maids-a-Milking, seven Swans-a-Swimming, six Geese-a-Laying, five Gold Rings, four Calling Birds, three French Hens, two Turtle Doves, and a Partridge in a Pear Tree."
  } else if verse == 9 {
    s = "On the ninth day of Christmas my true love gave to me: nine Ladies Dancing, eight Maids-a-Milking, seven Swans-a-Swimming, six Geese-a-Laying, five Gold Rings, four Calling Birds, three French Hens, two Turtle Doves, and a Partridge in a Pear Tree."
  } else if verse == 10 {
    s = "On the tenth day of Christmas my true love gave to me: ten Lords-a-Leaping, nine Ladies Dancing, eight Maids-a-Milking, seven Swans-a-Swimming, six Geese-a-Laying, five Gold Rings, four Calling Birds, three French Hens, two Turtle Doves, and a Partridge in a Pear Tree."
  } else if verse == 11 {
    s = "On the eleventh day of Christmas my true love gave to me: eleven Pipers Piping, ten Lords-a-Leaping, nine Ladies Dancing, eight Maids-a-Milking, seven Swans-a-Swimming, six Geese-a-Laying, five Gold Rings, four Calling Birds, three French Hens, two Turtle Doves, and a Partridge in a Pear Tree."
  } else if verse == 12 {
    s = "On the twelfth day of Christmas my true love gave to me: twelve Drummers Drumming, eleven Pipers Piping, ten Lords-a-Leaping, nine Ladies Dancing, eight Maids-a-Milking, seven Swans-a-Swimming, six Geese-a-Laying, five Gold Rings, four Calling Birds, three French Hens, two Turtle Doves, and a Partridge in a Pear Tree."
  }
  return(s)
}

func Song() string{
  /*
  Verse
    returns a string of the whole song 'twleve days of christmas'

    return: s string
  */
  expected := ""

  for i := 1; i <= 12; i++ {
    expected += Verse(i) + "\n"
  }
  return expected
}
