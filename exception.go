package echochamber

type HttpException interface {
	GetStatus() int
}
