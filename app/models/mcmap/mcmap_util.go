package mcmap

import (
    "gohub/pkg/app"
    "gohub/pkg/database"
    "gohub/pkg/paginator"

    "github.com/gin-gonic/gin"
)

func Get(idstr string) (mcmap Mcmap) {
    database.DB.Where("id", idstr).First(&mcmap)
    return
}

func GetBy(field, value string) (mcmap Mcmap) {
    database.DB.Where("? = ?", field, value).First(&mcmap)
    return
}

func All() (mcmaps []Mcmap) {
    database.DB.Find(&mcmaps)
    return 
}

func IsExist(field, value string) bool {
    var count int64
    database.DB.Model(Mcmap{}).Where("? = ?", field, value).Count(&count)
    return count > 0
}

func Paginate(c *gin.Context, perPage int) (mcmaps []Mcmap, paging paginator.Paging) {
    paging = paginator.Paginate(
        c,
        database.DB.Model(Mcmap{}),
        &mcmaps,
        app.V1URL(database.TableName(&Mcmap{})),
        perPage,
    )
    return
}