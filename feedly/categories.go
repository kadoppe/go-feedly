package feedly

type CategoriesService struct {
	client *Client
}

type Category struct {
	Id    *string `json:"id,omitempty"`
	Label *string `json:"label,omitempty"`
}

func (s *CategoriesService) List() ([]Category, *Response, error) {
	req, err := s.client.NewRequest("GET", "categories", nil)
	if err != nil {
		return nil, nil, err
	}

	categories := new([]Category)
	resp, err := s.client.Do(req, categories)
	if err != nil {
		return nil, resp, err
	}

	return *categories, resp, err
}
