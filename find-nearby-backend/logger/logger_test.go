package logger_test

import (
	"errors"
	"find-nearby-backend/logger"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	l := logger.New("debug", "plaintext")
	assert.NotNil(t, l)
}

func TestNoPanic(t *testing.T) {
	assert.NotPanics(t, func() {
		l := logger.New("debug", "plaintext")
		l.ErrorWithTag(errors.New("foo"), logger.Fields{"bar": "baz"})
		l.ErrorWithTag(errors.New("foo"), nil)
		l.ErrorWithTag(nil, logger.Fields{"bar": "baz"})
		l.Errorf("foo %d", 1)
		l.Infof("foo %d", 1)
		l.Debugf("foo %d", 1)
		l.Warnf("foo %d", 1)
	})
}

func TestPanic(t *testing.T) {
	assert.Panics(t, func() {
		l := logger.New("debug", "plaintext")
		l.Panicf("foo %d", 1)
	})
}
