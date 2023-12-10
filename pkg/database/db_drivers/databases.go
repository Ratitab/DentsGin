package db_drivers

type DB interface {
	Connect() error
}
