// GENERATED CODE - DO NOT EDIT
package routes

import "github.com/revel/revel"


type tGormController struct {}
var GormController tGormController


func (_ tGormController) SetUser(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("GormController.SetUser", args).Url
}

func (_ tGormController) Begin(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("GormController.Begin", args).Url
}

func (_ tGormController) Commit(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("GormController.Commit", args).Url
}

func (_ tGormController) Rollback(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("GormController.Rollback", args).Url
}


type tTestRunner struct {}
var TestRunner tTestRunner


func (_ tTestRunner) Index(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("TestRunner.Index", args).Url
}

func (_ tTestRunner) Run(
		suite string,
		test string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "suite", suite)
	revel.Unbind(args, "test", test)
	return revel.MainRouter.Reverse("TestRunner.Run", args).Url
}

func (_ tTestRunner) List(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("TestRunner.List", args).Url
}


type tJobs struct {}
var Jobs tJobs


func (_ tJobs) Status(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Jobs.Status", args).Url
}


type tStatic struct {}
var Static tStatic


func (_ tStatic) Serve(
		prefix string,
		filepath string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "prefix", prefix)
	revel.Unbind(args, "filepath", filepath)
	return revel.MainRouter.Reverse("Static.Serve", args).Url
}

func (_ tStatic) ServeModule(
		moduleName string,
		prefix string,
		filepath string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "moduleName", moduleName)
	revel.Unbind(args, "prefix", prefix)
	revel.Unbind(args, "filepath", filepath)
	return revel.MainRouter.Reverse("Static.ServeModule", args).Url
}


type tSecuredController struct {}
var SecuredController tSecuredController


func (_ tSecuredController) Check(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("SecuredController.Check", args).Url
}


type tApp struct {}
var App tApp


func (_ tApp) SetHeaders(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("App.SetHeaders", args).Url
}

func (_ tApp) Index(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("App.Index", args).Url
}

func (_ tApp) Auth(
		code string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "code", code)
	return revel.MainRouter.Reverse("App.Auth", args).Url
}

func (_ tApp) Logout(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("App.Logout", args).Url
}

func (_ tApp) Map(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("App.Map", args).Url
}

func (_ tApp) Terms(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("App.Terms", args).Url
}

func (_ tApp) Privacy(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("App.Privacy", args).Url
}

func (_ tApp) Tour(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("App.Tour", args).Url
}

func (_ tApp) Show(
		uuid string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "uuid", uuid)
	return revel.MainRouter.Reverse("App.Show", args).Url
}

func (_ tApp) Approve(
		uuid string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "uuid", uuid)
	return revel.MainRouter.Reverse("App.Approve", args).Url
}


type tProfile struct {}
var Profile tProfile


func (_ tProfile) Index(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Profile.Index", args).Url
}

func (_ tProfile) Save(
		user interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "user", user)
	return revel.MainRouter.Reverse("Profile.Save", args).Url
}


type tRequests struct {}
var Requests tRequests


func (_ tRequests) Index(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Requests.Index", args).Url
}

func (_ tRequests) Save(
		request interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "request", request)
	return revel.MainRouter.Reverse("Requests.Save", args).Url
}

func (_ tRequests) Recaptcha(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Requests.Recaptcha", args).Url
}


