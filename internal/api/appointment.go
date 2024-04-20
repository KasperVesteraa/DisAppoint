package api

import "github.com/google/uuid"

type Appointment struct {
	Id          string   `json:"id"`
	Title       string   `json:"title"`
	Location    string   `json:"location"`
	Description string   `json:"description"`
	StartTime   int      `json:"start_time"`
	EndTime     int      `json:"end_time"`
	AuthorId    string   `json:"author_id"`
	Parts_id    []string `json:"parts_id"`
}

func CreateAppointment(id, title, location, description string, startime, endtime int, author User, participant []User) (*Appointment, error) {

	return &Appointment{}, nil
}

func (a *Appointment) CreateUuid() {
	a.Id = uuid.New().String()
}
