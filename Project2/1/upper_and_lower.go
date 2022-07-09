package main

import (
	"fmt"
	"unicode"
)

func Detect_upper_and_lower(s string)(lower_case []rune, upper_case []rune){
        for _ , c := range(s){
                if unicode.IsLower(c){
                        lower_case = append(lower_case,c)
                }else if  unicode.IsUpper(c){
                        upper_case = append(upper_case, c)
                }else{
                        continue
                }
        }
        return

}

func main() {
        massage := "MaktabKhooneh is The Best platform for learning :) "
        
        lower_case, upper_case := Detect_upper_and_lower(massage)

        fmt.Printf("\n================= Lower case ====================== \n")
        for _, l :=range lower_case{
                fmt.Printf("[%c] ", rune(l))
        }
        fmt.Printf("\n\n================= Upper case ====================== \n")
        for _, u :=range upper_case{
                fmt.Printf("[%c] ", rune(u))
        }
        fmt.Printf("\n")
}