package output

import "github.com/lucasres/adr-gen/internal/engine"

type OutputBase interface {
	Write(path string, hits map[string]engine.ADRHit) error
	BuildOutput(engine.ADRHit) (string, error)
}
