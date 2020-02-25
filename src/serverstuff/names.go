package serverstuff

import (
  "net/http"
  "fmt"
  "io/ioutil"
  "encoding/json"
)

const NameEndpoint string = "https://uinames.com/api/"

type Name struct {
  Name string `json:"name"`
  Surname string `json:"surname"`
  Gender string `json:"gender"`
  Region string `json:"region"`
}

func FetchNameList(quantity int) error {
  fullUrl := fmt.Sprintf("%s?amount=%d", NameEndpoint, 500)

  resp, respErr := http.Get(fullUrl)
  if respErr != nil {
    return respErr
  }

  names, nameErr := ExtractBody(resp)
  if nameErr != nil {
    return nameErr
  }

  for _, name := range names {
    fmt.Printf("* %v\n", name)
  }
  return nil
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
