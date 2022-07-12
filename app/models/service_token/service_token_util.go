package service_token

import (
    "gohub/pkg/app"
    "gohub/pkg/database"
    "gohub/pkg/paginator"

    "github.com/gin-gonic/gin"
)

func Get(idstr string) (serviceToken ServiceToken) {
    database.DB.Where("id", idstr).First(&serviceToken)
    return
}

func GetBy(field, value string) (serviceToken ServiceToken) {
    database.DB.Where("? = ?", field, value).First(&serviceToken)
    return
}

func All() (serviceTokens []ServiceToken) {
    database.DB.Find(&serviceTokens)
    return 
}

func IsExist(field, value string) bool {
    var count int64
    database.DB.Model(ServiceToken{}).Where("? = ?", field, value).Count(&count)
    return count > 0
}

func Paginate(c *gin.Context, perPage int) (serviceTokens []ServiceToken, paging paginator.Paging) {
    paging = paginator.Paginate(
        c,
        database.DB.Model(ServiceToken{}),
        &serviceTokens,
        app.V1URL(database.TableName(&ServiceToken{})),
        perPage,
    )
    return
}