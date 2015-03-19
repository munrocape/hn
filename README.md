#A Golang Hacker News API Client

This is a wrapper client for the Hacker News API. It implements the most recent version, v0. It supports all resource requests.

##Usage
From a command line, run `$ go fetch github.com/munrocape/hn/client`

You can then include it in any go file.
```Go
import "github.com/munrocape/hn/client"
```


##API Specification and Code Examples
The API specification can be found [here.](https://github.com/HackerNews/API)
The below code examples are meant to demonstrate how to fetch resources as well as interact with them. 

```Go
package main
import ("fmt"
        hn "github.com/munrocape/hn/client"
)
func main() {
  client := hn.NewClient()
  mostPopular := client.GetTopStories(1) // You can request up to 500 of the top stories at once
  commentCount := mostPopular.Descendants // Descendants corresponds to all comments
  topLevelComments := len(mostPopular.Kids) // Kids corresponds to top level comments on a story
  fmt.Printf("The most popular story on HN: %s\nComment Count: %d\nTop Level Comment Count: %d", mostPopular.Title, commentCount, topLevelComments) // Gotta print out to make sure it compiles ;)
}
```

##Structs
In the HNverse of resources, there are Items and Users.

###Users
Users, as you may have guessed, represent user accounts.
```Go
type User struct {
	Id        string `json:"id"`
	Delay     int    `json:"delay"`
	Created   int    `json:"created"`
	Karma     int    `json:"karma"`
	About     string `json:"about"`
	Submitted []int  `json:"submitted"` // This represents the ids of all the items they have submitted
}
```

###Items
Items, on the other hand, may be slightly more vague. In fact, an Item represents a superset of attributes as follows:
```Go
type Item struct {
	Id          int    `json:"id"` // A UUID. if x > y then x was created after y.
	Deleted     bool   `json:"deleted"` // Whether or not the item was deleted
	Type        string `json:"type"` // one of {story, comment, poll, pollopt}
	By          string `json:"by"` // Account who submitted the item
	Time        int    `json:"time"` // UNIX time of submission
	Text        string `json:"text"` // Body of the submission
	Dead        bool   `json:"dead"` // Whether or not the item was killed
	Parent      int    `json:"parent"` // Parent - could be a comment or poll
	Kids        []int  `json:"kids"` // Top level comments
	Url         string `json:"url"` // External link. URL of item is https://news.ycombinator.com/item?id={Item.Id}
	Score       int    `json:"score"` // Current score of the item
	Title       string `json:"title"` // Title of the item
	Parts       []int  `json:"parts"` // Poll options
	Descendants int    `json:"descendants"` // Total comment count
}
```

Items represent all the attributes that can make up a story, comment, poll, or poll option on HN. The "Type" field corresponds to one of those four objects. They are outlined below.

```Go
type Story struct {
	By          string `json:"by"`
	Descendants int    `json:"descendants"`
	Id          int    `json:"id"`
	Kids        []int  `json:"kids"`
	Score       int    `json:"score"`
	Time        int    `json:"time"`
	Title       string `json:"title"`
	Url         string `json:"url"`
}
```

```Go
type Comment struct {
	By     string `json:"by"`
	Id     int    `json:"id"`
	Kids   []int  `json:"kids"`
	Parent int    `json:"parent"`
	Text   string `json:"text"`
	Time   int    `json:"time"`
}
```

```Go
type Poll struct {
	By          string `json:"by"`
	Descendants int    `json:"descendants"`
	Id          int    `json:"id"`
	Kids        []int  `json:"kids"`
	Parts       []int  `json:"parts"`
	Score       int    `json:"score"`
	Text        string `json:"text"`
	Time        int    `json:"time"`
	Title       string `json:"title"`
}
```

```Go
type PollOpt struct {
	By     string `json:"by"`
	Id     int    `json:"id"`
	Parent int    `json:"parent"`
	Score  int    `json:"score"`
	Text   string `json:"text"`
	Time   int    `json:"time"`
}
```
