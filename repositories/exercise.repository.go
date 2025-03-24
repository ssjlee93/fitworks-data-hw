package repositories

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/ssjlee93/fitworks-data-hw/models"
	"log"
)

const (
	readOneExercise  = "SELECT * FROM exercises WHERE exercise_id=$1"
	readAllExercises = "SELECT * FROM exercises ORDER BY exercise_id DESC"
	createExercise   = "INSERT INTO exercises (exercise_name, description) VALUES ($1, $2)"
	updateExercise   = "UPDATE exercises SET exercise_name=$2, description=$3, updated=current_timestamp WHERE exercise_id=$1"
	deleteExercise   = "DELETE FROM exercises WHERE exercise_id=$1"
)

type ExerciseRepository struct {
	d *sql.DB
}

func NewExerciseRepository(db *sql.DB) *ExerciseRepository {
	return &ExerciseRepository{d: db}
}

func (exerciseRepository *ExerciseRepository) ReadAll() ([]models.Exercise, error) {
	log.Println("| - - ExerciseRepository.ReadAll")
	result := make([]models.Exercise, 0)
	// query
	rows, err := exerciseRepository.d.Query(readAllExercises)
	if err != nil {
		return nil, fmt.Errorf("error querying : %v", err)
	}
	defer rows.Close()
	// Loop through rows, using Scan to assign column data to struct fields.
	for rows.Next() {
		var exercise models.Exercise
		if err := rows.Scan(
			&exercise.ExerciseID,
			&exercise.ExerciseName,
			&exercise.Description,
			&exercise.Created,
			&exercise.Updated,
		); err != nil {
			return nil, fmt.Errorf("error scanning : %v", err)
		}
		result = append(result, exercise)
	}

	// handle any other errors
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error reading all : %v", err)
	}
	return result, nil
}

func (exerciseRepository *ExerciseRepository) ReadOne(id int64) (*models.Exercise, error) {
	log.Println("| - - ExerciseRepository.ReadOne")
	// query
	row := exerciseRepository.d.QueryRow(readOneExercise, id)
	var exercise models.Exercise
	if err := row.Scan(
		&exercise.ExerciseID,
		&exercise.ExerciseName,
		&exercise.Description,
		&exercise.Created,
		&exercise.Updated,
	); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("error scanning : 0 rows found %v", err)
		}
		return nil, fmt.Errorf("error scanning : %v", err)
	}
	return &exercise, nil
}

func (exerciseRepository *ExerciseRepository) Create(exercise models.Exercise) error {
	log.Println("| - - ExerciseRepository.Create")

	exec, err := exerciseRepository.d.Exec(createExercise,
		exercise.ExerciseName,
		exercise.Description,
	)

	if err != nil {
		log.Printf("error exec create : %v", err)
		return err
	}
	log.Println(exec.RowsAffected())
	return nil
}

func (exerciseRepository *ExerciseRepository) Update(exercise models.Exercise) error {
	log.Println("| - - ExerciseRepository.Update")
	exec, err := exerciseRepository.d.Exec(updateExercise,
		exercise.ExerciseID,
		exercise.ExerciseName,
		exercise.Description,
	)

	if err != nil {
		log.Printf("error exec update : %v", err)
		return err
	}
	log.Println(exec.RowsAffected())
	return nil
}

func (exerciseRepository *ExerciseRepository) Delete(id int64) error {
	log.Println("| - - - ExerciseRepository Delete", id)
	exec, err := exerciseRepository.d.Exec(deleteExercise, id)
	if err != nil {
		log.Printf("ExerciseRepository.Delete error : %v", err)
		return err
	}
	log.Println(exec.RowsAffected())
	return nil
}
