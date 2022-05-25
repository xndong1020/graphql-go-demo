package dtos

type BookInput struct {
    Title     string `json:"title"`
    Author    string `json:"author"`
    Publisher string `json:"publisher"`
}