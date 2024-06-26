// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"sagooiot/internal/dao/internal"
)

// internalDevDeviceDao is internal type for wrapping internal DAO implements.
type internalDevDeviceDao = *internal.DevDeviceDao

// devDeviceDao is the data access object for table dev_device.
// You can define custom methods on it to extend its functionality as you wish.
type devDeviceDao struct {
	internalDevDeviceDao
}

var (
	// DevDevice is globally public accessible object for table dev_device operations.
	DevDevice = devDeviceDao{
		internal.NewDevDeviceDao(),
	}
)

// Fill with you ideas below.
