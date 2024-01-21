package runnables

import (
	"context"
	"time"

	"github.com/go-logr/logr"
	"github.com/pkg/errors"

	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/medik8s/customized-user-remediation/api/v1alpha1"
	"github.com/medik8s/customized-user-remediation/pkg/utils"
)

type Creator struct {
	client.Client
	logger logr.Logger
}

func NewTemplateCreator(cl client.Client, logger logr.Logger) *Creator {
	return &Creator{
		Client: cl,
		logger: logger,
	}
}

func (c *Creator) Start(ctx context.Context) error {

	// dirty workaround for waiting until webhook is up and running
	time.Sleep(10 * time.Second)

	if err := c.newTemplatesIfNotExist(ctx); err != nil {
		c.logger.Error(err, "failed to create remediation templates")
		return err
	}
	return nil
}

// newTemplatesIfNotExist creates new SelfNodeRemediationTemplate objects
func (c *Creator) newTemplatesIfNotExist(ctx context.Context) error {
	ns, err := utils.GetDeploymentNamespace()
	if err != nil {
		return errors.Wrap(err, "unable to get the deployment namespace")
	}

	templates := v1alpha1.NewRemediationTemplates()

	for _, template := range templates {
		template.SetNamespace(ns)
		err = c.Create(ctx, template, &client.CreateOptions{})
		if err != nil && !apierrors.IsAlreadyExists(err) {
			return errors.Wrap(err, "failed to create CustomizedUserRemediationTemplate CR")
		}
	}
	return nil
}
