package todo

import "github.com/graphql-go/graphql"

// used for Todo Schema
var todoSchema = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Todo",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.Int,
			},
			"title": &graphql.Field{
				Type: graphql.String,
			},
			"completed": &graphql.Field{
				Type: graphql.Boolean,
			},
		},
	},
)

// used for Queries
var todoQueries = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "TodoQueries",
		Fields: graphql.Fields{
			"todos": &graphql.Field{
				Type: graphql.NewList(todoSchema),
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return GetTodos(), nil
				},
			},
			"todo": &graphql.Field{
				Type: todoSchema,
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					id, ok := p.Args["id"].(int)
					if ok {
						return GetTodo(id), nil
					}
					return nil, nil
				},
			},
		},
	},
)

// used for Mutations
var todoMutations = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "TodoMutations",
		Fields: graphql.Fields{
			"addTodo": &graphql.Field{
				Type: todoSchema,
				Args: graphql.FieldConfigArgument{
					"title": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					title, ok := p.Args["title"].(string)
					if ok {
						return AddTodo(title), nil
					}
					return nil, nil
				},
			},
			"updateTodo": &graphql.Field{
				Type: graphql.Boolean,
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
					"title": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"completed": &graphql.ArgumentConfig{
						Type: graphql.Boolean,
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					id, ok := p.Args["id"].(int)
					title, ok2 := p.Args["title"].(string)
					completed, ok3 := p.Args["completed"].(bool)
					if ok && ok2 && ok3 {
						return UpdateTodo(id, title, completed), nil
					}
					return false, nil
				},
			},
			"deleteTodo": &graphql.Field{
				Type: graphql.Boolean,
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					id, ok := p.Args["id"].(int)
					if ok {
						return DeleteTodo(id), nil
					}
					return false, nil
				},
			},
		},
	},
)

// used for Root Schema
var TodoRootSchema, _ = graphql.NewSchema(
	graphql.SchemaConfig{
		Query:    todoQueries,
		Mutation: todoMutations,
	},
)
