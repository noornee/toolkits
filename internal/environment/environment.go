package environment

import (
	"os"

	"github.com/joho/godotenv"
)

// Env represents environmental variable instance
type Env struct{}

// New creates a new instance of Env and returns an error if any occurs
func New(filename string) (*Env, error) {
	err := godotenv.Load(filename)
	if err != nil {
		return nil, err
	}

	ev := &Env{}
	return ev, nil
}

// Get retrieves the string value of an environmental variable
func (e *Env) Get(key string) string {
	return os.Getenv(key)
}
