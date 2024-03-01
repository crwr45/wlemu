package main

import "github.com/crwr45/wlemu/cmd"

func main() {
	cmd.Execute()
}

// TODO:
/*
Output to k8s-compatible YAML
	prepend any extra resources the updater creates to the file
Finalise Naming
	some package names might read better verb-ised?
Unit Tests
	most packages should be quite simple
Handle more types of input resource
	pod-based ones first
Make rule-matching more flexible
	regex, wildcards, labels, etc.
	not just container names
more types of Updater
	do a non-trivial case next
consider a split between container-type Updaters and non-container-type Updaters
consider implications of multiple uses against the same resource with different rulesets
	chaining this way is limited but could add flexibility in some cases
*/
