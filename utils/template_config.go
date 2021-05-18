package utils

var TemplateTest = Test{
	URL: "testURL",
	Methods: map[string]Method{
		"testMethod": TemplateMethod,
	},
}

var TemplateMethod = Method{
	Path:   "testPath",
	Method: "testMethod",
	Query: map[string]interface{}{
		"testQa": "def",
	},
	Body: map[string]interface{}{
		"testQb": "ghi",
	},
}
