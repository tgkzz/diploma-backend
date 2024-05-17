package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Course struct {
	ID          primitive.ObjectID `json:"_id" bson:"_id"`
	Name        string             `json:"name" bson:"name"`
	ModuleCount string             `json:"module_count" bson:"module_count"`
	Modules     []Module           `json:"modules" bson:"modules"`
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
