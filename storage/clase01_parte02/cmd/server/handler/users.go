package handler

import (
	"fmt"
	"goweb/internal/domain"
	"goweb/internal/users"
	"goweb/pkg/web"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type request struct {
	Id int 				`json:"id"`
	Name string			`json:"name"`
	LastName string		`json:"lastname"`			
	Email string		`json:"email"`
	Age int				`json:"age"`
	Height float32		`json:"height"`
	Active bool			`json:"active"`
	CreatedAt string	`json:"createdat"`
}

type User struct {
	service users.Service
}

func NewUser(u users.Service) *User{
	return &User{
		service: u,
	}
}

// GetAllUsers List
// @Summary List of all platform users
// @Tags Users
// @Description get all platform users that exist on the platform
// @Accept  json
// @Produce  json
// @Param token header string true "token"
// @Success 200 {object} web.Response
// @Failure 401 {object} web.Response
// @Failure 404 {object} web.Response
// @Router /users [get]
func (c *User) GetAllUsers() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// valido token
		/* token := ctx.GetHeader("token")
		if token != os.Getenv("TOKEN") {
			ctx.JSON(http.StatusUnauthorized, web.NewResponse(http.StatusUnauthorized, nil, "Token inválido"))
			return
		} */
		u, err := c.service.GetAllUsers(ctx)
		if err != nil {
			ctx.JSON(404, web.NewResponse(404, nil, err.Error()))
			return
		}
		ctx.JSON(200, web.NewResponse(200, u, ""))
	}
}

// GetUserById
// @Summary Get one user by Id
// @Tags Users
// @Description Get one user by Id
// @Accept  json
// @Produce  json
// @Param token header string true "token"
// @Param id path int true "id"
// @Success 200 {object} web.Response
// @Failure 401 {object} web.Response
// @Failure 404 {object} web.Response
// @Router /users/{id} [get]
func (c *User) GetUserById() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// valido token
		/* token := ctx.GetHeader("token")
		if token != os.Getenv("TOKEN") {
			ctx.JSON(http.StatusUnauthorized, web.NewResponse(http.StatusUnauthorized, nil, "Token inválido"))
			return
		} */
		id,_ := strconv.Atoi(ctx.Param("id"))
		u, err := c.service.GetUserById(ctx, id)
		if err != nil {
			ctx.JSON(404, web.NewResponse(404, nil, err.Error()))
			return
		}
		ctx.JSON(200, web.NewResponse(200, u, ""))
	}
}

func (c *User) GetUserByName() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// valido token
		/* token := ctx.GetHeader("token")
		if token != os.Getenv("TOKEN") {
			ctx.JSON(http.StatusUnauthorized, web.NewResponse(http.StatusUnauthorized, nil, "Token inválido"))
			return
		} */

		name := ctx.Param("name")
		u, err := c.service.GetUserByName(ctx, name)
		if err != nil {
			ctx.JSON(404, web.NewResponse(404, nil, err.Error()))
			return
		}
		ctx.JSON(200, web.NewResponse(200, u, ""))
	}
}


// StoreUsers
// @Summary Store new users in the database
// @Tags Users
// @Description store users
// @Accept  json
// @Produce  json
// @Param token header string true "token"
// @Param user body request true "User to store"
// @Success 200 {object} web.Response
// @Failure 400 {object} web.Response
// @Failure 401 {object} web.Response
// @Failure 404 {object} web.Response
// @Router /users [post]
func (c *User) StoreUser() gin.HandlerFunc {
	return func(ctx *gin.Context){
		// traigo los datos del post y los guardo en una variable del tipo struct request que generé arriba
		var userToStore domain.User
		if err := ctx.ShouldBindJSON(&userToStore); err !=nil{
			ctx.JSON(404, web.NewResponse(404, nil, err.Error()))
			return
		}

		// valido campos requeridos
		if v := validar(userToStore); v != "" {
			ctx.JSON(400, web.NewResponse(400, nil, v))
			return
		}

		newUser, err := c.service.StoreUser(ctx, userToStore.Name, userToStore.LastName, userToStore.Email, userToStore.Age, userToStore.Height, userToStore.Active)
		if err != nil{
			ctx.JSON(404, web.NewResponse(404, nil, err.Error()))
			return
		}
		ctx.JSON(200, web.NewResponse(200, newUser, ""))
	}
}

// UpdateUser
// @Summary Edit all the fields of an User by Id
// @Tags Users
// @Description You can change any user's information but it is necessary to complete all the fields
// @Accept  json
// @Produce  json
// @Param token header string true "token"
// @Param id path int true "id"
// @Param user body request true "User to update"
// @Success 200 {object} web.Response
// @Failure 400 {object} web.Response
// @Failure 401 {object} web.Response
// @Failure 404 {object} web.Response
// @Router /users/{id} [put]
func (c *User) UpdateTotal() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// valido token
		/* token := ctx.GetHeader("token")
		if token != os.Getenv("TOKEN") {
			ctx.JSON(http.StatusUnauthorized, web.NewResponse(http.StatusUnauthorized, nil, "Token inválido"))
			return
		}
		 */
		//var errores []string
		
		id,err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.JSON(400, web.NewResponse(400, nil, "invalid ID"))
 			return
		}

		// validaciones
		var userToUpdate domain.User
		if err:= ctx.ShouldBindJSON(&userToUpdate); err!=nil{
			ctx.JSON(404, web.NewResponse(404, nil, err.Error()))
			return
		}
		// valido campos requeridos
		if v := validar(userToUpdate); v != "" {
			ctx.JSON(400, web.NewResponse(400, nil, v))
			return
		}

		userUpdated, err:= c.service.UpdateTotal(ctx, id, userToUpdate.Name, userToUpdate.LastName, userToUpdate.Email, userToUpdate.Age, userToUpdate.Height, userToUpdate.Active, userToUpdate.CreatedAt)

		if err !=nil {
			ctx.JSON(404, web.NewResponse(404, nil, err.Error()))
			return
		}
		ctx.JSON(200, web.NewResponse(200, userUpdated, ""))
	}
}

// UpdateUser
// @Summary Edit user's age or lastname
// @Tags Users
// @Description You can change user's age or lastname
// @Accept  json
// @Produce  json
// @Param token header string true "token"
// @Param id path int true "id"
// @Param user body request true "User to update"
// @Success 200 {object} web.Response
// @Failure 400 {object} web.Response
// @Failure 401 {object} web.Response
// @Failure 404 {object} web.Response
// @Router /users/{id} [patch]
func (c *User) UpdatePartial() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// valido token
		token := ctx.GetHeader("token")
		if token != os.Getenv("TOKEN") {
			ctx.JSON(http.StatusUnauthorized, web.NewResponse(http.StatusUnauthorized, nil, "Token inválido"))
			return
		}

		var errores []string

		id,err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.JSON(400, web.NewResponse(400, nil, "invalid ID"))
 			return
		}

		// validaciones
		var req request
		if err:= ctx.ShouldBindJSON(&req); err!=nil{
			ctx.JSON(404, web.NewResponse(404, nil, err.Error()))
			return
		}
		if req.LastName == ""{
			errores = append(errores, "El apellido del usuario es requerido")
		}
		if req.Age == 0 {
			errores  =append(errores, "La edad del usuario es requerido")
		}
		if len(errores) > 0 {
			erroresStr := strings.Join(errores, ", ")
			ctx.JSON(400, web.NewResponse(400, nil, erroresStr))
			return
		}

		userUpdated, err:=	 c.service.UpdatePartial(id, req.LastName, req.Age)

		if err !=nil {
			ctx.JSON(404, web.NewResponse(404, nil, err.Error()))
			return
		}
		ctx.JSON(200, web.NewResponse(200, userUpdated, ""))
	}
}


// DeleteUser
// @Summary Delete any User by Id
// @Tags Users
// @Description You can delete any user in the database
// @Accept  json
// @Produce  json
// @Param token header string true "token"
// @Param id path int true "id"
// @Success 200 {object} web.Response
// @Failure 400 {object} web.Response
// @Failure 401 {object} web.Response
// @Failure 404 {object} web.Response
// @Router /users/{id} [delete]
func (c *User) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// valido token
		/* token := ctx.GetHeader("token")
		if token != os.Getenv("TOKEN") {
			ctx.JSON(http.StatusUnauthorized, web.NewResponse(http.StatusUnauthorized, nil, "Token inválido"))
			return
		} */
		
		id,err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.JSON(400, web.NewResponse(400, nil, "invalid ID"))
 			return
		}

		err=c.service.Delete(id)
		if err !=nil {
			ctx.JSON(404, web.NewResponse(404, nil, err.Error()))
			return
		}
		ctx.JSON(200, web.NewResponse(200, fmt.Sprintf("El producto %d ha sido eliminado",
		id), ""))
	}
}

func validar(us domain.User) string {
	var response []string
	if us.Name == ""{
		response = append(response, "Nombre") 
	}
	if us.LastName == ""{
		response = append(response, "Apellido")
	}
	if us.Email == ""{
		response = append(response, "Email")
	}
	if us.Age == 0 {
		response = append(response, "Edad")
	}
	if us.Height == 0 {
		response = append(response,"Altura")
	}
	if len(response) > 0 {
		errores := "Por favor completar los siguientes campos: "+strings.Join(response, ", ")
		return errores
	}
	return ""
}