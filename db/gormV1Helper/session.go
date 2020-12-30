package gormV1Helper

/*
CREATE TABLE `session` (
    `session_key` char(64) NOT NULL,
    `session_data` blob,
    `session_expiry` int(11) unsigned NOT NULL,
    PRIMARY KEY (`session_key`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8;


func init() {
    models.DB.Set("gorm:table_options", "ENGINE=MyISAM DEFAULT CHARSET=utf8").AutoMigrate(&Session{})
}
*/

type Session struct {
	SessionKey    string `gorm:"type:char(64);not null;primary_key;"`
	SessionData   string `gorm:"type:blob"`
	SessionExpiry int64  `gorm:"type:int(11) unsigned;not null"`
}

func (this *Session) TableName() string {
	return "session"
}
