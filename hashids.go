package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/speps/go-hashids"
)

const secret = "Df3FAnyzwoNUGQ6Rd63wBwURUBGesqs9YNMNHPLnVkLyk77qaV"

func encode(minLength int, ids ...int64) (string, error) {
	data := hashids.NewData()
	data.Salt = secret
	data.Alphabet = "abcdefghijklmnopqrstuvwxyz1234567890"
	if minLength > 0 {
		data.MinLength = minLength
	}
	hashid := hashids.NewWithData(data)
	encoded, err := hashid.EncodeInt64(ids)
	return encoded, err
}

func decode(str string) ([]int64, error) {
	data := hashids.NewData()
	data.Salt = secret
	data.Alphabet = "abcdefghijklmnopqrstuvwxyz1234567890"
	return hashids.NewWithData(data).DecodeInt64WithError(str)
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("no args")
		return
	}
	for _, str := range os.Args[1:] {
		id, err := strconv.ParseInt(str, 10, 64)
		if err == nil {
			encoded, err := encode(15, id)
			if err != nil {
				fmt.Printf("%s encode failed, err = %s", str, err.Error())
				return
			}
			fmt.Println(encoded)
		} else {
			decoded, err := decode(str)
			if err != nil {
				fmt.Printf("%s decode failed, err = %s", str, err.Error())
			}
			for _, d := range decoded {
				fmt.Println(d)
			}
		}
	}
}
