package hooks

/*
 * This file is only ever generated once on the first generation and then is free to be modified.
 * Any hooks you wish to add should be registered in the initHooks function. Feel free to define
 * your hooks in this file or in separate files in the hooks package.
 *
 * Hooks are registered per SDK instance, and are valid for the lifetime of the SDK instance.
 */

import (
	"fmt"

	"github.com/sirupsen/logrus"
)

// Define a simple hook
type SimpleHook struct{}

func (s *SimpleHook) BeforeRequest(req interface{}) error {
	fmt.Println("SimpleHook: BeforeRequest")
	return nil
}

func (s *SimpleHook) AfterSuccess(resp interface{}) error {
	fmt.Println("SimpleHook: AfterSuccess")
	return nil
}

func (s *SimpleHook) AfterError(err error) error {
	fmt.Println("SimpleHook: AfterError")
	return nil
}

func initHooks(h *Hooks) {
	// Initialize logger
	logger := logrus.New()
	logger.SetLevel(logrus.DebugLevel)

	// Register URL Rewrite Hook
	urlRewriteHook := NewURLRewriteHook(logger)
	h.registerBeforeRequestHook(urlRewriteHook)

	fmt.Println("Registered URLRewriteHook")
}
