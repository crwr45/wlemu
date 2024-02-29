package config

import (
	"gopkg.in/yaml.v3"
)

// https://github.com/go-yaml/yaml/issues/13#issuecomment-1586325414
type RawNode struct {
	*yaml.Node
}

func (n *RawNode) UnmarshalYAML(node *yaml.Node) error {
	n.Node = node
	return nil
}
