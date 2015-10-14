load test_helper

@test "passes through command output" {
  run $cacheout 1m echo hello

  [ "$output" = "hello" ]
  [ $status -eq 0 ]
}
