package utils

import "gorm.io/gorm"

func FilterByParams(db *gorm.DB, filters map[string]any) *gorm.DB {
	for key, val := range filters {
		if str, ok := val.(string); ok && str != "" {
			db = db.Where(key+" LIKE ?", "%"+str+"%")
		}
	}
	return db
}
