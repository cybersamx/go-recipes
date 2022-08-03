# Fake vs Mock

An example that shows the use of test doubles, ie. fakes and mocks, in Go unit testing.

## Definition

Fake and mock are collectively called test doubles.

* **Fake** - Fake doesn't modify the behavior of the system under test. There are 2 forms of fakes:
  * Fake data - Use non-real (fake) data eg. nobody@example.com
  * Fake system - Instead of using an actual system dependency in production, eg. Postgres or MySQL, the test suite uses a simpler system, eg. simple SQLite, to mimic the actual system.
* **Mock** - You have components in your system that call or depend on another dependency. A mock allows the tester to mock that dependency and assert that your system has interacted with the dependency. This recipe provides examples for 2 popular mock frameworks:
  * [Gomock](https://github.com/golang/mock) - This is a framework that generates mock code at build time.
  * [Testify Mock](https://godoc.org/github.com/stretchr/testify/mock) - This is a framework that relies on reflection to generate mocks at runtime.

Read [here](https://stackoverflow.com/questions/346372/whats-the-difference-between-faking-mocking-and-stubbing) for an explanation on fake vs mock.

## Project Description

This example is based on a forgot password web service, which generates a random password and replaces the existing password in the system. The main components are:

* [Account](model/account.go) - The representation of the user's account in the system.
* [AccountModel](model/account.go) - The data access layer that creates, fetches, and deletes a user's account.
* [AccountService](api/api.go) - The API endpoint, from which the user can call to reset his/her password.

In unit test [api_test.go](api/api_test.go), we see that the test is using both a fake and mock:

* Test using a fake - We emulate data access to a SQL database with in-code `map` data type. **Note**: it probably more relevant using SQLite to fake the SQL database access layer than Golang `map`
* Test using a mock - We mock the `AcccountModel` with [MockAccountModel](api/mocks/model_testify_mock.go).

## Setup

1. Install mockgen

   ```bash
   $ make install
   ```

2. Generate the mocks.

   ```bash
   $ make mockgen
   ```

5. Run the test

   ```bash
   $ go test -v ./api
   ```

6. Alternatively, just run everything with the following command:

   ```bash
   $ make install   # Optional: run this if you haven't installed mockgen
   $ make  # Run all the above commands
   ```

## Reference

* [GoMock](https://github.com/golang/mock)
* [Stretchr Testify mock package](https://godoc.org/github.com/stretchr/testify/mock)
