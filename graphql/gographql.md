# Example Go Graphql
Here is simple Go graphql example. In this example, we simulate how to create graphql schema, then insert data into schema, and print values to stdout:

**Note:** 
- Assume you already have setup your Golang environment in linux machine, if not refer to [this](https://golang.org/doc/install) article for setup your Golang environment, example code are written using [Go v1.15.5](https://golang.org/dl/go1.15.5.linux-amd64.tar.gz).
- Put source code inside `$GOPATH/src/<yourAppFolder>`. In this example, we put source code in path `/home/ubuntu/go/src/mygraphql`.
- To get to know where is your GOPATH folder, type in terminal `go env GOPATH`.

```go
/* file: main.go */

package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/graphql-go/graphql"
)

type Tutorial struct {
	ID int
	Title    string
	Author   Author
	Comments []Comment
}

type Author struct {
	Name      string
	Tutorials []int
}

type Comment struct {
	Body string
}

var (

	// field object configuration for data struct Comment
	commentType = graphql.NewObject(
		graphql.ObjectConfig{
			Name: "Comment",
			// we define the name and the fields of our
			// object. In this case, we have one solitary
			// field that is of type string
			Fields: graphql.Fields{
				"body": &graphql.Field{
					Type: graphql.String,
				},
			},
		},
	)

	// field object configuration for data struct Author
	authorType = graphql.NewObject(
		graphql.ObjectConfig{
			Name: "Author",
			Fields: graphql.Fields{
				"Name": &graphql.Field{
					Type: graphql.String,
				},
				"Tutorials": &graphql.Field{
					// we'll use NewList to deal with an array
					// of int values
					Type: graphql.NewList(graphql.Int),
				},
			},
		},
	)

	// field object configuration for data struct Tutorial
	tutorialType = graphql.NewObject(
		graphql.ObjectConfig{
			Name: "Tutorial",
			Fields: graphql.Fields{
				"id": &graphql.Field{
					Type: graphql.Int,
				},
				"title": &graphql.Field{
					Type: graphql.String,
				},
				"author": &graphql.Field{
					// here, we specify type as authorType
					// which we've already defined.
					// This is how we handle nested objects
					Type: authorType,
				},
				"comments": &graphql.Field{
					Type: graphql.NewList(commentType),
				},
			},
		},
	)

	//insert dummy data populate
	tutorials = populate()


	mutationType = graphql.NewObject(graphql.ObjectConfig{
		Name: "Mutation",
		Fields: graphql.Fields{
			"create": &graphql.Field{
				Type:        tutorialType,
				Description: "Create a new Tutorial",
				Args: graphql.FieldConfigArgument{
					"title": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"id": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					tutorial := Tutorial{
						ID: params.Args["id"].(int),
						Title: params.Args["title"].(string),
					}
					tutorials = append(tutorials, tutorial)
					return tutorial, nil
				},
			},
		},
	})
)

func main() {
	// build and configure schema
	schema := buildSchema()

	// insert data into schema
	insert(schema)

	// print data from schema
	print(schema)
}


func populate() []Tutorial {
	author := &Author{Name: "SRIN Sharing Session Book", Tutorials: []int{1}}
	tutorial := Tutorial{
		ID:     0,
		Title:  "Go GraphQL Tutorial",
		Author: *author,
		Comments: []Comment{
			Comment{Body: "First Comment"},
		},
	}

	var tutorials []Tutorial
	tutorials = append(tutorials, tutorial)

	return tutorials
}

func buildSchema() graphql.Schema {
	// schema fields configuration
	fields := graphql.Fields{
		"tutorial": &graphql.Field{
			Type:        tutorialType,

			// it's good form to add a description
			// to each field.
			Description: "Get Tutorial By ID",

			// We can define arguments that allow us to
			// pick specific tutorials. In this case
			// we want to be able to specify the ID of the
			// tutorial we want to retrieve
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.Int,
				},
			},

			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				// take in the ID argument
				id, ok := p.Args["id"].(int)
				if ok {
					// Parse our tutorial array for the matching id
					for _, tutorial := range tutorials {
						if int(tutorial.ID) == id {
							// return our tutorial
							return tutorial, nil
						}
					}
				}
				return nil, nil
			},
		},

		// this is our `list` endpoint which will return all
		// tutorials available
		"list": &graphql.Field{
			Type:        graphql.NewList(tutorialType),
			Description: "Get Tutorial List",
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				return tutorials, nil
			},
		},
	}


	// build schema configuration
	rootQuery := graphql.ObjectConfig{
		Name: "RootQuery",
		Fields: fields,
	}

	schemaConfig := graphql.SchemaConfig{
		Query: graphql.NewObject(rootQuery),
		Mutation: mutationType, //add new variable "mutationType" for data mutation
	}

	schema, err := graphql.NewSchema(schemaConfig)
	if err != nil {
		log.Fatalf("failed to create new schema, error: %v", err)
	}

	return schema
}

func insert(targetSchema graphql.Schema) {
	//prepare query mutation data
	queryMutation := `
    mutation {
        create(id: 1, title: "Hello World") {
			id
            title
        }
    }
	`
	//execute param mutation
	params := graphql.Params {
		Schema:        targetSchema,
		RequestString: queryMutation,
	}

	r := graphql.Do(params)
	if len(r.Errors) > 0 {
		log.Fatalf("failed to execute graphql operation, errors: %+v", r.Errors)
	}

	rJSON, _ := json.Marshal(r)
	fmt.Printf("created: %s \n", rJSON)
}

func print(targetSchema graphql.Schema) {

	//prepare query list data
	queryList := `
		{
			list {
				id
				title
			}
		}
    `

	//execute param list
	paramLists := graphql.Params {
		Schema:        targetSchema,
		RequestString: queryList,
	}

	r := graphql.Do(paramLists)
	if len(r.Errors) > 0 {
		log.Fatalf("failed to execute graphql operation, errors: %+v", r.Errors)
	}

	rJSON, _ := json.Marshal(r)
	fmt.Printf("contents: %s \n", rJSON)
}
```

Build source code using command :

```shell
    $ go mod init # initialize go modules, skip this step if already done previously
    $ go mod vendor # downloading dependencies third party modules
    $ go mod tidy # clean up unused dependencies if any 
    $ go build -o mygraphql
    $ ./mygraphql # run app
```

If app running successfully, it will showing output like:

```shell
created: {"data":{"create":{"id":1,"title":"Hello World"}}} 
contents: {"data":{"list":[{"id":0,"title":"Go GraphQL Tutorial"},{"id":1,"title":"Hello World"}]}} 
```

# Example graphql-api 

Example code to build graphql API are based on references from:

- https://medium.com/@bradford_hamilton/building-an-api-with-graphql-and-go-9350df5c9356
- https://github.com/bradford-hamilton/go-graphql-api

We have updated source code based on references above to be align for this sharing session. Go to source code and details explanation on how to configure, build, run and test via this [link](code/graphql-api/)
