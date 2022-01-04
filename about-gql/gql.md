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

* GraphQL uses its own language, SDL - Schema definition language.

#### SDL - Schema Definition Language

##### Type

* has a name and can implement one or more interfaces.

```
type Post implements Item {
  # ...
}
```

##### Field

* Has a name and a type.

age: Int

* Can include Int, Float, String, Boolean, ID.
* Non-nullable fields indicated with !
* Lists denoted by []

##### Enum

* Scalar value that has a specified set of possible values.

```
enum Category {
  PROGRAMMING_LANGUAGES
  API_DESIGN
}
```
##### Interface

Interface is a list of fields. GraphQL type must have the same fields as all the interfaces it implments and all of the interface fields must be of the same type.

```
interface Item {
  title: String!
}
```

##### Schema Directive

A directive allows you to attach arbitrary information to any other schema definition element. Directives are always placed behind the element they describe:

```
name: String! @defaultValue(value: "new blogpost")
```

* Directives don’t have intrinsic meaning. Each GraphQL implementation can define their own custom directives that add new functionality.

* GraphQL specifies built-in skip and include directives that can be used to include or exclude specific fields in queries, but these aren't used in the schema language.

#### Queries

* Instead of multiple endpoints with specific data structures, GraphQL has one specific endpoint with a completely flexible data structure. 
* The client itself must send more information to the server to express its data needs - this information is called a query.

##### Basic Query

```
{
  allPersons {
    name
  }
}
```

* allPersons field in this query is called the root field of the query. Everything that follows th root field is the *payload* of the query.

This query would return a list of persons currently stored in the database, here is an example response:

```
{
  "allPersons": [
    { "name": "Johnny" },
    { "name": "Sarah" },
    { "name": "Alice" }
  ]
}
```
* Notice that, "allPersons" has, "name" in the response, but the "age" was not returned by the server. This is because, "name," was the only field specified in this query.  You could add age by specifying this within the query:

```
{
  allPersons {
    name
    age
  }
}
```
##### Basic Query with Arguments

* You can add various arguments in within parenthesis, for example: "last: 2"

```
{
  allPersons(last: 2) {
    name
  }
}
```

##### Mutations

Data can be added into a database with mutations:

```
mutation {
  createPerson(name: "Bob", age: 36) {
    name
    age
  }
}
```

* This above of course adds, Person, "Bob" with age:36.
* Note the root field, "createPerson"

There should be a server response for this:

```
"createPerson": {
  "name": "Bob",
  "age": 36,
}
```
An often used pattern within GraphQL is to include Id's that are generated by the server when new objects are created.

```
type Person {
  id: ID!
  name: String!
  age: Int!
}
```
Thus, when a mutation is used, such as, "createPerson," then id can be queried upon the creation/mutation:

```
mutation {
  createPerson(name: "Alice", age: 36) {
    id
  }
}
```
##### Subscriptions

* Many applications must have a realtime connection to the server in order to get immediately informed about important events.
* GraphQL has subscriptions to cover this.
* When a client subscribes to an event, it will initiate and hold a steady connection to the server. Whenever that particular event happens, the server pushes the corresponding data to the client.
* Unlike queries and mutations which follow the request-response-cycle, it's a stream of data.

```
subscription {
  newPerson {
    name
    age
  }
}
```
After a client send this subscription information to a server, a connection is opened. Whenever a new mutation is performed that creates a newPerson, the following type of info is pushed to the client immediatley:

```
{
  "newPerson": {
    "name": "Jane",
    "age": 23
  }
}
```

##### Defining a Schema

* There are several root types.

```
type Query { ... }
type Mutation { ... }
type Subscription { ... }
```
* Query
* Mutation
* Subscription

...are all entrypoints for the requests sent by the client. To enable an allPersons-query, the Query type would need to be written as follows:

```
type Query {
  allPersons: [Person!]!
}
```

* allPersons is the root field.
* So if we wanted to add the last argument to the allPersons field, you would have to write the above Query as follows:

```
type Query {
  allPersons(last: Int): [Person!]!
}
```
* Similarly, to establish a mutation on this datatype, you would have to add a root field to the Mutation type, using "type":

```
type Mutation {
  createPerson(name: String!, age: Int!): Person!
}
```
* the same would be necessary to create a subscription.  The full schema would be all of the queries and mutations together as follows:

```
type Query {
  allPersons(last: Int): [Person!]!
  allPosts(last: Int): [Post!]!
}

type Mutation {
  createPerson(name: String!, age: Int!): Person!
  updatePerson(id: ID!, name: String!, age: String!): Person!
  deletePerson(id: ID!): Person!
}

type Subscription {
  newPerson: Person!
}

type Person {
  id: ID!
  name: String!
  age: Int!
  posts: [Post!]!
}

type Post {
  title: String!
  author: Person!
}
```

### Architecture

There are different use cases for GraphQL

1. GraphQL Server with a Connected Database
2. GraphQL server that is a thin layer in front a number of third party or legacy systems and integrates them into a GraphQL API.
3. Hybrid approach of the above.

#### GraphQL Server with Connected Database

* Single web server that implements GraphQL specification. Server reads query payloads and fetches required information, "resolving," the query, then constructs a response and sends it back.
* Transport layer agnostic...it could be used over TCP, WebSockets, SSH, etc.
* Database format not important, it could be NoSQL like MongoDB, or SQL like AWS Aurora or  whatever.

#### GraphQL Layer Integrating with Existing Systems

* GraphQL could be used to unify underlying systems to hide their complexity.
* There could be legacy systems, microservices, third party APIs, etc., all tied together with GraphQL.

#### Hybrid Approach

* Basically combine the two above approaches.

#### Resolver Functions

* How is flexibility gained?  
* queries and mutations consist of a set of fields. In the server implementation, each field corresponds to one function called a resolver. The sole purpose of the resolver is to fetch data for a particular field.

Query Example:

```
query {
    User(id: "hi")

}
```
Resolver:
```
User(id: String!): User
```
* The [Resolver](https://graphql.org/learn/execution/#root-fields-resolvers) is a type that represents all of the possible entry points into the GraphQL API, the root type or the query type.
* GraphQL servers can be written in different languages, and the resolver uses four arguments:

* obj - previous object
* args - arguments provided to the field in the GraphQL query.
* context - value provided to every resolver and holds important contextual information, like the ucrrently logged in user
* info - includes schema details.

Basically, a resolver [graphobjecttype](https://graphql.org/graphql-js/type/#graphqlobjecttype) must be written.

#### GraphQL Client Libraries

* Complexity is pushed to the server, as opposed to REST api's which send everything (imperative because it's imperative that we get everything).
* GraphQL uses declarative data fetching, rather than imperative data fetching.


### Main Tutorial

https://www.howtographql.com/choose/

Choosing React+Apollo

https://www.howtographql.com/react-apollo/0-introduction/

#### Intro

https://www.howtographql.com/react-apollo/0-introduction/

* Build a clone of HackerNews:

* Display links
* Search the list of links
* User auth
* Allow auth'd users to create new links
* Upvote/downvote links
* Realtime updates

Arch:

* React Library
* Apollo Client - caching GraphQL client.
* Apollo Server - GraphQL Server with focus on easy setup.
* Prisma - makes it easier to work with relational datbases. 
#### Getting Started

#### Queries

#### Mutations

#### Routing

#### Authentication

#### Mutations and Updating the Store

#### Filtering

#### Realtime Updates with GraphQL Subscriptions

#### Pagination

#### Summary


# Resources

* [1] [GraphQL Tutorial](https://graphql.org/learn/)
* [2] [How to GraphQL](https://www.howtographql.com/)