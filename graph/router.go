package graph

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/graphql-go/graphql"
	"github.com/rahul-yr/learn-go-grapql/todo"
)

// var dummySchema, _ = graphql.NewSchema(
// 	graphql.SchemaConfig{
// 		Query: graphql.NewObject(graphql.ObjectConfig{
// 			Name: "RootQuery",
// 		}),
// 		Mutation: graphql.NewObject(graphql.ObjectConfig{}),
// 	},
// )

type RequestParams struct {
	Query     string                 `json:"query"`
	Operation string                 `json:"operation"`
	Variables map[string]interface{} `json:"variables"`
}

func TodoGraphRouter(c *gin.Context) {
	var reqObj RequestParams
	if err := c.ShouldBindJSON(&reqObj); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// perform cleanup and custom validations
	// check authentication
	// check authorization
	// apply rate limiting
	// apply tracing
	// apply metrics
	// apply logging
	// perform business logic
	result := graphql.Do(graphql.Params{
		Context:        c,
		Schema:         todo.TodoRootSchema,
		RequestString:  reqObj.Query,
		VariableValues: reqObj.Variables,
		OperationName:  reqObj.Operation,
	})
	if len(result.Errors) > 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": result.Errors})
		return
	} else {
		c.JSON(http.StatusOK, result)
	}
}
