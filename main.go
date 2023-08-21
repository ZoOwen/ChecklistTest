package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/zoowen/postTsk/auth"
	"github.com/zoowen/postTsk/checklist"
	"github.com/zoowen/postTsk/checklistItem"
	"github.com/zoowen/postTsk/handler"
	"github.com/zoowen/postTsk/helper"
	"github.com/zoowen/postTsk/user"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:@tcp(127.0.0.1:3306)/db_checklist_test?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)

	}
	fmt.Println("success connect to database")

	//auth
	authService := auth.NewService()
	//user
	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository)
	userHandler := handler.NewUserHandler(userService, authService)

	//checklist
	checklistRepository := checklist.NewRepository(db)
	checklistService := checklist.NewService(checklistRepository)
	checklistHandler := handler.NewChecklistHandler(checklistService)

	//checklistItem
	checklistItemRepository := checklistItem.NewRepository(db)
	checklistItemService := checklistItem.NewService(checklistItemRepository)
	checklistItemHandler := handler.NewChecklistItemHandler(checklistItemService)

	//endpoint
	router := gin.Default()
	// router.Use(cors.Default())
	api := router.Group("/api")

	api.POST("/register", userHandler.RegisterUser)
	api.POST("/login", userHandler.Login)

	//checklist
	api.GET("/checklists", authMiddleware(authService, userService), checklistHandler.GetChecklists)
	api.POST("/checklist", authMiddleware(authService, userService), checklistHandler.CreateChecklist)
	api.DELETE("/checklist/:id", authMiddleware(authService, userService), checklistHandler.DeleteChecklist)

	//checklistItem
	api.GET("/checklist/:checklistId/item", authMiddleware(authService, userService), checklistItemHandler.GetChecklistsItem)
	api.POST("/checklist/:checklistId/item", authMiddleware(authService, userService), checklistItemHandler.CreateChecklistItem)
	// api.DELETE("/checklist/:id", authMiddleware(authService, userService), checklistItemHandler.DeleteChecklist)

	router.Run()
}
func authMiddleware(authService auth.Service, userService user.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if !strings.Contains(authHeader, "Bearer") {
			response := helper.ApiResponse("UnAuthorized", http.StatusUnauthorized, "Error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		//Bearer token
		tokenString := ""
		arrayToken := strings.Split(authHeader, " ")

		if len(arrayToken) == 2 {
			tokenString = arrayToken[1]
		}

		token, err := authService.ValidateToken(tokenString)
		if err != nil {
			response := helper.ApiResponse("UnAuthorized", http.StatusUnauthorized, "Error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		claim, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			response := helper.ApiResponse("UnAuthorized", http.StatusUnauthorized, "Error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}
		userID := int(claim["user_id"].(float64))

		user, err := userService.GetUserByID(userID)

		if err != nil {
			response := helper.ApiResponse("UnAuthorized", http.StatusUnauthorized, "Error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}
		c.Set("currentUser", user)
	}
}
