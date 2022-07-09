package main

import (
	"fmt"
	"unicode"
)

func Detect_letter_and_digit(s string)(letter []rune, digit []rune){
        for _ , c := range(s){
                if unicode.IsLetter(c){
                        letter = append(letter,c)
                }else if  unicode.IsDigit(c){
                        digit = append(digit, c)
                }else{
                        continue
                }
        }
        return

}

func main() {
        massage := " my phone is 09115980792 and my code is 20 "
        
        letters, digits := Detect_letter_and_digit(massage)

        fmt.Printf("\n================= Letters ====================== \n")
        for _, l :=range letters{
                fmt.Printf("[%c] ", rune(l))
        }
        fmt.Printf("\n\n================= Digits ====================== \n")
        for _, u :=range digits{
                fmt.Printf("[%c] ", rune(u))
        }
        fmt.Printf("\n")
}