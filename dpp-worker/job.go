package main

type ProxyJob struct {
	UrlBase      string `json:"base"`
	OriginalFile string `json:"file"`
	ResolutionW  int    `json:"w"`
	ResolutionH  int    `json:"h"`
}
