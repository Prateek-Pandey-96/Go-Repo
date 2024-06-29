package models

type Sitemap struct {
	Username     string   `json:"username" bson:"username"`
	Url          string   `json:"url" bson:"url"`
	Depth        int      `json:"depth" bson:"depth"`
	Destinations []string `bson:"destinations"`
}

type Response struct {
	Url          string `json:"url"`
	Destinations []string
}
