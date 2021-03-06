// This file was automatically generated by informer-gen

package v1alpha1

import (
	internalinterfaces "github.com/kanisterio/kanister/pkg/client/informers/externalversions/internalinterfaces"
)

// Interface provides access to all the informers in this group version.
type Interface interface {
	// ActionSets returns a ActionSetInformer.
	ActionSets() ActionSetInformer
	// Blueprints returns a BlueprintInformer.
	Blueprints() BlueprintInformer
}

type version struct {
	factory          internalinterfaces.SharedInformerFactory
	namespace        string
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// New returns a new Interface.
func New(f internalinterfaces.SharedInformerFactory, namespace string, tweakListOptions internalinterfaces.TweakListOptionsFunc) Interface {
	return &version{factory: f, namespace: namespace, tweakListOptions: tweakListOptions}
}

// ActionSets returns a ActionSetInformer.
func (v *version) ActionSets() ActionSetInformer {
	return &actionSetInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// Blueprints returns a BlueprintInformer.
func (v *version) Blueprints() BlueprintInformer {
	return &blueprintInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}
