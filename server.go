package main

import (
  "net/http"
  "encoding/json"
  "log"
  "io/ioutil"
  "html/template"
)

type DogImg struct {
  Image string `json:"message"`
}

func GetImg(w http.ResponseWriter, r *http.Request) {
  response, err := http.Get("https://dog.ceo/api/breeds/image/random")

  if err != nil {
    panic(err)
  }

  responseData,errData := ioutil.ReadAll(response.Body)

  if errData != nil {
    panic(errData)
  }

  var dogData DogImg

  json.Unmarshal(responseData, &dogData)

  tmp := template.Must(template.ParseFiles("index.html"))

  tmp.Execute(w, dogData)
}


func main() {

  http.HandleFunc("/", GetImg)

  log.Fatal(http.ListenAndServe(":8080", nil))
}
