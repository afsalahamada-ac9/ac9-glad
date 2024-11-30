/*
 * Copyright 2024 AboveCloud9.AI Products and Services Private Limited
 * All rights reserved.
 * This code may not be used, copied, modified, or distributed without explicit permission.
 */

package repository

import (
	"database/sql"
	"encoding/json"
	"time"

	"sudhagar/glad/entity"
	"sudhagar/glad/pkg/util"
)

// CoursePGSQL mysql repo
type CoursePGSQL struct {
	db *sql.DB
}

// NewCoursePGSQL create new repository
func NewCoursePGSQL(db *sql.DB) *CoursePGSQL {
	return &CoursePGSQL{
		db: db,
	}
}

// Insert creates a course
func (r *CoursePGSQL) Create(e *entity.Course) (entity.ID, error) {
	addressJSON, err := json.Marshal(e.Address)
	if err != nil {
		return e.ID, err
	}

	stmt, err := r.db.Prepare(`
		INSERT INTO course
			(
				id, tenant_id, ext_id, center_id, product_id, name, notes, timezone, address, status,
			 	mode, max_attendees, num_attendees, created_at
			)
		VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14)`)
	if err != nil {
		return e.ID, err
	}
	_, err = stmt.Exec(
		e.ID,
		e.TenantID,
		e.ExtID,
		e.CenterID,
		e.ProductID,
		e.Name,
		e.Notes,
		e.Timezone,
		string(addressJSON),
		e.Status,
		e.Mode,
		e.MaxAttendees,
		e.NumAttendees,
		time.Now().Format("2006-01-02"),
	)
	if err != nil {
		return e.ID, err
	}
	err = stmt.Close()
	if err != nil {
		return e.ID, err
	}
	return e.ID, nil
}

// Get retrieves a course
func (r *CoursePGSQL) Get(id entity.ID) (*entity.Course, error) {
	stmt, err := r.db.Prepare(`
		SELECT id, tenant_id, ext_id, center_id, product_id, name, notes, timezone, address,
		status, mode, max_attendees, num_attendees, created_at
		FROM course
		WHERE id = $1;`)
	if err != nil {
		return nil, err
	}
	var c entity.Course
	var ext_id sql.NullString
	var name, notes, timezone, address_json, status, mode sql.NullString
	err = stmt.QueryRow(id).Scan(&c.ID, &c.TenantID, &ext_id, &c.CenterID, &c.ProductID, &name, &notes, &timezone,
		&address_json, &status, &mode, &c.MaxAttendees, &c.NumAttendees, &c.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	if address_json.Valid && address_json.String != "" {
		err = json.Unmarshal([]byte(address_json.String), &c.Address)
		if err != nil {
			return nil, err
		}
	}

	c.ExtID = &ext_id.String
	c.Name = name.String
	c.Notes = notes.String
	c.Timezone = timezone.String
	c.Status = entity.CourseStatus(status.String)
	c.Mode = entity.CourseMode(mode.String)

	return &c, nil
}

// Update updates a course
func (r *CoursePGSQL) Update(e *entity.Course) error {
	e.UpdatedAt = time.Now()
	addressJSON, err := json.Marshal(e.Address)
	if err != nil {
		return err
	}

	_, err = r.db.Exec(`
		UPDATE course SET center_id = $1, name = $2, notes = $3, timezone = $4, address = $5,
			status = $6, mode = $7, max_attendees = $8, num_attendees = $9,
			updated_at = $10, product_id = $11
		WHERE id = $12;
		`,
		e.CenterID, e.Name, e.Notes, e.Timezone, string(addressJSON), (e.Status), (e.Mode),
		e.MaxAttendees, e.NumAttendees, e.UpdatedAt.Format("2006-01-02"), e.ProductID,
		e.ID)
	if err != nil {
		return err
	}
	return nil
}

// Search searches courses
func (r *CoursePGSQL) Search(tenantID entity.ID, q string, page, limit int) ([]*entity.Course, error) {
	query := `
		SELECT id, tenant_id, ext_id, center_id, product_id, name, notes, timezone, address,
		status, mode, max_attendees, num_attendees, created_at
		FROM course
		WHERE tenant_id = $1 AND name LIKE $2
	`

	if page > 0 && limit > 0 {
		offset := (page - 1) * limit
		query += ` LIMIT $3 OFFSET $4;`
		stmt, err := r.db.Prepare(query)
		if err != nil {
			return nil, err
		}

		rows, err := stmt.Query(tenantID, "%"+q+"%", limit, offset)
		if err != nil {
			return nil, err
		}

		defer rows.Close()
		return r.scanRows(rows)
	}

	stmt, err := r.db.Prepare(query + ";")
	if err != nil {
		return nil, err
	}

	rows, err := stmt.Query(tenantID, "%"+q+"%")
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	return r.scanRows(rows)
}

// List lists courses
func (r *CoursePGSQL) List(tenantID entity.ID, page, limit int) ([]*entity.Course, error) {
	query := `
		SELECT id, tenant_id, ext_id, center_id, product_id, name, notes, timezone, address,
		status, mode, max_attendees, num_attendees, created_at
		FROM course
		WHERE tenant_id = $1`

	if page > 0 && limit > 0 {
		offset := (page - 1) * limit
		query += ` LIMIT $2 OFFSET $3;`
		stmt, err := r.db.Prepare(query)
		if err != nil {
			return nil, err
		}

		rows, err := stmt.Query(tenantID, limit, offset)
		if err != nil {
			return nil, err
		}

		defer rows.Close()
		return r.scanRows(rows)
	}

	stmt, err := r.db.Prepare(query + ";")
	if err != nil {
		return nil, err
	}

	rows, err := stmt.Query(tenantID)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	return r.scanRows(rows)
}

// Delete deletes a course
func (r *CoursePGSQL) Delete(id entity.ID) error {
	res, err := r.db.Exec(`DELETE FROM course WHERE id = $1;`, id)
	if err != nil {
		return err
	}

	if cnt, _ := res.RowsAffected(); cnt == 0 {
		return sql.ErrNoRows
	}

	return nil
}

// Get total courses
func (r *CoursePGSQL) GetCount(tenantID entity.ID) (int, error) {
	stmt, err := r.db.Prepare(`SELECT count(*) FROM course WHERE tenant_id = $1;`)
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

func (r *CoursePGSQL) scanRows(rows *sql.Rows) ([]*entity.Course, error) {
	var courses []*entity.Course
	for rows.Next() {
		var course entity.Course
		var ext_id, name, notes, timezone, address_json, status, mode sql.NullString
		err := rows.Scan(
			&course.ID,
			&course.TenantID,
			&ext_id,
			&course.CenterID,
			&course.ProductID,
			&name,
			&notes,
			&timezone,
			&address_json,
			&status,
			&mode,
			&course.MaxAttendees,
			&course.NumAttendees,
			&course.CreatedAt,
		)
		if err != nil {
			return nil, err
		}

		course.ExtID = &ext_id.String
		course.Name = name.String
		course.Notes = notes.String
		course.Timezone = timezone.String
		course.Status = entity.CourseStatus(status.String)
		course.Mode = entity.CourseMode(mode.String)

		if address_json.Valid && address_json.String != "" {
			err = json.Unmarshal([]byte(address_json.String), &course.Address)
			if err != nil {
				return nil, err
			}
		}

		courses = append(courses, &course)
	}
	return courses, nil
}

// --------------------------------------------------------------------------------
// Course Organizer
// --------------------------------------------------------------------------------
// InsertCourseOrganizer creates course to organizer mapping
// TODO: We should map error value appropriately to the API client. Ex. foreign key violation implies
// that request is a bad request, etc. or, maybe we don't. Something to think about.
func (r *CoursePGSQL) InsertCourseOrganizer(courseID entity.ID, cos []*entity.CourseOrganizer) error {
	values := func(index int) []interface{} {
		return []interface{}{
			courseID,
			cos[index].ID,
			util.DBTimeNow(),
		}
	}

	query, valueArgs := util.GenBulkInsertPGSQL(
		"course_organizer",
		[]string{"course_id", "organizer_id", "updated_at"},
		len(cos),
		values,
	)

	stmt, err := r.db.Prepare(query)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(valueArgs...)
	if err != nil {
		return err
	}

	err = stmt.Close()
	return err
}

// GetCourseOrganizer gets course organizer for the given course id
func (r *CoursePGSQL) GetCourseOrganizer(courseID entity.ID) ([]*entity.CourseOrganizer, error) {
	query := `SELECT organizer_id FROM course_organizer where course_id = $1;`

	stmt, err := r.db.Prepare(query)
	if err != nil {
		return nil, err
	}

	rows, err := stmt.Query(courseID)
	if err != nil {
		return nil, err
	}

	var cos []*entity.CourseOrganizer
	for rows.Next() {
		var co entity.CourseOrganizer

		err := rows.Scan(&co.ID)
		if err != nil {
			return nil, err
		}

		cos = append(cos, &co)
	}

	defer rows.Close()
	return cos, err
}

// UpdateCourseOrganizer updates course organizer for the given course id and the organizer
func (r *CoursePGSQL) UpdateCourseOrganizer(courseID entity.ID, cos []*entity.CourseOrganizer) error {
	// Note: It's possible we could use reflection and make this
	// split of add and remove items more generic.
	currentCOs, err := r.GetCourseOrganizer(courseID)
	if err != nil {
		return err
	}

	mapOrgID := make(map[entity.ID]bool)
	for _, co := range currentCOs {
		mapOrgID[co.ID] = true
	}

	var addCOs []*entity.CourseOrganizer
	var rmCOs []*entity.CourseOrganizer
	for _, co := range cos {
		if _, exists := mapOrgID[co.ID]; exists {
			delete(mapOrgID, co.ID)
		} else {
			addCOs = append(addCOs, co)
		}
	}

	for id := range mapOrgID {
		co := entity.CourseOrganizer{
			ID: id,
		}
		rmCOs = append(rmCOs, &co)
	}

	if len(addCOs) > 0 {
		err = r.InsertCourseOrganizer(courseID, addCOs)
		if err != nil {
			return err
		}
	}

	if len(rmCOs) > 0 {
		err = r.DeleteCourseOrganizer(courseID, rmCOs)
		if err != nil {
			return err
		}

	}

	return err
}

// DeleteCourseOrganizer deletes the given course organizers
func (r *CoursePGSQL) DeleteCourseOrganizer(courseID entity.ID, cos []*entity.CourseOrganizer) error {
	values := func(index int) []interface{} {
		return []interface{}{
			courseID,
			cos[index].ID,
		}
	}

	query, valueArgs := util.GenBulkDeletePGSQL(
		"course_organizer",
		[]string{"course_id", "organizer_id"},
		len(cos),
		values,
	)

	stmt, err := r.db.Prepare(query)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(valueArgs...)
	if err != nil {
		return err
	}

	err = stmt.Close()
	return err
}

// DeleteCourseOrganizerByCourse deletes course organizers using course id
func (r *CoursePGSQL) DeleteCourseOrganizerByCourse(courseID entity.ID) error {
	res, err := r.db.Exec(`DELETE FROM course_organizer WHERE course_id = $1;`, courseID)
	if err != nil {
		return err
	}

	if cnt, _ := res.RowsAffected(); cnt == 0 {
		return sql.ErrNoRows
	}

	return nil
}

// --------------------------------------------------------------------------------
// Course Teacher
// --------------------------------------------------------------------------------
// InsertCourseTeacher creates course to teacher mapping
func (r *CoursePGSQL) InsertCourseTeacher(courseID entity.ID, cts []*entity.CourseTeacher) error {
	values := func(index int) []interface{} {
		return []interface{}{
			courseID,
			cts[index].ID,
			cts[index].IsPrimary,
			util.DBTimeNow(),
		}
	}

	query, valueArgs := util.GenBulkInsertPGSQL(
		"course_teacher",
		[]string{"course_id", "teacher_id", "is_primary", "updated_at"},
		len(cts),
		values,
	)

	stmt, err := r.db.Prepare(query)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(valueArgs...)
	if err != nil {
		return err
	}

	err = stmt.Close()
	return err
}

// --------------------------------------------------------------------------------
// Course Contact
// --------------------------------------------------------------------------------
// InsertCourseContact creates course to contact mapping
func (r *CoursePGSQL) InsertCourseContact(courseID entity.ID, ccs []*entity.CourseContact) error {
	values := func(index int) []interface{} {
		return []interface{}{
			courseID,
			ccs[index].ID,
			util.DBTimeNow(),
		}
	}

	query, valueArgs := util.GenBulkInsertPGSQL(
		"course_contact",
		[]string{"course_id", "contact_id", "updated_at"},
		len(ccs),
		values,
	)

	stmt, err := r.db.Prepare(query)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(valueArgs...)
	if err != nil {
		return err
	}

	err = stmt.Close()
	return err
}

// --------------------------------------------------------------------------------
// Course Notify
// --------------------------------------------------------------------------------
// InsertCourseNotify creates course to notify mapping
func (r *CoursePGSQL) InsertCourseNotify(courseID entity.ID, cns []*entity.CourseNotify) error {
	values := func(index int) []interface{} {
		return []interface{}{
			courseID,
			cns[index].ID,
			util.DBTimeNow(),
		}
	}

	query, valueArgs := util.GenBulkInsertPGSQL(
		"course_notify",
		[]string{"course_id", "notify_id", "updated_at"},
		len(cns),
		values,
	)

	stmt, err := r.db.Prepare(query)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(valueArgs...)
	if err != nil {
		return err
	}

	err = stmt.Close()
	return err
}
