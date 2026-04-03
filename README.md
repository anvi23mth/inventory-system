Week 5:
Verified that every logical branch, including success paths and error handling, is executed during tests.
Adopted the standard Go idiomatic pattern to test multiple scenarios (valid data, empty fields, negative values) in a single, organized test suite.
Developed MockRepository and MockCategoryRepository with a ShouldFail boolean to simulate database crashes and verify graceful error propagation.
Utilized go test -coverprofile to generate detailed coverage reports.
