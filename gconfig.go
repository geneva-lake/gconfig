package gconfig

import (
	"bytes"
	"encoding/json"
	"io"
	"os"
	"strings"

	"gopkg.in/yaml.v2"
)

// Generic type containing certain configuration
type Config[T any] struct {
	Config *T
}

// Configuration constructor
func NewConfig[T any]() *Config[T] {
	return &Config[T]{new(T)}
}

// Private settings reader for chain organizing
type config[T any] struct {
	object *Config[T]
	reader io.Reader
	Error  error
}

// io.Closer realization
func (this *config[T]) Close() error {
	if this != nil && this.reader != nil {
		if closer, ok := this.reader.(io.Closer); ok {
			return closer.Close()
		}
	}
	return nil
}

// Settings reader constructor
func (this *Config[T]) from(reader io.Reader, err error) *config[T] {
	return &config[T]{
		object: this,
		reader: reader,
		Error:  err,
	}
}

// From* - a group of functions that start the chain of creating configuration through deserialization

// Set file
func (this *Config[T]) FromFile(name string) *config[T] {
	return this.from(os.Open(name))
}

// Set reader
func (this *Config[T]) FromReader(reader io.Reader) *config[T] {
	return this.from(reader, nil)
}

// Set bytes
func (this *Config[T]) FromBytes(b []byte) *config[T] {
	return this.from(bytes.NewReader(b), nil)
}

// Set string
func (this *Config[T]) FromString(s string) *config[T] { return this.from(strings.NewReader(s), nil) }

// Parse yaml file
func (this *config[T]) Yaml() (*T, error) {
	if this.Error != nil {
		return nil, this.Error
	}
	if this.reader == nil {
		return nil, nil
	}
	return this.object.Config, yaml.NewDecoder(this.reader).Decode(this.object.Config)
}

// Parse json file
func (this config[T]) JSON() (*T, error) {
	defer this.Close()
	if this.Error != nil {
		return nil, this.Error
	}
	if this.reader == nil {
		return nil, nil
	}
	return this.object.Config, json.NewDecoder(this.reader).Decode(this.object.Config)
}
