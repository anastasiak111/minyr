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
            fmt.Println("den gamle filen vil bli overskrevet, fortsette? (y/n)")
            scanner.Scan() // Read user input again
            confirm := scanner.Text()
            if confirm == "y" {
                fmt.Println("Konverterer alle malingene gitt i grader Celsius til grader Fahrenheit")
                // funksjon som opner fil, leser linjer, gjor endringer og lagrer nye linjer i en ny fil
                yr.ProcessLines()
            } else if confirm == "n" {
                fmt.Println("Conversion cancelled.")
            } else {
                fmt.Println("Invalid input, try again")
            }
        } else if input == "avarage" {
		 reader := bufio.NewReader(os.Stdin)

            fmt.Println("Press 'c' to calculate the average temperature in Celsius, or 'f' to calculate it in Fahrenheit:")
            unit, _ := reader.ReadString('\n')
            unit = strings.TrimSpace(unit)

            if unit == "c" {
                yr.ProcessLines("kjevik-temp-celsius-20220318-20230318.csv", "kjevik-temp-celsius-20220318-20230318-converted.csv", conv.AverageCelsius)
            } else if unit == "f" {
                yr.ProcessLines("kjevik-temp-fahr-20220318-20230318.csv", "kjevik-temp-fahr-20220318-20230318-converted.csv", conv.AverageFahrenheit)
            } else {
                fmt.Println("Invalid input. Please press 'c' or 'f' to calculate the average temperature.")
            }
        } else {
            fmt.Println("Venligst velg convert, average eller exit:")
        }
    }
}
