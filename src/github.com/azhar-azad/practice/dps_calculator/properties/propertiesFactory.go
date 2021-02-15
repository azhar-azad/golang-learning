package properties

import "strings"

func GetProperties(packageName string) IProperty {
	if strings.ToUpper(packageName) == "DEMO" {
		return &demoProperty{}
	}
	return nil
}
