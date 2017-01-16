// Package glogr implements github.com/thockin/logr.Logger in terms of
// github.com/golang/glog.
package glogr

import (
	lr "github.com/Sirupsen/logrus"
	"github.com/bketelsen/logr"
	"github.com/surge/glog"
)

// New returns a logr.Logger which is implemented by glog.
func New() (logr.Logger, error) {
	return logrus{
		level:  0,
		prefix: "",
	}, nil
}

type field struct {
	name  string
	value interface{}
}
type logrus struct {
	level  int
	prefix string
	fields []field
}

func prepend(prefix interface{}, args []interface{}) []interface{} {
	return append([]interface{}{prefix}, args...)
}

func (l logrus) Info(args ...interface{}) {
	if l.Enabled() {
		lr.WithFields(l.getFields()).Info(args)
	}
}

func (l logrus) Infof(format string, args ...interface{}) {
	if l.Enabled() {

		lr.WithFields(l.getFields()).Infof(format, args)

	}
}

func (l logrus) Enabled() bool {
	return bool(glog.V(glog.Level(l.level)))
}

func (l logrus) Error(args ...interface{}) {

	lr.WithFields(l.getFields()).Error(args)
}

func (l logrus) Errorf(format string, args ...interface{}) {
	lr.WithFields(l.getFields()).Errorf(format, args)
}

func (l logrus) V(level int) logr.InfoLogger {
	return logrus{
		level:  level,
		prefix: l.prefix,
	}
}

func (l logrus) NewWithPrefix(prefix string) logr.Logger {
	return logrus{
		level:  l.level,
		prefix: prefix,
	}
}

func (l logrus) WithField(name string, value interface{}) logr.Logger {
	return logrus{
		level:  l.level,
		prefix: l.prefix,
		fields: append(l.fields, field{name: name, value: value}),
	}
}

func (l logrus) getFields() lr.Fields {
	fields := make(map[string]interface{}, len(l.fields))
	for _, f := range l.fields {
		fields[f.name] = f.value
	}
	return fields
}

var _ logr.Logger = logrus{}
var _ logr.InfoLogger = logrus{}
