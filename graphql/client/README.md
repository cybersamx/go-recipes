# GraphQL Client in Go

An example of implementing a simple GraphQL client in Go to query against a public GraphQL service <https://countries.trevorblades.com/>. The client makes 2 calls:

* GraphQL request without any filter.
* GraphQL request with a filter passed as a parameter.

We can construct a GraphQL query as a string and pass it as plain text when making the API call. But for the construction of more complex queries, including support for passed arguments, we may want to use a GraphQL package, like the [machinebox/graphql](https://github.com/machinebox/graphql) package.

## Setup

1. Run the program

   ```bash
   $ make
   ```

## Querying from the Playground

GraphQL is a querying language that can be embedded in a programming language like Go and a transport stack like HTTP. The following steps show how the same GraphQL queries can be made from the service GraphQL playground.

1. Go to <https://countries.trevorblades.com/>
1. Run the following query to return all countries.

   ```
   query {
     countries {
       code,
       name,
       languages {
         name
       }
     }
   }
   ```

1. Return all countries that matches country.code == "US" (case sensitive).

   ```
   query {
	 countries(filter: {
       code: { eq: "US" }
     }) {
       code,
       name,
       languages {
         name
       }
     }
   }
   ```

## Reference

* [Machinebox GraphQL Package](https://github.com/machinebox/graphql)
* [Introduction to GraphQL](https://graphql.org/learn/)
* [Passing Arguments in GraphQL](https://graphql.org/graphql-js/passing-arguments/)
