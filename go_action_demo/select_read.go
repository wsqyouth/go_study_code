package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {

	fmt.Println("hello")
}

type DpCreativeElements struct {
	TitlePosition string `dpelement:""`
}

func (e *DpCreativeElements) parseElements(ctx context.Context, creativeElements []*storage.CreativeElement) error {
	v := reflect.Indirect(reflect.ValueOf(e))

	for i := 0; i < v.NumField(); i++ {
		typeField := v.Type().Field(i)
		dpelementTag, ok := typeField.Tag.Lookup("dpelement")
		if !ok {
			continue
		}

		fieldValue := v.Field(i)
		fieldName := v.Type().Field(i).Name
		var elementName string
		if dpelementTag == "" {
			elementName = fmt.Sprintf("/%s", strcase.ToSnake(fieldName))
		} else {
			elementName = fieldName
		}

		err := e.setStructField(ctx, fieldValue, elementName, creativeElements, typeField.Tag)
		if err != nil {
			return err
		}
	}

	return nil
}

func (e *DpCreativeElements) setStructField(ctx context.Context,
	fieldValue reflect.Value, elementName string,
	creativeElements []*storage.CreativeElement, fieldTag reflect.StructTag) error {
	if !(fieldValue.IsValid() && fieldValue.CanSet()) {
		return nil
	}

	fieldValueStr, ok := e.getContentByName(ctx, elementName, creativeElements)
	if !ok {
		// 有默认值的话使用默认值
		dpDefaultTag, hasDefault := fieldTag.Lookup("dpdefault")
		if hasDefault {
			fieldValueStr = dpDefaultTag
		} else {
			return nil
		}
	}

	switch fieldValue.Kind() {
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		iValue, err := strconv.ParseUint(fieldValueStr, 10, 64)
		if err != nil {
			return errors.Errorf("strconv.ParseUint err: %v, content: %s",
				err, fieldValueStr)
		}
		fieldValue.SetUint(iValue)
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		iValue, err := strconv.ParseInt(fieldValueStr, 10, 64)
		if err != nil {
			return errors.Errorf("strconv.ParseInt err: %v, content: %s",
				err, fieldValueStr)
		}
		fieldValue.SetInt(iValue)
	case reflect.String:
		fieldValue.SetString(fieldValueStr)
	case reflect.Bool:
		fieldValue.SetBool(ElementContent2Bool(fieldValueStr))
	default:
		return errors.Errorf("unsupport struct field type: %v", fieldValue.Kind())
	}

	return nil
}

func (e *DpCreativeElements) getContentByName(ctx context.Context,
	elementName string, creativeElements []*storage.CreativeElement) (string, bool) {
	for _, each := range creativeElements {
		if each.ElementName == elementName {
			return each.Content, true
		}
	}

	return "", false
}
