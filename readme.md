this repo is a source for learning test in Golang.

first follow this channel: https://t.me/golandchannel ... all videos storage here.

then see the videos by following ralated tag .

this articles can be useful for you , read these articles in order :
https://ieftimov.com/posts/testing-in-go-first-principles/
https://ieftimov.com/posts/testing-in-go-failing-tests/
https://ieftimov.com/posts/testing-in-go-writing-practical-failure-messages/
https://ieftimov.com/posts/testing-in-go-go-test/
https://ieftimov.com/posts/testing-in-go-table-driven-tests/
https://ieftimov.com/posts/testing-in-go-subtests/
https://ieftimov.com/posts/testing-in-go-fixtures/
https://ieftimov.com/posts/testing-in-go-dependency-injection/
https://ieftimov.com/posts/testing-in-go-test-doubles-by-example/
https://ieftimov.com/posts/testing-in-go-golden-files/
https://ieftimov.com/posts/testing-in-go-clean-tests-using-t-cleanup/
https://ieftimov.com/posts/testing-in-go-testing-http-servers/
https://ieftimov.com/posts/testing-in-go-websockets/
https://ieftimov.com/posts/testing-in-go-stop-leaking-files/



see this test in golang:
unit test
integration test
end to end test


in golang it is a convention that every file which has _test.go in last part , run as a test file by compliler.

another convention is When we are writting a test for a function or a class, we have to add Test prefix to it and the input must be T from type *testing.T

as well we must set _test in our package name. becuase in test file we need some new functions.
this functions are usable just for testing file.