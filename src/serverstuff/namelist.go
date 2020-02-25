package serverstuff

var storedNames []Name

func CreateList(names []Name) error {
  if len(storedNames) == 0 {
    storedNames = names
  }

  return nil
}

//Returns a copy of the name list
func GetList() []Name {
  return storedNames
}

func GetMarshalledList(){}
