package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strconv"

	"github.com/pkg/errors"
)

func main() {
	log.Println(reg.String())
	foo()
}

func foo() {
	path, err := filepath.Abs("./upstart.log.20180410_2")
	if err != nil {
		log.Fatal(errors.Wrapf(err, "file path abs failed"))
	}
	f, err := os.OpenFile(path, os.O_RDONLY, os.ModePerm)
	if err != nil {
		log.Fatal(errors.Wrapf(err, "open file failed"))
	}
	defer f.Close()
	reader := bufio.NewReader(f)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatal(err)
		}
		bar(line)
	}
	// log.Println("done.")
}

var reg = regexp.MustCompile(`^.*subsidy user period (?P<uuid>\w+), recharge (?P<recharge>\d+) and deduct (?P<deduct>[0-9\-]+)`)

func bar(str string) {
	if reg.MatchString(str) {
		groups := reg.FindStringSubmatch(str)
		uuid := groups[1]
		recharge, err := strconv.ParseInt(groups[2], 10, 64)
		if err != nil {
			log.Fatal(err)
		}
		deduct, err := strconv.ParseInt(groups[3], 10, 64)
		if err != nil {
			log.Fatal(err)
		}
		// log.Printf("uuid = %s, recharge = %d, deduct = %d, amountToChange = %d", uuid, recharge, deduct, -(deduct + recharge))
		fmt.Println(fmt.Sprintf(`"%s": %d,`, uuid, -(deduct + recharge)))
	}
}
