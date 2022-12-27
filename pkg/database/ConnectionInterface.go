package database

type DB interface {
	Init() error
	Close() error
}
