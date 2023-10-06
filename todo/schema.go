package todo

import (
	"errors"
	"github.com/graphql-go/graphql"
)

// create variabel TodoRepository
var todoRepo = NewTodoRepository()

// create variabel dto response
var todoResponse = graphql.NewObject(graphql.ObjectConfig{
	Name: "todo",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Name: "id",
			Type: graphql.Int,
		},
		"title": &graphql.Field{
			Name: "title",
			Type: graphql.String,
		},
		"completed": &graphql.Field{
			Name: "completed",
			Type: graphql.Boolean,
		},
	},
})

// create query type
var todoQueryType = graphql.NewObject(graphql.ObjectConfig{
	Name: "todoQuery",
	Fields: graphql.Fields{
		"todos": &graphql.Field{
			Name:        "todos",
			Description: "get all todos data",
			Type:        graphql.NewList(todoResponse),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				// get all data from repository
				return todoRepo.GetAllTodos()
			},
		},
		"todo": &graphql.Field{
			Name:        "todo",
			Description: "get data todo by Id",
			Type:        todoResponse,
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.Int,
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				id, ok := p.Args["id"].(int) // ambil value dari argument request 'id', kemudian convert ke integer
				if !ok {                     // jika gagal diconvert ke integer
					return nil, errors.New("cant convert id to int") // tampilkan error gagal convert value id ke integer
				}

				return todoRepo.GetTodoById(id)
			},
		},
	},
})

// create mutation type
var todoMutationType = graphql.NewObject(graphql.ObjectConfig{
	Name: "todoMutation",
	Fields: graphql.Fields{
		"add_todo": &graphql.Field{
			Name:        "add_todo",
			Description: "insert new data todo",
			Type:        todoResponse,
			Args: graphql.FieldConfigArgument{
				"title": &graphql.ArgumentConfig{
					Description: "value property title",
					Type:        graphql.String,
				},
			},
		},
	},
})
