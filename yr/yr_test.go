package yr
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

//tester antall linjer 
func TestLineCount(t *testing.T) {
        f, err := os.Open("")
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




func TestConversion(t *testing.T) {
	type test struct {
		input string
		want  string
	}

	tests := []test{
		{input: "Kjevik;SN39040;18.03.2022 01:50;6", want: "Kjevik;SN39040;18.03.2022 01:50;42.8"},
		{input: "Kjevik;SN39040;07.03.2023 18:20;0", want: "Kjevik;SN39040;07.03.2023 18:20;32.0"},
		{input: "Kjevik;SN39040;08.03.2023 02:20;-11", want: "Kjevik;SN39040;08.03.2023 02:20;12.2"},
	} 

	for _, tc := range tests {
		got := kjevik-temp-fahr-20220318-20230318.csv(tc.input)
		if !withinTolerance(tc.want, got, 1e-2) {
			t.Errorf("expected: %v, got: %v", tc.want, got)
		}
	}
}


