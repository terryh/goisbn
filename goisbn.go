package goisbn

import (
	"errors"
	"strconv"
	"strings"
)

//type ISBN struct {
//Isbn string
//}
type ISBN string

const (
	dash         = "-"
	space        = " "
	Isbn13Prefix = "978"
)

func CalculateCheckSum(isbnNoCheckSum string) string {

	isbnLen := len(isbnNoCheckSum)
	var sum = 0

	// ISBN 10
	if isbnLen == 9 {
		for i := 0; i < 9; i++ {
			toInt, _ := strconv.Atoi(string(isbnNoCheckSum[i]))
			sum += (10 - i) * toInt
		}

		countResult := 11 - sum%11
		if countResult == 10 {
			return "X"
		} else if countResult == 11 {
			return "0"
		} else {
			return strconv.Itoa(countResult)
		}
		// ISBN 13
	} else {

		for i := 0; i < 12; i++ {

			toInt, _ := strconv.Atoi(string(isbnNoCheckSum[i]))
			if i%2 == 0 {
				sum += toInt
			} else {

				sum += toInt * 3
			}
		}

		countResult := 10 - sum%10
		if countResult == 10 {
			return "0"
		} else {
			return strconv.Itoa(countResult)
		}
	}

}

func Cleanup(isn string) (string, error) {

	if strings.Contains(isn, dash) {
		isn = strings.Replace(isn, dash, "", -1)
	}
	if strings.Contains(isn, space) {
		isn = strings.Replace(isn, space, "", -1)
	}

	// zero fill
	if len(isn) == 9 {
		isn = "0" + isn
	}

	// 10 or 13
	if len(isn) != 10 && len(isn) != 13 {
		return "", errors.New("ISBN must be either 10 or 13 characters long")
	}

	//except last char, must be digit
	digitString := isn[:len(isn)-1]
	_, err := strconv.Atoi(digitString)

	if err != nil {
		return "", errors.New("ISBN must be digit without checksum")
	}

	return isn, nil
}

// translate isbn 10 to isbn 13, or reverse
func Convert(isn string) (string, error) {
	var isbn string
	if len(isn) == 10 {
		// convert to ISBN 13
		isbn = Isbn13Prefix + string(isn[0:9])
		return isbn + CalculateCheckSum(isbn), nil

	} else {
		// convert to ISBN 10
		if strings.HasPrefix(isn, Isbn13Prefix) {
			isbn = string(isn[3:12])
			return isbn + CalculateCheckSum(isbn), nil
		} else {
			return "", errors.New("Only ISBN13 with 978 Bookland code can be convert to ISBN10")

		}

	}

}

func ToISBN(isbn string) (ISBN, error) {
	var toisbn ISBN
	isn, err := Cleanup(isbn)

	if err != nil {
		return toisbn, err
	}

	toisbn = ISBN(isn)
	isbnLen := len(isn)

	// ISBN 10
	if isbnLen == 10 {
		if string(isn[9]) == CalculateCheckSum(string(isn[0:9])) {
			return toisbn, nil
		} else {
			return toisbn, errors.New("ISBN checksum wrong")
		}
		// ISBN 13
	} else {
		if string(isn[12]) == CalculateCheckSum(string(isn[0:12])) {
			return toisbn, nil
		} else {

			return toisbn, errors.New("ISBN checksum wrong")
		}
	}
}

//Get the ISBN 10 string
func (isbn ISBN) ISBN10() string {
	// 10
	if len(isbn) == 10 {
		return string(isbn)
	} else {
		isn, err := Convert(string(isbn))
		if err == nil {
			return isn
		}
	}
	return ""
}

func (isbn ISBN) ISBN13() string {
	// 13
	if len(isbn) == 13 {
		return string(isbn)
	} else {
		isn, err := Convert(string(isbn))
		if err == nil {
			return isn
		}
	}
	return ""
}
