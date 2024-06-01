package main

import (
	"fmt"
	"golinkparser"
	"strings"
)

// more examples can be picked from htmls directory
var htmlString string = `
<html>
<head>
  <link rel="stylesheet" href="https://maxcdn.bootstrapcdn.com/font-awesome/4.7.0/css/font-awesome.min.css">
</head>
<body>
  <h1>Social stuffs</h1>
  <div>
    <a href="https://www.twitter.com/joncalhoun">Check me out on twitter <i class="fa fa-twitter" aria-hidden="true"></i></a>
    <a href="https://github.com/gophercises">Gophercises is on Github!</a>
  </div>
</body>
</html>
`

func main() {
	reader := strings.NewReader(htmlString)
	links, err := golinkparser.Parse(reader)
	if err != nil {
		panic(err)
	}
	fmt.Print(links)
}
