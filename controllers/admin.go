package controllers

import (
	"cloud.google.com/go/firestore"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/trend-ai/TrendAI_mobile_backend/models"
	"github.com/trend-ai/TrendAI_mobile_backend/services/databases"
	"google.golang.org/api/iterator"
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
	for _, cat := range models.GetCategories() {
		catSnapshot, err := categoryCollection.Where("slug", "==", cat.Slug).Documents(ctx).Next()
		if err != nil {
			// If category doesn't exists, add it
			catRef, _, err := categoryCollection.Add(ctx, cat.BaseCategory)
			if err != nil {
				logs.Error("Add: Couldn't add new category:", err)
				return
			}
			// Add sub category to a sub collection
			subCollection := catRef.Collection("sub_categories")
			for _, s := range cat.SubCategories {
				_, _, err := subCollection.Add(ctx, s)
				if err != nil {
					logs.Error("Add: Couldn't add sub category:", err)
					return
				}
			}
		} else {
			// If category already exists, update it
			_, err := catSnapshot.Ref.Update(ctx, []firestore.Update{
				{Path: "name", Value: cat.Name},
			})
			if err != nil {
				logs.Error("Update: Couldn't update category:", err)
				return
			}
			// Get sub categories collection
			subCollection := catSnapshot.Ref.Collection("sub_categories")

			// Store all available categories (slug)
			availableCategories := make(map[string]bool, 0)

			// Loop new sub categories to update them in database
			for _, subCat := range cat.SubCategories {
				subCatSnapshot, err := subCollection.Where("slug", "==", subCat.Slug).Documents(ctx).Next()
				if err != nil {
					// If this (sub) category doesn't exists in old list, add it
					_, _, err := subCollection.Add(ctx, subCat)
					if err != nil {
						logs.Error("Update: Couldn't add sub category:", err)
						return
					}
				} else {
					// If this (sub) category already exists in old list, update it
					_, err := subCatSnapshot.Ref.Update(ctx, []firestore.Update{
						{Path: "name", Value: subCat.Name},
					})
					if err != nil {
						logs.Error("Update: Couldn't update sub category:", err)
						return
					}
				}
				// All this slug to list
				availableCategories[subCat.Slug] = true
			}

			// Get all sub categories in database
			subCategories := subCollection.Documents(ctx)

			// Loop all categories in database to remove categories unnecessary
			for {
				subCatSnapshot, err := subCategories.Next()
				if err == iterator.Done {
					break
				}
				if err != nil {
					logs.Error("Update: Loop sub categories error:", err)
					subCategories.Stop()
					return
				}
				var subCat models.SubCategory
				err = subCatSnapshot.DataTo(&subCat)
				if err != nil {
					logs.Error("Update: Cast sub category to struct failed:", err)
					subCategories.Stop()
					return
				}

				// Delete categories which not needed
				if !availableCategories[subCat.Slug] {
					_, err := subCatSnapshot.Ref.Delete(ctx)
					if err != nil {
						logs.Error("Update: Couldn't delete sub category:", err)
						return
					}
				}
			}
		}
	}
	res = "ok"
}

// @Title PushInterestCategoriesV2
// @Description Push interest categories to database
// @Success 200 {object} models.AuthenticationResponse
// @Failure 400 {object} models.ResponseWithError
// @router /interest_categories/v2 [put]
func (o *AdminController) PushInterestCategoriesV2() {
	var res models.Response
	defer func() {
		o.Data["json"] = res
		o.ServeJSON()
	}()

	ctx := databases.Context
	categoryCollection := models.GetCategoryCollection()
	for _, cat := range models.GetCategories() {
		var err error
		var catRef *firestore.DocumentRef
		catSnapshot, err := categoryCollection.Where("slug", "==", cat.Slug).Documents(ctx).Next()
		if err != nil {
			// If category doesn't exists, add it
			catRef = categoryCollection.Doc(cat.Slug)
			_, err = categoryCollection.Doc(cat.Slug).Set(ctx, cat.ToFirestoreCategory(nil))
			if err != nil {
				logs.Error("Add: Couldn't add new category:", err)
				return
			}
		} else {
			catRef = catSnapshot.Ref
		}

		for _, subCat := range cat.SubCategories {
			subId := cat.Slug + "-" + subCat.Slug
			subSlug := cat.Slug + "/" + subCat.Slug
			_, err := categoryCollection.Where("slug", "==", subSlug).Documents(ctx).Next()
			if err != nil {
				// Category doesn't exists in remote list, add it
				_, err := categoryCollection.Doc(subId).Set(ctx, subCat.ToFirestoreCategory(catRef))
				if err != nil {
					logs.Error("Add: Couldn't add new sub category:", err)
					return
				}
			}
		}
	}

	res = "ok"
}
