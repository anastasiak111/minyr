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

)

func main() {
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
                                        log.Println(string(linebuf)) //hvis ny linje gjor denne linjen til en string

                                        elementArray := strings.Split(string(linebuf), ";") //splitter opp stringen til elementer
                                        if len(elementArray) > 3 { //hvis storre enn 3
                                        celsius, err := strconv.ParseFloat(elementArray[3], 64) // 4. element (index 3) er parsed til float64
                                                if err != nil {
                                                                                                
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

}

