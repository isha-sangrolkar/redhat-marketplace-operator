// Copyright 2020 IBM Corp.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package meter_definition

import (
	"context"
	"reflect"

	"github.com/go-logr/logr"
	monitoringv1 "github.com/prometheus-operator/prometheus-operator/pkg/apis/monitoring/v1"
	rhmclient "github.com/redhat-marketplace/redhat-marketplace-operator/v2/pkg/client"
	"github.com/redhat-marketplace/redhat-marketplace-operator/v2/pkg/utils"
	. "github.com/redhat-marketplace/redhat-marketplace-operator/v2/pkg/utils/reconcileutils"
	corev1 "k8s.io/api/core/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// ServiceProcessor will update the meter definition
// status with the objects that matched it.
type ServiceProcessor struct {
	log logr.Logger
	cc  ClientCommandRunner
}

// NewServiceProcessor is the provider that creates
// the processor.
func NewServiceProcessor(
	log logr.Logger,
	cc ClientCommandRunner,
) *ServiceProcessor {
	return &ServiceProcessor{
		log: log,
		cc:  cc,
	}
}

// Start will register it's listener and execute the function.
func (u *ServiceProcessor) New(store *MeterDefinitionStore) Processor {
	p := NewProcessor("serviceProcessor", u.log, u.cc, store, u)
	return p
}

var serviceType = reflect.TypeOf(&corev1.Service{})

// Process will receive a new ObjectResourceMessage and find and update the metere
// definition associated with the object. To prevent gaps, it bulk retrieves the
// resoruces and checks it against the status.
func (u *ServiceProcessor) Process(ctx context.Context, inObj *ObjectResourceMessage) error {
	log := u.log.WithValues("process", "serviceProcessor")

	if inObj == nil {
		return nil
	}

	if inObj.Action == DeleteMessageAction {
		return nil
	}

	if reflect.TypeOf(inObj.Object) != serviceType {
		return nil
	}

	log.Info("checking for servicemonitor for service", "uid", inObj.UID)
	list := &monitoringv1.ServiceMonitorList{}

	result, _ := u.cc.Do(ctx,
		HandleResult(
			ListAction(list, client.MatchingFields{
				rhmclient.IndexOwnerRefContains: string(inObj.UID),
			}),
			OnContinue(Call(func() (ClientAction, error) {
				actions := []ClientAction{}

				log.Info("found servicemonitors", "len", len(list.Items))

				for _, sm := range list.Items {
					if !utils.HasMapKey(sm.ObjectMeta.Labels, utils.MeteredAnnotation) {
						if sm.ObjectMeta.Labels == nil {
							sm.ObjectMeta.Labels = make(map[string]string)
						}

						utils.SetMapKeyValue(sm.ObjectMeta.Labels, utils.MeteredAnnotation)

						log.Info("found servicemonitor to label", "sm", sm, "labels", sm.Labels)
						actions = append(actions, HandleResult(
							UpdateAction(sm),
							OnRequeue(ContinueResponse())))
					}
				}

				if len(actions) == 0 {
					return nil, nil
				}

				return Do(actions...), nil
			})),
		))

	if result.Is(NotFound) {
		return nil
	}

	if result.Is(Error) {
		log.Error(result, "failed to update")
		return result
	}

	return nil
}
