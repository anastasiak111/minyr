package main

import (
	"bufio"
	"fmt"
	"os"
	"github.com/anastasiak111/minyr/yr"
)

func main() {
    var input string
    scanner := bufio.NewScanner(os.Stdin)

    for scanner.Scan() {
        input = scanner.Text()
        if input == "q" || input == "exit" {
            fmt.Println("exit")
            os.Exit(0)
        } else if input == "convert" {
            fmt.Println("den gamle filen vil bli overskrevet, fortsette? (y/n)") //HUSK lage testen om den finnes
            scanner.Scan() // Read user input again
            confirm := scanner.Text()
            if confirm == "y" {
                fmt.Println("Konverterer alle malingene gitt i grader Celsius til grader Fahrenheit")
                yr.ProcessLines()
            } else if confirm == "n" {
                fmt.Println("Conversion cancelled.")
            } else {
                fmt.Println("Invalid input, try again")
            }
	}else if input == "average" {
               yr.AverageTemp()
        } else {
            fmt.Println("Venligst velg convert, average eller exit:")
        }
    }
}
