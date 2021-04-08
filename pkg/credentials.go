package cloudrunner

type Credentials struct {
	Account     string   `json:"account" gorm:"primary_key"`
	ProjectID   string   `json:"projectID"`
	ReadGroups  []string `json:"readGroups" gorm:"-"`
	WriteGroups []string `json:"writeGroups" gorm:"-"`
}
