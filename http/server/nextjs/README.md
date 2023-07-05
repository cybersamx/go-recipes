# Run Next.js (Client) on Go Server

Build a client-side Next.js web application and deploy it to a Go server.

## Setup

1. Run the application.

   ```shell
   cd web
   npm install
   npm run dev
   ```
   
1. Build and export the application.

   ```shell
   cd web
   npm run build
   ```
   
1. Build the go server.

   ```shell
   mkdir bin
   go build -o bin/server main.go
   cd bin
   ./server
   ```
   
   Alternatively, just run `make run`.

1. Navigate to <http://localhost:8000>.

## Reference

* [Nextjs Docs: Deploying application](https://nextjs.org/docs/pages/building-your-application/deploying)
* [Nextjs Docs: Static Exports](https://nextjs.org/docs/app/building-your-application/deploying/static-exports)
