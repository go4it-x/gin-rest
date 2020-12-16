package model

// Model Basic field
type Model struct {
	Id          uint64 `gorm:"type:int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID';primary_key;" json:"id"`
	Status      uint8  `gorm:"type:tinyint(1) unsigned NOT NULL;default:1;comment:'状态：#1启用 #2禁用'" json:"status"`
	IsDelete    uint8  `gorm:"type:tinyint(1) unsigned NOT NULL;default:2;comment:'是否删除：#1是 #2否'" json:"isDelete"`
	CreatedTime int64  `gorm:"type:int(10);default:0;comment:'创建时间'" json:"createdTime"`
	UpdatedTime int64  `gorm:"type:int(10);default:0;comment:'修改时间'" json:"updatedTime"`
}
