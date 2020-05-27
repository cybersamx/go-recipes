# Simple Web Form Submission and JWT Auth

This recipe implements web form submissionin Go. This is a login application where the user enters the username and password for authentication. Once a user has been authenticated, the application will persist the session using JWT.

 This application has the following functions:

1. Reads a file from the filesystem and streams the static file content to the web as a web form for the user to log in.
1. Handles the submitted form.
1. Persists a JWT in a cookie when the user is authenticated successfully.
1. Redirects the user to another static web page called welcome page, which is only accessible to authorized users. The application examines if a valid JWT is in the request.
1. Disables browser caching on the welcome page.

## Setup

1. Run the server

   ```bash
   $ make run
   ```

2. Launch a web browser and navigate to <http://localhost:8000>.
