package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/trend-ai/TrendAI_mobile_backend/models"
	"gopkg.in/mgo.v2/bson"
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
		Name               string    `json:"name"`
		Gender             string    `json:"gender"`
		Dob                time.Time `json:"dob"`
		Education          string    `json:"education"`
		InterestCategories []string  `json:"interest_categories"`
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

	// Check if they want to update interest categories
	if _, exists := rawPacket["interest_categories"]; exists {
		categoryCollection := models.GetCategoryCollection()
		interestCategories := make([]bson.ObjectId, 0)
		for _, categoryId := range packet.InterestCategories {
			if !bson.IsObjectIdHex(categoryId) {
				o.Ctx.Output.SetStatus(http.StatusBadRequest)
				res = models.NewResponseWithError("category_not_found", "Categories doesn't exists")
				return
			}
			n, err := categoryCollection.FindId(bson.ObjectIdHex(categoryId)).Count()
			if err != nil {
				logs.Error("Find category by ID failed", err)
				o.Ctx.Output.SetStatus(http.StatusInternalServerError)
				res = models.NewResponseWithError("update_failed", "Couldn't update user")
				return
			}
			if n == 0 {
				o.Ctx.Output.SetStatus(http.StatusBadRequest)
				res = models.NewResponseWithError("category_not_found", "Categories doesn't exists")
				return
			}
			interestCategories = append(interestCategories, bson.ObjectIdHex(categoryId))
		}
		user.InterestCategories = interestCategories
	}

	// Save data
	userCollection := models.GetUserCollection()
	err := userCollection.UpdateId(user.Id, user)
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

	var categories []models.Category
	categoriesResponse := make([]models.CategoryResponse, 0)
	categoryCollection := models.GetCategoryCollection()

	// Get parent categories from database
	err := categoryCollection.Find(bson.M{"parent": nil}).All(&categories)
	if err != nil {
		logs.Error("Couldn't get categories")
		o.Ctx.Output.SetStatus(http.StatusInternalServerError)
		res = models.NewResponseWithError("get_failed", "Couldn't get categories")
		return
	}

	// Loop all snapshots to get data
	for _, category := range categories {
		catResponse, err := category.ToResponse()
		if err != nil {
			logs.Error("Couldn't get category response", err)
			o.Ctx.Output.SetStatus(http.StatusInternalServerError)
			res = models.NewResponseWithError("get_failed", "Couldn't get categories")
			return
		}
		categoriesResponse = append(categoriesResponse, *catResponse)
	}

	res = categoriesResponse
}
