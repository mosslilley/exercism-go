/*
If the given name is "Alice", the result should be "One for Alice, one for me."
If no name is given, the result should be "One for you, one for me."
*/

package twofer

// ShareWith should have a comment documenting it.
func ShareWith(name string) string {

	if name == ""{
		return "One for you, one for me."
	}else {
		return "One for " + name + ", one for me."
	}
}
