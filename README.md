Example pact with go
====================

This is a port of the example provider written in Ruby to Google Go. The Pact is copied verbatim from the ruby project.

To run the provider:

    $ go get github.com/paulbellamy/mango
    $ go run src/main.go

You can then direct your browser to http://localhost:3000/producer.json

To run the pact verification test:

    $ go get github.com/bmizerany/assert
    $ go get github.com/drewolson/testflight
    $ go test -v src/pact/producer_test.go

