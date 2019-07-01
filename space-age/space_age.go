/*
Space Age Exercism

Using age in earth seconds calculates the age of someone on a planet

- Earth: orbital period 365.25 Earth days, or 31557600 seconds
- Mercury: orbital period 0.2408467 Earth years
- Venus: orbital period 0.61519726 Earth years
- Mars: orbital period 1.8808158 Earth years
- Jupiter: orbital period 11.862615 Earth years
- Saturn: orbital period 29.447498 Earth years
- Uranus: orbital period 84.016846 Earth years
- Neptune: orbital period 164.79132 Earth years

*/
package space

import "math"

type Planet string

func Age(s float64, p Planet) float64 {

	earth := s/31557600

  time := 0.00

  if p == "Earth" {
    time = earth
  } else if p == "Mercury" {
    time = earth / 0.2408467
  } else if p == "Venus" {
    time = earth / 0.61519726
  } else if p == "Mars" {
    time = earth / 1.8808158
  } else if p == "Jupiter" {
    time = earth / 11.862615
  } else if p == "Saturn" {
    time = earth / 29.447498
  } else if p == "Uranus" {
    time = earth / 84.016846
  } else if p == "Neptune" {
    time = earth / 164.79132
  }

  return math.Round(time*100)/100
}
