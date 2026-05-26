package Client

import "fmt"

type Stringer interface {
	ToString() string
}

type Person struct {
	Name string
	Age  int64
}

func (p Person) ToString() string {
	return fmt.Sprintf("Person: %s ", p.Name)
}

func (p Person) Speak() {
	fmt.Printf("Hi my name is %s! i'm %v years old!", p.Name, p.Age)
}

type Avatar struct {
	URL  string
	SIZE string
}

type Client struct {
	ID int64
	Person
	Avatar
}

func (c Client) HasAvatar() bool {
	if c.URL != "" {
		return true
	}
	return false
}

func (c *Client) UpdateAvatar() {
	c.URL = "google.com"
}

func (c Client) GetName() string {
	return c.Name
}
