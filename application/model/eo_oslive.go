package model

// EeoSmallProgramWeixinUser [...]
type EeoSmallProgramWeixinUser struct {
	ID         int    `gorm:"primary_key;column:id;type:int(10) unsigned;not null"` // 微信前台用户表主键
	OpenID     string `gorm:"unique;column:open_id;type:varchar(50);not null"`      // 微信登录openid
	NickName   string `gorm:"column:nick_name;type:varchar(35);not null"`           // 用户昵称
	AvatarURL  string `gorm:"column:avatar_url;type:varchar(200);not null"`         // 用户头像地址
	Gender     int8   `gorm:"column:gender;type:tinyint(3) unsigned;not null"`      // 用户性别
	City       string `gorm:"column:city;type:varchar(20);not null"`                // 所在城市
	Province   string `gorm:"column:province;type:varchar(20);not null"`            // 所在省份
	Country    string `gorm:"column:country;type:varchar(20);not null"`             // 所在国家
	UId        int64  `gorm:"column:uid;type:bigint(20) unsigned;not null"`         // 用户ID
	Telephone  string `gorm:"column:telephone;type:varchar(50);not null"`           // 用户绑定手机号码
	Createtime int    `gorm:"column:createtime;type:int(10) unsigned;not null"`     // 创建时间
	Status     int8   `gorm:"column:status;type:tinyint(3) unsigned;not null"`      // 用户状态1可用 0不可用
}

// EeoSmallProgramWeixinUserCourse [...]
type EeoSmallProgramWeixinUserCourse struct {
	ID           int    `gorm:"primary_key;column:id;type:int(10) unsigned;not null"`            // 用户课程订阅表主键ID
	CourseID     int64  `gorm:"column:course_id;type:bigint(20) unsigned;not null"`              // 课程ID
	OpenID       string `gorm:"column:open_id;type:varchar(50);not null"`                        // 用户openid
	Looktime     int    `gorm:"column:looktime;type:int(10) unsigned;not null"`                  // 用户查看时间
	WxUserid     int    `gorm:"index:wx_userid;column:wx_userid;type:int(10) unsigned;not null"` // 用户表主键ID
	IsSubCourse  uint8  `gorm:"column:is_sub_course;type:tinyint(3) unsigned;not null"`          // 是否订阅课程 0不订阅 1订阅
	SubClasslist string `gorm:"column:sub_classlist;type:text;not null"`                         // 课节列表(序列化)
	Createtime   int    `gorm:"column:createtime;type:int(10) unsigned;not null"`                // 创建时间
	Updatetime   int    `gorm:"column:updatetime;type:int(10) unsigned;not null"`                // 修改时间
	IsDelete     uint8  `gorm:"column:is_delete;type:tinyint(3) unsigned;not null"`              // 是否删除
}

// EeoWebcastCastInfo 直播聊天室相关信息表
type EeoWebcastCastInfo struct {
	ID            int   `gorm:"primary_key;column:id;type:int(10) unsigned;not null"`     // ID
	ClassID       int64 `gorm:"unique;column:class_id;type:bigint(20) unsigned;not null"` // 课节ID
	ClientClassID int64 `gorm:"column:client_class_id;type:bigint(20) unsigned;not null"` // 客户端课节ID
	Status        uint8 `gorm:"column:status;type:tinyint(3) unsigned;not null"`          // 聊天室状态
	PlayNum       int   `gorm:"column:play_num;type:int(10) unsigned;not null"`           // 播放次数
	LikeNum       int   `gorm:"column:like_num;type:int(10) unsigned;not null"`           // 点赞次数
	EndTimestamp  int   `gorm:"column:end_timestamp;type:int(10) unsigned;not null"`      // 课节结束时间
}

// EeoWebcastGoods 大直播商品表
type EeoWebcastGoods struct {
	ID         int64  `gorm:"primary_key;column:id;type:bigint(20) unsigned;not null"`          // 主键ID
	URL        string `gorm:"column:url;type:varchar(1024);not null"`                           // 商品链接
	Name       string `gorm:"column:name;type:varchar(100);not null"`                           // 商品名称
	Content    string `gorm:"column:content;type:varchar(100);not null"`                        // 商品描述
	ImgURL     string `gorm:"column:img_url;type:varchar(250);not null"`                        // 图片地址
	IsOn       int8   `gorm:"column:is_on;type:tinyint(3) unsigned;not null"`                   // 是否上架 0不上架 1上架 默认0
	LookCount  int    `gorm:"column:look_count;type:int(10) unsigned;not null"`                 // 浏览次数
	CreateTime int    `gorm:"column:create_time;type:int(10) unsigned;not null"`                // 创建时间
	UpdateTime int    `gorm:"column:update_time;type:int(10) unsigned;not null"`                // 更新时间
	ClassID    int64  `gorm:"index:class_id;column:class_id;type:bigint(20) unsigned;not null"` // 课节ID
	CourseID   int64  `gorm:"column:course_id;type:bigint(20) unsigned;not null"`               // 课程ID
}

// EeoWebcastGoodsRecord 大直播商品浏览记录表
type EeoWebcastGoodsRecord struct {
	ID       int64  `gorm:"primary_key;column:id;type:bigint(20) unsigned;not null"`                   // 主键ID
	Account  string `gorm:"column:account;type:varchar(50);not null"`                                  // 浏览账号
	Nickname string `gorm:"column:nickname;type:varchar(32);not null"`                                 // 浏览昵称
	LookTime int    `gorm:"column:look_time;type:int(10) unsigned;not null"`                           // 浏览时间
	GoodsID  int64  `gorm:"index:class_goods_index;column:goods_id;type:bigint(20) unsigned;not null"` // 商品ID
	ClassID  int64  `gorm:"index:class_goods_index;column:class_id;type:bigint(20) unsigned;not null"` // 课节ID
}

// EeoWebcastGoodsTemplate [...]
type EeoWebcastGoodsTemplate struct {
	ID         int64  `gorm:"primary_key;column:id;type:bigint(20) unsigned;not null"`              // 主键ID
	URL        string `gorm:"column:url;type:varchar(250);not null"`                                // 模板url
	Name       string `gorm:"column:name;type:varchar(100);not null"`                               // 模板名称
	Content    string `gorm:"column:content;type:varchar(100);not null"`                            // 模板描述
	ImgURL     string `gorm:"column:img_url;type:varchar(250);not null"`                            // 图片url
	ClassID    int64  `gorm:"column:class_id;type:bigint(20) unsigned;not null"`                    // 课节ID
	CourseID   int64  `gorm:"column:course_id;type:bigint(20) unsigned;not null"`                   // 课程ID
	CreateTime int    `gorm:"column:create_time;type:int(10) unsigned;not null"`                    // 创建时间
	IsDel      int8   `gorm:"column:is_del;type:tinyint(3) unsigned;not null"`                      // 是否删除 0不删除 1删除
	UpdateTime int    `gorm:"column:update_time;type:int(10) unsigned;not null"`                    // 删除时间
	SchoolUId  int64  `gorm:"index:school_uid;column:school_uid;type:bigint(20) unsigned;not null"` // 机构UID
}

// EeoWebcastLikeInfo 直播聊天室点赞信息表
type EeoWebcastLikeInfo struct {
	ID         int    `gorm:"primary_key;column:id;type:int(10) unsigned;not null"`                                                                    // ID
	ClassID    int64  `gorm:"index:eeo_webcast_like_infoindexA;index:eeo_webcast_like_infoindexB;column:class_id;type:bigint(20) unsigned;not null"` // 课节ID
	Account    string `gorm:"index:eeo_webcast_like_infoindexB;column:account;type:varchar(50);not null"`                                             // 帐号
	LikeStatus int8   `gorm:"index:eeo_webcast_like_infoindexA;column:like_status;type:tinyint(3) unsigned;not null"`                                 // 点赞状态
}

// EeoWebcastLoginRoomRecord [...]
type EeoWebcastLoginRoomRecord struct {
	ID          int    `gorm:"primary_key;column:id;type:int(10) unsigned;not null"`                  // 自增主键
	Account     string `gorm:"index:account_index;column:account;type:varchar(50);not null"`          // 用户账号
	ClassID     int64  `gorm:"index:classId_index;column:class_id;type:bigint(20) unsigned;not null"` // 课节ID
	CreateTime  int    `gorm:"column:create_time;type:int(10) unsigned;not null"`                     // 添加时间
	IsForbidden int8   `gorm:"column:is_forbidden;type:tinyint(3) unsigned;not null"`                 // 是否禁言 0没有禁言 1禁言 默认0
}

// EeoWebcastNote 大直播公告表
type EeoWebcastNote struct {
	ID          int64  `gorm:"primary_key;column:id;type:bigint(20) unsigned;not null"`          // 主键ID
	NoteContent string `gorm:"column:note_content;type:text;not null"`                           // 公告内容
	CreateTime  int    `gorm:"column:create_time;type:int(10) unsigned;not null"`                // 添加时间
	UpdateTime  int    `gorm:"column:update_time;type:int(10) unsigned"`                         // 更新时间
	ClassID     int64  `gorm:"index:class_id;column:class_id;type:bigint(20) unsigned;not null"` // 课节ID
	CourseID    int64  `gorm:"column:course_id;type:bigint(20) unsigned;not null"`               // 课程ID
}

// EeoWebcastRemindSubscribe 大直播课节提醒表
type EeoWebcastRemindSubscribe struct {
	ID        int64  `gorm:"primary_key;column:id;type:bigint(20) unsigned;not null"` // 主键ID
	ClassID   int64  `gorm:"column:class_id;type:bigint(20) unsigned;not null"`       // 课节ID
	Telephone string `gorm:"column:telephone;type:varchar(50);not null"`              // 手机号
	UId       int64  `gorm:"column:uid;type:bigint(20) unsigned;not null"`            // uid
	SubType   int8   `gorm:"column:sub_type;type:tinyint(3) unsigned;not null"`       // 订阅类型 0课程 1课节 (短信发送链接)
	SubTime   int    `gorm:"column:sub_time;type:int(10) unsigned;not null"`          // 订阅时间
}

// EeoWebcastUserRecord 直播聊天室成员进出表
type EeoWebcastUserRecord struct {
	ID             int    `gorm:"primary_key;column:id;type:int(10) unsigned;not null"`                                     // ID
	ClassID        int64  `gorm:"index:class_id;index:class_cookie_time;column:class_id;type:bigint(20) unsigned;not null"` // 课节ID
	Cookie         string `gorm:"index:class_cookie_time;column:cookie;type:varchar(50);not null"`                          // cookie值
	Account        string `gorm:"index:account;column:account;type:varchar(50);not null"`                                   // 帐号
	AccountType    uint8  `gorm:"column:account_type;type:tinyint(3) unsigned;not null"`                                    // 帐号类型
	Nickname       string `gorm:"column:nickname;type:varchar(32);not null"`                                                // 昵称
	Source         uint8  `gorm:"column:source;type:tinyint(3) unsigned;not null"`                                          // 来源
	Client         uint8  `gorm:"column:client;type:tinyint(3) unsigned;not null"`                                          // 客户端来源
	InTime         int    `gorm:"index:out_in_time;column:in_time;type:int(11);not null"`                                   // 进入时间
	OutTime        int    `gorm:"index:out_in_time;index:class_cookie_time;column:out_time;type:int(11);not null"`          // 离开时间
	ForbiddenSpeak int8   `gorm:"column:forbidden_speak;type:tinyint(3) unsigned;not null"`                                 // 是否禁言 0不禁言 1禁言 (废弃)
	ClientIP       string `gorm:"column:client_ip;type:varchar(50);not null"`                                               // 客户端IP
}
