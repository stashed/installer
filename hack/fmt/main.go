/*
Copyright AppsCode Inc. and Contributors

Licensed under the AppsCode Community License 1.0.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    https://github.com/appscode/licenses/raw/1.0.0/AppsCode-Community-1.0.0.md

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"bytes"
	"encoding/json"
	goflag "flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"stash.appscode.dev/installer/hack/fmt/templates"

	"github.com/Masterminds/sprig"
	flag "github.com/spf13/pflag"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"kmodules.xyz/client-go/tools/parser"
	"sigs.k8s.io/yaml"
)

type Location struct {
	App     string
	Version string
}

type TypedObject struct {
	Type     string // backup, restore
	Resource *unstructured.Unstructured
}

func main() {
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	flag.StringVar(&dir, "dir", dir, "Path to directory where the stashed/installer project resides (default is set o current directory)")

	flag.CommandLine.AddGoFlagSet(goflag.CommandLine)
	flag.Parse()

	resources, err := ListResources(filepath.Join(dir, "catalog", "raw"))
	if err != nil {
		panic(err)
	}

	store := map[Location][]TypedObject{}

	for _, obj := range resources {
		// remove labels
		obj.SetNamespace("")
		obj.SetLabels(nil)
		obj.SetAnnotations(nil)

		parts := strings.SplitN(obj.GetName(), "-", 3)

		loc := Location{
			App:     parts[0],
			Version: parts[2],
		}
		store[loc] = append(store[loc], TypedObject{
			Type:     parts[1],
			Resource: obj,
		})
	}

	for loc, objects := range store {
		for _, obj := range objects {
			data, err := yaml.Marshal(obj.Resource)
			if err != nil {
				panic(err)
			}

			filename := filepath.Join(dir, "catalog", "new_raw", loc.App, loc.Version, fmt.Sprintf("%s-%s-%s.yaml", loc.App, obj.Type, strings.ToLower(obj.Resource.GetKind())))
			err = os.MkdirAll(filepath.Dir(filename), 0755)
			if err != nil {
				panic(err)
			}

			err = ioutil.WriteFile(filename, data, 0644)
			if err != nil {
				panic(err)
			}
		}
	}

	{
		// move new_raw to raw
		err = os.RemoveAll(filepath.Join(dir, "catalog", "raw"))
		if err != nil {
			panic(err)
		}
		err = os.Rename(filepath.Join(dir, "catalog", "new_raw"), filepath.Join(dir, "catalog", "raw"))
		if err != nil {
			panic(err)
		}
	}

	{
		for loc, objects := range store {
			for _, r := range objects {
				if r.Resource.GetKind() == "Function" {
					img, _, err := unstructured.NestedString(r.Resource.Object, "spec", "image")
					if err != nil {
						panic(err)
					}
					if img != "" && strings.HasPrefix(img, "stashed/") {
						nuimg := "{{ .Values.image.registry }}/" + strings.TrimPrefix(img, "stashed/")
						err = unstructured.SetNestedField(r.Resource.Object, nuimg, "spec", "image")
						if err != nil {
							panic(err)
						}
					}

					args, _, err := unstructured.NestedStringSlice(r.Resource.Object, "spec", "args")
					if err != nil {
						panic(err)
					}

					for i := range args {
						if strings.HasPrefix(args[i], "--wait-timeout=") {
							args[i] = `--wait-timeout=${waitTimeout:={{ .Values.waitTimeout}}}`
						}

						switch loc.App {
						case "elasticsearch":
							if strings.HasPrefix(args[i], "--es-args=") {
								args[i] = fmt.Sprintf(`--es-args=${args:={{ .Values.%s.%s.args }}}`, loc.App, r.Type)
							}
						case "mariadb":
							if strings.HasPrefix(args[i], "--mariadb-args=") {
								args[i] = fmt.Sprintf(`--mariadb-args=${args:={{ .Values.%s.%s.args }}}`, loc.App, r.Type)
							}

						case "mongodb":
							if strings.HasPrefix(args[i], "--mongo-args=") {
								args[i] = fmt.Sprintf(`--mongo-args=${args:={{ .Values.%s.%s.args }}}`, loc.App, r.Type)
							}
							// --max-concurrency=${MAX_CONCURRENCY:={{ .Values.maxConcurrency}}}
							if strings.HasPrefix(args[i], "--max-concurrency=") {
								args[i] = fmt.Sprintf(`--max-concurrency=${MAX_CONCURRENCY:={{ .Values.%s.maxConcurrency}}}`, loc.App)
							}

						case "mysql":
							if strings.HasPrefix(args[i], "--mysql-args=") {
								args[i] = fmt.Sprintf(`--mysql-args=${args:={{ .Values.%s.%s.args }}}`, loc.App, r.Type)
							}

						case "perconaxtradb":
							if strings.HasPrefix(args[i], "--xtradb-args=") {
								args[i] = fmt.Sprintf(`--xtradb-args=${args:={{ .Values.%s.%s.args }}}`, loc.App, r.Type)
							}
							//
							if strings.HasPrefix(args[i], "--target-app-replicas=") {
								args[i] = fmt.Sprintf(`--target-app-replicas=${TARGET_APP_REPLICAS:={{ .Values.%s.restore.targetAppReplicas }}}`, loc.App)
							}

						case "postgres":
							if strings.HasPrefix(args[i], "--pg-args=") {
								args[i] = fmt.Sprintf(`--pg-args=${args:={{ .Values.%s.%s.args }}}`, loc.App, r.Type)
							}
						}
					}

					err = unstructured.SetNestedStringSlice(r.Resource.Object, args, "spec", "args")
					if err != nil {
						panic(err)
					}
				}

				var buf bytes.Buffer
				data := map[string]interface{}{
					"app":    loc.App,
					"object": r.Resource.Object,
				}
				funcMap := sprig.TxtFuncMap()
				funcMap["toYaml"] = toYAML
				funcMap["toJson"] = toJSON
				tpl := template.Must(template.New("").Funcs(funcMap).Parse(templates.Object))
				err = tpl.Execute(&buf, &data)
				if err != nil {
					panic(err)
				}

				filename := filepath.Join(dir, "charts", "stash-catalog", "new_templates", loc.App, loc.Version, fmt.Sprintf("%s-%s-%s.yaml", loc.App, r.Type, strings.ToLower(r.Resource.GetKind())))
				err = os.MkdirAll(filepath.Dir(filename), 0755)
				if err != nil {
					panic(err)
				}

				err = ioutil.WriteFile(filename, buf.Bytes(), 0644)
				if err != nil {
					panic(err)
				}
			}
		}

		{
			// move new_templates to templates
			err = os.Rename(filepath.Join(dir, "charts", "stash-catalog", "templates", "_helpers.tpl"), filepath.Join(dir, "charts", "stash-catalog", "new_templates", "_helpers.tpl"))
			if err != nil {
				panic(err)
			}
			err = os.RemoveAll(filepath.Join(dir, "charts", "stash-catalog", "templates"))
			if err != nil {
				panic(err)
			}
			err = os.Rename(filepath.Join(dir, "charts", "stash-catalog", "new_templates"), filepath.Join(dir, "charts", "stash-catalog", "templates"))
			if err != nil {
				panic(err)
			}
		}
	}
}

func ListResources(dir string) ([]*unstructured.Unstructured, error) {
	var resources []*unstructured.Unstructured

	err := parser.ProcessDir(dir, func(obj *unstructured.Unstructured) error {
		obj.SetNamespace("")
		resources = append(resources, obj)
		return nil
	})
	if err != nil {
		return nil, err
	}

	return resources, nil
}

// toYAML takes an interface, marshals it to yaml, and returns a string. It will
// always return a string, even on marshal error (empty string).
//
// This is designed to be called from a template.
func toYAML(v interface{}) string {
	data, err := yaml.Marshal(v)
	if err != nil {
		// Swallow errors inside of a template.
		return ""
	}
	return strings.TrimSuffix(string(data), "\n")
}

// toJSON takes an interface, marshals it to json, and returns a string. It will
// always return a string, even on marshal error (empty string).
//
// This is designed to be called from a template.
func toJSON(v interface{}) string {
	data, err := json.Marshal(v)
	if err != nil {
		// Swallow errors inside of a template.
		return ""
	}
	return string(data)
}
