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
	r.FieldsPerRecord = -1 // ignorer antall felter, dermed siste linjen
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

        expectedLineCount := 16756 // antall linjer forventet
        if lineCount != expectedLineCount {
                t.Errorf("Incorrect line count. Got %d, expected %d", lineCount, expectedLineCount)
        }
}



// Tester konvertering av linjer
func TestProcessLines(t *testing.T) {
        type test struct {
        input string
        want string
	err error
        }

        tests := []test{
        {input: "Kjevik;SN39040;18.03.2022 01:50;6", want: "Kjevik;SN39040;18.03.2022 01:50;42.8"},
        {input: "Kjevik;SN39040;07.03.2023 18:20;0", want: "Kjevik;SN39040;07.03.2023 18:20;32.0"},
        {input: "Kjevik;SN39040;08.03.2023 02:20;-11", want: "Kjevik;SN39040;08.03.2023 02:20;12.2"},
        }

		for _, tc := range tests {
		want := tc.want
		got, err := ProcessLines()
		if err != nil && err.Error() != tc.err.Error() {
			t.Errorf("unexpected error: got %v, want %v", err, tc.err)
		}
		if got != want {
			t.Errorf("unexpected output: got %s, want %s", got, tc.want)
		}
		}

}

