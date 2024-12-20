/*
 * Copyright 2024 AboveCloud9.AI Products and Services Private Limited
 * All rights reserved.
 * This code may not be used, copied, modified, or distributed without explicit permission.
 */

package repository

import (
	"database/sql"
	"log"
	"time"

	"ac9/glad/entity"
	"ac9/glad/pkg/id"
	l "ac9/glad/pkg/logger"
	"ac9/glad/pkg/util"
)

// CourseTimingPGSQL pgsql repo
type CourseTimingPGSQL struct {
	db *sql.DB
}

// NewCourseTimingPGSQL create new repository
func NewCourseTimingPGSQL(db *sql.DB) *CourseTimingPGSQL {
	return &CourseTimingPGSQL{
		db: db,
	}
}

// Create creates a course timing
func (r *CourseTimingPGSQL) Create(ct *entity.CourseTiming) (id.ID, error) {
	stmt, err := r.db.Prepare(`
		INSERT INTO course_timing
			(
				id, course_id, ext_id, course_date, start_time, end_time, created_at
			)
		VALUES($1, $2, $3, $4, $5, $6, $7)`)
	if err != nil {
		return ct.ID, err
	}
	_, err = stmt.Exec(
		ct.ID,
		ct.CourseID,
		ct.ExtID,
		ct.DateTime.Date,
		ct.DateTime.StartTime,
		ct.DateTime.EndTime,
		time.Now().Format("2006-01-02"),
	)
	if err != nil {
		return ct.ID, err
	}
	err = stmt.Close()
	if err != nil {
		return ct.ID, err
	}
	return ct.ID, nil
}

// Get retrieves a course
func (r *CourseTimingPGSQL) Get(id id.ID) (*entity.CourseTiming, error) {
	stmt, err := r.db.Prepare(`
		SELECT 
			course_id, ext_id, course_date, start_time, end_time, created_at
		FROM course_timing
		WHERE id = $1;`)
	if err != nil {
		return nil, err
	}
	var ct entity.CourseTiming
	var ext_id sql.NullString
	var course_date, start_time, end_time sql.NullString
	err = stmt.QueryRow(id).Scan(&ct.CourseID, &ext_id, &course_date, &start_time, &end_time,
		&ct.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	log.Printf("Course timing: id=%#v course_date=%v start_time=%v end_time=%v",
		id, course_date, start_time, end_time)
	ct.ID = id
	ct.ExtID = &ext_id.String
	ct.DateTime = entity.CourseDateTime{
		Date:      course_date.String,
		StartTime: start_time.String,
		EndTime:   end_time.String,
	}

	log.Printf("Course timing: %#v", ct)

	return &ct, nil
}

// Update updates a course timing
func (r *CourseTimingPGSQL) Update(ct *entity.CourseTiming) error {
	ct.UpdatedAt = time.Now()

	_, err := r.db.Exec(`
		UPDATE course_timing SET
			course_id = $1,
			ext_id = $2,
			course_date = $3,
			start_time = $4,
			end_time = $5,
			updated_at = $6
		WHERE id = $7;
		`,
		ct.CourseID,
		ct.ExtID,
		ct.DateTime.Date,
		ct.DateTime.StartTime,
		ct.DateTime.EndTime,
		ct.UpdatedAt.Format("2006-01-02"),
		ct.ID,
	)
	if err != nil {
		return err
	}
	return nil
}

// GetByCourse retrieves course timing by course id
func (r *CourseTimingPGSQL) GetByCourse(courseID id.ID) ([]*entity.CourseTiming, error) {
	query := `
		SELECT 
			id, course_id, ext_id, course_date, start_time, end_time, created_at
		FROM course_timing
		WHERE course_id = $1;
	`

	stmt, err := r.db.Prepare(query)
	if err != nil {
		return nil, err
	}

	rows, err := stmt.Query(courseID)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	return r.scanRows(rows)
}

// Delete deletes a course timing
func (r *CourseTimingPGSQL) Delete(id id.ID) error {
	res, err := r.db.Exec(`DELETE FROM course_timing WHERE id = $1;`, id)
	if err != nil {
		return err
	}

	if cnt, _ := res.RowsAffected(); cnt == 0 {
		return sql.ErrNoRows
	}

	return nil
}

// GetCount gets total items in course timing table
func (r *CourseTimingPGSQL) GetCount() (int, error) {
	stmt, err := r.db.Prepare(`SELECT count(*) FROM course_timing;`)
	if err != nil {
		return 0, err
	}

	var count int
	err = stmt.QueryRow().Scan(&count)
	if err != nil {
		return 0, err
	}

	return count, nil
}

func (r *CourseTimingPGSQL) scanRows(rows *sql.Rows) ([]*entity.CourseTiming, error) {
	var cts []*entity.CourseTiming
	for rows.Next() {
		ct, err := r.scanRow(rows)
		if err != nil {
			return nil, err
		}
		cts = append(cts, ct)
	}
	return cts, nil
}

func (r *CourseTimingPGSQL) scanRow(rows *sql.Rows) (*entity.CourseTiming, error) {
	var ct entity.CourseTiming
	// id, course_id, ext_id, course_date, start_time, end_time, created_at
	var ext_id, course_date, start_time, end_time sql.NullString
	err := rows.Scan(
		&ct.ID,
		&ct.CourseID,
		&ext_id,
		&course_date,
		&start_time,
		&end_time,
		&ct.CreatedAt,
	)
	if err != nil {
		l.Log.Warnf("err=%v", err)
		return nil, err
	}

	ct.ExtID = &ext_id.String
	ct.DateTime.Date = course_date.String
	ct.DateTime.StartTime = start_time.String
	ct.DateTime.EndTime = end_time.String

	return &ct, err
}

// MultiGetCourseTiming gets course timing for the given course ids
func (r *CourseTimingPGSQL) MultiGetCourseTiming(courseIDList []id.ID,
) ([][]*entity.CourseTiming, error) {
	values := func(index int) []interface{} {
		return []interface{}{
			courseIDList[index],
		}
	}

	query := `
		SELECT 
			id, course_id, ext_id, course_date, start_time, end_time, created_at
		FROM course_timing
	`
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

	ctList := make([][]*entity.CourseTiming, len(courseIDList))
	for rows.Next() {
		ct, err := r.scanRow(rows)
		if err != nil {
			l.Log.Warnf("err=%v", err)
			return nil, err
		}

		ctList[m[ct.CourseID]] = append(ctList[m[ct.CourseID]], ct)
	}

	defer rows.Close()
	return ctList, err
}
