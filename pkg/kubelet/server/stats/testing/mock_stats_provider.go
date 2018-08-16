/*
Copyright 2016 The Kubernetes Authors.

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

package testing

import "k8s.io/kubernetes/pkg/kubelet/cm"
import corev1 "k8s.io/api/core/v1"
import "github.com/stretchr/testify/mock"

import "k8s.io/apimachinery/pkg/types"
import "github.com/google/cadvisor/info/v1"
import "k8s.io/kubernetes/pkg/kubelet/apis/stats/v1alpha1"
import "k8s.io/kubernetes/pkg/volume"

// GENERATED BY mockery

// StatsProvider is an autogenerated mock type for the StatsProvider type
type StatsProvider struct {
	mock.Mock
}

// GetCgroupStats provides a mock function with given fields: cgroupName, updateStats
func (_m *StatsProvider) GetCgroupStats(cgroupName string, updateStats bool) (*v1alpha1.ContainerStats, *v1alpha1.NetworkStats, error) {
	ret := _m.Called(cgroupName, updateStats)

	var r0 *v1alpha1.ContainerStats
	if rf, ok := ret.Get(0).(func(string, bool) *v1alpha1.ContainerStats); ok {
		r0 = rf(cgroupName, updateStats)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1alpha1.ContainerStats)
		}
	}

	var r1 *v1alpha1.NetworkStats
	if rf, ok := ret.Get(1).(func(string, bool) *v1alpha1.NetworkStats); ok {
		r1 = rf(cgroupName, updateStats)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*v1alpha1.NetworkStats)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(string, bool) error); ok {
		r2 = rf(cgroupName, updateStats)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// GetPodByCgroupfs provides the pod that maps to the specified cgroup, as well
// as whether the pod was found.
func (_m *StatsProvider) GetPodByCgroupfs(cgroupfs string) (*corev1.Pod, bool) {
	return nil, false
}

// GetContainerInfo provides a mock function with given fields: podFullName, uid, containerName, req
func (_m *StatsProvider) GetContainerInfo(podFullName string, uid types.UID, containerName string, req *v1.ContainerInfoRequest) (*v1.ContainerInfo, error) {
	ret := _m.Called(podFullName, uid, containerName, req)

	var r0 *v1.ContainerInfo
	if rf, ok := ret.Get(0).(func(string, types.UID, string, *v1.ContainerInfoRequest) *v1.ContainerInfo); ok {
		r0 = rf(podFullName, uid, containerName, req)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1.ContainerInfo)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, types.UID, string, *v1.ContainerInfoRequest) error); ok {
		r1 = rf(podFullName, uid, containerName, req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetNode provides a mock function with given fields:
func (_m *StatsProvider) GetNode() (*corev1.Node, error) {
	ret := _m.Called()

	var r0 *corev1.Node
	if rf, ok := ret.Get(0).(func() *corev1.Node); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*corev1.Node)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetNodeConfig provides a mock function with given fields:
func (_m *StatsProvider) GetNodeConfig() cm.NodeConfig {
	ret := _m.Called()

	var r0 cm.NodeConfig
	if rf, ok := ret.Get(0).(func() cm.NodeConfig); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(cm.NodeConfig)
	}

	return r0
}

// GetPodCgroupRoot provides a mock function with given fields:
func (_m *StatsProvider) GetPodCgroupRoot() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// GetPodByName provides a mock function with given fields: namespace, name
func (_m *StatsProvider) GetPodByName(namespace string, name string) (*corev1.Pod, bool) {
	ret := _m.Called(namespace, name)

	var r0 *corev1.Pod
	if rf, ok := ret.Get(0).(func(string, string) *corev1.Pod); ok {
		r0 = rf(namespace, name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*corev1.Pod)
		}
	}

	var r1 bool
	if rf, ok := ret.Get(1).(func(string, string) bool); ok {
		r1 = rf(namespace, name)
	} else {
		r1 = ret.Get(1).(bool)
	}

	return r0, r1
}

// GetPods provides a mock function with given fields:
func (_m *StatsProvider) GetPods() []*corev1.Pod {
	ret := _m.Called()

	var r0 []*corev1.Pod
	if rf, ok := ret.Get(0).(func() []*corev1.Pod); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*corev1.Pod)
		}
	}

	return r0
}

// GetRawContainerInfo provides a mock function with given fields: containerName, req, subcontainers
func (_m *StatsProvider) GetRawContainerInfo(containerName string, req *v1.ContainerInfoRequest, subcontainers bool) (map[string]*v1.ContainerInfo, error) {
	ret := _m.Called(containerName, req, subcontainers)

	var r0 map[string]*v1.ContainerInfo
	if rf, ok := ret.Get(0).(func(string, *v1.ContainerInfoRequest, bool) map[string]*v1.ContainerInfo); ok {
		r0 = rf(containerName, req, subcontainers)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]*v1.ContainerInfo)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, *v1.ContainerInfoRequest, bool) error); ok {
		r1 = rf(containerName, req, subcontainers)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ImageFsStats provides a mock function with given fields:
func (_m *StatsProvider) ImageFsStats() (*v1alpha1.FsStats, error) {
	ret := _m.Called()

	var r0 *v1alpha1.FsStats
	if rf, ok := ret.Get(0).(func() *v1alpha1.FsStats); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1alpha1.FsStats)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListPodStats provides a mock function with given fields:
func (_m *StatsProvider) ListPodStats() ([]v1alpha1.PodStats, error) {
	ret := _m.Called()

	var r0 []v1alpha1.PodStats
	if rf, ok := ret.Get(0).(func() []v1alpha1.PodStats); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]v1alpha1.PodStats)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListVolumesForPod provides a mock function with given fields: podUID
func (_m *StatsProvider) ListVolumesForPod(podUID types.UID) (map[string]volume.Volume, bool) {
	ret := _m.Called(podUID)

	var r0 map[string]volume.Volume
	if rf, ok := ret.Get(0).(func(types.UID) map[string]volume.Volume); ok {
		r0 = rf(podUID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]volume.Volume)
		}
	}

	var r1 bool
	if rf, ok := ret.Get(1).(func(types.UID) bool); ok {
		r1 = rf(podUID)
	} else {
		r1 = ret.Get(1).(bool)
	}

	return r0, r1
}

// RootFsStats provides a mock function with given fields:
func (_m *StatsProvider) RootFsStats() (*v1alpha1.FsStats, error) {
	ret := _m.Called()

	var r0 *v1alpha1.FsStats
	if rf, ok := ret.Get(0).(func() *v1alpha1.FsStats); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1alpha1.FsStats)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RlimitStats provides a mock function with given fields:
func (_m *StatsProvider) RlimitStats() (*v1alpha1.RlimitStats, error) {
	ret := _m.Called()

	var r0 *v1alpha1.RlimitStats
	if rf, ok := ret.Get(0).(func() *v1alpha1.RlimitStats); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1alpha1.RlimitStats)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
