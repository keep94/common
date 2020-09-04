// Copyright 2014 Travis Keep. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file or
// at http://opensource.org/licenses/BSD-3-Clause.

// Package common provides common routines for github.com/keep94.
// Package common is for internal use only.
package common

import (
	"reflect"
)

// Join implements the composite patten.
// Let T be an arbitrary interface.
// aSlice is a []T consisting of the T values to be joined.
// aggregate is an arbitrary value of type U that implements T and has an
// underlying type []T.
// Caller usually passes nil for aggregate as aggregate is only used for
// reflection.
// none is of some arbitrary type V that implements T and represents no T
// values.
// Join returns some value of type T.
// Join returns none if aSlice is empty or consists only of values of type V.
// If aSlice contains only one T value that is not of type V, Join returns that
// value. Otherwise Join returns a value of type U containing all the T values
// in aSlice flattened out. That is, the returned U value will not contain
// any additional U values or V values.
func Join(
	aSlice interface{},
	aggregate interface{},
	none interface{}) interface{} {
	// -1 = no not none index; -2 = more than one not none index
	notNoneIdx := -1
	computedLength := 0
	sliceVal := reflect.ValueOf(aSlice)
	noneType := reflect.TypeOf(none)
	aggregateType := reflect.TypeOf(aggregate)
	sliceLength := sliceVal.Len()
	for i := 0; i < sliceLength; i++ {
		aVal := sliceVal.Index(i).Elem()
		aType := aVal.Type()
		if aType != noneType {
			if notNoneIdx == -1 {
				notNoneIdx = i
			} else {
				notNoneIdx = -2
			}
			if aType == aggregateType {
				computedLength += aVal.Len()
			} else {
				computedLength++
			}
		}
	}
	if notNoneIdx == -1 {
		return none
	}
	if notNoneIdx >= 0 {
		return sliceVal.Index(notNoneIdx).Interface()
	}
	result := reflect.MakeSlice(
		aggregateType, computedLength, computedLength)
	idx := 0
	for i := 0; i < sliceLength; i++ {
		aiVal := sliceVal.Index(i)
		aVal := aiVal.Elem()
		aType := aVal.Type()
		if aType != noneType {
			if aType == aggregateType {
				idx += reflect.Copy(result.Slice(idx, computedLength), aVal)
			} else {
				result.Index(idx).Set(aiVal)
				idx++
			}
		}
	}
	return result.Interface()
}
