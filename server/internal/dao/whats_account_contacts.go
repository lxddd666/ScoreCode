// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"hotgo/internal/dao/internal"
)

// internalWhatsAccountContactsDao is internal type for wrapping internal DAO implements.
type internalWhatsAccountContactsDao = *internal.WhatsAccountContactsDao

// whatsAccountContactsDao is the data access object for table whats_account_contacts.
// You can define custom methods on it to extend its functionality as you wish.
type whatsAccountContactsDao struct {
	internalWhatsAccountContactsDao
}

var (
	// WhatsAccountContacts is globally public accessible object for table whats_account_contacts operations.
	WhatsAccountContacts = whatsAccountContactsDao{
		internal.NewWhatsAccountContactsDao(),
	}
)

// Fill with you ideas below.
