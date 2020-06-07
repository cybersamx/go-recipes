# Fake vs Mock

A recipe on how to use test doubles, ie. fakes and mocks, in unit tests.

## Definition

Fake and mock are collectively called test doubles. The definitions of the terms are:

* **Fake** - Fake doesn't modify the behavior of the system under test. There are 2 forms of fakes:
  * Fake data - use non-real (fake) data eg. nobody@example.com
  * Fake system - instead of using an actual system dependency in production, eg. Postgres or MySQL, the test suite uses a simpler system, eg. simple SQLite, to mimic the actual system.
* **Mock** - You have components in your system that calls or depends on a dependency. A mock allows the tester to mock that dependency and assert that your system has interacted with the dependency. This recipe provides examples for 2 popular mock frameworks:
  * [Gomock](https://github.com/golang/mock) - this is a framework that generates mock code at build time.
  * [Testify Mock](https://godoc.org/github.com/stretchr/testify/mock) - this is a framework that relies on reflection to generate mocks at runtime.

The best explanation to this perennial question of fake vs mock can be read [here](https://stackoverflow.com/questions/346372/whats-the-difference-between-faking-mocking-and-stubbing).

## Project Description

This recipe emulates a forgot password web service, which generates a random password and replaces the existing password in the system. The main components are:

* [Account](model.go) - The representation of the user's account in the system.
* [AccountModel](model.go) - The data access layer that creates, fetches, and deletes a user's account.
* [AccountService](api.go) - The layer that implements the API endpoint, from which the user can call to reset his/her password.

The unit tests in this recipe are used to test `AccountService`. When `AccountService` calls a function in `AccountModel`, we need to either fake or mock that function in `AccountModel`. In unit test <api_test.go>, we see that the test is using both a fake and mock:

* Test using a fake - We emulate data access to a SQL database with in-code `map` data type. **Note**: it probably more relevant using SQLite to fake the SQL database access layer than Golang `map`
* Test using a mock - We mock the `AcccountModel` thru [MockAccountModel](mock_model.go).

## Setup

1. Install mockgen

   ```bash
   $ mockgen -source=./model/account.go -package=mock_recipe AccountModel > api/mock_recipe/model_mockgen.go
   ```

1. Run the test

   ```bash
   $ go test -v ./api
   ```

1. Alternatively, just run everything with the following command:

   ```bash
   $ make
   ```

## Reference

* [GoMock](https://github.com/golang/mock)
* [Stretchr Testify mock package](https://godoc.org/github.com/stretchr/testify/mock)
