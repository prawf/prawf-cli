package utils

// TemplateTest is the test with which the created config file is initialised
var TemplateTest = Test{
	URL: "https://jsonplaceholder.typicode.com",
	Methods: []Method{
		TemplateGETMethodPassing,
		TemplateGETMethodFailing,
		TemplatePOSTMethodPassing,
		TemplatePUTMethodPassing,
	},
}

var TemplateGETMethodPassing = Method{
	Name:   "get-pass",
	Path:   "/posts",
	Method: "get",
	Query: map[string]interface{}{
		"id": 1,
	},
	Expect: Expect{
		Contain: map[string]interface{}{
			"id": 1,
		},
	},
}

var TemplateGETMethodFailing = Method{
	Name:   "get-fail",
	Path:   "/posts",
	Method: "get",
	Query: map[string]interface{}{
		"id": 3,
	},
	Expect: Expect{
		Keys: []string{"uuid"},
	},
}

var TemplatePOSTMethodPassing = Method{
	Name:   "post-pass",
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
	Expect: Expect{
		Equal: map[string]interface{}{
			"title":  "prawf is amazing!",
			"body":   "If you haven't already, check out prawf to test your REST API endpoints",
			"userId": 1,
		},
	},
}

var TemplatePUTMethodPassing = Method{
	Name:   "put-pass",
	Path:   "/posts/1",
	Method: "put",
	Header: map[string]interface{}{
		"Content-type": "application/json; charset=UTF-8",
	},
	Body: map[string]interface{}{
		"id":     1,
		"title":  "prawf is awesome!",
		"body":   "Give us a star on GitHub/prawf!",
		"userId": 1,
	},
	Expect: Expect{
		Contain: map[string]interface{}{
			"title": "prawf is awesome!",
		},
	},
}
