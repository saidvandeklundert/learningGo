package main

/*
https://github.com/kubernetes-sigs/yaml
*/
import (
	"fmt"
	"log"

	"sigs.k8s.io/yaml"
)

var slice_of_test_cases = `
- database: APLL_DB
  redis_key: "LLDP_LOC_CHASSIS"
  hash_key: "lldp_loc_sys_desc"
  expected_substring: "SONiC Software Version"

- database: APLL_DB
  redis_key: "LLDP_LOC_CHASSIS"
  hash_key: "lldp_loc_sys_name"
  expected_substring: "acc-sw"
`

var single_test_cases = `
database: APLL_DB
redis_key: "LLDP_LOC_CHASSIS"
hash_key: "lldp_loc_sys_desc"
expected_substring: "SONiC Software Version"
`

type AssertContainsTestCase struct {
	Database          string `json:"database"`
	RedisKey          string `json:"redis_key"`
	HashKey           string `json:"hash_key"`
	ExpectedSubstring string `json:"expected_substring"`
}

func main() {
	fmt.Println("Marshal YAML into a struct:")
	var singleTestCase AssertContainsTestCase
	err := yaml.Unmarshal([]byte(single_test_cases), &singleTestCase)
	if err != nil {
		log.Fatalf("error: %v\n", err)
	}
	fmt.Printf("single singleTestCase:\n%v\n", singleTestCase)

	var testCases []AssertContainsTestCase
	err = yaml.Unmarshal([]byte(slice_of_test_cases), &testCases)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	fmt.Printf("testCases:\n%v\n", testCases)

}
