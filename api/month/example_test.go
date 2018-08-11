// Copyright (c) 2018, Bruno M V Souza <github@b.bmvs.io>. All rights reserved.
// Use of this source code is governed by a BSD-2-Clause license that can be
// found in the LICENSE file.

package month_test

import (
	"fmt"
	"reflect"

	"go.bmvs.io/ynab"
	"go.bmvs.io/ynab/api"
)

func ExampleService_GetMonth() {
	c := ynab.NewClient("<valid_ynab_access_token>")
	d, _ := api.DateFromString("2010-01-01")
	m, _ := c.Month().GetMonth("<valid_budget_id>", d)
	fmt.Println(reflect.TypeOf(m))

	// Output: *month.Month
}

func ExampleService_GetMonths() {
	c := ynab.NewClient("<valid_ynab_access_token>")
	months, _ := c.Month().GetMonths("<valid_budget_id>")
	fmt.Println(reflect.TypeOf(months))

	// Output: []*month.Summary
}
