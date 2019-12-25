package main

import (
  "net/http"
  "encoding/json"
  "log"
  "io/ioutil"
  "html/template"
  "strings"
)

type DogImg struct {
  Image string `json:"message"`
}

type DogData struct {
  Name string
  DogImage string
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

  var dogImage DogImg

  json.Unmarshal(responseData, &dogImage)

  var dogData DogData

  nameStartPos := strings.Index(dogImage.Image, "breeds/") + 7

  nameEndPos := strings.LastIndex(dogImage.Image, "/")

  nameOfDog := dogImage.Image[nameStartPos:nameEndPos]

  dogData = DogData {
    Name: nameOfDog,
    DogImage: dogImage.Image,
  }

  tmp := template.Must(template.ParseFiles("index.html"))

  tmp.Execute(w, dogData)
}


func main() {

  http.HandleFunc("/", GetImg)

  log.Fatal(http.ListenAndServe(":8080", nil))
}
