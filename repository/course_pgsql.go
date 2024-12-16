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

	"ac9/glad/entity"
	"ac9/glad/pkg/common"
	"ac9/glad/pkg/id"
	l "ac9/glad/pkg/logger"
	"ac9/glad/pkg/util"
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
func (r *CoursePGSQL) Create(e *entity.Course) (id.ID, error) {
	jsonAddress, err := json.Marshal(e.Address)
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
		string(jsonAddress),
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
func (r *CoursePGSQL) Get(id id.ID) (*entity.Course, error) {
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
	var name, notes, timezone, jsonAddress, status, mode sql.NullString
	err = stmt.QueryRow(id).Scan(&c.ID, &c.TenantID, &ext_id, &c.CenterID, &c.ProductID, &name, &notes, &timezone,
		&jsonAddress, &status, &mode, &c.MaxAttendees, &c.NumAttendees, &c.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	if jsonAddress.Valid && jsonAddress.String != "" {
		err = json.Unmarshal([]byte(jsonAddress.String), &c.Address)
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
	jsonAddress, err := json.Marshal(e.Address)
	if err != nil {
		return err
	}

	_, err = r.db.Exec(`
		UPDATE course SET center_id = $1, name = $2, notes = $3, timezone = $4, address = $5,
			status = $6, mode = $7, max_attendees = $8, num_attendees = $9,
			updated_at = $10, product_id = $11
		WHERE id = $12;
		`,
		e.CenterID, e.Name, e.Notes, e.Timezone, string(jsonAddress), (e.Status), (e.Mode),
		e.MaxAttendees, e.NumAttendees, e.UpdatedAt.Format("2006-01-02"), e.ProductID,
		e.ID)
	if err != nil {
		return err
	}
	return nil
}

// Search searches courses
func (r *CoursePGSQL) Search(tenantID id.ID, q string, page, limit int) ([]*entity.Course, error) {
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
func (r *CoursePGSQL) List(tenantID id.ID, page, limit int) ([]*entity.Course, error) {
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

// GetByAccount retrieves courses by account id
func (r *CoursePGSQL) GetByAccount(tenantID id.ID,
	accountID id.ID,
	page, limit int,
) (int, []*entity.Course, error) {

	// Get total count that matches this query
	count, err := r.getCountByAccount(tenantID, accountID)
	if err != nil {
		l.Log.Warnf("%v", err)
		return count, nil, err
	}

	// Get all courses limit X
	courseList, err := r.getByAccount(tenantID, accountID, page, limit)
	if err != nil {
		l.Log.Warnf("%v", err)
		return count, nil, err
	}

	return count, courseList, err
}

// Delete deletes a course
func (r *CoursePGSQL) Delete(id id.ID) error {
	res, err := r.db.Exec(`DELETE FROM course WHERE id = $1;`, id)
	if err != nil {
		l.Log.Warnf("%v", err)
		return err
	}

	if cnt, _ := res.RowsAffected(); cnt == 0 {
		return sql.ErrNoRows
	}

	return nil
}

// GetCount total courses
func (r *CoursePGSQL) GetCount(tenantID id.ID) (int, error) {
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

// getCountByAccount total courses
func (r *CoursePGSQL) getCountByAccount(tenantID id.ID,
	accountID id.ID,
) (int, error) {
	stmt, err := r.db.Prepare(`
		SELECT count(*)
		FROM course
		WHERE
			tenant_id = $1 AND
			id IN
			(
				SELECT course_id FROM course_organizer
					WHERE organizer_id = $2
					UNION
				SELECT course_id FROM course_teacher
					WHERE teacher_id = $2
					UNION
				SELECT course_id FROM course_contact
					WHERE contact_id = $2
					UNION
				SELECT course_id FROM course_notify
					WHERE notify_id = $2
			)
		;
	`)
	if err != nil {
		l.Log.Warnf("err=%v", err)
		return 0, err
	}

	var count int
	err = stmt.QueryRow(tenantID, accountID).Scan(&count)
	if err != nil {
		l.Log.Warnf("err=%v", err)
		return 0, err
	}

	return count, nil
}

// getByAccount retrieves courses that match account id
func (r *CoursePGSQL) getByAccount(tenantID id.ID,
	accountID id.ID,
	page, limit int,
) ([]*entity.Course, error) {
	query := `
		SELECT
			id, tenant_id, ext_id, center_id, product_id, name,
			notes, timezone, address, status, mode, max_attendees,
			num_attendees, created_at
		FROM course
		WHERE
			tenant_id = $1
		AND
			id IN
			(
				SELECT course_id FROM course_organizer
					WHERE organizer_id = $2
					UNION
				SELECT course_id FROM course_teacher
					WHERE teacher_id = $2
					UNION
				SELECT course_id FROM course_contact
					WHERE contact_id = $2
					UNION
				SELECT course_id FROM course_notify
					WHERE notify_id = $2
	`

	offset := 0
	if page > 0 {
		offset = (page - 1) * limit
	}
	query += ` LIMIT $3 OFFSET $4) ORDER BY updated_at DESC;`
	stmt, err := r.db.Prepare(query)
	if err != nil {
		l.Log.Warnf("err=%v", err)
		return nil, err
	}

	rows, err := stmt.Query(tenantID, accountID, limit, offset)
	if err != nil {
		l.Log.Warnf("err=%v", err)
		return nil, err
	}

	defer rows.Close()
	return r.scanRows(rows)
}

func (r *CoursePGSQL) scanRows(rows *sql.Rows) ([]*entity.Course, error) {
	var courses []*entity.Course
	for rows.Next() {
		var course entity.Course
		var ext_id, name, notes, timezone, jsonAddress, status, mode sql.NullString
		err := rows.Scan(
			&course.ID,
			&course.TenantID,
			&ext_id,
			&course.CenterID,
			&course.ProductID,
			&name,
			&notes,
			&timezone,
			&jsonAddress,
			&status,
			&mode,
			&course.MaxAttendees,
			&course.NumAttendees,
			&course.CreatedAt,
		)
		if err != nil {
			l.Log.Warnf("err=%v", err)
			return nil, err
		}

		course.ExtID = &ext_id.String
		course.Name = name.String
		course.Notes = notes.String
		course.Timezone = timezone.String
		course.Status = entity.CourseStatus(status.String)
		course.Mode = entity.CourseMode(mode.String)

		if jsonAddress.Valid && jsonAddress.String != "" {
			err = json.Unmarshal([]byte(jsonAddress.String), &course.Address)
			if err != nil {
				l.Log.Warnf("err=%v", err)
				return nil, err
			}
		}

		courses = append(courses, &course)
	}
	return courses, nil
}

// Upsert inserts or updates the course and returns the id
func (r *CoursePGSQL) Upsert(e *entity.Course) (id.ID, error) {
	stmt, err := r.db.Prepare(`
		WITH upsert AS (
			INSERT INTO course (
				id, ext_id, tenant_id, product_id, center_id,
				name, notes, timezone,
				address, status, mode,
				max_attendees, num_attendees,
				url, checkout_url,
				created_at, updated_at
			)
			VALUES
				($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17) 
			ON CONFLICT (ext_id)
			DO UPDATE
				SET tenant_id = $3, product_id = $4, center_id = $5, name = $6, notes = $7,
				 	timezone = $8, address = $9, status = $10, mode = $11,
					max_attendees = $12, num_attendees = $13, url = $14, checkout_url = $15,
					created_at = $16, updated_at = $17
			WHERE course.updated_at <= $17
			RETURNING id
		)
		SELECT id FROM upsert
		UNION ALL
		SELECT id FROM course WHERE ext_id = $2 AND NOT EXISTS (SELECT 1 FROM upsert);
	`)
	if err != nil {
		l.Log.Warnf("err=%v", err)
		return id.IDInvalid, err
	}

	// TODO: Retrieve productID, centerID using the extID for product & center
	// + accounts as well

	jsonAddress, err := json.Marshal(e.Address)
	if err != nil {
		l.Log.Warnf("err=%v", err)
		return id.IDInvalid, err
	}
	l.Log.Debugf("Address=%#v jsonAddress=%v", e.Address, jsonAddress)

	var courseId id.ID
	err = stmt.QueryRow(
		e.ID,
		e.ExtID,
		e.TenantID,
		e.ProductID,
		e.CenterID,
		e.Name,
		e.Notes,
		e.Timezone,
		jsonAddress,
		string(e.Status),
		string(e.Mode),
		e.MaxAttendees,
		e.NumAttendees,
		e.URL,
		e.CheckoutURL,
		e.CreatedAt.Format(common.DBFormatDateTimeMS),
		e.UpdatedAt.Format(common.DBFormatDateTimeMS),
	).Scan(&courseId)
	if err != nil {
		l.Log.Warnf("err=%v", err)
		return id.IDInvalid, err
	}

	return courseId, nil
}

// --------------------------------------------------------------------------------
// Course Organizer
// --------------------------------------------------------------------------------
// InsertCourseOrganizer creates course to organizer mapping
// TODO: We should map error value appropriately to the API client. Ex. foreign key violation implies
// that request is a bad request, etc. or, maybe we don't. Something to think about.
func (r *CoursePGSQL) InsertCourseOrganizer(courseID id.ID, cos []*entity.CourseOrganizer) error {
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
func (r *CoursePGSQL) GetCourseOrganizer(courseID id.ID) ([]*entity.CourseOrganizer, error) {
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
func (r *CoursePGSQL) UpdateCourseOrganizer(courseID id.ID, cos []*entity.CourseOrganizer) error {
	// Note: It's possible we could use reflection and make this
	// split of add and remove items more generic.
	currentCOs, err := r.GetCourseOrganizer(courseID)
	if err != nil {
		return err
	}

	mapOrgID := make(map[id.ID]bool)
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

// DeleteCourseOrganizer deletes the given course organizer
func (r *CoursePGSQL) DeleteCourseOrganizer(courseID id.ID, cos []*entity.CourseOrganizer) error {
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

// DeleteCourseOrganizerByCourse deletes course organizer using course id
func (r *CoursePGSQL) DeleteCourseOrganizerByCourse(courseID id.ID) error {
	res, err := r.db.Exec(`DELETE FROM course_organizer WHERE course_id = $1;`, courseID)
	if err != nil {
		return err
	}

	if cnt, _ := res.RowsAffected(); cnt == 0 {
		return sql.ErrNoRows
	}

	return nil
}

// MultiGetCourseOrganizer gets course organizer for the given course ids
func (r *CoursePGSQL) MultiGetCourseOrganizer(courseIDList []id.ID,
) ([][]*entity.CourseOrganizer, error) {
	values := func(index int) []interface{} {
		return []interface{}{
			courseIDList[index],
		}
	}

	query := `SELECT course_id, organizer_id FROM course_organizer`
	whereIn, valueArgs := util.BuildQueryWhereClauseIn(
		"course_id",
		len(courseIDList),
		values,
	)

	stmt, err := r.db.Prepare(query + whereIn)
	if err != nil {
		l.Log.Warnf("err=%v", err)
		return nil, err
	}

	l.Log.Debugf("courseIDList=%v, query=%v, values=%v", courseIDList, query+whereIn, valueArgs)
	rows, err := stmt.Query(valueArgs...)
	if err != nil {
		l.Log.Warnf("err=%v", err)
		return nil, err
	}

	m := make(map[id.ID]int)
	for i, cID := range courseIDList {
		m[cID] = i
	}

	cosList := make([][]*entity.CourseOrganizer, len(courseIDList))
	for rows.Next() {
		var co entity.CourseOrganizer
		var courseID id.ID

		err := rows.Scan(&courseID, &co.ID)
		if err != nil {
			l.Log.Warnf("err=%v", err)
			return nil, err
		}

		cosList[m[courseID]] = append(cosList[m[courseID]], &co)
	}

	defer rows.Close()
	return cosList, err
}

// --------------------------------------------------------------------------------
// Course Teacher
// --------------------------------------------------------------------------------
// InsertCourseTeacher creates course to teacher mapping
func (r *CoursePGSQL) InsertCourseTeacher(courseID id.ID, cts []*entity.CourseTeacher) error {
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

// GetCourseTeacher gets course teacher for the given course id
func (r *CoursePGSQL) GetCourseTeacher(courseID id.ID) ([]*entity.CourseTeacher, error) {
	query := `SELECT teacher_id, is_primary FROM course_teacher WHERE course_id = $1;`

	stmt, err := r.db.Prepare(query)
	if err != nil {
		return nil, err
	}

	rows, err := stmt.Query(courseID)
	if err != nil {
		return nil, err
	}

	var cts []*entity.CourseTeacher
	for rows.Next() {
		var ct entity.CourseTeacher

		err := rows.Scan(&ct.ID, &ct.IsPrimary)
		if err != nil {
			return nil, err
		}

		cts = append(cts, &ct)
	}

	defer rows.Close()
	return cts, err
}

// UpdateCourseTeacher updates course teacher for the given course id and the teacher
func (r *CoursePGSQL) UpdateCourseTeacher(courseID id.ID, cts []*entity.CourseTeacher) error {
	currentCTs, err := r.GetCourseTeacher(courseID)
	if err != nil {
		return err
	}

	mapTeacherID := make(map[id.ID]*entity.CourseTeacher)
	for _, ct := range currentCTs {
		mapTeacherID[ct.ID] = ct
	}

	var addCTs []*entity.CourseTeacher
	var rmCTs []*entity.CourseTeacher
	for _, ct := range cts {
		if _, exists := mapTeacherID[ct.ID]; exists {
			// TODO: There are 2 cases: is_primary is changed and unchanged.
			// If is_primary is changed, then we need to update that entry instead
			// of deleting it.
			delete(mapTeacherID, ct.ID)
		} else {
			addCTs = append(addCTs, ct)
		}
	}

	for id := range mapTeacherID {
		ct := entity.CourseTeacher{
			ID: id,
		}
		rmCTs = append(rmCTs, &ct)
	}

	if len(addCTs) > 0 {
		err = r.InsertCourseTeacher(courseID, addCTs)
		if err != nil {
			return err
		}
	}

	if len(rmCTs) > 0 {
		err = r.DeleteCourseTeacher(courseID, rmCTs)
		if err != nil {
			return err
		}
	}

	return err
}

// DeleteCourseTeacher deletes the given course teacher
func (r *CoursePGSQL) DeleteCourseTeacher(courseID id.ID, cts []*entity.CourseTeacher) error {
	values := func(index int) []interface{} {
		return []interface{}{
			courseID,
			cts[index].ID,
		}
	}

	query, valueArgs := util.GenBulkDeletePGSQL(
		"course_teacher",
		[]string{"course_id", "teacher_id"},
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

// DeleteCourseTeacherByCourse deletes course teachers using course id
func (r *CoursePGSQL) DeleteCourseTeacherByCourse(courseID id.ID) error {
	res, err := r.db.Exec(`DELETE FROM course_teacher WHERE course_id = $1;`, courseID)
	if err != nil {
		return err
	}

	if cnt, _ := res.RowsAffected(); cnt == 0 {
		return sql.ErrNoRows
	}
	return nil
}

// MultiGetCourseTeacher gets course teacher for the given course ids
func (r *CoursePGSQL) MultiGetCourseTeacher(courseIDList []id.ID,
) ([][]*entity.CourseTeacher, error) {
	values := func(index int) []interface{} {
		return []interface{}{
			courseIDList[index],
		}
	}

	query := `SELECT course_id, teacher_id, is_primary FROM course_teacher`
	whereIn, valueArgs := util.BuildQueryWhereClauseIn(
		"course_id",
		len(courseIDList),
		values,
	)

	stmt, err := r.db.Prepare(query + whereIn)
	if err != nil {
		l.Log.Warnf("err=%v", err)
		return nil, err
	}

	l.Log.Debugf("courseIDList=%v, query=%v, values=%v", courseIDList, query+whereIn, valueArgs)
	rows, err := stmt.Query(valueArgs...)
	if err != nil {
		l.Log.Warnf("err=%v", err)
		return nil, err
	}

	m := make(map[id.ID]int)
	for i, cID := range courseIDList {
		m[cID] = i
	}

	ctsList := make([][]*entity.CourseTeacher, len(courseIDList))
	for rows.Next() {
		var ct entity.CourseTeacher
		var courseID id.ID

		err := rows.Scan(&courseID, &ct.ID, &ct.IsPrimary)
		if err != nil {
			l.Log.Warnf("err=%v", err)
			return nil, err
		}

		ctsList[m[courseID]] = append(ctsList[m[courseID]], &ct)
	}

	defer rows.Close()
	return ctsList, err
}

// --------------------------------------------------------------------------------
// Course Contact
// --------------------------------------------------------------------------------
// InsertCourseContact creates course to contact mapping
func (r *CoursePGSQL) InsertCourseContact(courseID id.ID, ccs []*entity.CourseContact) error {
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

// GetCourseContact gets course contact for the given course id
func (r *CoursePGSQL) GetCourseContact(courseID id.ID) ([]*entity.CourseContact, error) {
	query := `SELECT contact_id FROM course_contact WHERE course_id = $1;`

	stmt, err := r.db.Prepare(query)
	if err != nil {
		return nil, err
	}

	rows, err := stmt.Query(courseID)
	if err != nil {
		return nil, err
	}

	var ccs []*entity.CourseContact
	for rows.Next() {
		var cc entity.CourseContact

		err := rows.Scan(&cc.ID)
		if err != nil {
			return nil, err
		}

		ccs = append(ccs, &cc)
	}

	defer rows.Close()
	return ccs, err
}

// UpdateCourseContact updates course contact for the given course id and the contact
func (r *CoursePGSQL) UpdateCourseContact(courseID id.ID, ccs []*entity.CourseContact) error {
	currentCCs, err := r.GetCourseContact(courseID)
	if err != nil {
		return err
	}

	mapContactID := make(map[id.ID]bool)
	for _, cc := range currentCCs {
		mapContactID[cc.ID] = true
	}

	var addCCs []*entity.CourseContact
	var rmCCs []*entity.CourseContact

	for _, cc := range ccs {
		if _, exists := mapContactID[cc.ID]; exists {
			delete(mapContactID, cc.ID)
		} else {
			addCCs = append(addCCs, cc)
		}
	}

	for id := range mapContactID {
		cc := entity.CourseContact{
			ID: id,
		}
		rmCCs = append(rmCCs, &cc)
	}

	if len(addCCs) > 0 {
		err = r.InsertCourseContact(courseID, addCCs)
		if err != nil {
			return err
		}
	}

	if len(rmCCs) > 0 {
		err = r.DeleteCourseContact(courseID, rmCCs)
		if err != nil {
			return err
		}
	}

	return err
}

// DeleteCourseContact deletes the given course contacts
func (r *CoursePGSQL) DeleteCourseContact(courseID id.ID, ccs []*entity.CourseContact) error {
	values := func(index int) []interface{} {
		return []interface{}{
			courseID,
			ccs[index].ID,
		}
	}

	query, valueArgs := util.GenBulkDeletePGSQL(
		"course_contact",
		[]string{"course_id", "contact_id"},
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

// DeleteCourseContactByCourse deletes course contacts using course id
func (r *CoursePGSQL) DeleteCourseContactByCourse(courseID id.ID) error {
	res, err := r.db.Exec(`DELETE FROM course_contact WHERE course_id = $1;`, courseID)
	if err != nil {
		return err
	}

	if cnt, _ := res.RowsAffected(); cnt == 0 {
		return sql.ErrNoRows
	}
	return nil
}

// MultiGetCourseContact gets course contact for the given course ids
func (r *CoursePGSQL) MultiGetCourseContact(courseIDList []id.ID,
) ([][]*entity.CourseContact, error) {
	values := func(index int) []interface{} {
		return []interface{}{
			courseIDList[index],
		}
	}

	query := `SELECT course_id, contact_id FROM course_contact`
	whereIn, valueArgs := util.BuildQueryWhereClauseIn(
		"course_id",
		len(courseIDList),
		values,
	)

	stmt, err := r.db.Prepare(query + whereIn)
	if err != nil {
		l.Log.Warnf("err=%v", err)
		return nil, err
	}

	l.Log.Debugf("courseIDList=%v, query=%v, values=%v", courseIDList, query+whereIn, valueArgs)
	rows, err := stmt.Query(valueArgs...)
	if err != nil {
		l.Log.Warnf("err=%v", err)
		return nil, err
	}

	m := make(map[id.ID]int)
	for i, cID := range courseIDList {
		m[cID] = i
	}

	ccsList := make([][]*entity.CourseContact, len(courseIDList))
	for rows.Next() {
		var cc entity.CourseContact
		var courseID id.ID

		err := rows.Scan(&courseID, &cc.ID)
		if err != nil {
			l.Log.Warnf("err=%v", err)
			return nil, err
		}

		ccsList[m[courseID]] = append(ccsList[m[courseID]], &cc)
	}

	defer rows.Close()
	return ccsList, err
}

// --------------------------------------------------------------------------------
// Course Notify
// --------------------------------------------------------------------------------
// InsertCourseNotify creates course to notify mapping
func (r *CoursePGSQL) InsertCourseNotify(courseID id.ID, cns []*entity.CourseNotify) error {
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

// GetCourseNotify gets course notify for the given course id
func (r *CoursePGSQL) GetCourseNotify(courseID id.ID) ([]*entity.CourseNotify, error) {
	query := `SELECT notify_id FROM course_notify WHERE course_id = $1;`

	stmt, err := r.db.Prepare(query)
	if err != nil {
		return nil, err
	}

	rows, err := stmt.Query(courseID)
	if err != nil {
		return nil, err
	}

	var cns []*entity.CourseNotify
	for rows.Next() {
		var cn entity.CourseNotify

		err := rows.Scan(&cn.ID)
		if err != nil {
			return nil, err
		}

		cns = append(cns, &cn)
	}

	defer rows.Close()
	return cns, err
}

// UpdateCourseNotify updates course notify for the given course id and the notify
func (r *CoursePGSQL) UpdateCourseNotify(courseID id.ID, cns []*entity.CourseNotify) error {
	currentCNs, err := r.GetCourseNotify(courseID)
	if err != nil {
		return err
	}

	mapNotifyID := make(map[id.ID]bool)
	for _, cn := range currentCNs {
		mapNotifyID[cn.ID] = true
	}

	var addCNs []*entity.CourseNotify
	var rmCNs []*entity.CourseNotify

	for _, cn := range cns {
		if _, exists := mapNotifyID[cn.ID]; exists {
			delete(mapNotifyID, cn.ID)
		} else {
			addCNs = append(addCNs, cn)
		}
	}

	for id := range mapNotifyID {
		cn := entity.CourseNotify{
			ID: id,
		}
		rmCNs = append(rmCNs, &cn)
	}

	if len(addCNs) > 0 {
		err = r.InsertCourseNotify(courseID, addCNs)
		if err != nil {
			return err
		}
	}

	if len(rmCNs) > 0 {
		err = r.DeleteCourseNotify(courseID, rmCNs)
		if err != nil {
			return err
		}
	}

	return err
}

// DeleteCourseNotify deletes the given course notify
func (r *CoursePGSQL) DeleteCourseNotify(courseID id.ID, cns []*entity.CourseNotify) error {
	values := func(index int) []interface{} {
		return []interface{}{
			courseID,
			cns[index].ID,
		}
	}

	query, valueArgs := util.GenBulkDeletePGSQL(
		"course_notify",
		[]string{"course_id", "notify_id"},
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

// DeleteCourseNotifyByCourse deletes course notify using course id
func (r *CoursePGSQL) DeleteCourseNotifyByCourse(courseID id.ID) error {
	res, err := r.db.Exec(`DELETE FROM course_notify WHERE course_id = $1;`, courseID)
	if err != nil {
		return err
	}

	if cnt, _ := res.RowsAffected(); cnt == 0 {
		return sql.ErrNoRows
	}
	return nil
}

// MultiGetCourseNotify gets course notify for the given course ids
func (r *CoursePGSQL) MultiGetCourseNotify(courseIDList []id.ID,
) ([][]*entity.CourseNotify, error) {
	values := func(index int) []interface{} {
		return []interface{}{
			courseIDList[index],
		}
	}

	query := `SELECT course_id, notify_id FROM course_notify`
	whereIn, valueArgs := util.BuildQueryWhereClauseIn(
		"course_id",
		len(courseIDList),
		values,
	)

	stmt, err := r.db.Prepare(query + whereIn)
	if err != nil {
		l.Log.Warnf("err=%v", err)
		return nil, err
	}

	rows, err := stmt.Query(valueArgs...)
	if err != nil {
		l.Log.Warnf("err=%v", err)
		return nil, err
	}

	m := make(map[id.ID]int)
	for i, cID := range courseIDList {
		m[cID] = i
	}

	cnsList := make([][]*entity.CourseNotify, len(courseIDList))
	for rows.Next() {
		var cn entity.CourseNotify
		var courseID id.ID

		err := rows.Scan(&courseID, &cn.ID)
		if err != nil {
			l.Log.Warnf("err=%v", err)
			return nil, err
		}

		cnsList[m[courseID]] = append(cnsList[m[courseID]], &cn)
	}

	defer rows.Close()
	return cnsList, err
}
