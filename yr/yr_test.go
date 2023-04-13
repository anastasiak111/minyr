package test
import (
        "os"
        //for open og lage filer
        "io"
        //for open lese fil og skrive inn i fil
	"encoding/csv"
	//for a definere reader
	"testing"
	//tester 
)

func TestFahrtableLineCount(t *testing.T) {
        f, err := os.Open("kjevik-temp-fahr-20220318-20230318.csv")
        if err != nil {
                t.Errorf("Error opening file: %v", err)
        }
        defer f.Close()

        r := csv.NewReader(f)
        lineCount := 0
        for {  
                _, err := r.Read()
                if err == io.EOF {
                        break
                } else if err != nil {
                        t.Errorf("Error reading file: %v", err)
                }
                lineCount++
        }

        expectedLineCount := 16756 // update this with the expected number of lines
        if lineCount != expectedLineCount {
                t.Errorf("Incorrect line count. Got %d, expected %d", lineCount, expectedLineCount)
        }
}
