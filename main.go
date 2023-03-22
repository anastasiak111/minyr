package main 

import (
"os"
"log"
)

Func main () {
Scr,err := os.Open( “home/anastasiak111/minyr/kjevik-temp-celsius-20220318-20230318.csv”)
If err != nil {
Log.Fatal(err)
}
Defer src.Close()
Log.Println(scr)

Var buffer []byte 
Buffer = make ([]byte, 1)
For ; n != 0{
n, err = scr.Read (buffer (I retur))
if err != nil {
log.Fatal(err)
}
}
Log. Println(string(buffer([:n]))
}

