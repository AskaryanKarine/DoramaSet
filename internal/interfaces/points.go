package interfaces

type IPointsController interface {
	earnPointForLogin(username string) error
	purgePoint(username string, point int) error
}
