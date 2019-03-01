package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/trend-ai/TrendAI_mobile_backend/models"
	"github.com/trend-ai/TrendAI_mobile_backend/services/databases"
	"net/http"
	"time"
)

// Operations about Users
type UserController struct {
	beego.Controller
}

// @Title Get
// @Description Get current user
// @Success 200 {object} models.UserResponse
// @Failure 401 {object} models.ResponseWithError
// @router / [get]
func (o *UserController) Get() {
	var res models.Response
	defer func() {
		o.Data["json"] = res
		o.ServeJSON()
	}()
	user := o.Ctx.Input.GetData("user").(models.User)
	res = user.ToResponse()
}

// @Title Patch
// @Description Update specific fields for current user
// @Success 200 {object} models.UserResponse
// @Failure 400 {object} models.ResponseWithError
// @Failure 500 {object} models.ResponseWithError
// @router / [patch]
func (o *UserController) Patch() {
	var res models.Response
	defer func() {
		o.Data["json"] = res
		o.ServeJSON()
	}()

	var packet struct {
		Name      string    `json:"name"`
		Gender    string    `json:"gender"`
		Dob       time.Time `json:"dob"`
		Education string    `json:"education"`
	}

	var rawPacket map[string]interface{}

	// Cast request to packet
	errPacket := json.Unmarshal(o.Ctx.Input.RequestBody, &packet)

	// Cast request to raw packet
	errRawPacket := json.Unmarshal(o.Ctx.Input.RequestBody, &rawPacket)

	// Check if casting error
	if errPacket != nil || errRawPacket != nil {
		logs.Error("Parse request error", errPacket, errRawPacket)
		o.Ctx.Output.SetStatus(http.StatusBadRequest)
		res = models.NewResponseWithError("request_invalid", "Couldn't parse your request")
		return
	}

	// Get current user from request
	user := o.Ctx.Input.GetData("user").(models.User)

	// Check if they want to update gender
	if _, exists := rawPacket["name"]; exists && len(packet.Name) > 0 {
		user.Name = packet.Name
	}

	// Check if they want to update gender
	if _, exists := rawPacket["gender"]; exists {
		if gender, err := models.UserGenderToInt(packet.Gender); err == nil {
			user.Gender = gender
		}
	}

	// Check if they want to update dob
	if _, exists := rawPacket["dob"]; exists && packet.Dob.Before(time.Now().UTC()) {
		user.Dob = packet.Dob
	}

	// Check if they want to update education
	if _, exists := rawPacket["education"]; exists {
		user.Education = packet.Education
	}

	ctx := databases.Context
	// Save data
	userCollection := models.GetUserCollection()
	userRef := userCollection.Doc(user.Id)
	_, err := userRef.Set(ctx, user)
	if err != nil {
		logs.Error("Couldn't update user:", err)
		o.Ctx.Output.SetStatus(http.StatusInternalServerError)
		res = models.NewResponseWithError("update_failed", "Couldn't update user")
		return
	}

	res = user.ToResponse()
}

// @Title GetCategories
// @Description Get all categories available
// @Success 200 {object} []models.Category
// @Failure 500 {object} models.ResponseWithError
// @router /categories [get]
func (o *UserController) GetCategories() {
	var res models.Response
	defer func() {
		o.Data["json"] = res
		o.ServeJSON()
	}()

	var err error
	ctx := databases.Context
	categories := make([]models.Category, 0)
	categoryCollection := models.GetCategoryCollection()

	// Get parent categories from database
	catSnapshots, err := categoryCollection.Where("parent", "==", "").Documents(ctx).GetAll()
	if err != nil {
		logs.Error("Couldn't get parent categories", err)
		o.Ctx.Output.SetStatus(http.StatusInternalServerError)
		res = models.NewResponseWithError("get_failed", "Couldn't get categories")
		return
	}

	logs.Info("catSnapshots", catSnapshots)

	//// Loop all snapshots to get data
	//for _, catSnapshot := range catSnapshots {
	//	var catData models.FirestoreCategory
	//	err = catSnapshot.DataTo(&catData)
	//	if err != nil {
	//		logs.Error("Couldn't convert category data to struct", err)
	//		o.Ctx.Output.SetStatus(http.StatusInternalServerError)
	//		res = models.NewResponseWithError("get_failed", "Couldn't get categories")
	//		return
	//	}
	//	child := make([]models.SubCategory, 0)
	//	for _, subRef := range catData.Child {
	//		subSnap, err := subRef.Get(ctx)
	//		if err != nil {
	//			logs.Error("Couldn't get sub-category", err)
	//			o.Ctx.Output.SetStatus(http.StatusInternalServerError)
	//			res = models.NewResponseWithError("get_failed", "Couldn't get categories")
	//			return
	//		}
	//		var subData models.FirestoreCategory
	//		err = subSnap.DataTo(&subData)
	//		if err != nil {
	//			logs.Error("Couldn't convert sub-category to struct", err)
	//			o.Ctx.Output.SetStatus(http.StatusInternalServerError)
	//			res = models.NewResponseWithError("get_failed", "Couldn't get categories")
	//			return
	//		}
	//		child = append(child, models.SubCategory{
	//			Id: subRef.ID,
	//			Name: subData.Name,
	//			Slug: subData.Slug,
	//		})
	//	}
	//	categories = append(categories, models.Category{
	//		BaseCategory: catData.BaseCategory,
	//		Child: child,
	//	})
	//}

	res = categories
}
