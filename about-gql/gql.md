# GraphQL

Original Project(s):

* https://graphql.org
* https://www.gqlstandards.org/

# About

* Similar to SQL, but rather than querying across a table, you query across a graph.
* Created by Facebook originally.
* Fields are nested within other fields within JSON format.
# Tutorial

## Installing GraphQL on Debian

To use as a CLI, you have to install graphql-cli.

* [GraphQL CLI](https://www.graphql-cli.com/introduction/)

```
apt install nodejs npm

npm install --global yarn

yarn global add graphql-cli graphql


```
> Each command in GraphQL CLI is treated as a plugin. In order to use the command, you have to install it first. Each command's package name follows this pattern: @graphql-cli/[COMMAND-NAME]. So to install the init command we used above, we would run


```
yarn global add @graphql-cli/init

graphql init

root@a93522e30f3d:/# graphql init
? Select the best option for you (Use arrow keys)
❯ I want to create a new project from a GraphQL CLI Project Template.
  I have an existing project using OpenAPI/Swagger Schema Definition.
  I have an existing project using GraphQL and want to add GraphQL CLI (run from
 project root).

 ? Choose a template to bootstrap
  apollo-fullstack-react-postgres-ts
❯ apollo-fullstack-react-mongo-ts

🚀  GraphQL CLI project successfully initialized:
helloworld
Next Steps:
- Change directory to the project folder - cd /helloworld
- Run yarn install to install dependencies
- (Optional) Initialize your git repo. git init.
- Follow the instructions in README.md to continue...

```

You can then use GraphQL as a CLI.


## Basics - https://graphql.org/learn/

* Query language for API, server-side runtime for executing queries using a type system.

* GraphQL service tells you who the logged in user is (me) as well as that name might look like:

```
type Query {
  me: User
}

type User {
  id: ID
  name: String
}
```
Functions for each field on each type:

```
function Query_me(request) {
  return request.auth.user;
}

function User_name(user) {
  return user.getName();
}
```

## HowToGraphQL

* General Notes
* REST has been the fuzzy standard for designing web API's, includes stateless servers and structured access to resources. However, they are too inflexible.
* GraphQL copdes with the need for more flexibility and efficiency.

For example, you need to display the titles and posts of a specific user. Screen also displays the last 3 followers of a user.

### REST Solution

With REST API, you would do an endpoint:

```
/users/<id>/posts
```
or

```
/users/<id>/followers
```

* There would be an 1) HTTP GET request, grabbing the user ID.
* 2) /posts/ would be returned as a JSON list
* 3) /followers/ would be returned as a JSON list

There would be 3 requests to 3 different endpoints to fetch the required data.  The posts return additional data that is not needed, so there is overfetching.

### GraphQL Solution

Send a query request as an HTTP POST request:

```
query{
    User(id: "blah") {
        name
        posts {
            title
        }
        followers(last:3){
            name
        }
    }
}
```

Then the response:

```
{
    "data": {
        "User": {
            "name": "Mary",
            "posts: [
                {title: "Learn GraphQL Today" }
            ],
            "followers:" [
                { name: "John" },
                { name: "Alice" },
                { name: "Sam" },                                
            ]
        }
    }

}
```

Basically it returns specifically and exactly what was requested without the extra information as the API would do.


### Core Concepts



# Resources

* [1] [GraphQL Tutorial](https://graphql.org/learn/)
* [2] [How to GraphQL](https://www.howtographql.com/)