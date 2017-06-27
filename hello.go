package main

//import "fmt"
import "log"

func main() {
    dbConnect := DbConnect{}

    db := dbConnect.connect()

    rows, err := db.Query("SELECT * FROM usuarios_clientes")
    if err != nil {
        log.Fatal(err)
    }

    results := dbConnect.getResult(rows)

    for _, result := range results {
        log.Print(result["id"])
    }

}

