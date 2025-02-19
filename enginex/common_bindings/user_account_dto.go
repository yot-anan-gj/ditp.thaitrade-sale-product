package common_bindings

type UserAccountDTO struct {
	Version    int64 `json:"version"`
	VersionOld int64 `json:"version_old"`

	AccountId string `json:"account_id"`
	NameEN    string `json:"name_en"`
	SurnameEN string `json:"surname_en"`
	Tel       string `json:"tel"`
	Email     string `json:"email"`
	Country   string `json:"country"`
	Gender    string `json:"gender_id"`

	BusinessRole string `json:"business_role"`
	BusinessID   string `json:"business_id"`
	UserType     string `json:"user_type"`

	IsSubscribe         bool `json:"is_subscribe"`
	IsVerifyByThaitrade bool `json:"is_verify_by_thaitrade"`
	IsUserThaitrade     bool `json:"is_user_thaitrade"`

	ImageName         string `json:"image_name"`
	ImageType         string `json:"image_type"`
	ImageRequestID    string `json:"image_request_id"`
	ImagePublicURL    string `json:"image_public_url"`
	ImageCDNURL       string `json:"image_cdn_url"`
	ImageContextName  string `json:"image_context_name"`
	ImageFolder       string `json:"image_folder"`
	ImageFileName     string `json:"image_file_name"`
	ImageFileLocation string `json:"image_file_location"`
	ImageFileSize     int64  `json:"image_file_size"`

	Channel string `json:"channel"`

	GoogleID        string `json:"google_id"`
	GoogleFirstName string `json:"google_first_name"`
	GoogleLastName  string `json:"google_last_name"`
	GoogleFullName  string `json:"google_full_name"`
	GoogleEmail     string `json:"google_email"`
	GoogleAvatarURL string `json:"google_avatar_url"`

	FacebookID        string `json:"facebook_id"`
	FacebookFirstName string `json:"facebook_first_name"`
	FacebookLastName  string `json:"facebook_last_name"`
	FacebookFullName  string `json:"facebook_full_name"`
	FacebookEmail     string `json:"facebook_email"`
	FacebookAvatarURL string `json:"facebook_avatar_url"`

	CategoryInterest []string `json:"category_interest"`
	TimeZone         string   `json:"time_zone"`
}

type UserAccountConfirmEmailDTO struct {
	AccountId   string `json:"account_id"`
	Email       string `json:"email"`
	IsConfirmed bool   `json:"is_confirmed"`
	ConfirmDate int64  `json:"confirm_date"`
}

type UserAccountAddressDTO struct {
	Version    int64 `json:"version"`
	VersionOld int64 `json:"version_old"`

	AccountId string `json:"account_id"`

	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`

	RefCode       string `json:"ref_code"`
	Address       string `json:"address"`
	Tel           string `json:"tel"`
	Country       string `json:"country"`
	ProvinceInput string `json:"province_input"`
	StateInput    string `json:"state_input"`
	PostalCode    string `json:"postal_code"`

	IsBilling  bool `json:"is_billing"`
	IsShipping bool `json:"is_shipping"`
}
