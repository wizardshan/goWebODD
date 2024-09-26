// Code generated by ent, DO NOT EDIT.

package ent

import (
	"goWebODD/repository/ent/schema"
	"goWebODD/repository/ent/user"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescHashID is the schema descriptor for hash_id field.
	userDescHashID := userFields[1].Descriptor()
	// user.DefaultHashID holds the default value on creation for the hash_id field.
	user.DefaultHashID = userDescHashID.Default.(string)
	// userDescMobile is the schema descriptor for mobile field.
	userDescMobile := userFields[2].Descriptor()
	// user.DefaultMobile holds the default value on creation for the mobile field.
	user.DefaultMobile = userDescMobile.Default.(string)
	// userDescAge is the schema descriptor for age field.
	userDescAge := userFields[4].Descriptor()
	// user.DefaultAge holds the default value on creation for the age field.
	user.DefaultAge = userDescAge.Default.(int)
	// userDescLevel is the schema descriptor for level field.
	userDescLevel := userFields[5].Descriptor()
	// user.DefaultLevel holds the default value on creation for the level field.
	user.DefaultLevel = userDescLevel.Default.(int)
}
