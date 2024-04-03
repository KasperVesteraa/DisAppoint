package api

type Appointment struct {
	Id           string
	Title        string
	Location     string
	Description  string
	StartTime    int
	EndTime      int
	Author       User
	Participants []User
}

func CreateAppointment(id, title, location, description string, startime, endtime int, author User, participant []User) (*Appointment, error) {

	return &Appointment{}, nil
}
