package conv

import (
    "os"
    "log"
    "strconv"
    "strings"
    "github.com/anastasiak111/funtemps/conv"
)
// åpner og leser fil en, mulig feil byte atm
func main() {
    src,err := os.Open("^`^|home/anastasiak111/minyr/kjevik-temp-celsius-20220318-202")
    if err != nil {
        log.Fatal(err)
    }
    defer src.Close()
    log.Println(src)

    var buffer []byte
    buffer = make([]byte, 1)
    var n int
    for ; n != 0; {
        n, err = src.Read(buffer)
        if err != nil {
            log.Fatal(err)
        }
    }
    log.Println(string(buffer[:n]))
}

// Gjør endringer og lagrer nye linjer i en ny fil
	var newLines []string
	for i, line := range lines {
		if i == 0 {
			newLines = append(newLines, line)
			continue

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
	}
