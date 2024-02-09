package course

import "content/internal/model"

func (c CourseRepo) CreateCourse(course model.Course) error {
	query := "INSERT INTO course (name, description, link, cost) VALUES ($1, $2, $3, $4)"

	if _, err := c.DB.Exec(query, course.Name, course.Description, course.Link, course.Cost); err != nil {
		return err
	}

	return nil
}

func (c CourseRepo) GetCourseByName(courseName string) (model.Course, error) {
	query := "SELECT id, name, description, link, cost FROM course WHERE name = $1"

	var course model.Course

	if err := c.DB.QueryRow(query, courseName).Scan(&course.Id, &course.Name, &course.Description, &course.Link, &course.Cost); err != nil {
		return model.Course{}, err
	}

	return course, nil
}

func (c CourseRepo) GetCourseById(id int) (model.Course, error) {
	query := "SELECT id, name, description, link, cost FROM course WHERE id = $1"

	var course model.Course

	if err := c.DB.QueryRow(query, id).Scan(&course.Id, &course.Name, &course.Description, &course.Link, &course.Cost); err != nil {
		return model.Course{}, err
	}

	return course, nil
}

func (c CourseRepo) GetAllCourse() ([]model.Course, error) {
	res := []model.Course{}

	query := "SELECT id, name, description, link, cost FROM course"

	rows, err := c.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var s model.Course
		err := rows.Scan(&s.Id, &s.Name, &s.Description, &s.Link, &s.Cost)
		if err != nil {
			return nil, err
		}
		res = append(res, s)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return res, nil
}

// DeleteCourseByName TODO: think of reasons for such method
func (c CourseRepo) DeleteCourseByName(courseName string) error {
	return nil
}
