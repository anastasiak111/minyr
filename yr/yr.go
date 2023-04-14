package yr

import (
        "os"
        //for open og lage filer
        "log"
        //for open logge feilmeldinger
        "io"
        //for open lese fil og skrive inn i fil
        "strings"
        //splitte linjer og string ned til elementer
        "strconv"
        //for open konvertere temperatur verdier til floats
        "encoding/csv"
        //csv package som lag w.write metoden funke
        "github.com/anastasiak111/funtemps/conv"
        //min funksjon for konvertering
        "fmt"
        //formatere F verdier til string og printe 
	"bufio"
	//til funksjonen AvarageTemp

)

func ProcessLines() {
        //leser filen
        src, err := os.Open("kjevik-temp-celsius-20220318-20230318.csv")
                if err != nil {
                log.Fatal(err)
                }
                defer src.Close() //sikrer at filen er stengt ved retur verdi

        // lager ny fil hvor nye returverier skal legges inn
        dst, err := os.Create("kjevik-temp-fahr-20220318-20230318.csv")
                if err != nil {
                log.Fatal(err)
                }
                defer dst.Close()


        w := csv.NewWriter(dst) //definere writer


        //buffer og linebuf er byte slices som holder verdiene
        var buffer []byte //variabelen leser 1 og 1 byte av input og lagrer verdier i seg selv
        var linebuf []byte // variabelen lagrer bytes fra hver linje mens de er lest, satt til "nil" som gjor at den ikke alokerer minne


        buffer = make([]byte, 1) //setter 1 byte for lesing
        bytesCount := 0 //startverdi
	
                        for {
                                _, err := src.Read(buffer) // metode for da lese data fra "src" og returnerer mengde bytes, til det blir tom for linjer
                                if err != nil && err != io.EOF {
                                        log.Fatal(err)
                                        }
                                        

                                        bytesCount++ //endrer verdi av bytescount med 1 hver gang en byte er lest
                                        if buffer[0] == 0x0A { //sjekker om det er en ny linje

                                        elementArray := strings.Split(string(linebuf), ";") //splitter opp stringen til elementer
                                        if len(elementArray) > 4 { //hvis storre enn 3

                                        celsius, err := strconv.ParseFloat(elementArray[3], 64) // 4. element (index 3) er parsed til float64
                                                
						if err != nil {
                                                log.Fatal(err)
                                                }

                                        fahr := conv.CelsiusToFarhenheit(celsius) //konverterer verdi
					
                                        elementArray[3] = fmt.Sprintf("%.1f", fahr) } //lager ny variabel "fahr" med F verdi med 2 desimaler og legger tilbake i 4. plassering (index 3)                                                
                                                if err := w.Write(elementArray); err != nil { //skriver inn den konverterte linjen i dst filen?                                              
                                                log.Fatal(err)
                                                }
 
                                

                                linebuf = nil //etter hver iterasjon linebuf settes til null for so gjenbruke buffer igjen

                                } else {
                                linebuf = append(linebuf, buffer[0]) // append funksjon legger lagret verdi fra slicebuf i slutten av hver linje                                 
                        }

                                if err == io.EOF {
                                break  //src.read returnerer feilmelding hvis ingen flere linjer to read, dermed loopen avsluttes
                                }
                        }

                // Flush any remaining writes to the output file
                w.Flush()


		// Replace the last line with "hi and goodbye"
	lastLinePos, err := dst.Seek(0, 2) // get the position of the last byte in the file
	if err != nil {
	    log.Fatal(err)
	}

	for i := 1; i <= len("\nhi and goodbye")*3; i++ {
	    _, err = dst.WriteAt([]byte{'\x00'}, lastLinePos-int64(i)) // overwrite the last line with null bytes
	    if err != nil {
	        log.Fatal(err)
	    }
	}

	_, err = dst.WriteAt([]byte("hi and goodbye\n"), lastLinePos-int64(len("\nhi and goodbye"))*3) // write the new last line
	if err != nil {
	    log.Fatal(err)
}


}


func AverageTemp() {
    src, err := os.Open("kjevik-temp-celsius-20220318-20230318.csv")
    if err != nil {
        log.Fatal(err)
    }
    defer src.Close()

    var sumCelsius float64
    var sumFahr float64
    var count int

    scanner := bufio.NewScanner(src)
    isFirstLine := true // flag to skip the first line
    for scanner.Scan() {
        if isFirstLine {
            isFirstLine = false
            continue
        }

	 fields := strings.Split(scanner.Text(), ";")
        if len(fields) >= 4 && fields[3] != ""{
            celsius, err := strconv.ParseFloat(fields[3], 64)
            if err != nil {
                log.Fatal(err)
            }
            sumCelsius += celsius

            fahr := conv.CelsiusToFarhenheit(celsius)
            sumFahr += fahr

            count++
        }
    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }

    avgCelsius := sumCelsius / float64(count)
    avgFahr := sumFahr / float64(count)


    scanner = bufio.NewScanner(os.Stdin) // declare new scanner for user input

    fmt.Println("Vennligst velg 'c' for Celsius eller 'f' for Fahrenheits (c/f):")
    scanner.Scan() // Read user input again
    input := scanner.Text()
    if input == "c" {
        fmt.Printf("Gjennomsnittlig temperatur i Celsius er %.1f\n", avgCelsius)
    } else if input == "f" {
        fmt.Printf("Gjennomsnittlig temperatur i Fahrenheit er %.1f\n", avgFahr)
    }
}
