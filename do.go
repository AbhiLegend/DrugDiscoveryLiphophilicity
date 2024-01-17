package main

import (
    "bufio"
    "bytes"
    "encoding/json"
    "fmt"
    "io/ioutil"
    "net/http"
    "os"
)

type PredictionRequest struct {
    Smiles string `json:"smiles"`
}

func makePredictionRequest(apiURL, smiles string) {
    requestPayload := PredictionRequest{Smiles: smiles}
    jsonData, err := json.Marshal(requestPayload)
    if err != nil {
        fmt.Println("Error marshalling input data:", err)
        return
    }

    resp, err := http.Post(apiURL, "application/json", bytes.NewBuffer(jsonData))
    if err != nil {
        fmt.Println("Error making request:", err)
        return
    }
    defer resp.Body.Close()

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        fmt.Println("Error reading response body:", err)
        return
    }

    fmt.Println("Response from Flask API:", string(body))
}

func main() {
    flaskAPIURL := "http://localhost:5000/predict"
    scanner := bufio.NewScanner(os.Stdin)

    fmt.Println("Enter a SMILES string for lipophilicity prediction:")
    scanner.Scan()
    smilesString := scanner.Text()

    fmt.Println("Sending prediction request to Flask API...")
    makePredictionRequest(flaskAPIURL, smilesString)
}
