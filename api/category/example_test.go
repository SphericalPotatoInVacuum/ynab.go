package category_test

import (
	"fmt"
	"reflect"

	"go.bmvs.io/ynab"
)

func ExampleService_GetCategory() {
	c := ynab.NewClient("<valid_ynab_access_token>")
	category, _ := c.Category().GetCategory("<valid_budget_id>", "<valid_category_id>")
	fmt.Println(reflect.TypeOf(category))

	// Output: *category.Category
}

func ExampleService_GetCategories() {
	c := ynab.NewClient("<valid_ynab_access_token>")
	categories, _ := c.Category().GetCategories("<valid_budget_id>")
	fmt.Println(reflect.TypeOf(categories))

	// Output: []*category.GroupWithCategories
}
