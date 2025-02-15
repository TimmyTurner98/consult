package modules

type Client struct {
	Id          int    `db:"id"`
	Name        string `json:"name" binding:"omitempty"`
	Number      string `json:"number" binding:"required"`
	Companyname string `json:"companyname" binding:"omitempty"`
	Companybin  string `json:"companybin" binding:"omitempty"`
	ClientType  string `json:"client_type" binding:"omitempty"`
}
