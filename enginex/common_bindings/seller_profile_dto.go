package common_bindings

type SellerProfileRegisterDTO struct {
	Version    int64 `json:"version"`
	VersionOld int64 `json:"version_old"`

	SellerCode string `json:"seller_code"`
	AccountID  string `json:"account_id"`

	CreateTime  int64  `json:"create_time"`
	CreateBy    string `json:"create_by"`
	CreateCode  string `json:"create_code"`
	UpdateTime  int64  `json:"update_time"`
	UpdateBy    string `json:"update_by"`
	UpdateCode  string `json:"update_code"`
	SubmitTime  int64  `json:"submit_time"`
	SubmitBy    string `json:"submit_by"`
	SubmitCode  string `json:"submit_code"`
	ApproveTime int64  `json:"approve_time"`
	ApproveBy   string `json:"approve_by"`
	ApproveCode string `json:"approve_code"`

	Status string `json:"status"`

	SellerNameEN     string   `json:"seller_name_en"`
	BusinessID       string   `json:"business_id"`
	BusinessType     []string `json:"business_type"`
	ExportMarket     []string `json:"export_market"`
	SellerDesc       string   `json:"seller_desc"`
	StoreUrl         string   `json:"store_url"`
	Website          string   `json:"website"`
	Email            string   `json:"email"`
	ReturnPolicy     string   `json:"return_policy"`
	GuaranteesPolicy string   `json:"guarantees_policy"`

	CategoryInterest []string `json:"category_interest"`

	SellerIdMagento string `json:"seller_id_magento"`
	SellerIdWCM     string `json:"seller_id_wcm"`

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

	Address      string `json:"address"`
	Country      string `json:"country"`
	Province     string `json:"province"`
	PostalCode   string `json:"postal_code"`
	Mobile       string `json:"mobile"`
	CompanyPhone string `json:"company_phone"`
	Fax          string `json:"fax"`
	MapLinkUrl   string `json:"map_link_url"`

	AddressImageName         string `json:"address_image_name"`
	AddressImageType         string `json:"address_image_type"`
	AddressImageRequestID    string `json:"address_image_request_id"`
	AddressImagePublicURL    string `json:"address_image_public_url"`
	AddressImageCDNURL       string `json:"address_image_cdn_url"`
	AddressImageContextName  string `json:"address_image_context_name"`
	AddressImageFolder       string `json:"address_image_folder"`
	AddressImageFileName     string `json:"address_image_file_name"`
	AddressImageFileLocation string `json:"address_image_file_location"`
	AddressImageFileSize     int64  `json:"address_image_file_size"`

	DocImportant     []SellerProfileRegisterDocImportantDTO     `json:"doc_important"`
	DocCertification []SellerProfileRegisterDocCertificationDTO `json:"doc_certification"`
	DocAward         []SellerProfileRegisterDocAwardDTO         `json:"doc_award"`

	TotalDocSize int64 `json:"total_doc_size"`
}

type SellerProfileRegisterDocImportantDTO struct {
	Version      int64  `json:"version"`
	VersionOld   int64  `json:"version_old"`
	RefCode      string `json:"ref_code"`
	StatusAction string `json:"status_action"`

	DocType  string                              `json:"doc_type"`
	DocImage []SellerProfileRegisterDocUploadDTO `json:"doc_image"`
}

type SellerProfileRegisterDocCertificationDTO struct {
	Version      int64  `json:"version"`
	VersionOld   int64  `json:"version_old"`
	RefCode      string `json:"ref_code"`
	StatusAction string `json:"status_action"`

	DocType      string                              `json:"doc_type"`
	DocNameOther string                              `json:"doc_name_other"`
	DocImage     []SellerProfileRegisterDocUploadDTO `json:"doc_image"`
}

type SellerProfileRegisterDocAwardDTO struct {
	Version      int64  `json:"version"`
	VersionOld   int64  `json:"version_old"`
	RefCode      string `json:"ref_code"`
	StatusAction string `json:"status_action"`

	Year        string                              `json:"year"`
	AwardDetail string                              `json:"award_detail"`
	DocImage    []SellerProfileRegisterDocUploadDTO `json:"doc_image"`
}

type SellerProfileRegisterDocUploadDTO struct {
	StatusAction string `json:"status_action"`

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
}

type SellerProfileWarehourseAddressDTO struct {
	Version    int64 `json:"version"`
	VersionOld int64 `json:"version_old"`

	AccountId  string `json:"account_id"`
	SellerCode string `json:"seller_code"`

	RefCode       string `json:"ref_code"`
	Address       string `json:"address"`
	Tel           string `json:"tel"`
	Country       string `json:"country"`
	ProvinceInput string `json:"province_input"`
	StateInput    string `json:"state_input"`
	PostalCode    string `json:"postal_code"`

	IsDefault bool `json:"is_default"`
}
