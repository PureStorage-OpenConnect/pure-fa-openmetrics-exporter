package client

type FileSystem struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type Member struct {
	Id           string `json:"id"`
	Name         string `json:"name"`
	ResourceType string `json:"resource_type"`
}

type Policy struct {
	Id           string `json:"id"`
	Name         string `json:"name"`
	ResourceType string `json:"resource_type"`
}

type LimitedBy struct {
	Member Member `json:"member"`
	Policy Policy `json:"policy"`
}

type Directory struct {
	Id            string     `json:"id"`
	Name          string     `json:"name"`
	Created       int        `json:"created"`
	Destroyed     bool       `json:"destroyed"`
	DirectoryName string     `json:"directory_name"`
	FileSystem    FileSystem `json:"file_system"`
	Path          string     `json:"path"`
	Space         Space      `json:"space"`
	TimeRemaining int        `json:"time_remaining"`
	LimitedBy     LimitedBy  `json:"limited_by"`
}

type DirectoriesList struct {
	ContinuationToken  string      `json:"continuation_token"`
	TotalItemCount     int         `json:"total_item_count"`
	MoreItemsRemaining bool        `json:"more_items_remaining"`
	Items              []Directory `json:"items"`
	Total              []Directory `json:"total"`
}

func (fa *FAClient) GetDirectories() *DirectoriesList {
	uri := "/directories"
	result := new(DirectoriesList)
	res, err := fa.RestClient.R().
		SetResult(&result).
		Get(uri)
	if err != nil {
		fa.Error = err
	}

	if res.StatusCode() == 401 {
		fa.RefreshSession()
		_, err = fa.RestClient.R().
			SetResult(&result).
			Get(uri)
		if err != nil {
			fa.Error = err
		}
	}

	return result
}
