package utils

import "github.com/gruntwork-io/terratest/modules/random"

func RandomUniqueId() string {
	return random.UniqueId()
}
