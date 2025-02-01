package modules

type Client struct {
	Id         int    `db:"id"`
	Name       string `json:"name" binding:"omitempty"`
	Number     int    `json:"number" binding:"required"`
	Company    string `json:"company" binding:"omitempty"`
	CompanyId  int    `json:"company_id" binding:"omitempty"`
	ClientType string `json:"client_type" binding:"omitempty"`
}
