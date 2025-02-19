package common_bindings

type CompanyProfileSellerRequestDTO struct {
	SellerCode []string `json:"seller_code"`
}

type SellerCompanyProfileRequestDTO struct {
	AccountId string `json:"account_id"`
}

type CompanyProfileSellerResponseDTO struct {
	Success bool     `json:"success"`
	Data    []string `json:"data"`
}
