package feedly

type ProfileService struct {
	client *Client
}

type Profile struct {
	Id                  *string `json:"id,omitempty"`
	Email               *string `json:"email,omitempty"`
	GivenName           *string `json:"givenName,omitempty"`
	FamilyName          *string `json:"familyName,omitempty"`
	FullName            *string `json:"fullName,omitempty"`
	Picture             *string `json:"picture,omitempty"`
	Gender              *string `json:"gender,omitempty"`
	Locale              *string `json:"locale,omitempty"`
	Google              *string `json:"google,omitempty"`
	Reader              *string `json:"reader,omitempty"`
	TwitterUserId       *string `json:"twitterUserId,omitempty"`
	FacebookUserId      *string `json:"facebookUserId,omitempty"`
	WordPressId         *string `json:"wordPressId,omitempty"`
	WindowsLiveid       *string `json:"windowsLiveId,omitempty"`
	Wave                *string `json:"wave,omitempty"`
	Client              *string `json:"client,omitempty"`
	Source              *string `json:"source,omitempty"`
	Created             *int    `json:"created,omitempty"`
	Product             *string `json:"product,omitempty"`
	ProductExpiration   *int    `json:"productExpiration,omitempty"`
	SubscriptionStatus  *int    `json:"subscriptionStatus,omitempty"`
	IsEvernoteConnected *bool   `json:"isEvernoteConnected,omitempty"`
	IsPocketConnected   *bool   `json:"isPocketConnected,omitempty"`
}

func (s *ProfileService) Get() (*Profile, *Response, error) {
	req, err := s.client.NewRequest("GET", "profile", nil)
	if err != nil {
		return nil, nil, err
	}

	profile := new(Profile)
	resp, err := s.client.Do(req, profile)
	if err != nil {
		return nil, resp, err
	}

	return profile, resp, err
}
