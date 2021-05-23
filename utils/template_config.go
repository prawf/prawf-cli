package utils

// TemplateTest is the test with which the created config file is initialised
var TemplateTest = Test{
	URL: "https://jsonplaceholder.typicode.com",
	Methods: []Method{
		TemplateGETMethod,
		TemplatePOSTMethod,
	},
}

var TemplateGETMethod = Method{
	Name:   "get-post",
	Path:   "/posts",
	Method: "get",
	Query: map[string]interface{}{
		"id": 1,
	},
	Expect: Expect{
		Contain: map[string]interface{}{
			"id": 22,
		},
		Keys: []string{"id"},
		Equal: map[string]interface{}{
			"id": 22,
		},
	},
}

var TemplatePOSTMethod = Method{
	Name:   "post-post",
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
