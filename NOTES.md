#### Reference Videos:

- Complete REST API in Go – Build an Event App (GIn, JWT, SQLite, Swagger): https://www.youtube.com/watch?v=ERZadn9artM (~2hr 3mins)
  <br />
  <br />

#### Documentation and Quick Links:

- Go Playground: https://go.dev/play/
- Go User Manual: https://go.dev/doc/
- Go Spec: https://go.dev/ref/spec
- Effective Go: https://go.dev/doc/effective_go
- Standard library: https://pkg.go.dev/std
- Search packages or symbols: https://pkg.go.dev/
- Learn: https://go.dev/learn/ - Go by example: https://gobyexample.com/ - Web Dev: https://gowebexamples.com/ - Books
  <br />
  <br />

#### Others:

- Go GitHub Repositories
  - golang-clean-web-api: https://github.com/naeemaei/golang-clean-web-api/
- Poushak's auth api server in Go - https://github.com/poushak/auth
  <br />
  <br />

### NOTES:

#### Gin:

- gin.H is just a type alias for map[string]interface{}.
  ```go
  // The definition from Gin’s source code:
  // H is a shortcut for map[string]interface{}
  // It is used for JSON responses.
  type H map[string]interface{}
  ```
