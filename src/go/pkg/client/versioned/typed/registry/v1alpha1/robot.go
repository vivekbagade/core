// Copyright 2019 The Google Cloud Robotics Authors
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

// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	"time"

	v1alpha1 "github.com/googlecloudrobotics/core/src/go/pkg/apis/registry/v1alpha1"
	scheme "github.com/googlecloudrobotics/core/src/go/pkg/client/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// RobotsGetter has a method to return a RobotInterface.
// A group's client should implement this interface.
type RobotsGetter interface {
	Robots(namespace string) RobotInterface
}

// RobotInterface has methods to work with Robot resources.
type RobotInterface interface {
	Create(*v1alpha1.Robot) (*v1alpha1.Robot, error)
	Update(*v1alpha1.Robot) (*v1alpha1.Robot, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.Robot, error)
	List(opts v1.ListOptions) (*v1alpha1.RobotList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.Robot, err error)
	RobotExpansion
}

// robots implements RobotInterface
type robots struct {
	client rest.Interface
	ns     string
}

// newRobots returns a Robots
func newRobots(c *RegistryV1alpha1Client, namespace string) *robots {
	return &robots{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the robot, and returns the corresponding robot object, and an error if there is any.
func (c *robots) Get(name string, options v1.GetOptions) (result *v1alpha1.Robot, err error) {
	result = &v1alpha1.Robot{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("robots").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Robots that match those selectors.
func (c *robots) List(opts v1.ListOptions) (result *v1alpha1.RobotList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.RobotList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("robots").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested robots.
func (c *robots) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("robots").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a robot and creates it.  Returns the server's representation of the robot, and an error, if there is any.
func (c *robots) Create(robot *v1alpha1.Robot) (result *v1alpha1.Robot, err error) {
	result = &v1alpha1.Robot{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("robots").
		Body(robot).
		Do().
		Into(result)
	return
}

// Update takes the representation of a robot and updates it. Returns the server's representation of the robot, and an error, if there is any.
func (c *robots) Update(robot *v1alpha1.Robot) (result *v1alpha1.Robot, err error) {
	result = &v1alpha1.Robot{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("robots").
		Name(robot.Name).
		Body(robot).
		Do().
		Into(result)
	return
}

// Delete takes name of the robot and deletes it. Returns an error if one occurs.
func (c *robots) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("robots").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *robots) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("robots").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched robot.
func (c *robots) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.Robot, err error) {
	result = &v1alpha1.Robot{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("robots").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
