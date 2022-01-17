package application

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

const (
	swaggerUrn = "/swagger/doc.json"

	HealthUrn = "/health"

	GroupAuthPath = "auth"

	AuthUrnNew     = "/new"
	AuthUrnLogin   = "/login"
	AuthUrnRefresh = "/refresh"
	AuthUrnLogout  = "/logout"

	GroupUsersPath = "users"
	UserUrn        = ""

	GroupTasksPath    = "tasks"
	TasksUrn          = ""
	TasksPreCreateUrn = "/preCreate"
	TasksObjectsUrn   = "/objects"
	TasksActionsUrn   = "/actions"
	TasksStatusUrn    = "/status"

	GroupKnowledgePath                   = "/knowledge"
	KnowledgeUrnInstruments              = "/instruments"
	KnowledgeUrnFloorTypes               = "/floorTypes"
	KnowledgeUrnRoomTypes                = "/roomTypes"
	KnowledgeUrnSpecializations          = "/specializations"
	KnowledgeUrnSkills                   = "/skills"
	KnowledgeUrnJobKinds                 = "/jobKinds"
	KnowledgeUrnWorkSections             = "/workSections"
	KnowledgeUrnWorkStages               = "/workStages"
	KnowledgeUrnActions                  = "/actions"
	KnowledgeUrnActionParams             = "/actionParams"
	KnowledgeUrnActionParamOptions       = "/actionParamOptions"
	KnowledgeUrnMaterialKinds            = "/materialKinds"
	KnowledgeUrnMaterialKindParams       = "/materialKindParams"
	KnowledgeUrnMaterialKindParamOptions = "/materialKindParamOptions"
	KnowledgeUrnMeasurementUnits         = "/measurementUnits"
	KnowledgeUrnTechnologies             = "/technologies"
	KnowledgeUrnOperations               = "/operations"
	KnowledgeUrnDevelopments             = "/developments"
	KnowledgeUrnFileNameTypes            = "/fileNameTypes"

	GroupObjectsPath              = "/objects"
	ObjectsUrn                    = ""
	ObjectsUrnGlobal              = "/global"
	ObjectsUrnPattern             = "/patterns"
	ObjectsUrnFloors              = "/floors"
	ObjectsUrnTasks               = "/tasks"
	ObjectsUrnMasters             = "/masters"
	ObjectsUrnGraph               = "/graph"
	ObjectsUrnGraphTechnologies   = "/graphTechnologies"
	ObjectsUrnGraphHasGraph       = "/hasGraph"
	ObjectsUrnMaterials           = "/materials"
	ObjectsUrnWorkTypes           = "/workTypes"
	ObjectsUrnEstimates           = "/estimates"
	ObjectsUrnEstimatesApprove    = "/approve-estimates"
	ObjectsUrnEstimatesDisapprove = "/disapprove-estimates"
	ObjectsUrnQCEs                = "/qce"

	GroupFilesPath = "/files"
	FilesUrnInfo   = "/info"
)

func (a *Application) buildPublicApi() API {
	router := gin.New()
	router.Use(a.MiddlewareResponse)
	router.HandleMethodNotAllowed = true
	router.Use(middlewareCORS())
	router.NoRoute(a.HandlerNotFound)
	router.NoMethod(a.HandlerMethodNotAllowed)
	router.MaxMultipartMemory = a.cfg.MaxMultipartMemory

	router.GET("/", a.HandlerHealth)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, ginSwagger.URL(swaggerUrn)))
	router.GET(HealthUrn, a.HandlerHealth)
	api := router.Group("/api")
	{
		v1 := api.Group("/v1")
		v1.Use(a.MiddlewareAccess)
		{

			//auth := v1.Group(GroupAuthPath)
			//{
			//	auth.POST(AuthUrnNew, a.HandlerNewAuth)
			//	auth.POST(AuthUrnLogin, a.HandlerLogin)
			//	auth.POST(AuthUrnRefresh, a.HandlerRefresh)
			//	auth.Use(a.MiddlewareAuth)
			//	{
			//		auth.POST(AuthUrnLogout, a.HandlerLogout)
			//	}
			//}
			//users := v1.Group(GroupUsersPath)
			//{
			//	users.Use(a.MiddlewareAuth)
			//	{
			//		users.GET(UserUrn, a.HandlerUser)
			//		users.POST(UserUrn, a.HandlerSetUser)
			//	}
			//}
			//knowledge := v1.Group(GroupKnowledgePath)
			//{
			//	knowledge.Use(a.MiddlewareAuth)
			//	{
			//		knowledge.GET(KnowledgeUrnMeasurementUnits, a.HandlerMeasurementUnits)
			//		knowledge.GET(KnowledgeUrnFloorTypes, a.HandlerFloorTypes)
			//		knowledge.POST(KnowledgeUrnFloorTypes, a.HandlerSetFloorTypes)
			//		knowledge.GET(KnowledgeUrnRoomTypes, a.HandlerRoomTypes)
			//		knowledge.POST(KnowledgeUrnRoomTypes, a.HandlerSetRoomTypes)
			//		knowledge.GET(KnowledgeUrnSpecializations, a.HandlerSpecializations)
			//		knowledge.POST(KnowledgeUrnSpecializations, a.HandlerSetSpecializations)
			//		knowledge.GET(KnowledgeUrnSkills, a.HandlerSkills)
			//		knowledge.POST(KnowledgeUrnSkills, a.HandlerSetSkills)
			//		knowledge.GET(KnowledgeUrnJobKinds, a.HandlerJobKinds)
			//		knowledge.POST(KnowledgeUrnJobKinds, a.HandlerSetJobKinds)
			//		knowledge.GET(KnowledgeUrnWorkStages, a.HandlerWorkStages)
			//		knowledge.POST(KnowledgeUrnWorkStages, a.HandlerSetWorkStages)
			//		knowledge.GET(KnowledgeUrnWorkSections, a.HandlerWorkSections)
			//		knowledge.POST(KnowledgeUrnWorkSections, a.HandlerSetWorkSections)
			//		knowledge.GET(KnowledgeUrnActions, a.HandlerActions)
			//		knowledge.POST(KnowledgeUrnActions, a.HandlerSetActions)
			//		knowledge.GET(KnowledgeUrnActionParams, a.HandlerActionParams)
			//		knowledge.POST(KnowledgeUrnActionParams, a.HandlerSetActionParams)
			//		knowledge.GET(KnowledgeUrnActionParamOptions, a.HandlerActionParamOptions)
			//		knowledge.POST(KnowledgeUrnActionParamOptions, a.HandlerSetActionParamOptions)
			//		knowledge.GET(KnowledgeUrnMaterialKinds, a.HandlerMaterialKinds)
			//		knowledge.POST(KnowledgeUrnMaterialKinds, a.HandlerSetMaterialKinds)
			//		knowledge.GET(KnowledgeUrnMaterialKindParams, a.HandlerMaterialKindParams)
			//		knowledge.POST(KnowledgeUrnMaterialKindParams, a.HandlerSetMaterialKindParams)
			//		knowledge.GET(KnowledgeUrnMaterialKindParamOptions, a.HandlerMaterialKindParamOptions)
			//		knowledge.POST(KnowledgeUrnMaterialKindParamOptions, a.HandlerSetMaterialKindParamOptions)
			//		knowledge.GET(KnowledgeUrnInstruments, a.HandlerInstruments)
			//		knowledge.POST(KnowledgeUrnInstruments, a.HandlerSetInstruments)
			//		knowledge.GET(KnowledgeUrnTechnologies, a.HandlerTechnologies)
			//		knowledge.GET(fmt.Sprintf("%s/:%s", KnowledgeUrnTechnologies, DTO.PathParamUUID), a.HandlerTechnologies)
			//		knowledge.POST(KnowledgeUrnTechnologies, a.HandlerSetTechnologies)
			//		knowledge.DELETE(KnowledgeUrnTechnologies, a.HandlerDeleteTechnologies)
			//		knowledge.GET(KnowledgeUrnOperations, a.HandlerOperations)
			//		knowledge.GET(fmt.Sprintf("%s/:%s", KnowledgeUrnOperations, DTO.PathParamUUID), a.HandlerOperations)
			//		knowledge.POST(KnowledgeUrnOperations, a.HandlerSetOperations)
			//		knowledge.DELETE(KnowledgeUrnOperations, a.HandlerDeleteOperations)
			//		knowledge.POST(KnowledgeUrnFileNameTypes, a.HandlerSetFileNameTypes)
			//		knowledge.GET(KnowledgeUrnFileNameTypes, a.HandlerFileNameTypes)
			//
			//		knowledge.GET(KnowledgeUrnDevelopments, a.HandlerDevelopments)
			//	}
			//}
			//objects := v1.Group(GroupObjectsPath)
			//{
			//	objects.Use(a.MiddlewareAuth)
			//	{
			//		// Global objects
			//		objects.GET(ObjectsUrnGlobal, a.HandlerGlobalObjects)
			//		objects.GET(fmt.Sprintf("%s/:%s", ObjectsUrnGlobal, DTO.PathParamUUID), a.HandlerGlobalObjects)
			//		objects.POST(ObjectsUrnGlobal, a.HandlerSetGlobalObjects)
			//		objects.DELETE(ObjectsUrnGlobal, a.HandlerDeleteGlobalObject)
			//		objects.DELETE(ObjectsUrn, a.HandlerDeleteObject)
			//
			//		// Hourly Rate
			//		objects.GET(fmt.Sprintf("%s/:%s", ObjectsUrnGlobal, DTO.PathParamUUID)+"/hourly-rate", a.HandlerGetHourlyRate)
			//
			//		// Patterns
			//		objects.GET(fmt.Sprintf("%s/:%s", ObjectsUrnGlobal, DTO.PathParamUUID)+ObjectsUrnPattern, a.HandlerPatternsForGlobalObject)
			//
			//		objects.GET(ObjectsUrnPattern, a.HandlerPatternObjects)
			//		objects.GET(fmt.Sprintf("%s/:%s", ObjectsUrnPattern, DTO.PathParamUUID), a.HandlerPatternObjects)
			//		objects.POST(ObjectsUrnPattern, a.HandlerSetPatternObjects)
			//		objects.DELETE(ObjectsUrnPattern, a.HandlerDeletePatternObject)
			//
			//		// Floors
			//		objects.GET(fmt.Sprintf("%s/:%s", ObjectsUrnGlobal, DTO.PathParamUUID)+ObjectsUrnFloors, a.HandlerGetFloors)
			//
			//		objects.POST(ObjectsUrnFloors, a.HandlerSetFloors)
			//		objects.DELETE(ObjectsUrnFloors, a.HandlerDeleteFloors)
			//
			//		// Tasks
			//		objects.GET(fmt.Sprintf("%s/:%s", ObjectsUrnGlobal, DTO.PathParamUUID)+ObjectsUrnTasks, a.HandlerTasksList)
			//
			//		// Masters
			//		objects.GET(fmt.Sprintf("%s/:%s", ObjectsUrnGlobal, DTO.PathParamUUID)+ObjectsUrnMasters,
			//			a.HandlerGetMasters)
			//
			//		// Graph
			//		objects.GET(fmt.Sprintf("%s/:%s", ObjectsUrnGlobal, DTO.PathParamUUID)+ObjectsUrnGraphTechnologies,
			//			a.HandlerGetGraphTechnologies)
			//
			//		objects.GET(fmt.Sprintf("%s/:%s", ObjectsUrnGlobal, DTO.PathParamUUID)+ObjectsUrnGraphHasGraph,
			//			a.HandlerHasGraph)
			//
			//		objects.POST(fmt.Sprintf("%s/:%s", ObjectsUrnGlobal, DTO.PathParamUUID)+ObjectsUrnGraph, a.HandlerSetGraph)
			//		objects.GET(fmt.Sprintf("%s/:%s", ObjectsUrnGlobal, DTO.PathParamUUID)+ObjectsUrnGraph, a.HandlerGetGraph)
			//
			//		// Materials
			//		objects.GET(fmt.Sprintf("%s/:%s", ObjectsUrnGlobal, DTO.PathParamUUID)+ObjectsUrnMaterials, a.HandlerMaterialsList)
			//		objects.POST(fmt.Sprintf("%s/:%s", ObjectsUrnGlobal, DTO.PathParamUUID)+ObjectsUrnMaterials, a.HandlerSetMaterials)
			//
			//		// WorkTypes
			//		objects.GET(fmt.Sprintf("%s/:%s", ObjectsUrnGlobal, DTO.PathParamUUID)+ObjectsUrnWorkTypes,
			//			a.HandlerGetWorkTypes)
			//
			//		// Estimates
			//		objects.GET(fmt.Sprintf("%s/:%s", ObjectsUrnGlobal, DTO.PathParamUUID)+ObjectsUrnEstimates+"/sections", a.HandlerGetEstimatesSections)
			//		objects.GET(fmt.Sprintf("%s/:%s", ObjectsUrnGlobal, DTO.PathParamUUID)+ObjectsUrnEstimates, a.HandlerGetEstimates)
			//		objects.POST(fmt.Sprintf("%s/:%s", ObjectsUrnGlobal, DTO.PathParamUUID)+ObjectsUrnEstimatesApprove, a.HandlerApproveEstimates)
			//		objects.POST(fmt.Sprintf("%s/:%s", ObjectsUrnGlobal, DTO.PathParamUUID)+ObjectsUrnEstimatesDisapprove, a.HandlerDisapproveEstimates)
			//
			//		//QCEs
			//		objects.GET(fmt.Sprintf("%s/:%s", ObjectsUrnGlobal, DTO.PathParamUUID)+ObjectsUrnQCEs, a.HandlerSearchQCE)
			//	}
			//}
			//files := v1.Group(GroupFilesPath)
			//{
			//	files.Use(a.MiddlewareAuth)
			//	{
			//		files.POST("", a.HandlerUploadFile)
			//		files.GET("/:"+DTO.PathParamUUID, a.HandlerFile)
			//		files.GET(fmt.Sprintf("/:%s", DTO.PathParamUUID)+FilesUrnInfo, a.HandlerFileInfo)
			//	}
			//}
			//tasks := v1.Group(GroupTasksPath)
			//{
			//	tasks.Use(a.MiddlewareAuth)
			//	{
			//		tasks.POST(TasksPreCreateUrn, a.HandlerPreCreateTask)
			//		tasks.POST(TasksUrn, a.HandlerSetTask)
			//		tasks.GET(fmt.Sprintf("%s/:%s", TasksUrn, DTO.PathParamUUID), a.HandlerTasks)
			//		tasks.GET(TasksObjectsUrn, a.HandlerTasksObjects)
			//		tasks.GET(TasksActionsUrn, a.HandlerTaskActions)
			//		tasks.POST(TasksActionsUrn, a.HandlerApplyTaskActions)
			//	}
			//}
		}
	}

	return router
}
