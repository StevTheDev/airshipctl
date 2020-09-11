/*
 Licensed under the Apache License, Version 2.0 (the "License");
 you may not use this file except in compliance with the License.
 You may obtain a copy of the License at

     https://www.apache.org/licenses/LICENSE-2.0

 Unless required by applicable law or agreed to in writing, software
 distributed under the License is distributed on an "AS IS" BASIS,
 WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 See the License for the specific language governing permissions and
 limitations under the License.
*/

package ifc

import (
	"opendev.org/airship/airshipctl/pkg/api/v1alpha1"
	"opendev.org/airship/airshipctl/pkg/cluster/clustermap"
)

// Phase provides a way to interact with a phase
// TODO add render method
type Phase interface {
	Validate() error
	Run(RunOptions) error
	DocumentRoot() string
	Details() (string, error)
	Executor() (Executor, error)
}

// ID uniquely identifies the phase
type ID struct {
	Name      string
	Namespace string
}

// Client is a phase client that can be used by command line or ui packages
type Client interface {
	PhaseByID(ID) (Phase, error)
	PhaseByAPIObj(*v1alpha1.Phase) (Phase, error)
	ClusterMap() (clustermap.ClusterMap, error)
}