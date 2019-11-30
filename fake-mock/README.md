# Fake vs Mock

An example of how to use test doubles ie. fakes and mocks, and their differences.

## Definition

Fake and mock are collectively called test doubles. The definitions of the terms are:

* **Fake** - Fake doesn't modify the behavior of the system under test. There are 2 forms of fakes:
  * Fake data - use not-real (fake) data eg. nobody@example.com
  * Fake system - instead of using an actual system dependency in production eg. Postgres or MySQL, the test system uses a simple SQLite
* **Mock** - You have components in your system that calls or depends on a dependency. You mock that dependency and assert that your system has interacted with the dependency.

The best explanation for this perennial question is available [here](https://stackoverflow.com/questions/346372/whats-the-difference-between-faking-mocking-and-stubbing).

## Project Description

This example emulates a forgot password web service, which generates a random password and replaces the existing password in the system that is associated to the user. The main components in this example are:

* [Account](model.go) - The user's account in the system.
* [AccountModel](model.go) - The layer that create, get, and delete a user's account.
* [AccountService](api.go) - Represents the API, from which the user can call to reset his/her password.

The unit test is for testing `AccountService`. As `AccountService` calls a function in `AccountModel`, we need to either fake or mock the `AccountModel` function, on which `AccountService` depends. This example demonstrates both technique - see <api_test.go>.

* Test using a fake - We emulate data access to a SQL database with in-code `map` data type. **Note**: it probably more relevant using SQLite to fake the SQL database access layer than Golang `map`
* Test using a mock - We mock the `AcccountModel` thru [MockAccountModel](mock_model.go).

## Setup

1. Run the test

   ```bash
   $ go test .
   ```

## Reference

* [Stretchr Testify mock package](https://godoc.org/github.com/stretchr/testify/mock)
