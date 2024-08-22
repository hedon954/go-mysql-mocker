package examples

const DBName = "gmm_db"

type CertificationInfo struct {
	ID       int `gorm:"index;primaryKey;autoIncrement"`
	Username string
	Password string
}

func (receiver CertificationInfo) TableName() string {
	return "certification_info"
}
