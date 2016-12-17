package schema

import (
	"github.com/elazarl/go-bindata-assetfs"
	"github.com/xeipuuv/gojsonschema"
)

var AssetFS = &assetfs.AssetFS{Asset: Asset, AssetDir: AssetDir, AssetInfo: AssetInfo}

func LoadSchema(source string) (*gojsonschema.Schema, error) {
	schemaLoader := NewBindataAssetLoader(source)
	return gojsonschema.NewSchema(schemaLoader)
}

func MustLoadSchema(source string) *gojsonschema.Schema {
	schema, err := LoadSchema(source)
	if err != nil {
		panic(err)
	}
	return schema
}

func NewBindataAssetLoader(source string) gojsonschema.JSONLoader {
	return gojsonschema.NewReferenceLoaderFileSystem(source, AssetFS)
}
