package api

import (
	"net/http"

	_ "crmapi/api/docs" //for swagger
	"crmapi/api/handler"
	"crmapi/config"
	"crmapi/pkg/grpc_client"
	"crmapi/pkg/logger"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// Config ...
type Config struct {
	Logger     logger.Logger
	GrpcClient *grpc_client.GrpcClient
	Cfg        config.Config
}

// @title           CRM API
// @version         1.0
// @description     This is a sample server celler server.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8082
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/

func New(cnf Config) *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())

	r.Use(gin.Recovery())

	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowHeaders = append(config.AllowHeaders, "*")

	r.Use(cors.New(config))

	handler := handler.New(&handler.HandlerConfig{
		Logger:     cnf.Logger,
		GrpcClient: cnf.GrpcClient,
		Cfg:        cnf.Cfg,
	})

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "Api gateway"})
	})

	// super-admins
	r.POST("/api/v1/super-admin", handler.CreateSuperAdmin)
	r.GET("/api/v1/super-admin/:id", handler.GetByIdSuperAdmin)
	r.PUT("/api/v1/super-admin", handler.UpdateSuperAdmin)
	r.DELETE("/api/v1/super-admin/:id", handler.DeleteSuperAdmin)
	r.GET("/api/v1/super-admins", handler.GetAllSuperAdmin)

	// branches
	r.POST("/api/v1/branch", handler.CreateBranch)
	r.GET("/api/v1/branch/:id", handler.GetByIdBranch)
	r.PUT("/api/v1/branch", handler.UpdateBranch)
	r.DELETE("/api/v1/branch/:id", handler.DeleteBranch)
	r.GET("/api/v1/branches", handler.GetAllBranches)

	// admins
	r.POST("/api/v1/admin", handler.CreateAdmin)
	r.GET("/api/v1/admin/:id", handler.GetByIdAdmin)
	r.PUT("/api/v1/admin", handler.UpdateAdmin)
	r.DELETE("/api/v1/admin/:id", handler.DeleteAdmin)
	r.GET("/api/v1/admins", handler.GetAllAdmins)


	// events
	r.POST("/api/v1/event", handler.CreateEvent)
	r.GET("/api/v1/event/:id", handler.GetByIdEvent)
	r.PUT("/api/v1/event", handler.UpdateEvent)
	r.DELETE("/api/v1/event/:id", handler.DeleteEvent)
	r.GET("/api/v1/events", handler.GetAllEvents)

	// groups
	r.POST("/api/v1/group", handler.CreateGroup)
	r.GET("/api/v1/group/:id", handler.GetByIdGroup)
	r.PUT("/api/v1/group", handler.UpdateGroup)
	r.DELETE("/api/v1/group/:id", handler.DeleteGroup)
	r.GET("/api/v1/groups", handler.GetAllGroups)

	// join-event
	r.POST("/api/v1/join-event", handler.CreateJoinEvent)
	r.GET("/api/v1/join-event/:id", handler.GetByIdEventJoin)
	r.DELETE("/api/v1/join-event/:id", handler.DeleteEventJoin)

	// managers
	r.POST("/api/v1/manager", handler.CreateManager)
	r.GET("/api/v1/manager/:id", handler.GetByIdManager)
	r.PUT("/api/v1/manager", handler.UpdateManager)
	r.DELETE("/api/v1/manager/:id", handler.DeleteManager)
	r.GET("/api/v1/managers", handler.GetAllManagers)

	// students
	r.POST("/api/v1/student", handler.CreateStudent)
	r.GET("/api/v1/student/:id", handler.GetByIdStudent)
	r.PUT("/api/v1/student", handler.UpdateStudent)
	r.DELETE("/api/v1/student/:id", handler.DeleteStudent)
	r.GET("/api/v1/students", handler.GetAllStudents)

	// support-teachers
	r.POST("/api/v1/support-teacher", handler.CreateSupportTeacher)
	r.GET("/api/v1/support-teacher/:id", handler.GetByIdSupportTeacher)
	r.PUT("/api/v1/support-teacher", handler.UpdateSupportTeacher)
	r.DELETE("/api/v1/support-teacher/:id", handler.DeleteSupportTeacher)
	r.GET("/api/v1/support-teachers", handler.GetAllSupportTeacher)

	// teachers
	r.POST("/api/v1/teacher", handler.CreateTeacher)
	r.GET("/api/v1/teacher/:id", handler.GetByIdTeacher)
	r.PUT("/api/v1/teacher", handler.UpdateTeacher)
	r.DELETE("/api/v1/teacher/:id", handler.DeleteTeacher)
	r.GET("/api/v1/teachers", handler.GetAllTeachers)

	// schedule
	r.POST("/api/v1/schedule", handler.CreateSchedule)
	r.GET("/api/v1/schedule/:id", handler.GetByIdSchedule)
	r.PUT("/api/v1/schedule", handler.UpdateSchedule)
	r.DELETE("/api/v1/schedule/:id", handler.DeleteSchedule)
	r.GET("/api/v1/schedules", handler.GetAllSchedules)

	// tasks
	r.POST("/api/v1/task", handler.CreateTask)
	r.GET("/api/v1/task/:id", handler.GetByIdTask)
	r.PUT("/api/v1/task", handler.UpdateTask)
	r.DELETE("/api/v1/task/:id", handler.DeleteTask)
	r.GET("/api/v1/tasks", handler.GetAllTasks)

	// lessons
	r.POST("/api/v1/lesson", handler.CreateLesson)
	r.GET("/api/v1/lesson/:id", handler.GetByIdLesson)
	r.PUT("/api/v1/lesson", handler.UpdateLesson)
	r.GET("/api/v1/lessons", handler.GetAllLessons)

	// attendances
	r.POST("/api/v1/attendance", handler.CreateAttendance)
	r.GET("/api/v1/attendance/:id", handler.GetByIdAttendance)
	r.PUT("/api/v1/attendance", handler.UpdateAttendance)
	r.DELETE("/api/v1/attendance/:id", handler.DeleteAttendance)
	r.GET("/api/v1/attendances", handler.GetAllAttendances)


	// Shipper endpoints
	url := ginSwagger.URL("swagger/doc.json") // The url pointing to API definition
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	return r
}
