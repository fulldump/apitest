# Apitest

Easy way to test golax apis

## Getting started

```go
my_api := golax.NewApi()

// build `my_api`...

testserver := apitest.New(my_api)

r := testserver.Request("POST", "/users/23/items").
    WithHeader("Content-Type", "application/json").
    WithCookie("sess_id", "123123213213213"),
    WithBodyString(`
        {
            "name": "pencil",
            "description": "Blah blah..."
        }
    `).
    Do()

r.StatusCode // Check this
r.BodyString() // Check this
```

## Sending body JSON

```go
r := testserver.Request("POST", "/users/23/items").
    WithBodyJson(map[string]interface{}{
        "name": "pencil",
        "description": "Blah blah",
    }).
    Do()
```

## Reading body JSON

```go
r := testserver.Request("GET", "/users/23").
    Do()
    
body := r.BodyJson()
```

## Asynchronous request

```go
func Test_Example(t *testing.T) {

	a := golax.NewApi()

	a.Root.Node("users").Method("GET", func(c *golax.Context) {
		fmt.Fprint(c.Response, "John")
	})

	s := apitest.New(a)

	w := &sync.WaitGroup{}

	for i := 0; i < 10; i++ {
		w.Add(1)
		n := i
		go s.Request("GET", "/users").DoAsync(func(r *apitest.Response) {

			if http.StatusOK != r.StatusCode {
				t.Error("Expected status code is 200")
			}

			fmt.Println(r.BodyString(), n)

			w.Done()
		})
	}

	w.Wait()
}
```

