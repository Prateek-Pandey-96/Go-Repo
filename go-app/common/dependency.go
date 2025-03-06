package common

import "github.com/gin-gonic/gin"

// any dependency common across application like db_connection, redis_connection, logger etx
type Dependency struct {
	Router *gin.Engine
}
