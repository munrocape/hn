package main

import (
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
	fmt.Printf("requesting: %s\n", url)
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	} else {
		defer response.Body.Close()
		contents, err := ioutil.ReadAll(response.Body)
		if err != nil {
			return nil, err
		}
		return contents, nil
	}
}

func (c *Client) GetItem(id int) {
	url := fmt.Sprintf(c.ItemUrl, id)
	rep, _ := c.GetResource(url)
	fmt.Printf("%s\n", rep)
}

func (c *Client) GetUser(username string) {
	url := fmt.Sprintf(c.UserUrl, username)
	rep, _ := c.GetResource(url)
	fmt.Printf("%s\n", rep)
}

func main() {
	c := NewClient()
	c.GetItem(8863)
	c.GetUser("munrocape")
}
