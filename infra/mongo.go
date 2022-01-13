package infra

func NewMongoService(connString string) *mongoService {
	return &mongoService{
		connectionString: connString,
	}
}

//mongoService Service responsible for interacting with mongoDB
type mongoService struct {
	connectionString string
}

func (ms mongoService) addEvent() {
	//TODO implement me
	panic("implement me")
}

func (ms mongoService) updateEvent() {
	//TODO implement me
	panic("implement me")
}

func (ms mongoService) removeEvent() {
	//TODO implement me
	panic("implement me")
}

func (ms mongoService) healthCheck() {
	//TODO implement me
	panic("implement me")
}
