package helper

import (
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	"reflect"
	"runtime"
	"strconv"
	"testing"

	"github.com/carwow/terraform-provider-hirefire/client"
	"github.com/carwow/terraform-provider-hirefire/provider"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func PreCheck(t *testing.T) func() {
	return func() {
		if os.Getenv("HIREFIRE_API_KEY") == "" {
			t.Fatal("HIREFIRE_API_KEY must be set")
		}
	}
}

func ProviderFactories() map[string]func() (*schema.Provider, error) {
	return map[string]func() (*schema.Provider, error){
		"hirefire": func() (*schema.Provider, error) { return provider.Provider(), nil },
	}
}

func Client() *client.Client {
	return client.New(os.Getenv("HIREFIRE_API_KEY"))
}

func CheckResourceAttributes(name string, attributes map[string]string) resource.TestCheckFunc {
	checks := []resource.TestCheckFunc{}
	for attr, value := range attributes {
		checks = append(checks, resource.TestCheckResourceAttr(name, attr, value))
	}
	return resource.ComposeAggregateTestCheckFunc(checks...)
}

func GetResourceID(s *terraform.State, name string) (string, error) {
	rs, ok := s.RootModule().Resources[name]
	if !ok {
		return "", fmt.Errorf("Not found: %s", name)
	}
	return rs.Primary.ID, nil
}

func Equals(exp, act interface{}) error {
	if !reflect.DeepEqual(exp, act) {
		_, file, line, _ := runtime.Caller(1)
		return fmt.Errorf("\033[31m%s:%d:\n\n\texp: %#v\n\n\tgot: %#v\033[39m\n\n", filepath.Base(file), line, exp, act)
	}
	return nil
}

func StringOrEmpty(str *string) string {
	if str == nil {
		return ""
	}
	return *str
}

func ItoaOrZero(n *int) string {
	if n == nil {
		return "0"
	}
	return strconv.Itoa(*n)
}

func BoolOrFalse(b *bool) string {
	if b == nil {
		return "false"
	}
	return strconv.FormatBool(*b)
}

func RandString(size int) string {
	return acctest.RandString(size)
}

func RandHex(size int) string {
	return acctest.RandStringFromCharSet(size, CharSetHex)
}

const CharSetHex = "012346789abcdef"

func RandBool() bool {
	return rand.Intn(2) == 0
}

func RandInt(from, to int) int {
	return from + rand.Intn((to-from)+1)
}
