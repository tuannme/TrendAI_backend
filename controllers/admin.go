package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/trend-ai/TrendAI_mobile_backend/models"
	"gopkg.in/mgo.v2/bson"
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

	// Get category collection
	categoryCollection := models.GetCategoryCollection()

	// Loop all raw categories in our model to update to database
	for _, rawCategory := range models.GetRawCategories() {
		var category models.Category

		// Get category by slug
		err := categoryCollection.Find(bson.M{"slug": category.Slug}).One(&category)
		if err != nil {
			// If category doesn't exists, add it
			category = rawCategory.ToCategory()
			category.Id = bson.NewObjectId()
			err := categoryCollection.Insert(category)
			if err != nil {
				logs.Error("Add: Couldn't add new category:", err)
				return
			}
		}

		// Loop sub categories and add or update it
		var child []bson.ObjectId
		for _, rawSubCat := range rawCategory.Child {
			var subCategory models.Category
			// Generate new sub-category's slug by parent slug
			subSlug := category.Slug + "/" + rawSubCat.Slug
			// Get sub category by slug
			err := categoryCollection.Find(bson.M{"slug": subSlug}).One(&subCategory)
			if err != nil {
				// If sub category doesn't exists, add it
				subCategory = rawSubCat.ToCategory()
				subCategory.Id = bson.NewObjectId()
				subCategory.Parent = category.Id
				subCategory.Slug = subSlug
				err := categoryCollection.Insert(subCategory)
				if err != nil {
					logs.Error("Add: Couldn't add new sub-category:", err)
					return
				}
			}
			// Add sub category to child list
			child = append(child, subCategory.Id)
		}

		// Update parent category's child list
		category.Child = child

		// Update parent category with new data
		err = categoryCollection.UpdateId(category.Id, category)
		if err != nil {
			logs.Error("Add: Couldn't save category:", err)
			return
		}
	}

	res = "ok"
}
