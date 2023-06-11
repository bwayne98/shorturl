package shorturl

//go:generate mockgen -destination=../../mockdb/shorturl.go -package=mockdb shorturl/m/db/store/shorturl ShorturlQuery

type ShorturlQuery interface {
	Querier
}
