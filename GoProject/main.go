package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// ErrorInfo 定义JSON数据结构
type ErrorInfo struct {
	ApplicationType    *string `json:"application_type"`
	BusinessType       *string `json:"business_type"`
	FunctionalCategory *string `json:"functional_category"`
	URL                *string `json:"url"`
	ErrorCode          *int    `json:"error_code"`
	ErrorMsg           *string `json:"error_msg"`
	NewErrorCode       *int    `json:"new_error_code"`
	IfFeUsed           *string `json:"if_fe_used"`
	ErrorCategory      *string `json:"error_category"`
	UrgencyDegreeType  *string `json:"urgency_degree_type"`
	ErrorReasonList    *string `json:"error_reason_list"`
	ErrorAlarmStrategy *string `json:"error_alarm_strategy"`
	ErrorAlarmSop      *string `json:"error_alarm_sop"`
	ErrorExtraInfo     *string `json:"error_extra_info"`
}

// RecordErrorInfoTab 定义数据库表结构
type RecordErrorInfoTab struct {
	Id                 uint64 `gorm:"column:id;type:bigint(20) unsigned;primary_key;AUTO_INCREMENT;comment:id" json:"id"`
	ApplicationType    string `gorm:"column:application_type;type:varchar(128);default:'';comment:application type like fms;NOT NULL" json:"application_type"`
	BusinessType       string `gorm:"column:business_type;type:varchar(128);default:'';comment:business type like login;NOT NULL" json:"business_type"`
	FunctionalCategory string `gorm:"column:functional_category;type:varchar(128);default:'';comment:functional category like goggle login;NOT NULL" json:"functional_category"`
	Url                string `gorm:"column:url;type:varchar(128);default:'';comment:url info;NOT NULL" json:"url"`
	ErrorCode          int    `gorm:"column:error_code;type:int(10) unsigned;default:0;comment:error_code info;NOT NULL" json:"error_code"`
	ErrorMsg           string `gorm:"column:error_msg;type:varchar(512);default:'';comment:error_msg info;NOT NULL" json:"error_msg"`
	NewErrorCode       int    `gorm:"column:new_error_code;type:int(10) unsigned;default:0;comment:new error_code info;NOT NULL" json:"new_error_code"`
	IfFeUsed           string `gorm:"column:if_fe_used;type:varchar(128);default:'';comment:fe if use error code,0:not use,1:used;NOT NULL" json:"if_fe_used"`
	ErrorCategory      string `gorm:"column:error_category;type:varchar(128);default:'';comment:error category like db error,buiness error;NOT NULL" json:"error_category"`
	UrgencyDegreeType  int    `gorm:"column:urgency_degree_type;type:int(3) unsigned;default:0;comment:urgency degree type like high;NOT NULL" json:"urgency_degree_type"`
	ErrorReasonList    string `gorm:"column:error_reason_list;type:varchar(512);default:'';comment:error source;NOT NULL" json:"error_reason_list"`
	ErrorAlarmStrategy string `gorm:"column:error_alarm_strategy;type:varchar(512);default:'';comment:alarm strategy;NOT NULL" json:"error_alarm_strategy"`
	ErrorAlarmSop      string `gorm:"column:error_alarm_sop;type:varchar(512);default:'';comment:alarm aop;NOT NULL" json:"error_alarm_sop"`
	ErrorExtraInfo     string `gorm:"column:error_extra_info;type:varchar(512);default:'';comment:extra add error information;NOT NULL" json:"error_extra_info"`
	Ctime              int64  `gorm:"column:ctime;type:int(10) unsigned;default:0;comment:citme;NOT NULL" json:"ctime"`
	Mtime              int64  `gorm:"column:mtime;type:int(10) unsigned;default:0;comment:mtime;NOT NULL" json:"mtime"`
}

func main() {
	// 数据库连接配置
	dsn := "user:password@tcp(127.0.0.1:3306)/database_name?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// 读取JSON文件
	jsonFile, err := os.ReadFile("Account告警准确性治理_20250423_115543.json")
	if err != nil {
		log.Fatalf("Failed to read JSON file: %v", err)
	}

	// 解析JSON数据到ErrorInfo结构体
	var errorInfos []ErrorInfo
	err = json.Unmarshal(jsonFile, &errorInfos)
	if err != nil {
		log.Fatalf("Failed to parse JSON: %v", err)
	}

	// 转换ErrorInfo到RecordErrorInfoTab
	var records []RecordErrorInfoTab
	now := time.Now().Unix()

	for _, info := range errorInfos {
		record := RecordErrorInfoTab{
			Ctime: now,
			Mtime: now,
		}

		// 设置字段值，处理指针类型
		if info.ApplicationType != nil {
			// 这里需要根据实际业务逻辑将字符串转换为对应的int值
			record.ApplicationType = *info.ApplicationType
		}
		if info.BusinessType != nil {
			record.BusinessType = *info.BusinessType
		}
		if info.FunctionalCategory != nil {
			record.FunctionalCategory = *info.FunctionalCategory
		}
		if info.URL != nil {
			record.Url = *info.URL
		}
		if info.ErrorCode != nil {
			record.ErrorCode = *info.ErrorCode
		}
		if info.ErrorMsg != nil {
			record.ErrorMsg = *info.ErrorMsg
		}
		if info.NewErrorCode != nil {
			record.NewErrorCode = *info.NewErrorCode
		}
		if info.IfFeUsed != nil {
			record.IfFeUsed = *info.IfFeUsed
		}
		if info.ErrorCategory != nil {
			record.ErrorCategory = *info.ErrorCategory
		}
		if info.UrgencyDegreeType != nil {
			record.UrgencyDegreeType = 1 // 示例值，需要根据实际业务逻辑修改
		}
		if info.ErrorReasonList != nil {
			record.ErrorReasonList = *info.ErrorReasonList
		}
		if info.ErrorAlarmStrategy != nil {
			record.ErrorAlarmStrategy = *info.ErrorAlarmStrategy
		}
		if info.ErrorAlarmSop != nil {
			record.ErrorAlarmSop = *info.ErrorAlarmSop
		}
		if info.ErrorExtraInfo != nil {
			record.ErrorExtraInfo = *info.ErrorExtraInfo
		}

		records = append(records, record)
	}

	// 批量插入数据
	result := db.Create(&records)
	if result.Error != nil {
		log.Fatalf("Failed to insert records: %v", result.Error)
	}

	fmt.Printf("Successfully inserted %d records\n", result.RowsAffected)
}
