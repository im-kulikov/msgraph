package msgraph

import (
	"fmt"
	"net/url"
	"strings"
)

type GraphQueryOption interface {
	Name() string
	Value() string
}

type GraphQuery struct {
	options []GraphQueryOption
}

func (q *GraphQuery) AddOption(option GraphQueryOption) *GraphQuery {
	q.options = append(q.options, option)
	return q
}

func (q GraphQuery) String() string {
	var opts url.Values
	for _, val := range q.options {
		opts.Add(val.Name(), val.Value())
	}
	return opts.Encode()
}

// SelectQuery is used to specify a subset of properties to return.
type SelectQuery struct {
	Properties []string
}

func (q *SelectQuery) AddProperty(property string) {
	q.Properties = append(q.Properties, property)
}

func (q SelectQuery) Name() string {
	return "$select"
}

func (q SelectQuery) Value() string {
	return strings.Join(q.Properties, ",")
}

// NewSelectQuery creates and returns a new SelectQuery.
func NewSelectQuery(properties []string) SelectQuery {
	return SelectQuery{properties}
}

/*
type ExpandQuery commaSeparatedPropertyList

func NewExpandQuery(properties []string) ExpandQuery {
	return ExpandQuery{properties}
}
*/

// OrderByDirection specifies the sorting direction for an OrderBy
// query.
type OrderByDirection int

const (
	// Sort results in ascending order.
	OrderByAscending OrderByDirection = iota

	// Sort results in descending order.
	OrderByDescending
)

// OrderByQuery specifies the sort order of items returned from the API,
// and which properties should be used to sort the results.
type OrderByQuery struct {
	Properties map[string]OrderByDirection
}

func (q OrderByQuery) Name() string {
	return "$orderby"
}

func (q OrderByQuery) Value() string {
	var result []string
	for k, v := range q.Properties {
		result = append(result, fmt.Sprintf("%s %s", k, v))
	}
	return strings.Join(result, ",")
}

func (q *OrderByQuery) AddProperty(property string, direction OrderByDirection) {
}

/*
// NewOrderByQuery creates and returns a new OrderByQuery.
func NewOrderByQuery(properties []string, direction OrderByDirection) OrderByQuery {
	return OrderByQuery{properties, direction}
}
*/

type FilterQuery struct{}

type LogicalFilterOperator string

func (l LogicalFilterOperator) String() string {
	return string(l)
}

const (
	AndFilter = "and"
	OrFilter  = "or"
)

type LogicalFilterQuery struct {
	Operator LogicalFilterOperator
	Queries  []FilterQuery
}

func (q LogicalFilterQuery) String() string {
	var result []string
	for _, query := range q.Queries {
		result = append(result, fmt.Sprintf("%s", query))
	}
	return fmt.Sprintf("(%s)", strings.Join(result, q.Operator.String()))
}

// TopQuery specifies the maximum number of items to return in a result
// set.
type TopQuery int

func (q TopQuery) Name() string {
	return "$top"
}
func (q TopQuery) Value() string {
	return fmt.Sprintf("%d", q)
}

// SkipQuery specifies the number of items to skip.
type SkipQuery int

func (q SkipQuery) Name() string {
	return "$skip"
}
func (q SkipQuery) Value() string {
	return fmt.Sprintf("%d", q)
}

// SkipToken is an opaque string used to specify the next set of results
// to return during a paged query operation.
type SkipToken string

func (q SkipToken) Name() string {
	return "$skipToken"
}
func (q SkipToken) Value() string {
	return string(q)
}

// CountQuery specifies whether to return the number of items in a
// collection as well as the items in the collection.
type CountQuery bool

func (q CountQuery) Name() string {
	return "$count"
}
func (q CountQuery) Value() string {
	return fmt.Sprintf("%d", q)
}
