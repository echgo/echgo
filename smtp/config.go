package smtp

// Access is to config the smtp access data
type Access struct {
	Host     string
	Port     int
	Username string
	Password string
}

// Data is to structure the data
type Data struct {
	Email   string
	Subject string
	File    *string
	HTML    string
}
