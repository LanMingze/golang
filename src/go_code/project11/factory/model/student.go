package model
type student struct{
	name string
	score float64
}
func NewStudent(name string,score float64)*student{
	return &student{
		name: name,
		score: score,
	}
}
func(s *student)GetScore()float64{
	return s.score
}
func(s *student)GetName()string{
	return s.name
}