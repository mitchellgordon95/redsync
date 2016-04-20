package redsync

import "time"

// Redsync provides a simple method for creating distributed mutexes using multiple Redis connection pools.
type Redsync interface {
	NewMutex(name string, options ...Option) Mutex
}

// redsyncImp is an implementation of redsync
type redsyncImp struct {
	pools []Pool
}

// New creates and returns a new redsyncImp instance from given Redis connection pools.
func New(pools []Pool) Redsync {
	return &redsyncImp{
		pools: pools,
	}
}

// NewMutex returns a new distributed mutex with given name.
func (r *redsyncImp) NewMutex(name string, options ...Option) Mutex {
	m := &mutex{
		name:   name,
		expiry: 8 * time.Second,
		tries:  32,
		delay:  500 * time.Millisecond,
		factor: 0.01,
		quorum: len(r.pools)/2 + 1,
		pools:  r.pools,
	}
	for _, o := range options {
		o.Apply(m)
	}
	return m
}

// An Option configures a mutex.
type Option interface {
	Apply(*mutex)
}

// OptionFunc is a function that configures a mutex.
type OptionFunc func(*mutex)

// Apply calls f(*mutex)
func (f OptionFunc) Apply(mutex *mutex) {
	f(mutex)
}

// SetExpiry can be used to set the expiry of a mutex to the given value.
func SetExpiry(expiry time.Duration) Option {
	return OptionFunc(func(m *mutex) {
		m.expiry = expiry
	})
}

// SetTries can be used to set the number of times lock acquire is attempted.
func SetTries(tries int) Option {
	return OptionFunc(func(m *mutex) {
		m.tries = tries
	})
}

// SetRetryDelay can be used to set the amount of time to wait between retries.
func SetRetryDelay(delay time.Duration) Option {
	return OptionFunc(func(m *mutex) {
		m.delay = delay
	})
}

// SetDriftFactor can be used to set the clock drift factor.
func SetDriftFactor(factor float64) Option {
	return OptionFunc(func(m *mutex) {
		m.factor = factor
	})
}
