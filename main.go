package main



import (
    "bufio"
    "os"
    "fmt"
    "log"
    "strconv"
    "strings"

"github.com/anastasiak111/funtemps/conv"
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
	
	//  åpne fil
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

		// Gjor endringer og lagrer nye linjer i en ny fil
			var newLines []string
			for i, line := range lines {
				if i == 0 {
					newLines = append(newLines, line)
					continue

				}
			}

		fields := strings.Split(line, ";")
				if len(fields) == 4 {
					celsius, err := strconv.ParseFloat(fields[3], 64)
					if err != nil {
						log.Fatal(err)
					}
					fahrenheit := conv.CelsiusToFahrenheit(celsius)
					fields[3] = fmt.Sprintf("%.2f", fahrenheit)
					newLine := strings.Join(fields, ";")
					newLines = append(newLines, newLine)
				} else {
					fmt.Printf("Error: line %d has %d fields: %v\n", i, len(fields), fields)
					newLines = append(newLines, line)
				}



            // flere else-if setninger
        } else {
            fmt.Println("Venligst velg convert, average eller exit:")
        }
    }
}

