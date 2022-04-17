package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {

	usec := flag.Bool("c", false, "help message for flagname")
	used := flag.Bool("d", false, "help message for flagname")
	useu := flag.Bool("u", false, "help message for flagname")
	usei := flag.Bool("i", false, "help message for flagname")

	flag.Parse()

	//var str string
	var vvod int

	file, err := os.Open("test.txt")
	if err != nil {
		vvod = 1
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	type astr struct {
		st string
		n  int
	}

	var anstr []astr
	var vstr []string

	if vvod == 1 {

		q := 0
		vstr = append(vstr, "")
		fmt.Println("Введи строку, end является концом ввода")
		for vstr[q] != "end" {

			var a string
			fmt.Fscan(os.Stdin, &a)
			vstr = append(vstr, a)
			q++
		}

		for _, k := range vstr {
			fl := 0
			if len(anstr) == 0 {
				anstr = append(anstr, astr{"", 0})
				anstr[0].st = k
				anstr[0].n++
				anstr = append(anstr, astr{"", 0})
			} else {
				for i := 0; i < len(anstr); i++ {

					if anstr[i].st == k {

						anstr[i].n++
						fl++
						break
					}

				}
				if fl == 0 {
					anstr = append(anstr, astr{k, 1})
				}

			}
		}

	} else {

		for scanner.Scan() {
			fl := 0
			if len(anstr) == 0 {
				anstr = append(anstr, astr{"", 0})
				anstr[0].st = scanner.Text()
				anstr[0].n++
				anstr = append(anstr, astr{"", 0})
			} else {
				for i := 0; i < len(anstr); i++ {

					if anstr[i].st == scanner.Text() {

						anstr[i].n++
						fl++
						break
					}

				}
				if fl == 0 {
					anstr = append(anstr, astr{scanner.Text(), 1})
				}

			}
		}
		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}
	}

	if *usec {
		for i := 0; i < len(anstr); i++ {
			fmt.Println(anstr[i].n, " ", anstr[i].st)
		}
	} else if *used {
		for i := 0; i < len(anstr); i++ {
			fmt.Println(anstr[i].st)
		}
	} else if *useu {
		for i := 0; i < len(anstr); i++ {
			if anstr[i].n == 1 {
				fmt.Println(anstr[i].st)
			}
		}
	} else if *usei {
		var anstr1 []astr
		anstr1 = append(anstr1, astr{"", 0})
		for i := 0; i < len(anstr); i++ {
			anstr[i].st = strings.ToLower(anstr[i].st)
		}
		var n_str_an []string

		for i := 0; i < len(anstr); i++ {
			anstr1 = append(anstr1, astr{"", 0})

			for j := 0; j < len(anstr); j++ {
				var fl int
				if anstr[i].st == anstr[j].st && i == j {
					for _, k := range n_str_an {
						if k == anstr[i].st {
							fl++
						}
					}
					if fl == 0 {
						n_str_an = append(n_str_an, anstr[i].st)

						anstr1[i].st = anstr[i].st
						anstr1[i].n = anstr[i].n
					}

				} else if anstr[i].st == anstr[j].st && i != j && fl != 0 {

					anstr1[i].n = anstr[i].n + anstr[j].n
				}

			}

		}
		for i := 0; i < len(anstr); i++ {
			if anstr1[i].n > 0 {
				fmt.Println(anstr1[i].st)
			}
		}

	} else {
		for i := 0; i < len(anstr); i++ {
			fmt.Println(anstr[i].st)
		}

	}
	flag.Parse()
}
