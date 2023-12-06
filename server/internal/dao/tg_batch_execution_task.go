// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"hotgo/internal/dao/internal"
)

// internalTgBatchExecutionTaskDao is internal type for wrapping internal DAO implements.
type internalTgBatchExecutionTaskDao = *internal.TgBatchExecutionTaskDao

// tgBatchExecutionTaskDao is the data access object for table tg_batch_execution_task.
// You can define custom methods on it to extend its functionality as you wish.
type tgBatchExecutionTaskDao struct {
	internalTgBatchExecutionTaskDao
}

var (
	// TgBatchExecutionTask is globally public accessible object for table tg_batch_execution_task operations.
	TgBatchExecutionTask = tgBatchExecutionTaskDao{
		internal.NewTgBatchExecutionTaskDao(),
	}
)

// Fill with you ideas below.
