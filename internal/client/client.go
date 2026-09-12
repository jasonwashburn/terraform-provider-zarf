// Package client provides a client for performing zarf operations.
package client

import (
	"context"

	"github.com/zarf-dev/zarf/src/pkg/packager"
)

type Client interface {
	InspectPackage(ctx context.Context, source string) (PackageData, error)
}

type PackageData struct {
	Metadata ZarfPackageMetadata
	Digest   string
}

type ZarfPackageMetadata struct {
	Name         string
	Description  string
	Version      string
	URL          string
	Architecture string
}

func NewClient() (Client, error) {
	return &client{}, nil
}

type client struct{}

func (c *client) InspectPackage(ctx context.Context, source string) (PackageData, error) {
	opts := packager.LoadOptions{}
	layout, err := packager.LoadPackage(ctx, source, opts)
	if err != nil {
		return PackageData{}, err
	}
	defer func() {
		if err := layout.Cleanup(); err != nil {
			// Log the error but don't return it, as we want to return the package data if possible
			// TODO: Use a proper logging mechanism instead of printing to stdout
			println("Error cleaning up package layout:", err.Error())
		}
	}()
	metadata := layout.PackageDefinition.AsV1alpha1().Metadata
	return PackageData{
		Metadata: ZarfPackageMetadata{
			Name:         metadata.Name,
			Description:  metadata.Description,
			Version:      metadata.Version,
			URL:          metadata.URL,
			Architecture: metadata.Architecture,
		},
		Digest: layout.Digest(),
	}, nil
}
