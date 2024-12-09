/*
 * Copyright 2024 AboveCloud9.AI Products and Services Private Limited
 * All rights reserved.
 * This code may not be used, copied, modified, or distributed without explicit permission.
 */

package repository

import (
	"database/sql"

	"ac9/glad/pkg/id"
	"ac9/glad/pkg/util"
	"ac9/glad/services/pushd/entity"
)

// DevicePGSQL mysql repo
type DevicePGSQL struct {
	db *sql.DB
}

// NewDevicePGSQL create new repository
func NewDevicePGSQL(db *sql.DB) *DevicePGSQL {
	return &DevicePGSQL{
		db: db,
	}

}

// Insert creates a device
func (r *DevicePGSQL) Create(d *entity.Device) (id.ID, error) {
	stmt, err := r.db.Prepare(`
		INSERT INTO device
			(
				id, tenant_id, account_id, push_token, revoke_id, app_version,
				device_info, platform_info,
			 	created_at, updated_at
			)
		VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)`)
	if err != nil {
		return d.ID, err
	}
	_, err = stmt.Exec(
		d.ID,
		d.TenantID,
		d.AccountID,
		d.PushToken,
		d.RevokeID,
		d.AppVersion,
		d.DeviceInfo,
		d.PlatformInfo,
		util.DBTimeNow(),
		util.DBTimeNow(),
	)
	if err != nil {
		return d.ID, err
	}
	err = stmt.Close()
	if err != nil {
		return d.ID, err
	}
	return d.ID, nil
}

// GetByAccount retrieves a device using account id
func (r *DevicePGSQL) GetByAccount(tenantID id.ID, accountID id.ID) ([]*entity.Device, error) {
	stmt, err := r.db.Prepare(`
		SELECT
			id, tenant_id, account_id, push_token, revoke_id,
			app_version, device_info, platform_info
			created_at, updated_at
		FROM device
		WHERE tenant_id = $1 AND account_id = $2;`)
	if err != nil {
		return nil, err
	}

	rows, err := stmt.Query(tenantID, accountID)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	return r.scanRows(rows)
}

// // GetTokenByAccount retrieves device tokens using account id
// func (r *DevicePGSQL) GetTokenByAccount(tenantID id.ID, accountID id.ID) ([]*entity.Device, error) {
// 	stmt, err := r.db.Prepare(`
// 		SELECT
// 			id, token
// 		FROM device
// 		WHERE tenant_id = $1 AND account_id = $2;`)
// 	if err != nil {
// 		return nil, err
// 	}

// 	rows, err := stmt.Query(tenantID, accountID)
// 	if err != nil {
// 		return nil, err
// 	}

// 	defer rows.Close()
// 	return r.scanRows(rows)
// }

// Delete deletes a device
func (r *DevicePGSQL) Delete(id id.ID) error {
	res, err := r.db.Exec(`DELETE FROM device WHERE id = $1;`, id)
	if err != nil {
		return err
	}

	if cnt, _ := res.RowsAffected(); cnt == 0 {
		return sql.ErrNoRows
	}

	return nil
}

// Get total devices
func (r *DevicePGSQL) GetCount(tenantID id.ID) (int, error) {
	stmt, err := r.db.Prepare(`SELECT count(*) FROM device WHERE tenant_id = $1;`)
	if err != nil {
		return 0, err
	}

	var count int
	err = stmt.QueryRow(tenantID).Scan(&count)
	if err != nil {
		return 0, err
	}

	return count, nil
}

func (r *DevicePGSQL) scanRows(rows *sql.Rows) ([]*entity.Device, error) {
	var devices []*entity.Device

	for rows.Next() {
		var device entity.Device
		var pushToken, revokeID, appVersion, deviceInfo, platformInfo sql.NullString
		err := rows.Scan(
			&device.ID,
			&device.TenantID,
			&device.AccountID,
			&pushToken,
			&revokeID,
			&appVersion,
			&deviceInfo,
			&platformInfo,
			&device.CreatedAt,
			&device.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		device.PushToken = pushToken.String
		device.RevokeID = &revokeID.String
		device.AppVersion = appVersion.String
		device.DeviceInfo = &deviceInfo.String
		device.PlatformInfo = &platformInfo.String

		devices = append(devices, &device)
	}
	return devices, nil
}
