package models

//User struct is used for users schema in db
type User struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

//Scan struct is used for scan schema in db
type Scan struct {
	Id     int64  `json:"scan_id"`
	Person string `json:"person"`
	Url    string `json:"url"`
	Result string `json:"result"`
}

//Result struct is used for Scan schema
type Result struct {
	Request struct {
		URL   string `json:"url"`
		Depth int    `json:"depth"`
	} `json:"request"`
	Person string `json:"person"`
	Divs   []struct {
		Div         string `json:"Div"`
		Wc111Aria6  string `json:"Wc111Aria6"`
		Wc111Aria10 string `json:"Wc111Aria10"`
	} `json:"divs"`
	Buttons interface{} `json:"buttons"`
	Inputs  []struct {
		Input       string `json:"Input"`
		Wc111Aria6  string `json:"Wc111Aria6"`
		Wc111Aria10 string `json:"Wc111Aria10"`
	} `json:"inputs"`
	Images []struct {
		Img         string `json:"Img"`
		Wc111Aria6  string `json:"Wc111Aria6"`
		Wc111Aria10 string `json:"Wc111Aria10"`
		Wc111G94    string `json:"Wc111G94"`
		Wc111H2     string `json:"Wc111H2"`
	} `json:"images"`
	Videos    interface{} `json:"videos"`
	Audios    interface{} `json:"audios"`
	TextAreas interface{} `json:"textAreas"`
	Selects   interface{} `json:"selects"`
	Paras     interface{} `json:"paras"`
	Iframes   interface{} `json:"iframes"`
	Links     []struct {
		Anchor      string `json:"Anchor"`
		Wc111Aria6  string `json:"Wc111Aria6"`
		Wc111Aria10 string `json:"Wc111Aria10"`
	} `json:"links"`
	Areas   interface{} `json:"areas"`
	Objects interface{} `json:"objects"`
	Embeds  interface{} `json:"embeds"`
	Tracks  interface{} `json:"tracks"`
}

//Response format
type Response struct {
	ID      int64  `json:"id,omitempty"`
	Message string `json:"message,omitempty"`
}
