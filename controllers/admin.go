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
	for _, cat := range models.GetCategories() {
		var err error
		var catRef *firestore.DocumentRef
		var catData = cat.ToFirestoreCategory(nil)

		// Get category by slug
		catSnapshot, err := categoryCollection.Where("slug", "==", cat.Slug).Documents(ctx).Next()
		if err != nil {
			// If category doesn't exists, add it
			catRef, _, err = categoryCollection.Add(ctx, catData)
			if err != nil {
				logs.Error("Add: Couldn't add new category:", err)
				return
			}
		} else {
			catRef = catSnapshot.Ref
		}

		// Loop sub categories and add or update it
		var child []*firestore.DocumentRef
		for _, subCat := range cat.Child {
			var subRef *firestore.DocumentRef
			// Generate new sub-category's slug
			subCat.Slug = cat.Slug + "/" + subCat.Slug
			// Get sub category by slug
			subSnap, err := categoryCollection.Where("slug", "==", subCat.Slug).Documents(ctx).Next()
			if err != nil {
				// If sub category doesn't exists, add it
				subRef, _, err = categoryCollection.Add(ctx, subCat.ToFirestoreCategory(catRef))
				if err != nil {
					logs.Error("Add: Couldn't add new sub category:", err)
					return
				}
			} else {
				subRef = subSnap.Ref
			}
			// Add sub category to child list
			child = append(child, subRef)
		}

		// Update parent category's child list
		catData.Child = child

		// Update parent category with new data
		_, err = catRef.Set(ctx, catData)
		if err != nil {
			logs.Error("Add: Couldn't save category:", err)
			return
		}
	}

	res = "ok"
}
