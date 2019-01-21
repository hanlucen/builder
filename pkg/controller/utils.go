package controller

import (
	"fmt"

	"github.com/drycc/builder/pkg/conf"
	drycc "github.com/drycc/controller-sdk-go"
	"github.com/drycc/pkg/log"
)

// New creates a new SDK client configured as the builder.
func New(host, port string) (*drycc.Client, error) {

	client, err := drycc.New(true, fmt.Sprintf("http://%s:%s/", host, port), "")
	if err != nil {
		return client, err
	}
	client.UserAgent = "drycc-builder"

	builderKey, err := conf.GetBuilderKey()
	if err != nil {
		return client, err
	}
	client.HooksToken = builderKey

	return client, nil
}

// CheckAPICompat checks for API compatibility errors and warns about them.
func CheckAPICompat(c *drycc.Client, err error) error {
	if err == drycc.ErrAPIMismatch {
		log.Info("WARNING: SDK and Controller API versions do not match. SDK: %s Controller: %s",
			drycc.APIVersion, c.ControllerAPIVersion)

		// API mismatch isn't fatal, so after warning continue on.
		return nil
	}

	return err
}
