package noti

import (
	"encoding/json"
	"testing"
)

func TestAutomaticChatTemplateTab_JSONMethods(t *testing.T) {
	template := &AutomaticChatTemplateTab{
		ID:             1,
		TemplateName:   "测试模板",
		ChannelType:    ChannelTypeChat,
		RecipientTypes: json.RawMessage(`[1,2]`),
		ServiceTypes:   json.RawMessage(`[1001,1002]`),
		NssShopIds:     json.RawMessage(`[100,200]`),
	}

	// 测试获取接收者类型
	recipientTypes, err := template.GetRecipientTypes()
	if err != nil {
		t.Errorf("GetRecipientTypes failed: %v", err)
	}
	if len(recipientTypes) != 2 || recipientTypes[0] != 1 || recipientTypes[1] != 2 {
		t.Errorf("GetRecipientTypes returned wrong values: %v", recipientTypes)
	}

	// 测试设置接收者类型
	newTypes := []int{3, 4}
	err = template.SetRecipientTypes(newTypes)
	if err != nil {
		t.Errorf("SetRecipientTypes failed: %v", err)
	}

	// 验证设置结果
	recipientTypes, err = template.GetRecipientTypes()
	if err != nil {
		t.Errorf("GetRecipientTypes after set failed: %v", err)
	}
	if len(recipientTypes) != 2 || recipientTypes[0] != 3 || recipientTypes[1] != 4 {
		t.Errorf("GetRecipientTypes after set returned wrong values: %v", recipientTypes)
	}

	// 测试获取服务类型
	serviceTypes, err := template.GetServiceTypes()
	if err != nil {
		t.Errorf("GetServiceTypes failed: %v", err)
	}
	if len(serviceTypes) != 2 || serviceTypes[0] != 1001 || serviceTypes[1] != 1002 {
		t.Errorf("GetServiceTypes returned wrong values: %v", serviceTypes)
	}

	// 测试获取NSS店铺ID
	shopIds, err := template.GetNssShopIds()
	if err != nil {
		t.Errorf("GetNssShopIds failed: %v", err)
	}
	if len(shopIds) != 2 || shopIds[0] != 100 || shopIds[1] != 200 {
		t.Errorf("GetNssShopIds returned wrong values: %v", shopIds)
	}
}

func TestAutoSendChatLogTab_JSONMethods(t *testing.T) {
	log := &AutoSendChatLogTab{
		ID:            1,
		MessageId:     "msg_001",
		SendStatus:    SendStatusPending,
		ChannelType:   ChannelTypeSMS,
		RecipientType: RecipientTypeBuyer,
		Recipient:     "13800138000",
		ServiceTypes:  json.RawMessage(`[2001,2002]`),
		Parameters:    json.RawMessage(`{"name":"张三","order":"12345"}`),
		TraceId:       "trace_001",
	}

	// 测试获取服务类型
	serviceTypes, err := log.GetServiceTypes()
	if err != nil {
		t.Errorf("GetServiceTypes failed: %v", err)
	}
	if len(serviceTypes) != 2 || serviceTypes[0] != 2001 || serviceTypes[1] != 2002 {
		t.Errorf("GetServiceTypes returned wrong values: %v", serviceTypes)
	}

	// 测试获取参数
	params, err := log.GetParameters()
	if err != nil {
		t.Errorf("GetParameters failed: %v", err)
	}
	if params["name"] != "张三" || params["order"] != "12345" {
		t.Errorf("GetParameters returned wrong values: %v", params)
	}

	// 测试设置参数
	newParams := map[string]interface{}{
		"name":  "李四",
		"order": "67890",
	}
	err = log.SetParameters(newParams)
	if err != nil {
		t.Errorf("SetParameters failed: %v", err)
	}

	// 验证设置结果
	params, err = log.GetParameters()
	if err != nil {
		t.Errorf("GetParameters after set failed: %v", err)
	}
	if params["name"] != "李四" || params["order"] != "67890" {
		t.Errorf("GetParameters after set returned wrong values: %v", params)
	}
}

func TestAutomaticChatTemplateTabLog_JSONMethods(t *testing.T) {
	log := &AutomaticChatTemplateTabLog{
		ID:             1,
		OperationType:  OperationTypeInsert,
		Operator:       "admin",
		TemplateId:     100,
		ChannelType:    ChannelTypeEmail,
		RecipientTypes: json.RawMessage(`[1]`),
		ServiceTypes:   json.RawMessage(`[3001]`),
		NssShopIds:     json.RawMessage(`[500,600]`),
	}

	// 测试获取接收者类型
	recipientTypes, err := log.GetRecipientTypes()
	if err != nil {
		t.Errorf("GetRecipientTypes failed: %v", err)
	}
	if len(recipientTypes) != 1 || recipientTypes[0] != 1 {
		t.Errorf("GetRecipientTypes returned wrong values: %v", recipientTypes)
	}

	// 测试获取服务类型
	serviceTypes, err := log.GetServiceTypes()
	if err != nil {
		t.Errorf("GetServiceTypes failed: %v", err)
	}
	if len(serviceTypes) != 1 || serviceTypes[0] != 3001 {
		t.Errorf("GetServiceTypes returned wrong values: %v", serviceTypes)
	}

	// 测试获取NSS店铺ID
	shopIds, err := log.GetNssShopIds()
	if err != nil {
		t.Errorf("GetNssShopIds failed: %v", err)
	}
	if len(shopIds) != 2 || shopIds[0] != 500 || shopIds[1] != 600 {
		t.Errorf("GetNssShopIds returned wrong values: %v", shopIds)
	}
}

func TestConstants(t *testing.T) {
	// 测试渠道类型常量
	if ChannelTypeSMS != 1 {
		t.Errorf("ChannelTypeSMS should be 1, got %d", ChannelTypeSMS)
	}
	if ChannelTypeEmail != 2 {
		t.Errorf("ChannelTypeEmail should be 2, got %d", ChannelTypeEmail)
	}
	if ChannelTypePush != 3 {
		t.Errorf("ChannelTypePush should be 3, got %d", ChannelTypePush)
	}
	if ChannelTypeChat != 12 {
		t.Errorf("ChannelTypeChat should be 12, got %d", ChannelTypeChat)
	}

	// 测试接收者类型常量
	if RecipientTypeBoth != 0 {
		t.Errorf("RecipientTypeBoth should be 0, got %d", RecipientTypeBoth)
	}
	if RecipientTypeSeller != 1 {
		t.Errorf("RecipientTypeSeller should be 1, got %d", RecipientTypeSeller)
	}
	if RecipientTypeBuyer != 2 {
		t.Errorf("RecipientTypeBuyer should be 2, got %d", RecipientTypeBuyer)
	}

	// 测试模板状态常量
	if TemplateStatusAvailable != 1 {
		t.Errorf("TemplateStatusAvailable should be 1, got %d", TemplateStatusAvailable)
	}
	if TemplateStatusUnavailable != 2 {
		t.Errorf("TemplateStatusUnavailable should be 2, got %d", TemplateStatusUnavailable)
	}

	// 测试发送状态常量
	if SendStatusPending != 0 {
		t.Errorf("SendStatusPending should be 0, got %d", SendStatusPending)
	}
	if SendStatusSuccess != 1 {
		t.Errorf("SendStatusSuccess should be 1, got %d", SendStatusSuccess)
	}
	if SendStatusFail != 3 {
		t.Errorf("SendStatusFail should be 3, got %d", SendStatusFail)
	}

	// 测试操作类型常量
	if OperationTypeInsert != 1 {
		t.Errorf("OperationTypeInsert should be 1, got %d", OperationTypeInsert)
	}
	if OperationTypeUpdate != 2 {
		t.Errorf("OperationTypeUpdate should be 2, got %d", OperationTypeUpdate)
	}
	if OperationTypeDelete != 3 {
		t.Errorf("OperationTypeDelete should be 3, got %d", OperationTypeDelete)
	}
	if OperationTypeUpdateStatus != 4 {
		t.Errorf("OperationTypeUpdateStatus should be 4, got %d", OperationTypeUpdateStatus)
	}
}
