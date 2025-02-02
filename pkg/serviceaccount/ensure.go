/*
SPDX-License-Identifier: Apache-2.0

Copyright Contributors to the Submariner project.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package serviceaccount

import (
	"context"

	"github.com/submariner-io/admiral/pkg/resource"
	resourceutil "github.com/submariner-io/subctl/pkg/resource"
	"github.com/submariner-io/submariner-operator/pkg/embeddedyamls"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/client-go/kubernetes"
)

// EnsureFromYAML creates the given service account.
// nolint:wrapcheck // No need to wrap errors here.
func EnsureFromYAML(kubeClient kubernetes.Interface, namespace, yaml string) (bool, error) {
	sa := &corev1.ServiceAccount{}

	err := embeddedyamls.GetObject(yaml, sa)
	if err != nil {
		return false, err
	}

	return Ensure(kubeClient, namespace, sa)
}

// nolint:wrapcheck // No need to wrap errors here.
func Ensure(kubeClient kubernetes.Interface, namespace string, sa *corev1.ServiceAccount) (bool, error) {
	return resourceutil.CreateOrUpdate(context.TODO(), resource.ForServiceAccount(kubeClient, namespace), sa)
}
