//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// Code generated by operator-sdk. DO NOT EDIT.

package v1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OnePasswordItem) DeepCopyInto(out *OnePasswordItem) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OnePasswordItem.
func (in *OnePasswordItem) DeepCopy() *OnePasswordItem {
	if in == nil {
		return nil
	}
	out := new(OnePasswordItem)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *OnePasswordItem) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OnePasswordItemCondition) DeepCopyInto(out *OnePasswordItemCondition) {
	*out = *in
	in.LastTransitionTime.DeepCopyInto(&out.LastTransitionTime)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OnePasswordItemCondition.
func (in *OnePasswordItemCondition) DeepCopy() *OnePasswordItemCondition {
	if in == nil {
		return nil
	}
	out := new(OnePasswordItemCondition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OnePasswordItemList) DeepCopyInto(out *OnePasswordItemList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]OnePasswordItem, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OnePasswordItemList.
func (in *OnePasswordItemList) DeepCopy() *OnePasswordItemList {
	if in == nil {
		return nil
	}
	out := new(OnePasswordItemList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *OnePasswordItemList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OnePasswordItemSpec) DeepCopyInto(out *OnePasswordItemSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OnePasswordItemSpec.
func (in *OnePasswordItemSpec) DeepCopy() *OnePasswordItemSpec {
	if in == nil {
		return nil
	}
	out := new(OnePasswordItemSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OnePasswordItemStatus) DeepCopyInto(out *OnePasswordItemStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]OnePasswordItemCondition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Ready != nil {
		in, out := &in.Ready, &out.Ready
		*out = new(bool)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OnePasswordItemStatus.
func (in *OnePasswordItemStatus) DeepCopy() *OnePasswordItemStatus {
	if in == nil {
		return nil
	}
	out := new(OnePasswordItemStatus)
	in.DeepCopyInto(out)
	return out
}
