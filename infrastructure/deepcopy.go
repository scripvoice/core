package infrastructure

import (
	"reflect"
)

func DeepCopyObject(obj interface{}) interface{} {
	// Get the type of the object
	objType := reflect.TypeOf(obj)

	// Create a new instance of the object
	copyObj := reflect.New(objType).Elem()

	// Iterate over each field of the object
	for i := 0; i < objType.NumField(); i++ {
		field := objType.Field(i)

		// Get the value of the field in the original object
		fieldValue := reflect.ValueOf(obj).Field(i)

		// Perform a deep copy of the field value if it's a reference type
		if fieldValue.Kind() == reflect.Ptr {
			// Create a new instance of the field type
			newFieldValue := reflect.New(field.Type.Elem())

			// Perform a deep copy of the field value by recursively calling DeepCopyObject
			deepCopy := DeepCopyObject(fieldValue.Elem().Interface())

			// Set the deep copied value to the field in the new object
			newFieldValue.Elem().Set(reflect.ValueOf(deepCopy))

			// Set the field value in the new object
			copyObj.Field(i).Set(newFieldValue)
		} else {
			// Copy the field value directly if it's not a reference type
			copyObj.Field(i).Set(fieldValue)
		}
	}

	// Return the deeply copied object
	return copyObj.Interface()
}
