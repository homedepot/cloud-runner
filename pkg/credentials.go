package cloudrunner

type Credentials struct {
	Account     string   `json:"account" gorm:"primary_key"`
	ProjectID   string   `json:"projectID" binding:"required"`
	ReadGroups  []string `json:"readGroups,omitempty" gorm:"-"`
	WriteGroups []string `json:"writeGroups,omitempty" gorm:"-"`
}
