//Copyright (c) 2015, Neil Moore
//See LICENSE.txt for license information.

package formats

type Format interface {
	MakeWith(str string, loc string) (string, error)
}
