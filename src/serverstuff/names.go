package serverstuff

import (
  "net/http"
  "fmt"
  "io/ioutil"
  "encoding/json"
)

const NameEndpoint string = "https://uinames.com/api/"
const DefaultAmount int = 500

type Name struct {
  Name string `json:"name"`
  Surname string `json:"surname"`
  Gender string `json:"gender"`
  Region string `json:"region"`
}

func FetchNameList(quantity int) ([]Name, error) {
  rounds := getNumberOfRounds(quantity)
  amount := DefaultAmount
  allNames := make([]Name, 0)

  for i := 0; i < rounds; i++ {
    fullUrl := fmt.Sprintf("%s?amount=%d", NameEndpoint, amount)
    resp, respErr := http.Get(fullUrl)
    if respErr != nil {
      return allNames, respErr
    }

    names, nameErr := ExtractBody(resp)
    if nameErr != nil {
      return allNames, nameErr
    }

    allNames = append(allNames, names...)
    if quantity - amount < DefaultAmount {
      amount = quantity - amount
    }
  }

  //Leaving the for loop in place for debugging
  /*for i, name := range allNames {
    fmt.Printf("%d) %v\n", i + 1, name)
  }*/
  return allNames, nil
}

func ExtractBody(resp *http.Response) ([]Name, error) {
  names := make([]Name, 0)

  bodyBytes, readErr := ioutil.ReadAll(resp.Body)
  if readErr != nil {
    return names, readErr
  }

  err := json.Unmarshal(bodyBytes, &names)
  if err != nil {
    return names, err
  }

  return names, nil
}

func getNumberOfRounds(quantity int) int {
  rounds := 1
  if quantity > DefaultAmount {
    switch quantity % DefaultAmount > 0 {
    case true:
      rounds = (quantity / DefaultAmount) + 1
    case false:
      rounds = (quantity / DefaultAmount)
    }
  }

  return rounds
}
