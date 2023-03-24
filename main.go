package main

import (
    "bufio"
    "os"
    "fmt"
    "log"
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
            fmt.Println("Konverterer alle ligningene gitt i grader Celsius til grader Fahrenheit.")
            // funksjon som åpner fil, leser linjer, gjør endringer og lagrer nye linjer i en ny fil
	
	//opne fil
			file, err := os.Open("table.csv")
			if err != nil {
				log.Fatal(err)
			}
			defer file.Close()

	
	// Les linjer
			scanner := bufio.NewScanner(file)
			var lines []string
			for scanner.Scan() {
				line := scanner.Text()
				if len(line) == 0 {
					continue // skip empty lines
				}
				lines = append(lines, line)
				}

	

            // flere else-if setninger
        } else {
            fmt.Println("Venligst velg convert, average eller exit:")
        }
    }
}

