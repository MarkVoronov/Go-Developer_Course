module main

go 1.20

require (
	hellomodule_v0 v1.0.0
    hellomodule_v1 v1.1.0
)

replace (
	hellomodule_v0 v1.0.0 => github.com/markvoronov/hellomodule v1.0.0
	hellomodule_v1 v1.1.0 => github.com/markvoronov/hellomodule v1.1.0
)