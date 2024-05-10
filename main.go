package main

import (
	"fmt"
	"os"
	"path"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	myargs := os.Args[1:]

	/// read the Files & hundlle the reding file Errors ///

	if len(myargs) < 1 {
		fmt.Println("There is no File")
		return
	}

	text, err := os.ReadFile(myargs[0])

	if len(myargs) > 2 {
		fmt.Println("You enter more then tree argumments!")
		return
	}

	if len(myargs) < 2 {
		fmt.Println("You need to enter the result file!")
		return
	}

	if err != nil {
		fmt.Println("Error!", myargs[0], "file not found !")
		return
	}

	if path.Ext(myargs[1]) != ".txt" {
		fmt.Println("Error!! result file Extension must be (.txt) !!!")
		return
	}

	words := strings.Fields(string(text))

	for i := 0; i < len(words); i++ {

		if words[i] == "(cap)" {
			if i-1 >= 0 {
				words[i-1] = Capitalize(words[i-1])
			}
			words = append(words[:i], words[i+1:]...)
			i--

		} else if words[i] == "(up)" {
			if i-1 >= 0 {
				words[i-1] = strings.ToUpper(words[i-1])
			}
			words = append(words[:i], words[i+1:]...)
			i--
		} else if words[i] == "(low)" {
			if i-1 >= 0 {
				words[i-1] = strings.ToLower(words[i-1])
			}

			words = append(words[:i], words[i+1:]...)
			i--
		} else if words[i] == "(hex)" {
			if i-1 >= 0 {
				words[i-1] = hexadecimaltoInt(words[i-1])
			}
			words = append(words[:i], words[i+1:]...)
			i--
		} else if words[i] == "(bin)" {
			if i-1 >= 0 {
				words[i-1] = binarytoInt(words[i-1])
			}

			words = append(words[:i], words[i+1:]...)

			/// hundlling cap/up/low with Flags ////

		} else if words[i] == "(up," {
			if i+1 < len(words) {
				if isNumbValid(words[i+1]) {
					v, _ := strconv.Atoi(string(words[i+1][:len(words[i+1])-1]))
					if v < 0 {
						fmt.Println("Error! You Enter a Negative Number!!")
					}
					for h := 0; h < v; h++ {
						if i-h > 0 {
							words[i-1-h] = strings.ToUpper(words[i-1-h])
						}
					}
					words = append(words[:i], words[i+2:]...)
					i--
				}
			}

		} else if words[i] == "(low," {

			if i+1 < len(words) {
				if isNumbValid(words[i+1]) {
					v, _ := strconv.Atoi(string(words[i+1][:len(words[i+1])-1]))
					if v < 0 {
						fmt.Println("Error! You Enter a Negative Number!!")
					}
					for h := 0; h < v; h++ {
						if i-h > 0 {
							words[i-1-h] = strings.ToLower(words[i-1-h])
						}
					}
					words = append(words[:i], words[i+2:]...)
					i--
				}

			}

		} else if words[i] == "(cap," {
			if i+1 < len(words) {
				if isNumbValid(words[i+1]) {

					v, _ := strconv.Atoi(string(words[i+1][:len(words[i+1])-1]))

					if v < 0 {
						fmt.Println("Error! You Enter a Negative Number!!")
					}
					for h := 0; h < v; h++ {
						if i-h > 0 {
							words[i-1-h] = Capitalize(words[i-1-h])
						}
					}
					words = append(words[:i], words[i+2:]...)
					i--
				}
			}
		}

		/// hndelling the vowels ///

		checkVowels(words)
	}

	forPunctuations := SliceToString(words)
	newSlice := strings.Split(forPunctuations, " ")

	//// punctuation position ; : : , . ///

	toString := strings.Join(newSlice, " ")

	///  For The Concatenat Ponctuations ///

	pattern := regexp.MustCompile(`(\w[.?!:;,]+)(\w)`)
	toString = pattern.ReplaceAllString(toString, "$1 $2")
	newSlice = strings.Split(toString, " ")

	for i := 0; i < len(newSlice); i++ {
		for strings.Index(newSlice[i], ".") == 0 || strings.Index(newSlice[i], ",") == 0 ||
			strings.Index(newSlice[i], "!") == 0 || strings.Index(newSlice[i], "?") == 0 ||
			strings.Index(newSlice[i], ":") == 0 || strings.Index(newSlice[i], ";") == 0 {
			newSlice = Punctuations(newSlice, i)

			// if punc is inn the first ///
			if i == 0 {
				break
			}
		}
	}

	result := SliceToString(newSlice)

	rr := SingleQuote(result)

	os.WriteFile(myargs[1], []byte(rr), 0777)

} /////////////////// the end of the main function ////////////////////////////////

/// slice to string ///

func SliceToString(mySlice []string) string {
	myString := ""
	for i := 0; i < len(mySlice); i++ {
		if mySlice[i] != "" {
			myString += mySlice[i]
		}
		if i == len(mySlice)-1 {
			break
		}
		//if i != len -1 && mySlice[i] != ""
		if mySlice[i] != "" && mySlice[i] != " " {
			myString += " "
		}
	}
	return myString
}

/// hundlling Capitalize ///

func Capitalize(s string) string {
	vl := 0
	new := []rune(s)
	for range s {
		vl++
	}
	for i := 0; i < vl; i++ {
		if new[i] >= 'A' && new[i] <= 'Z' {
			new[i] = (new[i] + 32)
		}
		if new[0] >= 'a' && new[0] <= 'z' {
			new[0] = (new[0] - 32)
		}
		/// Check if the previous character is not an apostrophe
		if i > 0 && new[i-1] != '\'' { 
			if (new[i-1] > 'Z' && new[i-1] < 'a') || (new[i-1] < 'A' && new[i-1] > '9') || new[i-1] < '0' || new[i-1] > 'z' {
				if new[i] >= 'a' && new[i] <= 'z' {
					new[i] = (new[i] - 32)
				}
			}
		}
	}
	str := string(new)
	return str
}

func isNumbValid(s string) bool {
	if s[len(s)-1:] == ")" {
		for _, char := range s {
			if !(char >= '0' && char <= '9') && char != ')' && !(char < '0') {
				return false
			}
		}
	}
	return true
}

/// Hex to int ///

func hexadecimaltoInt(hex string) string {
	nbr, err := strconv.ParseInt(hex, 16, 64) //
	//check(err)
	if err != nil {
		fmt.Println("You Enter Invalid input!")
		return hex
	}
	return fmt.Sprint(nbr)
}

/// Bin to int ///

func binarytoInt(bin string) string {
	nbr, err := strconv.ParseInt(bin, 2, 64)
	//check(err)
	if err != nil {
		fmt.Println("You Enter Invalid input!")
		return bin
	}
	return fmt.Sprint(nbr)
}

/// hundlling ponctuation ///

func Punctuations(slices []string, i int) []string {

	if i > 0 {
		if strings.Index(slices[i], ".") == 0 {
			slices[i-1] = slices[i-1] + "."
			slices[i] = slices[i][1:]
		} else if strings.Index(slices[i], "!") == 0 {
			slices[i-1] = slices[i-1] + "!"
			slices[i] = slices[i][1:]
		} else if strings.Index(slices[i], ",") == 0 {
			slices[i-1] = slices[i-1] + ","
			slices[i] = slices[i][1:]
		} else if strings.Index(slices[i], "?") == 0 {
			slices[i-1] = slices[i-1] + "?"
			slices[i] = slices[i][1:]
		} else if strings.Index(slices[i], ":") == 0 {
			slices[i-1] = slices[i-1] + ":"
			slices[i] = slices[i][1:]
		} else if strings.Index(slices[i], ";") == 0 {
			slices[i-1] = slices[i-1] + ";"
			slices[i] = slices[i][1:]
		}
	}

	kk := SliceToString(slices)
	newSlice1 := strings.Split(kk, " ")
	return newSlice1
}

/// hundel the Vowels ///

func checkVowels(mySlice []string) {
	vowels := "aeoiuhAEOIUH"

	for i := 0; i < len(mySlice); i++ {
		for j := 0; j < len(vowels); j++ {

			// check if the first char is a vowel
			if i > 0 && mySlice[i][0] == vowels[j] {
				if string(mySlice[i-1]) == "A" || string(mySlice[i-1]) == "a" {
					mySlice[i-1] += "n"
				}
				if len(mySlice[i-1]) == 2 {
					if string(mySlice[i-1][1]) == "a" {
						if string(mySlice[i-1][0]) == "(" || string(mySlice[i-1][0]) == "\"" || string(mySlice[i-1][0]) == "'" || string(mySlice[i-1][0]) == "{" || string(mySlice[i-1][0]) == "[" {
							mySlice[i-1] += "n"
						}
					}
				}
			}
		}
	}
}

/// hundlling the Apostroph ///

func SingleQuote(value string) string {

	myStr := stringToRune(value)
	mySlice := strings.Fields(myStr)

	//SingleQuote position ///
	firstQuote := false
	lastQuote := false
	openIndex := -1

	for i := 0; i < len(mySlice); i++ {
		// check if the singleQuote id concatenate ///

		if i < len(mySlice)-1 && mySlice[i] == "'" && !firstQuote && mySlice[i+1] != "'" {
			firstQuote = true
			openIndex = i

		} else if mySlice[i] == "'" && firstQuote && !lastQuote {
			lastQuote = true
			mySlice[openIndex+1] = mySlice[openIndex] + mySlice[openIndex+1]
			mySlice[i-1] += mySlice[i]
			mySlice[openIndex] = ""
			mySlice[i] = ""

			firstQuote = false
			lastQuote = false
		}
	}
	return SliceToString(mySlice)
}

/// convert string to rune ///

func stringToRune(s string) string {
	myRune := []rune(s)
	new := []rune{}

	for i, v := range myRune {
		///check if v is a single quote // d'dont
		if v == '\'' && i-1 >= 0 && i+1 < len(myRune) && unicode.IsLetter(myRune[i-1]) && unicode.IsLetter(myRune[i+1]) {
			new = append(new, v)
		} else if v == '\'' {
			new = append(new, ' ')
			new = append(new, v)
			new = append(new, ' ')
		} else {
			new = append(new, v)
		}
	}
	return string(new)
}
