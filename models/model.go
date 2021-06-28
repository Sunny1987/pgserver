package models

import "encoding/json"

//User struct is used for users schema in db
type User struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

//Scan struct is used for scan schema in db
type Scan struct {
	Id     int64           `json:"scan_id"`
	Person string          `json:"person"`
	Url    string          `json:"url"`
	Result json.RawMessage `json:"result"`
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
	Buttons []struct {
		Button      string `json:"button"`
		Wc111Aria6  string `json:"Wc111Aria6"`
		Wc111Aria10 string `json:"Wc111Aria10"`
	} `json:"buttons"`
	Inputs []struct {
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
	Videos []struct {
		Video       string `json:"video"`
		Wc111Aria6  string `json:"Wc111Aria6"`
		Wc111Aria10 string `json:"Wc111Aria10"`
	} `json:"videos"`
	Audios []struct {
		Audio       string `json:"audio"`
		Wc111Aria6  string `json:"Wc111Aria6"`
		Wc111Aria10 string `json:"Wc111Aria10"`
		Wc121H96    string `json:"Wc121h96"`
		Wc121ARIA   string `json:"Wc121aria"`
	} `json:"audios"`
	TextAreas []struct {
		Textarea    string `json:"textarea"`
		Wc111Aria6  string `json:"Wc111Aria6"`
		Wc111Aria10 string `json:"Wc111Aria10"`
	} `json:"textAreas"`
	Selects []struct {
		Select      string `json:"select"`
		Wc111Aria6  string `json:"Wc111Aria6"`
		Wc111Aria10 string `json:"Wc111Aria10"`
	} `json:"selects"`
	Paras []struct {
		Para        string `json:"para"`
		Wc111Aria6  string `json:"Wc111Aria6"`
		Wc111Aria10 string `json:"Wc111Aria10"`
	} `json:"paras"`
	Iframes []struct {
		Iframe      string `json:"iframe"`
		Wc111Aria6  string `json:"Wc111Aria6"`
		Wc111Aria10 string `json:"Wc111Aria10"`
	} `json:"iframes"`
	Links []struct {
		Anchor      string `json:"Anchor"`
		Wc111Aria6  string `json:"Wc111Aria6"`
		Wc111Aria10 string `json:"Wc111Aria10"`
	} `json:"links"`
	Areas []struct {
		Area     string `json:"area"`
		Wc111G94 string `json:"Wc111Aria6"`
		Wc111H2  string `json:"Wc111h2"`
	} `json:"areas"`
	Objects []struct {
		Object    string `json:"object"`
		Wc121H96  string `json:"wc_121_h_96"`
		Wc121ARIA string `json:"wc_121_aria"`
		Wc111H53  string `json:"wc_111_h_53"`
		Wc121G166 string `json:"wc_121_g_166"`
	} `json:"objects"`
	Embeds []struct {
		Embed     string `json:"embed"`
		Wc121H96  string `json:"wc_121_h_96"`
		Wc121ARIA string `json:"wc_121_aria"`
		Wc111H53  string `json:"wc_111_h_53"`
		Wc121G166 string `json:"wc_121_g_166"`
	} `json:"embeds"`
	Tracks []struct {
		Track      string `json:"track"`
		Wc121H96   string `json:"wc_121_h_96"`
		Wc121ARIA  string `json:"wc_121_aria"`
		Wc111G94   string `json:"wc_111_g_94"`
		Wc111H2    string `json:"wc_111_h_2s"`
		Wcag122G87 string `json:"wcag_122_g_87"`
	} `json:"tracks"`
}

//Response format
type Response struct {
	ID      int64  `json:"id,omitempty"`
	Message string `json:"message,omitempty"`
}
