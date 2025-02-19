package session

import (
	"github.com/gomodule/redigo/redis"
	"github.com/gorilla/sessions"
	"gitlab.com/ditp.thaitrade/enginex/redisstore"
)

type RedisStore interface {
	Store
}

// size: maximum number of idle connections.
// network: tcp or udp
// address: host:port
// password: redis-password
// Keys are defined in pairs to allow key rotation, but the common case is to set a single
// authentication key and optionally an encryption key.
//
// The first key in a pair is used for authentication and the second for encryption. The
// encryption key can be set to nil or omitted in the last pair, but the authentication key
// is required in all pairs.
//
// It is recommended to use an authentication key with 32 or 64 bytes. The encryption key,
// if set, must be either 16, 24, or 32 bytes to select AES-128, AES-192, or AES-256 modes.
func NewRedisStore(size int, maxAge int, maxLength int, network, address, password string, keyPairs ...[]byte) (RedisStore, error) {
	store, err := redisstore.NewRedisStore(size, network, address, password, keyPairs...)
	if err != nil {
		return nil, err
	}
	redisStore := &redisStore{store}
	redisStore.SetMaxAge(maxAge)
	redisStore.SetMaxLength(maxLength)
	return redisStore, nil
}

func NewRedisStoreWithSecret(size int, maxAge int, maxLength int, network, address, password string, httpOnly bool, secure bool, keyPairs ...[]byte) (RedisStore, error) {
	store, err := redisstore.NewRedisStore(size, network, address, password, keyPairs...)
	if err != nil {
		return nil, err
	}
	redisStore := &redisStore{store}
	redisStore.SetMaxAge(maxAge)
	redisStore.SetMaxLength(maxLength)
	redisStore.SetHttpOnly(httpOnly)
	redisStore.SetSecure(secure)
	return redisStore, nil
}

// pool: redis pool connections
// Keys are defined in pairs to allow key rotation, but the common case is to set a single
// authentication key and optionally an encryption key.
//
// The first key in a pair is used for authentication and the second for encryption. The
// encryption key can be set to nil or omitted in the last pair, but the authentication key
// is required in all pairs.
//
// It is recommended to use an authentication key with 32 or 64 bytes. The encryption key,
// if set, must be either 16, 24, or 32 bytes to select AES-128, AES-192, or AES-256 modes.
func NewRedisStoreWithPool(pool *redis.Pool, keyPairs ...[]byte) (RedisStore, error) {
	store, err := redisstore.NewRedisStoreWithPool(pool, keyPairs...)
	if err != nil {
		return nil, err
	}
	return &redisStore{store}, nil
}

// size: maximum number of idle connections.
// network: tcp or udp
// address: host:port
// password: redis-password
// Keys are defined in pairs to allow key rotation, but the common case is to set a single
// authentication key and optionally an encryption key.
// DB: database index
//
// The first key in a pair is used for authentication and the second for encryption. The
// encryption key can be set to nil or omitted in the last pair, but the authentication key
// is required in all pairs.
//
// It is recommended to use an authentication key with 32 or 64 bytes. The encryption key,
// if set, must be either 16, 24, or 32 bytes to select AES-128, AES-192, or AES-256 modes.
func NewRedisStoreWithDB(size int, network, address, password string, DB string, keyPairs ...[]byte) (RedisStore, error) {
	store, err := redisstore.NewRedisStoreWithDB(size, network, address, password, DB, keyPairs...)
	if err != nil {
		return nil, err
	}
	return &redisStore{store}, nil
}

type redisStore struct {
	*redisstore.RedisStore
}

func (c *redisStore) Options(options Options) {
	c.RedisStore.Options = &sessions.Options{
		Path:     options.Path,
		Domain:   options.Domain,
		MaxAge:   options.MaxAge,
		Secure:   options.Secure,
		HttpOnly: options.HttpOnly,
	}
}

// MaxAge restricts the maximum age, in seconds, of the session record
// both in database and a browser. This is to change session storage configuration.
// If you want just to remove session use your session `s` object and change it's
// `Options.MaxAge` to -1, as specified in
//    http://godoc.org/github.com/gorilla/sessions#Options
//
// Default is the one provided by github.com/boj/redistore package value - `sessionExpire`.
// Set it to 0 for no restriction.
// Because we use `MaxAge` also in SecureCookie crypting algorithm you should
// use this function to change `MaxAge` value.
func (c *redisStore) MaxAge(age int) {
	c.RedisStore.SetMaxAge(age)
}

// MaxLength sets RedisStore.maxLength if the `l` argument is greater or equal 0
// maxLength restricts the maximum length of new sessions to l.
// If l is 0 there is no limit to the size of a session, use with caution.
// The default for a new RedisStore is 4096. Redis allows for max.
// value sizes of up to 512MB (http://redis.io/topics/data-types)
// Default: 4096,
func (c *redisStore) MaxLength(length int) {
	c.RedisStore.SetMaxLength(length)
}
