/*
Copyright 2020 VMware, Inc.
SPDX-License-Identifier: Apache-2.0
*/

package testutil

import "github.com/suzerain-io/controller-go"

type ObservableWithInformerOption struct {
	informerToFilterMap map[controller.InformerGetter]controller.Filter
}

func NewObservableWithInformerOption() *ObservableWithInformerOption {
	return &ObservableWithInformerOption{
		informerToFilterMap: make(map[controller.InformerGetter]controller.Filter),
	}
}

func (i *ObservableWithInformerOption) WithInformer(
	getter controller.InformerGetter,
	filter controller.Filter,
	opt controller.InformerOption,
) controller.Option {
	i.informerToFilterMap[getter] = filter
	return controller.WithInformer(getter, filter, opt)
}

func (i *ObservableWithInformerOption) GetFilterForInformer(getter controller.InformerGetter) controller.Filter {
	return i.informerToFilterMap[getter]
}
