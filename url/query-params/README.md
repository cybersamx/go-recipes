# Setting Query Parameters in URL

What if we want to turn this url:

`https://localhost:8000/api/v3/location/3a/shelf/5ff`

to

`https://localhost:8000/api/v3/location/3a/shelf/5ff?index=3&page_size=10&sort=id&sort=mod_time`

* We create an `URL` object representing `https://localhost:8000/api/v3/location/3a/shelf/5ff`. But we can't just append the query parameters directly. And the method `Query()` isn't what you think it is. The way to set the query parameters here is a bit unusual.
* The `Query()` method is a getter that returns the `RawQuery` field in an `URL` object as `url.Values`. It's not a setter.
* So the trick to setting the query parameters in a `URL` object is to get an `url.Values` object and use its methods to construct query parameters. And finally setting the `RawQuery` field of the `URL` object directly with url-encoded query parameters.

