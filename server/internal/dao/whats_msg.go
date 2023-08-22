// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"hotgo/internal/dao/internal"
)

// internalWhatsMsgDao is internal type for wrapping internal DAO implements.
type internalWhatsMsgDao = *internal.WhatsMsgDao

// whatsMsgDao is the data access object for table whats_msg.
// You can define custom methods on it to extend its functionality as you wish.
type whatsMsgDao struct {
	internalWhatsMsgDao
}

var (
	// WhatsMsg is globally public accessible object for table whats_msg operations.
	WhatsMsg = whatsMsgDao{
		internal.NewWhatsMsgDao(),
	}
)

// Fill with you ideas below.