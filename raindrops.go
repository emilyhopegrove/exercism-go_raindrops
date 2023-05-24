package raindrops
import "strconv"

//If number % 3 == 0, concatenate 'Pling' to result
//If number % 5 == 0, concatenate 'Plang' to result
//If number % 7 == 0, concatenate 'Plong' to the result
//else, return number

func Convert(number int) string {
	result := ""

	if number % 3 == 0 {
		result += "Pling"
	}
	if number % 5 == 0 {
		result += "Plang"
	}
	if number % 7 == 0 {
		result += "Plong"
	 } 
	if result == "" {
		return strconv.Itoa(number)
	}

	return result
	
	
}
