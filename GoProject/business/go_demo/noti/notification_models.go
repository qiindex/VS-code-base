package noti

import (
	"encoding/json"
)

// ChannelType 渠道类型常量
const (
	ChannelTypeSMS   = 1  // SMS短信
	ChannelTypeEmail = 2  // 邮件
	ChannelTypePush  = 3  // 推送通知
	ChannelTypeChat  = 12 // 聊天消息（默认值）
)

// RecipientType 接收者类型常量
const (
	RecipientTypeSeller = 1 // 卖家
	RecipientTypeBuyer  = 2 // 买家
	RecipientTypeBoth   = 0 // 两者
)

// TemplateStatus 模板状态常量
const (
	TemplateStatusAvailable   = 1 // 可用
	TemplateStatusUnavailable = 2 // 不可用
)

// SendStatus 发送状态常量
const (
	SendStatusPending = 0 // 待发送
	SendStatusSuccess = 1 // 发送成功
	SendStatusFail    = 3 // 发送失败
)

// OperationType 操作类型常量
const (
	OperationTypeInsert       = 1 // 插入
	OperationTypeUpdate       = 2 // 更新
	OperationTypeDelete       = 3 // 删除
	OperationTypeUpdateStatus = 4 // 更新模板状态
)

// AutomaticChatTemplateTab 自动聊天模板表
type AutomaticChatTemplateTab struct {
	ID                          uint64          `json:"id" gorm:"id"`                                                         // 主键ID
	TemplateName                string          `json:"template_name" gorm:"template_name"`                                   // 聊天通知模板名称
	ChatMessageType             uint32          `json:"chat_message_type" gorm:"chat_message_type"`                           // 自动聊天业务类型
	Content                     string          `json:"content" gorm:"content"`                                               // 聊天模板内容
	TemplateStatus              uint32          `json:"template_status" gorm:"template_status"`                               // 模板状态：1-可用，2-不可用
	OrderDeliveryTrackingStatus uint32          `json:"order_delivery_tracking_status" gorm:"order_delivery_tracking_status"` // 订单配送跟踪状态：1-配送中，2-已配送，3-配送暂停
	AtlFlag                     uint32          `json:"atl_flag" gorm:"atl_flag"`                                             // ATL标志：0-两者，1-放门口，2-不放门口
	ConditionContent            string          `json:"condition_content" gorm:"condition_content"`                           // 聊天模板条件内容
	LanguageType                string          `json:"language_type" gorm:"language_type"`                                   // 语言类型
	ChannelType                 uint32          `json:"channel_type" gorm:"channel_type"`                                     // 渠道类型：1-SMS，2-邮件，3-推送，12-聊天
	RecipientTypes              json.RawMessage `json:"recipient_types" gorm:"recipient_types"`                               // 接收者类型数组，JSON格式
	ServiceTypes                json.RawMessage `json:"service_types" gorm:"service_types"`                                   // 服务类型数组，JSON格式
	NssShopIds                  json.RawMessage `json:"nss_shop_ids" gorm:"nss_shop_ids"`                                     // NSS店铺ID数组，JSON格式
	Ctime                       int64           `json:"ctime" gorm:"ctime"`                                                   // 创建时间戳
	Mtime                       int64           `json:"mtime" gorm:"mtime"`                                                   // 修改时间戳
}

// TableName 获取数据库表名
func (m *AutomaticChatTemplateTab) TableName() string {
	return "automatic_chat_template_tab"
}

// GetRecipientTypes 获取接收者类型数组
func (m *AutomaticChatTemplateTab) GetRecipientTypes() ([]int, error) {
	var types []int
	if len(m.RecipientTypes) == 0 {
		return []int{}, nil
	}
	err := json.Unmarshal(m.RecipientTypes, &types)
	return types, err
}

// SetRecipientTypes 设置接收者类型数组
func (m *AutomaticChatTemplateTab) SetRecipientTypes(types []int) error {
	data, err := json.Marshal(types)
	if err != nil {
		return err
	}
	m.RecipientTypes = data
	return nil
}

// GetServiceTypes 获取服务类型数组
func (m *AutomaticChatTemplateTab) GetServiceTypes() ([]int, error) {
	var types []int
	if len(m.ServiceTypes) == 0 {
		return []int{9999}, nil // 默认值
	}
	err := json.Unmarshal(m.ServiceTypes, &types)
	return types, err
}

// SetServiceTypes 设置服务类型数组
func (m *AutomaticChatTemplateTab) SetServiceTypes(types []int) error {
	data, err := json.Marshal(types)
	if err != nil {
		return err
	}
	m.ServiceTypes = data
	return nil
}

// GetNssShopIds 获取NSS店铺ID数组
func (m *AutomaticChatTemplateTab) GetNssShopIds() ([]int, error) {
	var ids []int
	if len(m.NssShopIds) == 0 {
		return []int{}, nil
	}
	err := json.Unmarshal(m.NssShopIds, &ids)
	return ids, err
}

// SetNssShopIds 设置NSS店铺ID数组
func (m *AutomaticChatTemplateTab) SetNssShopIds(ids []int) error {
	data, err := json.Marshal(ids)
	if err != nil {
		return err
	}
	m.NssShopIds = data
	return nil
}

// AutoSendChatLogTab 自动发送聊天日志表
type AutoSendChatLogTab struct {
	ID                          uint64          `json:"id" gorm:"id"`                                                         // 主键ID
	MessageId                   string          `json:"message_id" gorm:"message_id"`                                         // 唯一消息ID，核心字段，有索引
	SendStatus                  uint32          `json:"send_status" gorm:"send_status"`                                       // 发送状态：0-待发送，1-成功，3-失败
	MessageType                 uint32          `json:"message_type" gorm:"message_type"`                                     // 消息类型：0-配送，1-揽收，2-取件，4-退货
	OrderDeliveryTrackingStatus uint32          `json:"order_delivery_tracking_status" gorm:"order_delivery_tracking_status"` // 订单配送跟踪状态
	Reason                      string          `json:"reason" gorm:"reason"`                                                 // 原因
	DriverId                    uint32          `json:"driver_id" gorm:"driver_id"`                                           // 司机ID
	BuyerId                     uint64          `json:"buyer_id" gorm:"buyer_id"`                                             // 买家ID
	ShipmentId                  string          `json:"shipment_id" gorm:"shipment_id"`                                       // 运单ID
	ShipmentIds                 string          `json:"shipment_ids" gorm:"shipment_ids"`                                     // 运单ID列表
	Content                     string          `json:"content" gorm:"content"`                                               // 内容
	TemplateId                  uint64          `json:"template_id" gorm:"template_id"`                                       // 模板ID
	ChannelType                 uint32          `json:"channel_type" gorm:"channel_type"`                                     // 渠道类型
	RecipientType               uint32          `json:"recipient_type" gorm:"recipient_type"`                                 // 接收者类型：1-卖家，2-买家，0-两者
	Recipient                   string          `json:"recipient" gorm:"recipient"`                                           // 接收者（电话号码）
	ServiceTypes                json.RawMessage `json:"service_types" gorm:"service_types"`                                   // 服务类型数组，JSON格式
	Parameters                  json.RawMessage `json:"parameters" gorm:"parameters"`                                         // 模板参数占位符值，JSON格式
	TraceId                     string          `json:"trace_id" gorm:"trace_id"`                                             // 追踪ID
	Ctime                       int64           `json:"ctime" gorm:"ctime"`                                                   // 创建时间
	Mtime                       int64           `json:"mtime" gorm:"mtime"`                                                   // 修改时间
}

// TableName 获取数据库表名
func (m *AutoSendChatLogTab) TableName() string {
	return "auto_send_chat_log_tab"
}

// GetServiceTypes 获取服务类型数组
func (m *AutoSendChatLogTab) GetServiceTypes() ([]int, error) {
	var types []int
	if len(m.ServiceTypes) == 0 {
		return []int{9999}, nil // 默认值
	}
	err := json.Unmarshal(m.ServiceTypes, &types)
	return types, err
}

// SetServiceTypes 设置服务类型数组
func (m *AutoSendChatLogTab) SetServiceTypes(types []int) error {
	data, err := json.Marshal(types)
	if err != nil {
		return err
	}
	m.ServiceTypes = data
	return nil
}

// GetParameters 获取模板参数
func (m *AutoSendChatLogTab) GetParameters() (map[string]interface{}, error) {
	var params map[string]interface{}
	if len(m.Parameters) == 0 {
		return map[string]interface{}{}, nil
	}
	err := json.Unmarshal(m.Parameters, &params)
	return params, err
}

// SetParameters 设置模板参数
func (m *AutoSendChatLogTab) SetParameters(params map[string]interface{}) error {
	data, err := json.Marshal(params)
	if err != nil {
		return err
	}
	m.Parameters = data
	return nil
}

// AutomaticChatTemplateTabLog 自动聊天模板操作日志表
type AutomaticChatTemplateTabLog struct {
	ID                          uint64          `json:"id" gorm:"id"`                                                         // 主键ID
	OperationType               uint32          `json:"operation_type" gorm:"operation_type"`                                 // 操作类型：1-插入，2-更新，3-删除，4-更新模板状态
	Operator                    string          `json:"operator" gorm:"operator"`                                             // 操作者
	TemplateId                  uint64          `json:"template_id" gorm:"template_id"`                                       // 相关模板ID
	TemplateName                string          `json:"template_name" gorm:"template_name"`                                   // 聊天通知模板名称
	ChatMessageType             uint32          `json:"chat_message_type" gorm:"chat_message_type"`                           // 自动聊天业务类型
	Content                     string          `json:"content" gorm:"content"`                                               // 聊天模板内容
	TemplateStatus              uint32          `json:"template_status" gorm:"template_status"`                               // 模板状态：1-可用，2-不可用
	OrderDeliveryTrackingStatus uint32          `json:"order_delivery_tracking_status" gorm:"order_delivery_tracking_status"` // 订单配送跟踪状态
	AtlFlag                     uint32          `json:"atl_flag" gorm:"atl_flag"`                                             // ATL标志：0-两者，1-放门口，2-不放门口
	ConditionContent            string          `json:"condition_content" gorm:"condition_content"`                           // 聊天模板条件内容
	ChannelType                 uint32          `json:"channel_type" gorm:"channel_type"`                                     // 渠道类型
	RecipientTypes              json.RawMessage `json:"recipient_types" gorm:"recipient_types"`                               // 接收者类型数组，JSON格式
	ServiceTypes                json.RawMessage `json:"service_types" gorm:"service_types"`                                   // 服务类型数组，JSON格式
	NssShopIds                  json.RawMessage `json:"nss_shop_ids" gorm:"nss_shop_ids"`                                     // NSS店铺ID数组，JSON格式
	Ctime                       int64           `json:"ctime" gorm:"ctime"`                                                   // 创建时间
	Mtime                       int64           `json:"mtime" gorm:"mtime"`                                                   // 修改时间
}

// TableName 获取数据库表名
func (m *AutomaticChatTemplateTabLog) TableName() string {
	return "automatic_chat_template_tab_log"
}

// GetRecipientTypes 获取接收者类型数组
func (m *AutomaticChatTemplateTabLog) GetRecipientTypes() ([]int, error) {
	var types []int
	if len(m.RecipientTypes) == 0 {
		return []int{}, nil
	}
	err := json.Unmarshal(m.RecipientTypes, &types)
	return types, err
}

// SetRecipientTypes 设置接收者类型数组
func (m *AutomaticChatTemplateTabLog) SetRecipientTypes(types []int) error {
	data, err := json.Marshal(types)
	if err != nil {
		return err
	}
	m.RecipientTypes = data
	return nil
}

// GetServiceTypes 获取服务类型数组
func (m *AutomaticChatTemplateTabLog) GetServiceTypes() ([]int, error) {
	var types []int
	if len(m.ServiceTypes) == 0 {
		return []int{9999}, nil // 默认值
	}
	err := json.Unmarshal(m.ServiceTypes, &types)
	return types, err
}

// SetServiceTypes 设置服务类型数组
func (m *AutomaticChatTemplateTabLog) SetServiceTypes(types []int) error {
	data, err := json.Marshal(types)
	if err != nil {
		return err
	}
	m.ServiceTypes = data
	return nil
}

// GetNssShopIds 获取NSS店铺ID数组
func (m *AutomaticChatTemplateTabLog) GetNssShopIds() ([]int, error) {
	var ids []int
	if len(m.NssShopIds) == 0 {
		return []int{}, nil
	}
	err := json.Unmarshal(m.NssShopIds, &ids)
	return ids, err
}

// SetNssShopIds 设置NSS店铺ID数组
func (m *AutomaticChatTemplateTabLog) SetNssShopIds(ids []int) error {
	data, err := json.Marshal(ids)
	if err != nil {
		return err
	}
	m.NssShopIds = data
	return nil
}
