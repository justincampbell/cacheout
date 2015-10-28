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

@test "exposes environment variables to the command" {
  export FOO=bar
  run $cacheout 1m test/fixtures/echo_foo.sh

  [ "$output" = "bar" ]
  [ $status -eq 0 ]
}

@test "no arguments shows help" {
  run $cacheout

  echo $output | grep "cacheout"
  [ $status -eq 1 ]
}

@test "not enough arguments shows help" {
  run $cacheout 1m

  echo $output | grep "cacheout"
  [ $status -eq 1 ]
}
