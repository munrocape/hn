#A Golang Hacker News API Client

This is a wrapper client for the Hacker News API. It implements the most recent version, v0. It supports all resource requests.

##Usage
From a command line, run `$ go fetch github.com/munrocape/hn-client`

You can then include it in any go file.
```Go
import "github.com/munrocape/hn-client"
```


##API Specification and Code Examples
The API specification can be found [here.](https://github.com/HackerNews/API)
The below code examples are meant to demonstrate how to fetch resources as well as interact with them. 

```Go
package main
import ("fmt"
        hn "github.com/munrocape/hn-client"
)
func main() {
  client := hn.NewClient()
  mostPopular := client.GetTopStories(1) // You can request up to 500 of the top stories at once
  commentCount := mostPopular.Descendants // Descendants corresponds to all comments
  topLevelComments := len(mostPopular.Kids) // Kids corresponds to top level comments on a story
  fmt.Printf("The most popular story on HN: %s\nComment Count: %d\nTop Level Comment Count: %d", mostPopular.Title, commentCount, topLevelComments) // Gotta print out to make sure it compiles ;)
}
```
