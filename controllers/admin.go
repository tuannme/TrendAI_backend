package controllers

import (
	"cloud.google.com/go/firestore"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/trend-ai/TrendAI_mobile_backend/models"
	"github.com/trend-ai/TrendAI_mobile_backend/services/databases"
)

type AdminController struct {
	beego.Controller
}

// @Title PushInterestCategories
// @Description Push interest categories to database
// @Success 200 {object} models.AuthenticationResponse
// @Failure 400 {object} models.ResponseWithError
// @router /interest_categories [put]
func (o *AdminController) PushInterestCategories() {
	var res models.Response
	defer func() {
		o.Data["json"] = res
		o.ServeJSON()
	}()

	ctx := databases.Context
	categoryCollection := models.GetCategoryCollection()

	// Loop all categories in our model to update to database
	for _, rawCategory := range models.GetCategories() {
		var err error
		var categoryRef *firestore.DocumentRef
		var category = rawCategory.ToCategory(nil, nil)

		// Get category by slug
		catSnapshot, err := categoryCollection.Where("slug", "==", category.Slug).Documents(ctx).Next()
		if err != nil {
			// If category doesn't exists, add it
			categoryRef, _, err = categoryCollection.Add(ctx, category)
			if err != nil {
				logs.Error("Add: Couldn't add new category:", err)
				return
			}
		} else {
			categoryRef = catSnapshot.Ref
		}

		// Loop sub categories and add or update it
		var child []*firestore.DocumentRef
		for _, rawSubCat := range rawCategory.Child {
			var subCategory = rawSubCat.ToCategory(categoryRef, nil)
			var subCategoryRef *firestore.DocumentRef
			// Generate new sub-category's slug by parent slug
			subCategory.Slug = category.Slug + "/" + subCategory.Slug
			// Get sub category by slug
			subSnap, err := categoryCollection.Where("slug", "==", subCategory.Slug).Documents(ctx).Next()
			if err != nil {
				// If sub category doesn't exists, add it
				subCategoryRef, _, err = categoryCollection.Add(ctx, subCategory)
				if err != nil {
					logs.Error("Add: Couldn't add new sub category:", err)
					return
				}
			} else {
				subCategoryRef = subSnap.Ref
			}
			// Add sub category to child list
			child = append(child, subCategoryRef)
		}

		// Update parent category's child list
		category.Child = child

		// Update parent category with new data
		_, err = categoryRef.Set(ctx, category)
		if err != nil {
			logs.Error("Add: Couldn't save category:", err)
			return
		}
	}

	res = "ok"
}
