package main

import (
	"fmt"
	"net/http"
	"log"
	"io"
	"encoding/json"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

type SensorData struct {
	Light       float32 `json:"light"`
	Temperature float32 `json:"temperature"`
	Humidity    float32 `json:"humidity"`
}

func handleData(w http.ResponseWriter, r *http.Request) {
	// fmt.Printf("Received data from sensor!\n")

	defer r.Body.Close()
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	// fmt.Printf("Body: \n'%s'\n\n", body)

	var data SensorData
	err = json.Unmarshal(body, &data);
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("New sensor data %+v\n", data)
}

func main() {
	fmt.Printf("Sveika, pasaule!\n")

	http.HandleFunc("/sensor/data", handleData)
	go func() {
		log.Fatal(http.ListenAndServe(":5000", nil))
	}()

	db, err := sql.Open("mysql", "root:root@tcp(10.152.33.88:3307)/test")
	if err != nil	{
		log.Fatal(err)
	} else {
		fmt.Printf("All is ok\n")
	}
	defer db.Close()

	for {
		
	}
}
