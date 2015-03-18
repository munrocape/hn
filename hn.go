package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Client struct {
	BaseUrl string
	UserSuffix string
	ItemSuffix string
	MaxSuffix string
	TopSuffix string
	JobSuffix string
	AskSuffix string
	UpdateSuffix string
}

func NewClient() *Client {
	var c Client
	c.BaseUrl = "https://hacker-news.firebaseio.com/v0/"
	c.UserSuffix = "user/%s.json"
	c.ItemSuffix = "item/%d.json"
	c.MaxSuffix = "maxitem.json"
	c.TopSuffix = "topstories.json"
	c.NewSuffix = "newstories.json"
	c.JobSuffix = "jobstories.json"
	c.AskSuffix = "askstories.json"
	c.UpdateSuffix = "updates.json"
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

		return contents, err
	}
}

func (c *Client) GetItem(id int) (Item, error) {
	url := c.BaseUrl + fmt.Sprintf(c.ItemSuffix, id)
	rep, err := c.GetResource(url)

	var i Item
	if err != nil {
		return i, err
	}

	err = json.Unmarshal(rep, &i)
	return i, err
}

func (c *Client) GetUser(username string) (User, error) {
	url := c.BaseUrl + fmt.Sprintf(c.UserSuffix, username)
	rep, err := c.GetResource(url)

	var user User
	if err != nil {
		return user, err
	}

	err = json.Unmarshal(rep, &user)
	return user, err
}

// GetTopStories takes an int number and returns an array of up to number ints that represent the top stories.
func (c *Client) GetTopStories(number int) ([]int, error) {
	var top500 []int
	if number > 500 {
		return top500, fmt.Errorf("Number %d greater than maximum 500 items allowed", number)
	}
	
	url := c.BaseUrl + c.TopSuffix
	rep, err := c.GetResource(url)
	
	err = json.Unmarshal(rep, &top500)
	
	if err != nil {
		return nil, err
	}

	return top500[:number], nil
}

// GetNewStories takes an int number and returns an array of up to number ints that represent the newest stories.
func (c *Client) GetNewStories(number int) ([]int, error) {
	var top500 []int
	if number > 500 {
		return top500, fmt.Errorf("Number %d greater than maximum 500 items allowed", number)
	}
	
	url := c.BaseUrl + c.NewSuffix
	rep, err := c.GetResource(url)

	err = json.Unmarshal(rep, &top500)
	
	if err != nil {
		return nil, err
	}

	return top500[:number], nil
}

func main() {
	c := NewClient()

	story, _ := c.GetItem(8715529)
	fmt.Printf("%+v\n\n", story)

	user, _ := c.GetUser("munrocape")
	fmt.Printf("%+v\n\n", user)

	comment, _ := c.GetItem(8715677)
	fmt.Printf("%+v\n\n", comment)

	top10, _ := c.GetTop(10)
	fmt.Printf("%+v\n\n", top10)
}
