package cloudrunner

type CredentialsReadPermission struct {
	Account   string `json:"accountName"`
	ID        string `json:"-" gorm:"primary_key"`
	ReadGroup string `json:"readGroup"`
}

type CredentialsWritePermission struct {
	Account    string `json:"accountName"`
	ID         string `json:"-" gorm:"primary_key"`
	WriteGroup string `json:"writeGroup"`
}
