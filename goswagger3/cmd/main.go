package main

import "github.com/bcordobaq/swagger-go-example/goswagger3/pkg/thing"

// @version 1.0
// @title A thing server
// @description A thing server full description
// @termsofservice http://swagger.io/terms/

// @contactname Barbara Cordoba
// @contacturl https://github.com/bcordobaq
// @contactemail barbaracordobaq@gmail.com

// @licensename Apache 2.0
// @licenseurl http://www.apache.org/licenses/LICENSE-2.0.html

// @server localhost:8080 localhost
// @server localhost:9090 localhost2

func main() {
	server := thing.NewServer()
	server.ListenAndServe(":8080")
}
