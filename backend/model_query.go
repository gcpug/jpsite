// generated by qbg -usedatastorewrapper -output model_query.go .; DO NOT EDIT

package backend

import (
	"go.mercari.io/datastore"
)

// Plugin supply hook point for query constructions.
type Plugin interface {
	Init(typeName string)
	Ancestor(ancestor datastore.Key)
	KeysOnly()
	Start(cur datastore.Cursor)
	Offset(offset int)
	Limit(limit int)
	Filter(name, op string, value interface{})
	Asc(name string)
	Desc(name string)
}

// Plugger supply Plugin component.
type Plugger interface {
	Plugin() Plugin
}

// OrganizationQueryBuilder build query for Organization.
type OrganizationQueryBuilder struct {
	q             datastore.Query
	plugin        Plugin
	Name          *OrganizationQueryProperty
	URL           *OrganizationQueryProperty
	LogoURL       *OrganizationQueryProperty
	Order         *OrganizationQueryProperty
	CreatedAt     *OrganizationQueryProperty
	UpdatedAt     *OrganizationQueryProperty
	SchemaVersion *OrganizationQueryProperty
}

// OrganizationQueryProperty has property information for OrganizationQueryBuilder.
type OrganizationQueryProperty struct {
	bldr *OrganizationQueryBuilder
	name string
}

// NewOrganizationQueryBuilder create new OrganizationQueryBuilder.
func NewOrganizationQueryBuilder(client datastore.Client) *OrganizationQueryBuilder {
	return NewOrganizationQueryBuilderWithKind(client, "Organization")
}

// NewOrganizationQueryBuilderWithKind create new OrganizationQueryBuilder with specific kind.
func NewOrganizationQueryBuilderWithKind(client datastore.Client, kind string) *OrganizationQueryBuilder {
	q := client.NewQuery(kind)
	bldr := &OrganizationQueryBuilder{q: q}
	bldr.Name = &OrganizationQueryProperty{
		bldr: bldr,
		name: "Name",
	}
	bldr.URL = &OrganizationQueryProperty{
		bldr: bldr,
		name: "URL",
	}
	bldr.LogoURL = &OrganizationQueryProperty{
		bldr: bldr,
		name: "LogoURL",
	}
	bldr.Order = &OrganizationQueryProperty{
		bldr: bldr,
		name: "Order",
	}
	bldr.CreatedAt = &OrganizationQueryProperty{
		bldr: bldr,
		name: "CreatedAt",
	}
	bldr.UpdatedAt = &OrganizationQueryProperty{
		bldr: bldr,
		name: "UpdatedAt",
	}
	bldr.SchemaVersion = &OrganizationQueryProperty{
		bldr: bldr,
		name: "SchemaVersion",
	}

	if plugger, ok := interface{}(bldr).(Plugger); ok {
		bldr.plugin = plugger.Plugin()
		bldr.plugin.Init("Organization")
	}

	return bldr
}

// Ancestor sets parent key to ancestor query.
func (bldr *OrganizationQueryBuilder) Ancestor(parentKey datastore.Key) *OrganizationQueryBuilder {
	bldr.q = bldr.q.Ancestor(parentKey)
	if bldr.plugin != nil {
		bldr.plugin.Ancestor(parentKey)
	}
	return bldr
}

// KeysOnly sets keys only option to query.
func (bldr *OrganizationQueryBuilder) KeysOnly() *OrganizationQueryBuilder {
	bldr.q = bldr.q.KeysOnly()
	if bldr.plugin != nil {
		bldr.plugin.KeysOnly()
	}
	return bldr
}

// Start setup to query.
func (bldr *OrganizationQueryBuilder) Start(cur datastore.Cursor) *OrganizationQueryBuilder {
	bldr.q = bldr.q.Start(cur)
	if bldr.plugin != nil {
		bldr.plugin.Start(cur)
	}
	return bldr
}

// Offset setup to query.
func (bldr *OrganizationQueryBuilder) Offset(offset int) *OrganizationQueryBuilder {
	bldr.q = bldr.q.Offset(offset)
	if bldr.plugin != nil {
		bldr.plugin.Offset(offset)
	}
	return bldr
}

// Limit setup to query.
func (bldr *OrganizationQueryBuilder) Limit(limit int) *OrganizationQueryBuilder {
	bldr.q = bldr.q.Limit(limit)
	if bldr.plugin != nil {
		bldr.plugin.Limit(limit)
	}
	return bldr
}

// Query returns *datastore.Query.
func (bldr *OrganizationQueryBuilder) Query() datastore.Query {
	return bldr.q
}

// Filter with op & value.
func (p *OrganizationQueryProperty) Filter(op string, value interface{}) *OrganizationQueryBuilder {
	switch op {
	case "<=":
		p.LessThanOrEqual(value)
	case ">=":
		p.GreaterThanOrEqual(value)
	case "<":
		p.LessThan(value)
	case ">":
		p.GreaterThan(value)
	case "=":
		p.Equal(value)
	default:
		p.bldr.q = p.bldr.q.Filter(p.name+" "+op, value) // error raised by native query
	}
	if p.bldr.plugin != nil {
		p.bldr.plugin.Filter(p.name, op, value)
	}
	return p.bldr
}

// LessThanOrEqual filter with value.
func (p *OrganizationQueryProperty) LessThanOrEqual(value interface{}) *OrganizationQueryBuilder {
	p.bldr.q = p.bldr.q.Filter(p.name+" <=", value)
	if p.bldr.plugin != nil {
		p.bldr.plugin.Filter(p.name, "<=", value)
	}
	return p.bldr
}

// GreaterThanOrEqual filter with value.
func (p *OrganizationQueryProperty) GreaterThanOrEqual(value interface{}) *OrganizationQueryBuilder {
	p.bldr.q = p.bldr.q.Filter(p.name+" >=", value)
	if p.bldr.plugin != nil {
		p.bldr.plugin.Filter(p.name, ">=", value)
	}
	return p.bldr
}

// LessThan filter with value.
func (p *OrganizationQueryProperty) LessThan(value interface{}) *OrganizationQueryBuilder {
	p.bldr.q = p.bldr.q.Filter(p.name+" <", value)
	if p.bldr.plugin != nil {
		p.bldr.plugin.Filter(p.name, "<", value)
	}
	return p.bldr
}

// GreaterThan filter with value.
func (p *OrganizationQueryProperty) GreaterThan(value interface{}) *OrganizationQueryBuilder {
	p.bldr.q = p.bldr.q.Filter(p.name+" >", value)
	if p.bldr.plugin != nil {
		p.bldr.plugin.Filter(p.name, ">", value)
	}
	return p.bldr
}

// Equal filter with value.
func (p *OrganizationQueryProperty) Equal(value interface{}) *OrganizationQueryBuilder {
	p.bldr.q = p.bldr.q.Filter(p.name+" =", value)
	if p.bldr.plugin != nil {
		p.bldr.plugin.Filter(p.name, "=", value)
	}
	return p.bldr
}

// Asc order.
func (p *OrganizationQueryProperty) Asc() *OrganizationQueryBuilder {
	p.bldr.q = p.bldr.q.Order(p.name)
	if p.bldr.plugin != nil {
		p.bldr.plugin.Asc(p.name)
	}
	return p.bldr
}

// Desc order.
func (p *OrganizationQueryProperty) Desc() *OrganizationQueryBuilder {
	p.bldr.q = p.bldr.q.Order("-" + p.name)
	if p.bldr.plugin != nil {
		p.bldr.plugin.Desc(p.name)
	}
	return p.bldr
}
