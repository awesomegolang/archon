// +build !ignore_autogenerated

/*
Copyright 2017 The Kubernetes Authors.

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

// This file was autogenerated by deepcopy-gen. Do not edit it manually!

package apiextensions

import (
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
	reflect "reflect"
)

func init() {
	SchemeBuilder.Register(RegisterDeepCopies)
}

// RegisterDeepCopies adds deep-copy functions to the given scheme. Public
// to allow building arbitrary schemes.
func RegisterDeepCopies(scheme *runtime.Scheme) error {
	return scheme.AddGeneratedDeepCopyFuncs(
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_apiextensions_CustomResourceDefinition, InType: reflect.TypeOf(&CustomResourceDefinition{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_apiextensions_CustomResourceDefinitionCondition, InType: reflect.TypeOf(&CustomResourceDefinitionCondition{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_apiextensions_CustomResourceDefinitionList, InType: reflect.TypeOf(&CustomResourceDefinitionList{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_apiextensions_CustomResourceDefinitionNames, InType: reflect.TypeOf(&CustomResourceDefinitionNames{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_apiextensions_CustomResourceDefinitionSpec, InType: reflect.TypeOf(&CustomResourceDefinitionSpec{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_apiextensions_CustomResourceDefinitionStatus, InType: reflect.TypeOf(&CustomResourceDefinitionStatus{})},
	)
}

func DeepCopy_apiextensions_CustomResourceDefinition(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*CustomResourceDefinition)
		out := out.(*CustomResourceDefinition)
		*out = *in
		if newVal, err := c.DeepCopy(&in.ObjectMeta); err != nil {
			return err
		} else {
			out.ObjectMeta = *newVal.(*v1.ObjectMeta)
		}
		if err := DeepCopy_apiextensions_CustomResourceDefinitionSpec(&in.Spec, &out.Spec, c); err != nil {
			return err
		}
		if err := DeepCopy_apiextensions_CustomResourceDefinitionStatus(&in.Status, &out.Status, c); err != nil {
			return err
		}
		return nil
	}
}

func DeepCopy_apiextensions_CustomResourceDefinitionCondition(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*CustomResourceDefinitionCondition)
		out := out.(*CustomResourceDefinitionCondition)
		*out = *in
		if newVal, err := c.DeepCopy(&in.LastTransitionTime); err != nil {
			return err
		} else {
			out.LastTransitionTime = *newVal.(*v1.Time)
		}
		return nil
	}
}

func DeepCopy_apiextensions_CustomResourceDefinitionList(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*CustomResourceDefinitionList)
		out := out.(*CustomResourceDefinitionList)
		*out = *in
		if in.Items != nil {
			in, out := &in.Items, &out.Items
			*out = make([]CustomResourceDefinition, len(*in))
			for i := range *in {
				if err := DeepCopy_apiextensions_CustomResourceDefinition(&(*in)[i], &(*out)[i], c); err != nil {
					return err
				}
			}
		}
		return nil
	}
}

func DeepCopy_apiextensions_CustomResourceDefinitionNames(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*CustomResourceDefinitionNames)
		out := out.(*CustomResourceDefinitionNames)
		*out = *in
		if in.ShortNames != nil {
			in, out := &in.ShortNames, &out.ShortNames
			*out = make([]string, len(*in))
			copy(*out, *in)
		}
		return nil
	}
}

func DeepCopy_apiextensions_CustomResourceDefinitionSpec(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*CustomResourceDefinitionSpec)
		out := out.(*CustomResourceDefinitionSpec)
		*out = *in
		if err := DeepCopy_apiextensions_CustomResourceDefinitionNames(&in.Names, &out.Names, c); err != nil {
			return err
		}
		return nil
	}
}

func DeepCopy_apiextensions_CustomResourceDefinitionStatus(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*CustomResourceDefinitionStatus)
		out := out.(*CustomResourceDefinitionStatus)
		*out = *in
		if in.Conditions != nil {
			in, out := &in.Conditions, &out.Conditions
			*out = make([]CustomResourceDefinitionCondition, len(*in))
			for i := range *in {
				if err := DeepCopy_apiextensions_CustomResourceDefinitionCondition(&(*in)[i], &(*out)[i], c); err != nil {
					return err
				}
			}
		}
		if err := DeepCopy_apiextensions_CustomResourceDefinitionNames(&in.AcceptedNames, &out.AcceptedNames, c); err != nil {
			return err
		}
		return nil
	}
}