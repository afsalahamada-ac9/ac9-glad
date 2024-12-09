// package rds_export

// import (
// 	"log"
// 	"sudhagar/glad/entity"
// 	infra "sudhagar/glad/ops/db"
// 	"sudhagar/glad/repository"
// )

// func Export(course_id entity.ID) {
// 	db, err := infra.GetDB()
// 	if err != nil {
// 		log.Println("there was an error fetching the database")
// 	}
// 	sqldb, err := db.DB()
// 	if err != nil {
// 		log.Println("error fetching the sql db")
// 	}
// 	course_db := repository.NewCoursePGSQL(sqldb)
// 	log.Println(course_db)
// 	course, err := course_db.Get(course_id)
// 	if err != nil {
// 		log.Println("there was an error fetching the course based on this id", err)
// 	}
// 	log.Println(course)
// 	// todo: logic to send the data
// }

package rds_export

import (
	pkgentity "ac9/glad/pkg/id"
	service "ac9/glad/usecase/sf_export"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// func Export(courseID entity.ID) {
//     sfService, err := service.NewSFExportService()
//     if err != nil {
//         log.Printf("Failed to initialize SF export service: %v", err)
//         return
//     }

//     err = sfService.ExportToSF(courseID)
//     if err != nil {
//         log.Printf("Failed to export to SF: %v", err)
//         return
//     }

//     log.Println("Successfully exported course to SF")
// }

func Export(courseID pkgentity.ID) {
	sfService, err := service.NewSFExportService()
	if err != nil {
		log.Printf("Failed to initialize SF export service: %v", err)
		return
	}

	err = sfService.ExportToSF(courseID)
	if err != nil {
		log.Printf("Failed to export to SF: %v", err)
		return
	}

	log.Println("Successfully exported course to SF")
}

func ExportHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	log.Println(params)
	course_id := params["id"]
	id, err := strconv.Atoi(course_id)
	if err != nil {
		log.Println("unable to convert the parameter")
	}
	Export(pkgentity.ID(id))
}
