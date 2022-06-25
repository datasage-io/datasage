package classifiers

import "time"

type DPDatasource struct {
	ID                    int       `json:"id" gorm:"primary_key"`
	UserID                string    `json:"user_id" binding:"required" gorm:"not null"`
	WorkspaceID           int       `json:"workspace_id" binding:"required" gorm:"not null"`
	DataSourceDomain      string    `json:"ds_domain" binding:"required" gorm:"not null"`
	Name                  string    `json:"name" binding:"required" gorm:"not null"`
	DataSourceDescription string    `json:"ds_description" binding:"required" gorm:"not null"`
	DsType                string    `json:"ds_type" binding:"required" gorm:"not null"`
	DataSourceVersion     string    `json:"ds_version" binding:"required" gorm:"not null"`
	DsKey                 string    `json:"ds_key" gorm:"not null"`
	CreatedAt             time.Time `json:"created_at"`
	Deleted               time.Time `json:"deleted"`
}
