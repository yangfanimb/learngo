package controller

type blogItem struct {
	ID         int    `json:"id", int`
	Title      string `json:"title"`
	Content    string `json:"content"`
	CreateTime int64  `json:"createtime", int`
	Author     string `json:"author"`
}

func GetList(author, keyword string) (interface{}, error) {
	s := []blogItem{
		{
			ID:         0,
			Title:      "title A",
			Content:    "content A",
			CreateTime: 1560278459854,
			Author:     "zhangsan",
		},
		{
			ID:         1,
			Title:      "title B",
			Content:    "content B",
			CreateTime: 1560278467755,
			Author:     "lisi",
		},
	}

	return s, nil
}
