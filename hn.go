package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Client struct {
	UserUrl string
	ItemUrl string
}

func NewClient() *Client {
	var c Client
	c.UserUrl = "https://hacker-news.firebaseio.com/v0/user/%s.json"
	c.ItemUrl = "https://hacker-news.firebaseio.com/v0/item/%d.json"
	return &c
}

func (c *Client) GetResource(url string) ([]byte, error) {
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	} else {
		defer response.Body.Close()
		contents, err := ioutil.ReadAll(response.Body)
		if err != nil {
			return nil, err
		}
		var jsonMap map[string]interface{}
		err = json.Unmarshal(contents, &jsonMap)
		return contents, err
	}
}

func (c *Client) GetItem(id int) (Item, error) {
	url := fmt.Sprintf(c.ItemUrl, id)
	rep, err := c.GetResource(url)
	var i Item
	if err != nil {
		return i, err
	}
	err = json.Unmarshal(rep, &i)

	return i, err
}

func (c *Client) GetUser(username string) (User, error) {
	url := fmt.Sprintf(c.UserUrl, username)
	rep, err := c.GetResource(url)
	var user User
	if err != nil {
		return user, err
	}

	err = json.Unmarshal(rep, &user)

	return user, err
}

func main() {
	c := NewClient()

	story, _ := c.GetItem(8715529)
	fmt.Printf("%+v\n", story)

	user, _ := c.GetUser("munrocape")
	fmt.Printf("%+v\n", user)

	comment, _ := c.GetItem(8715677)
	fmt.Printf("%+v\n", comment)
}
