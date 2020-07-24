/*
Copyright 2020 VMware, Inc.
SPDX-License-Identifier: Apache-2.0
*/

package starter

import (
	"context"
	"fmt"
	"time"

	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/events"

	"github.com/suzerain-io/placeholder-name/internal/controller"
	examplecontroller "github.com/suzerain-io/placeholder-name/test/integration/examplecontroller/controller"
)

func StartExampleController(ctx context.Context, config *rest.Config, secretData string) error {
	kubeClient, err := kubernetes.NewForConfig(config)
	if err != nil {
		return fmt.Errorf("failed to build client: %w", err)
	}

	kubeInformers := informers.NewSharedInformerFactory(kubeClient, 20*time.Minute)

	recorder := events.NewEventBroadcasterAdapter(kubeClient).NewRecorder("example-controller")

	manager := controller.NewManager().
		WithController(
			examplecontroller.NewExampleCreatingController(
				kubeInformers.Core().V1().Services(),
				kubeInformers.Core().V1().Secrets(),
				kubeClient.CoreV1(),
				recorder,
				secretData,
			), 5,
		).
		WithController(
			examplecontroller.NewExampleUpdatingController(
				kubeInformers.Core().V1().Services(),
				kubeInformers.Core().V1().Secrets(),
				kubeClient.CoreV1(),
				recorder,
				secretData,
			), 5,
		)

	kubeInformers.Start(ctx.Done())
	go manager.Start(ctx)

	return nil
}