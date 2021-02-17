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

package reconcileutils

import (
	"context"
	"fmt"

	emperrors "emperror.dev/errors"
	"github.com/redhat-marketplace/redhat-marketplace-operator/v2/pkg/utils/codelocation"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
)

type getAction struct {
	BaseAction
	NamespacedName types.NamespacedName
	Object         runtime.Object
	getActionOptions
}

//go:generate go-options -option GetActionOption -prefix GetWith getActionOptions
type getActionOptions struct {
}

func GetAction(
	namespacedName types.NamespacedName,
	object runtime.Object,
	options ...GetActionOption,
) *getAction {
	opts, _ := newGetActionOptions(options...)
	return &getAction{
		NamespacedName:   namespacedName,
		Object:           object,
		getActionOptions: opts,
		BaseAction: BaseAction{
			codelocation: codelocation.New(1),
		},
	}
}

func (g *getAction) Bind(r *ExecResult) {
	g.lastResult = r
}

func (g *getAction) Exec(ctx context.Context, c *ClientCommand) (*ExecResult, error) {
	reqLogger := c.log.WithValues("file", g.codelocation, "action", "GetAction")

	if isNil(g.Object) {
		err := emperrors.New("object to get is nil")
		reqLogger.Error(err, "updatedObject is nil")
		return NewExecResult(Error, reconcile.Result{Requeue: true}, err), err
	}

	reqLogger = reqLogger.WithValues("requestType", fmt.Sprintf("%T", g.Object), "key", g.NamespacedName)

	err := c.client.Get(ctx, g.NamespacedName, g.Object)
	if err != nil {
		if errors.IsNotFound(err) {
			reqLogger.Info("not found")
			return NewExecResult(NotFound, reconcile.Result{}, err), nil
		}
		if err != nil {
			reqLogger.Error(err, "error getting")
			return NewExecResult(Error, reconcile.Result{Requeue: true}, err), emperrors.Wrap(err, "error during get")
		}
	}

	reqLogger.V(2).Info("found")
	return NewExecResult(Continue, reconcile.Result{}, err), nil
}
