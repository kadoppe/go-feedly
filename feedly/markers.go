package feedly

type MarkersService struct {
	client *Client
}

type UnreadCounts struct {
	UnreadCounts []UnreadCount `json:"unreadcounts,omitempty"`
}

type UnreadCount struct {
	Updated *int    `json:"updated,omitempty"`
	Count   *int    `json:"count,omitempty"`
	Id      *string `json:"id,omitempty"`
}

func (s *MarkersService) ListUnreadCounts() ([]UnreadCount, *Response, error) {
	req, err := s.client.NewRequest("GET", "markers/counts", nil)
	if err != nil {
		return nil, nil, err
	}

	counts := new(UnreadCounts)
	resp, err := s.client.Do(req, counts)
	if err != nil {
		return nil, resp, err
	}

	return counts.UnreadCounts, resp, err
}
