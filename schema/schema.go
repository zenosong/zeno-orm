package schema

import (
	"go/ast"
	"reflect"
	"zenoorm/dialect"
)

type Field struct {
	Name string
	Type string
	Tag  string
}

type Schema struct {
	Model      interface{}
	Name       string
	Fields     []*Field
	FieldNames []string
	fieldMap   map[string]*Field
}

func (schema *Schema) GetField(name string) *Field {
	return schema.fieldMap[name]
}

// 使用反射将模型实例解析为Schema实例
func Parse(dest interface{}, d dialect.Dialect) *Schema {
	// desc 入参为实例，需要ValueOf获取值（指针）后再用 Indirect 获取对应结构体
	modelType := reflect.Indirect(reflect.ValueOf(dest)).Type()
	schema := &Schema{
		Model:    dest,
		Name:     modelType.Name(),
		fieldMap: make(map[string]*Field),
	}
	// 遍历模型结构体字段
	for i := 0; i < modelType.NumField(); i++ {
		p := modelType.Field(i)
		// 非匿名且可导出字段，装入Schema
		if !p.Anonymous && ast.IsExported(p.Name) {
			field := &Field{
				Name: p.Name,
				Type: d.DataTypeOf(reflect.Indirect(reflect.New(p.Type))),
			}
			// 获取字段Tag
			if v, ok := p.Tag.Lookup("zenoorm"); ok {
				field.Tag = v
			}

			schema.Fields = append(schema.Fields, field)
			schema.FieldNames = append(schema.FieldNames, p.Name)
			schema.fieldMap[p.Name] = field
		}
	}

	return schema
}
