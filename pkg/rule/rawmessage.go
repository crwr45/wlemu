package rule

import (
	"gopkg.in/yaml.v3"
)

// postpone unmarshal of a YAML section so the correct type can be established
// https://github.com/go-yaml/yaml/issues/13#issuecomment-1586325414
type RawNode struct {
	*yaml.Node
}

func (n *RawNode) UnmarshalYAML(node *yaml.Node) error {
	n.Node = node
	return nil
}
