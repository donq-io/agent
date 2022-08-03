package node_exporter

import (
	"context"
	"fmt"
	"net/http"
	"sync"

	"github.com/go-kit/log"
	"github.com/grafana/agent/component"
	node_integration "github.com/grafana/agent/pkg/integrations/node_exporter"
	"github.com/prometheus/common/model"
)

func init() {
	component.Register(component.Registration{
		Name:    "integration.node_exporter",
		Args:    node_integration.Config{},
		Exports: Exports{},

		Build: func(opts component.Options, args component.Arguments) (component.Component, error) {
			return NewComponent(opts, args.(node_integration.Config))
		},
	})
}

// Target refers to a singular HTTP or HTTPS endpoint that will be used for scraping.
// Here, we're using a map[string]string instead of labels.Labels; if the label ordering
// is important, we can change to follow the upstream logic instead.
// TODO (@tpaschalis) Maybe the target definitions should be part of the
// Service Discovery components package. Let's reconsider once it's ready.
type Target map[string]string

type Exports struct {
	Output []Target `river:"output,attr"`
}

type Component struct {
	log  log.Logger
	opts component.Options

	mut sync.RWMutex
	cfg *node_integration.Config

	integration *node_integration.Integration
}

func NewComponent(o component.Options, args node_integration.Config) (*Component, error) {
	c := &Component{
		log:  o.Logger,
		cfg:  &args,
		opts: o,
	}

	// Call to Update() to set the output once at the start
	if err := c.Update(args); err != nil {
		return nil, err
	}

	return c, nil
}

// Run implements component.Component.
func (c *Component) Run(ctx context.Context) error {
	c.log.Log("Msg", "Running")
	<-ctx.Done()
	return nil
}

// Update implements component.Component.
func (c *Component) Update(args component.Arguments) error {
	c.log.Log("Msg", "Update")
	var err error
	c.integration, err = node_integration.New(c.log, c.cfg)
	targets := []Target{{
		model.AddressLabel:     "127.0.0.1:12345",
		model.SchemeLabel:      "http",
		model.MetricsPathLabel: fmt.Sprintf("/component/%s/metrics", c.opts.ID),
		"name":                 "node_exporter",
	}}
	c.opts.OnStateChange(Exports{
		Output: targets,
	})
	return err
}

func (c *Component) Handler() http.Handler {
	if c.integration != nil {
		// todo: handle
		h, _ := c.integration.MetricsHandler()
		return h
	}
	return nil
}
