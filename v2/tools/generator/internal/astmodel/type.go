/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package astmodel

import (
	"fmt"
	"strings"

	"github.com/dave/dst"
)

// Type represents something that is a Go type
type Type interface {
	// RequiredPackageReferences returns a set of packages imports required by this type
	RequiredPackageReferences() *PackageReferenceSet

	// References returns the names of all types that this type
	// references. For example, an Array of Persons references a
	// Person.
	References() TypeNameSet

	// AsType renders as a Go abstract syntax tree for a type
	// (yes this says dst.Expr but that is what the Go 'dst' package uses for types)
	AsType(codeGenerationContext *CodeGenerationContext) dst.Expr

	// AsDeclarations renders as a Go abstract syntax tree for a declaration
	AsDeclarations(codeGenerationContext *CodeGenerationContext, declContext DeclarationContext) []dst.Decl

	// AsZero renders an expression for the "zero" value of the type
	// definitions allows TypeName to resolve to the underlying type
	// ctx allows current imports to be correctly identified where needed
	AsZero(definitions TypeDefinitionSet, ctx *CodeGenerationContext) dst.Expr

	// Equals returns true if the passed type is the same as this one, false otherwise
	Equals(t Type, overrides EqualityOverrides) bool

	// Make sure all TypeDefinitionSet have a printable version for debugging/user info.
	// This doesn't need to be a full representation of the type.
	fmt.Stringer

	// WriteDebugDescription adds a description of the current type to the passed builder
	// builder receives the full description, including nested types
	// definitions is a dictionary for resolving named types
	WriteDebugDescription(builder *strings.Builder, definitions TypeDefinitionSet)
}

type EqualityOverrides struct {
	TypeName func(left TypeName, right TypeName) bool
}

// IgnoringErrors returns the type stripped of any ErroredType wrapper
func IgnoringErrors(t Type) Type {
	if errored, ok := t.(*ErroredType); ok {
		return errored.InnerType()
	}

	return t
}

// DeclarationContext represents some metadata about a specific declaration
type DeclarationContext struct {
	Name        TypeName
	Description []string
	Validations []KubeBuilderValidation
}

// TypeEquals decides if the types are the same and handles the `nil` case
// overrides can be passed to combe
func TypeEquals(left, right Type, overrides ...EqualityOverrides) bool {
	if left == nil {
		return right == nil
	}

	override := EqualityOverrides{}
	if len(overrides) > 0 {
		if len(overrides) > 1 {
			panic("can only pass one EqualityOverrides")
		}

		override = overrides[0]
	}

	return left.Equals(right, override)
}

func DebugDescription(t Type, definitions TypeDefinitionSet) string {
	var builder strings.Builder
	t.WriteDebugDescription(&builder, definitions)
	return builder.String()
}
