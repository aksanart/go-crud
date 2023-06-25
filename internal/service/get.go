package service

func (s *serviceStruct) GetService() (interface{}, error) {
	return s.Repo.Get()
}
