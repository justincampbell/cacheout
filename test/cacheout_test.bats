load test_helper

@test "passes through command output" {
  run $cacheout 1m echo hello

  [ "$output" = "hello" ]
  [ $status -eq 0 ]
}

@test "passes through exit codes" {
  run $cacheout 1m test/fixtures/exit123.sh

  [ "$output" = "" ]
  [ $status -eq 123 ]
}
