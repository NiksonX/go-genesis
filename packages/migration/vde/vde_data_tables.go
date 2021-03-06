package vde

var tablesDataSQL = `
INSERT INTO "%[1]d_tables" ("id", "name", "permissions","columns", "conditions") VALUES ('1', 'contracts', 
	'{"insert": "ContractConditions(\"MainCondition\")", "update": "ContractConditions(\"MainCondition\")", 
	  "new_column": "ContractConditions(\"MainCondition\")"}',
	'{"name": "false",
	  "value": "ContractConditions(\"MainCondition\")",
	  "conditions": "ContractConditions(\"MainCondition\")"}', 'ContractAccess("EditTable")'),
	('2', 'languages', 
	'{"insert": "ContractConditions(\"MainCondition\")", "update": "ContractConditions(\"MainCondition\")", 
	  "new_column": "ContractConditions(\"MainCondition\")"}',
	'{ "name": "ContractConditions(\"MainCondition\")",
	  "res": "ContractConditions(\"MainCondition\")",
	  "conditions": "ContractConditions(\"MainCondition\")"}', 'ContractAccess("EditTable")'),
	('3', 'menu', 
	'{"insert": "ContractConditions(\"MainCondition\")", "update": "ContractConditions(\"MainCondition\")", 
	  "new_column": "ContractConditions(\"MainCondition\")"}',
	'{"name": "ContractConditions(\"MainCondition\")",
"value": "ContractConditions(\"MainCondition\")",
"conditions": "ContractConditions(\"MainCondition\")"
	}', 'ContractAccess("EditTable")'),
	('4', 'pages', 
	'{"insert": "ContractConditions(\"MainCondition\")", "update": "ContractConditions(\"MainCondition\")", 
	  "new_column": "ContractConditions(\"MainCondition\")"}',
	'{"name": "ContractConditions(\"MainCondition\")",
"value": "ContractConditions(\"MainCondition\")",
"menu": "ContractConditions(\"MainCondition\")",
"conditions": "ContractConditions(\"MainCondition\")",
"validate_count": "ContractConditions(\"MainCondition\")",
"validate_mode": "ContractConditions(\"MainCondition\")",
"app_id": "ContractConditions(\"MainCondition\")"
	}', 'ContractAccess("EditTable")'),
	('5', 'blocks', 
	'{"insert": "ContractConditions(\"MainCondition\")", "update": "ContractConditions(\"MainCondition\")", 
	  "new_column": "ContractConditions(\"MainCondition\")"}',
	'{"name": "ContractConditions(\"MainCondition\")",
"value": "ContractConditions(\"MainCondition\")",
"conditions": "ContractConditions(\"MainCondition\")"
	}', 'ContractAccess("EditTable")'),
	('6', 'signatures', 
	'{"insert": "ContractConditions(\"MainCondition\")", "update": "ContractConditions(\"MainCondition\")", 
	  "new_column": "ContractConditions(\"MainCondition\")"}',
	'{"name": "ContractConditions(\"MainCondition\")",
"value": "ContractConditions(\"MainCondition\")",
"conditions": "ContractConditions(\"MainCondition\")"
	}', 'ContractAccess("EditTable")'),
	('7', 'cron',
	  '{"insert": "ContractConditions(\"MainCondition\")", "update": "ContractConditions(\"MainCondition\")",
		"new_column": "ContractConditions(\"MainCondition\")"}',
	  '{"owner": "ContractConditions(\"MainCondition\")",
	  "cron": "ContractConditions(\"MainCondition\")",
	  "contract": "ContractConditions(\"MainCondition\")",
	  "counter": "ContractConditions(\"MainCondition\")",
	  "till": "ContractConditions(\"MainCondition\")",
		"conditions": "ContractConditions(\"MainCondition\")"
	  }', 'ContractConditions("MainCondition")'),
	('8', 'binaries',
	  '{"insert": "ContractConditions(\"MainCondition\")", "update": "ContractConditions(\"MainCondition\")",
		  "new_column": "ContractConditions(\"MainCondition\")"}',
	  '{"app_id": "ContractConditions(\"MainCondition\")",
		  "member_id": "ContractConditions(\"MainCondition\")",
		  "name": "ContractConditions(\"MainCondition\")",
		  "data": "ContractConditions(\"MainCondition\")",
		  "hash": "ContractConditions(\"MainCondition\")",
		  "mime_type": "ContractConditions(\"MainCondition\")"}',
			'ContractConditions("MainCondition")'),
	('9', 'keys',
	  '{"insert": "ContractConditions(\"MainCondition\")", "update": "ContractConditions(\"MainCondition\")",
			"new_column": "ContractConditions(\"MainCondition\")"}',
		'{"pub": "ContractConditions(\"MainCondition\")",
			"multi": "ContractConditions(\"MainCondition\")",
			"deleted": "ContractConditions(\"MainCondition\")",
			"blocked": "ContractConditions(\"MainCondition\")"}',
		'ContractConditions("MainCondition")');	 
`
