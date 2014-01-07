#!/usr/bin/env roundup

describe "flint"

before(){
  rm -rf fixtures/ghost/
  mkdir -p fixtures/ghost
}

it_exits_nonzero_for_usage(){
  status=$(set +e ; ../flint -h >/dev/null ; echo $?)
  test 1 -eq $status
}

it_shows_usage(){
  expected="Checks a project for common sources of contributor friction"
  actual=$(../flint -h | sed -n "2p")
  test "$expected" "=" "$actual"
}

it_reports_no_readme(){
  expected="[ERROR] README not found"
  actual=$(../flint | grep "README" | sed -n "1p")
  test "$expected" "=" "$actual"
}

it_finds_a_readme(){
   echo "README" > fixtures/ghost/README
   status=$(set +e; ../flint | grep "README"; echo $?)
   test 1 -eq $status
}

it_reports_no_license(){
  expected="[ERROR] License file not found"
  actual=$(../flint | grep "License" | sed -n "1p")
  test "$expected" "=" "$actual"
}

it_finds_a_license(){
   echo "LICENSE" > fixtures/ghost/LICENSE.md
   status=$(set +e; ../flint | grep "License"; echo $?)
   test 1 -eq $status
}

it_reports_no_contributing_guide(){
  expected="[WARNING] Contributing guide not found"
  actual=$(../flint | grep "Contributing" | sed -n "1p")
  test "$expected" "=" "$actual"
}

it_finds_a_contributing_guide(){
  echo "CONTRIBUTING" > fixtures/ghost/CONTRIBUTING.md
  status=$(set +e; ../flint | grep "Contributing"; echo $?)
  test 1 -eq $status
}

it_reports_no_bootstrap(){
  expected="[WARNING] Bootstrap script not found"
  actual=$(../flint | grep "Bootstrap" | sed -n "1p")
  test "$expected" "=" "$actual"
}

it_finds_a_bootstrap(){
  mkdir -p fixtures/ghost/script
  echo "bootstrap" > fixtures/ghost/script/bootstrap
  status=$(set +e; ../flint | grep "Bootstrap"; echo $?)
  test 1 -eq $status
}

it_reports_no_test_script(){
  expected="[WARNING] Test script not found"
  actual=$(../flint | grep "Test" | sed -n "1p")
  test "$expected" "=" "$actual"
}

it_finds_a_test_script(){
  mkdir -p fixtures/ghost/script
  echo "testing" > fixtures/ghost/script/test
  status=$(set +e; ../flint | grep "Test"; echo $?)
  test 1 -eq $status
}
