package domain

type Product struct {
	ID   int    `gorm:"primary_key;auto_increment" json:"id"`
	Name string `json:"name"`
	Qty  int    `json:"qty"`

	AuditTable AuditTable `gorm:"embedded" json:"-"`
}
