package repository

import (
	sf_entity "ac9/glad/entity/sf_entity"
	pkgentity "ac9/glad/pkg/id"
	"database/sql"
)

type TimingPGSQL struct {
	db *sql.DB
}

func NewTimingPGSQL(db *sql.DB) *TimingPGSQL {
	return &TimingPGSQL{
		db: db,
	}
}

func (r *TimingPGSQL) GetByCourseID(courseID pkgentity.ID) ([]*sf_entity.Timing_value, error) {
	query := `
        SELECT id, ext_id, course_date, start_time, end_time
        FROM course_timing
        WHERE course_id = $1
    `

	rows, err := r.db.Query(query, courseID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var timings []*sf_entity.Timing_value
	for rows.Next() {
		var t sf_entity.Timing_value
		var extID sql.NullString
		err := rows.Scan(&t.Ext_id, &extID, &t.Course_date, &t.Start_time, &t.End_time)
		if err != nil {
			return nil, err
		}
		t.Ext_id = extID.String
		timings = append(timings, &t)
	}

	return timings, nil
}
