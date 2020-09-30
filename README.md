# YAML Extractor

This is a simple tool to extract inline chunks of yaml inside markdown
documents.

## Usage

The tool takes in a markdown file `M.md` as input, and outputs N yaml files
named `M-0.yml` through `M-N.yml` where N is the number of labeled, fenced, yaml
codeblocks inside the markdown.

### Examples

To use:
```bash
go build -o yamlextractor
./yamlextractor README.md
```

Fenced blocks of yaml that are labeled will be turned into `.yml` files:
```yaml
---
name: example-1
value:
  dog: bark
```
It will work if the inline code contains ytt as well
```yaml
#@ load("@ytt:overlay", "overlay")
#@ load("@ytt:json", "json")


#@ def modify_configmap(data):
#@   decoded = json.decode(data)
#@   decoded["animal"] = {
#@     "dog": {
#@       "sound": "bark"
#@     }
#@   }
#@   return json.encode(decoded)
#@ end

---
#@overlay/replace via=lambda a,_: modify_configmap(a)
value:
```
It will not pick up any codeblocks that aren't labeled as `yaml`, like this:
```
name: example-2
value:
 ignored: true
```

It will also ignore other labeled fenced blocks like this:
```bash
echo hello
```

If you were to run the program with this README as input, the output would be
two `.yml` files, try it out for yourself!
