package env

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"sync"

	"github.com/rocketbitz/set"
)

var (
	s = set.NewUnsafe()
	l sync.RWMutex
)

func init() {
	l.Lock()
	for _, pair := range os.Environ() {
		keyValue := strings.Split(pair, "=")
		s.Add(Var{Key: keyValue[0], Value: keyValue[1]})
	}
	l.Unlock()
}

// Var represents an environment variable,
// default or set by the environment itself
type Var struct {
	Key     string `json:"key"`
	Value   string `json:"value"`
	Default bool   `json:"default,omitempty"`
}

// SetDefault registers a string as a default
// value for an environment variable
func SetDefault(key, value string) {
	l.Lock()
	s.Add(Var{Key: key, Value: value, Default: true})
	l.Unlock()
}

// SetDefaults registers a map[string]string as
// default values for environment variables
func SetDefaults(defaults map[string]string) {
	l.Lock()
	for k, v := range defaults {
		s.Add(Var{Key: k, Value: v, Default: true})
	}
	l.Unlock()
}

// Set an environment variable, returns true
// if the value was already set by either
// default or the environment itself
func Set(key, value string) bool {
	l.Lock()
	defer l.Unlock()

	for _, v := range s.Slice() {
		ev := v.(Var)
		if ev.Key == key {
			return s.Replace(ev, Var{Key: key, Value: value})
		}
	}

	s.Add(Var{Key: key, Value: value})

	return false
}

// Get a given environment variable by name,
// returns a default value if the variable
// is not set
func Get(key string) (value string) {
	l.RLock()
	defer l.RUnlock()

	for _, v := range s.Slice() {
		ev := v.(Var)
		if ev.Key == key {
			return ev.Value
		}
	}

	return ""
}

// Count the number of environment variables
// set, default and explicit
func Count() int {
	return s.Len()
}

// DefaultCount is the number of environment
// variables, defaults only
func DefaultCount() (count int) {
	l.RLock()
	defer l.RUnlock()

	for _, v := range s.Slice() {
		if v.(Var).Default {
			count++
		}
	}

	return
}

// ExplicitCount is the number of environment
// variables, explicit only
func ExplicitCount() int {
	return Count() - DefaultCount()
}

// JSON representation of the environment,
// defaults included
func JSON() []byte {
	tmp := make([]Var, s.Len())

	for i, v := range s.Slice() {
		tmp[i] = v.(Var)
	}

	buf, _ := json.MarshalIndent(tmp, "", "\t")

	return buf
}

// String representation (in JSON) of the
// environment, defaults included
func String() string {
	return string(JSON())
}

// Print the environment, defaults included
func Print() {
	fmt.Println(String())
}
