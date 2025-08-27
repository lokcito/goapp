package controllers

import (
    "net/http"
    "strconv"

    "github.com/gin-gonic/gin"
    "gorm.io/gorm"

    "equisd.com/bichito/robotapp/models"
)

type RobotController struct {
    DB *gorm.DB
}

func NewRobotController(db *gorm.DB) *RobotController {
    return &RobotController{DB: db}
}

// Index - list robots
func (rc *RobotController) Index(c *gin.Context) {
    var robots []models.Robot
    rc.DB.Order("created_at desc").Find(&robots)
    c.HTML(http.StatusOK, "index.html", gin.H{
        "robots": robots,
    })
}

func (rc *RobotController) New(c *gin.Context) {
    c.HTML(http.StatusOK, "new.html", gin.H{})
}

func (rc *RobotController) Create(c *gin.Context) {
    nombre := c.PostForm("nombre")
    descripcion := c.PostForm("descripcion")
    robot := models.Robot{Nombre: nombre, Descripcion: descripcion}
    if err := rc.DB.Create(&robot).Error; err != nil {
        c.HTML(http.StatusInternalServerError, "new.html", gin.H{
            "error": err.Error(),
        })
        return
    }
    c.Redirect(http.StatusFound, "/robots")
}

func (rc *RobotController) Show(c *gin.Context) {
    id, _ := strconv.Atoi(c.Param("id"))
    var robot models.Robot
    if err := rc.DB.First(&robot, id).Error; err != nil {
        c.String(http.StatusNotFound, "Not found")
        return
    }
    c.HTML(http.StatusOK, "show.html", gin.H{
        "robot": robot,
    })
}

func (rc *RobotController) Edit(c *gin.Context) {
    id, _ := strconv.Atoi(c.Param("id"))
    var robot models.Robot
    if err := rc.DB.First(&robot, id).Error; err != nil {
        c.String(http.StatusNotFound, "Not found")
        return
    }
    c.HTML(http.StatusOK, "edit.html", gin.H{
        "robot": robot,
    })
}

func (rc *RobotController) Update(c *gin.Context) {
    id, _ := strconv.Atoi(c.Param("id"))
    var robot models.Robot
    if err := rc.DB.First(&robot, id).Error; err != nil {
        c.String(http.StatusNotFound, "Not found")
        return
    }
    robot.Nombre = c.PostForm("nombre")
    robot.Descripcion = c.PostForm("descripcion")
    if err := rc.DB.Save(&robot).Error; err != nil {
        c.HTML(http.StatusInternalServerError, "edit.html", gin.H{
            "error": err.Error(),
            "robot": robot,
        })
        return
    }
    c.Redirect(http.StatusFound, "/robots")
}

func (rc *RobotController) Delete(c *gin.Context) {
    id, _ := strconv.Atoi(c.Param("id"))
    if err := rc.DB.Delete(&models.Robot{}, id).Error; err != nil {
        c.String(http.StatusInternalServerError, err.Error())
        return
    }
    c.Redirect(http.StatusFound, "/robots")
}
