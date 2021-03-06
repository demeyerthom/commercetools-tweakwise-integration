/*
 * Commercetools Tweakwise Integration
 *
 * A service built for managing the Commercetools to Tweakwise integration
 *
 * API version: 0.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package server

type Configuration struct {
	Tweakwise ConfigurationTweakwise `json:"tweakwise"`

	Source ConfigurationSource `json:"source"`
}

// AssertConfigurationRequired checks if the required fields are not zero-ed
func AssertConfigurationRequired(obj Configuration) error {
	elements := map[string]interface{}{
		"tweakwise": obj.Tweakwise,
		"source":    obj.Source,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	if err := AssertConfigurationTweakwiseRequired(obj.Tweakwise); err != nil {
		return err
	}
	if err := AssertConfigurationSourceRequired(obj.Source); err != nil {
		return err
	}
	return nil
}

// AssertRecurseConfigurationRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of Configuration (e.g. [][]Configuration), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseConfigurationRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aConfiguration, ok := obj.(Configuration)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertConfigurationRequired(aConfiguration)
	})
}
