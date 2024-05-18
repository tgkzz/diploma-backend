package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Course struct {
	ID               primitive.ObjectID `json:"_id" bson:"_id"`
	Name             string             `json:"name" bson:"name"`
	ShortDescription string             `json:"short_description" bson:"short_description"`
	ImageUrl         string             `json:"image_url" bson:"image_url"`
	Cost             float64            `json:"cost" bson:"cost"`
	ModuleCount      string             `json:"module_count" bson:"module_count"`
	Modules          []Module           `json:"modules" bson:"modules"`
}

type Module struct {
	Name    string   `json:"module_name" bson:"module_name"`
	Lessons []Lesson `json:"lessons" bson:"lessons"`
}

type Lesson struct {
	LessonName    string    `json:"lesson_name" bson:"lesson_name"`
	LessonType    string    `json:"lesson_type" bson:"lesson_type"`
	LessonContent []Content `json:"lesson_content" bson:"lesson_content"`
	VideoPath     string    `json:"video_path" bson:"video_path"`
}

type Content struct {
	Paragraph string `json:"paragraph" bson:"paragraph"`
}

func (c Course) ToRespModel() (response GetCourseLimitedResponse) {
	return GetCourseLimitedResponse{
		Id:               c.ID,
		Name:             c.Name,
		ImageUrl:         c.ImageUrl,
		ModuleCount:      c.ModuleCount,
		Cost:             c.Cost,
		ShortDescription: c.ShortDescription,
	}
}
