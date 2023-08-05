package openapi3

type SchemaCollection struct {
	Title []string
	Type  []string
	// include arrays of all additional Schema fields
}

// Merge replaces objects under AllOf with a flattened equivalent
func Merge(schema Schema) (*Schema, error) {
	if !isListOfObjects(&schema) {
		return &schema, nil
	}
	if schema.AllOf != nil {

	}
	schema.AllOf = nil
	return &schema, nil
}

func collect(schemas []*Schema) SchemaCollection {
	collection := SchemaCollection{}
	for _, s := range schemas {
		collection.Title = append(collection.Title, s.Title)
	}
	// collect additional fields
	return collection
}

func mergeFields(schema *Schema, collection *SchemaCollection) (*Schema, error) {
	// resolve Title
	for _, title := range collection.Title {
		// create new title
		title := "new title"
		schema.Title = title
	}

	// merge additional fields ...

	return schema, nil
}

func mergeAllOf(allOf SchemaRefs) (Schema, error) {

	schemas := make([]*Schema, 0) // naming
	for _, schema := range allOf {
		merged, err := Merge(*schema.Value)
		if err != nil {
			return Schema{}, err
		}
		schemas = append(schemas, merged)
	}

	collection := collect(schemas)
	schema, err := mergeFields(schema, collection)
	if err != nil {
		return *schema, err
	}
	return *schema, nil
}

func isListOfObjects(schema *Schema) bool {
	if schema == nil || schema.AllOf == nil {
		return false
	}

	for _, subSchema := range schema.AllOf {
		if subSchema.Value.Type != "object" {
			return false
		}
	}
	return true
}
