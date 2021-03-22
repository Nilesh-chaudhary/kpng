/*
Copyright 2021 The Kubernetes Authors.

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

package global

import (
	"google.golang.org/grpc"
	"m.cluseau.fr/kpng/pkg/api/localnetv1"
	"m.cluseau.fr/kpng/pkg/proxystore"
)

func Setup(s grpc.ServiceRegistrar, store *proxystore.Store) {
	localnetv1.RegisterGlobalService(s, localnetv1.NewGlobalService(localnetv1.UnstableGlobalService(&Server{
		Store: store,
	})))
}
