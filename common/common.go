// common includes several misc methods and structs that are shared between different resources and endpoints
package common

type MALItem interface {
	GetID() int64
	GetType() string
}
