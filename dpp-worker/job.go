package main

type ProxyJob struct {
	UrlBase      string `json:"base"`
	OriginalFile string `json:"file"`
	ResolutionW  int    `json:"w"`
	ResolutionH  int    `json:"h"`
}

type ProxyJobResult struct {
	Url   string `json:"base"`
	Size  int64  `json:"size"`
	Error string `json:"error"`
}
