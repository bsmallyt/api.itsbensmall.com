module api/main

go 1.23.3

replace api/spellcheck => ../spellcheck

require (
	api/reload v0.0.0-00010101000000-000000000000
	api/spellcheck v0.0.0-00010101000000-000000000000
)

replace api/reload => ../reload
