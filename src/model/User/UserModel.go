package userModel

type Details struct {
  Name string
  Age int
}



func (user Details) PrintUserName() string {
  return user.Name
}

func (user Details) PrintUserAge() int {
  return user.Age
}

func (user Details) PrintUserDetails() Details {
  return user
}
