package interfaces

type IPointsController interface {
	EarnPointForLogin(username string) error
	PurgePoint(username string, point int) error
}
