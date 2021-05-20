package utils

// TemplateTest is the test with which the created config file is initialised
var TemplateTest = Test{
	URL: "https://jsonplaceholder.typicode.com",
	Methods: map[string]Method{
		"get-post":  TemplateGETMethod,
		"post-post": TemplatePOSTMethod,
	},
}

var TemplateGETMethod = Method{
	Path:   "/posts",
	Method: "get",
	Header: map[string]interface{}{
		"Content-type": "application/json; charset=UTF-8",
	},
	Query: map[string]interface{}{
		"id": 1,
	},
}

var TemplatePOSTMethod = Method{
	Path:   "/posts",
	Method: "post",
	Header: map[string]interface{}{
		"Content-type": "application/json; charset=UTF-8",
	},
	Body: map[string]interface{}{
		"title":  "prawf is amazing!",
		"body":   "If you haven't already, check out prawf to test your REST API endpoints",
		"userId": 1,
	},
}
