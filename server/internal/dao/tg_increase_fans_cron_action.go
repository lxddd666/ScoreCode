// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"hotgo/internal/dao/internal"
)

// internalTgIncreaseFansCronActionDao is internal type for wrapping internal DAO implements.
type internalTgIncreaseFansCronActionDao = *internal.TgIncreaseFansCronActionDao

// tgIncreaseFansCronActionDao is the data access object for table tg_increase_fans_cron_action.
// You can define custom methods on it to extend its functionality as you wish.
type tgIncreaseFansCronActionDao struct {
	internalTgIncreaseFansCronActionDao
}

var (
	// TgIncreaseFansCronAction is globally public accessible object for table tg_increase_fans_cron_action operations.
	TgIncreaseFansCronAction = tgIncreaseFansCronActionDao{
		internal.NewTgIncreaseFansCronActionDao(),
	}
)

// Fill with you ideas below.