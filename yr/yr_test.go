package yr
import (
        "os"
        //for opne og lage filer
        "io"
        //for opne lese fil og skrive inn i fil
	"encoding/csv"
	//for a definere reader
	"testing"
	//tester
)

//tester antall linjer
func TestLineCount(t *testing.T) {
        f, err := os.Open("../kjevik-temp-fahr-20220318-20230318.csv")
        if err != nil {
                t.Errorf("Error opening file: %v", err)
        }
        defer f.Close()

        r := csv.NewReader(f)
	r.FieldsPerRecord = -1 // ignore the number of fields per record
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


