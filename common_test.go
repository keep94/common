// Copyright 2014 Travis Keep. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file or
// at http://opensource.org/licenses/BSD-3-Clause.

package common_test

import (
  "fmt"
  "strings"
  "testing"
  "github.com/keep94/common"
)

var (
  kNone1 noneType = 1
  kNone2 noneType = 2
  kNone3 noneType = 3
  kSingle1 singleType = "one"
  kSingle2 singleType = "two"
  kSingle3 singleType = "three"
  kMulti456 = multiType{
      singleType("four"),
      singleType("five"),
      singleType("six"),
  }
  kMulti78 = multiType{
      singleType("seven"),
      singleType("eight"),
  }
)

func TestControl(t *testing.T) {
  arr := multiType{
      kNone2,
      kMulti78,
      kSingle1,
  }
  expected := "[NONE [seven eight] one]"
  assertStrEqual(t, expected, arr.ToString())
}

func TestNone(t *testing.T) {
  assertStrEqual(t, "NONE", join(nil))
  assertStrEqual(t, "NONE", join([]fooType{kNone1, kNone2, kNone3}))
}

func TestSingle(t *testing.T) {
  assertStrEqual(t, "one", join([]fooType{kSingle1}))
  assertStrEqual(t, "one", join([]fooType{kNone1, kSingle1, kNone2}))
}

func TestMulti(t *testing.T) {
  assertStrEqual(
      t,
      "[one three four five six two seven eight]",
      join([]fooType{
          kNone1,
          kSingle1,
          kSingle3,
          kNone2,
          kMulti456,
          kSingle2,
          kMulti78,
          kNone3,
      }))
  assertStrEqual(
      t,
      "[four five six]",
      join([]fooType{
          kNone1,
          kMulti456,
          kNone3,
      }))
  assertStrEqual(
      t,
      "[four five six]",
      join([]fooType{
          kMulti456,
      }))
}

func join(arr []fooType) string {
  var agg multiType
  var none noneType
  return common.Join(arr, agg, none).(fooType).ToString()
}

func assertStrEqual(t *testing.T, expected, actual string) {
  if expected != actual {
    t.Errorf("Expected %s, got %s", expected, actual)
  }
}

type fooType interface {
  ToString() string
}

type noneType int

func (n noneType) ToString() string {
  return "NONE"
}

type singleType string

func (s singleType) ToString() string {
  return string(s)
}

type multiType []fooType

func (m multiType) ToString() string {
  strs := make([]string, len(m))
  for i := range m {
    strs[i] = m[i].ToString()
  }
  return fmt.Sprintf("[%s]", strings.Join(strs, " "))
}
