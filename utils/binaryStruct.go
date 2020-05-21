package utils

import (
	"bufio"
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"

	"github.com/recraft/recraft-lib/types"
)

//Get the original field name by tag
func getField(count string, rt reflect.Type) (fieldName string) {
	if rt.Kind() != reflect.Struct {
		panic("bad type")
	}
	for c := 0; c < rt.NumField(); c++ {
		field := rt.Field(c)
		if strings.Split(field.Tag.Get("rcount"), ",")[0] == count {
			return field.Name
		}
	}
	return ""
}

//BinaryToStruct converts main data types from the Minecraft protcols directly to any struct.
//The order can be know by tags.
//Note: This function depends on reflect, any error can give a panic.
func BinaryToStruct(structt interface{}, reader *bufio.Reader) error {
	//Get the original struct
	valueOf := reflect.ValueOf(structt).Elem()
	typeOf := reflect.TypeOf(structt).Elem()

	//Check if type is struct
	if valueOf.Kind() != reflect.Struct {
		return errors.New("bad type")
	}
	//Get fields by couting using tag(s)
	for rcount := 1; rcount <= valueOf.NumField(); rcount++ {

		//Get the name of the field using a tag
		field := getField(strconv.Itoa(rcount), typeOf)
		fieldSettings := valueOf.FieldByName(field)

		switch fmt.Sprint(fieldSettings.Type()) {
		case "types.String":
			var str types.String
			err := str.Read(reader)
			if err != nil {
				return err
			}
			fieldSettings.SetString(string(str))
			break
		case "types.Boolean":
			var b types.Boolean
			err := b.Read(reader)
			if err != nil {
				return err
			}
			fieldSettings.SetBool(bool(b))

			break
		case "types.Short":
			var s types.Short
			err := s.Read(reader)
			if err != nil {
				return err
			}
			fieldSettings.SetInt(int64(s))
			break
		case "types.UnsignedShort":
			var us types.UnsignedShort
			err := us.Read(reader)
			if err != nil {
				return err
			}
			fieldSettings.SetUint(uint64(us))
			break
		case "types.Int":
			var b types.Int
			err := b.Read(reader)
			if err != nil {
				return err
			}
			fieldSettings.SetInt(int64(b))
			break
		case "types.Long":
			var l types.Long
			err := l.Read(reader)
			if err != nil {
				return err
			}
			fieldSettings.SetInt(int64(l))
			break
		case "types.VarLong":
			var vl types.VarLong
			err := vl.Read(reader)
			if err != nil {
				return err
			}
			fieldSettings.SetInt(int64(vl))
			break
		case "types.VarInt":
			var vi types.VarInt
			err := vi.Read(reader)
			if err != nil {
				return err
			}
			fieldSettings.SetInt(int64(vi))
			break
		default:
			return errors.New("Type of rcount" + strconv.Itoa(rcount) + " not supported")

		}
	}

	return nil
}

//StructToBinary acts like BinaryToStruct but doing the opposite thing, converting a struct into binary
func StructToBinary(structt interface{}, packetID int32) ([]byte, error) {

	//Get the original struct
	valueOf := reflect.ValueOf(structt).Elem()
	typeOf := reflect.TypeOf(structt).Elem()

	//Check if type is struct
	if valueOf.Kind() != reflect.Struct {
		return []byte{}, errors.New("bad type")
	}
	buf := make([]byte, 0)
	buf = append(buf, types.VarInt(packetID).GetBytes()...)

	//Get fields by couting using tag(s)
	for rcount := 1; rcount <= valueOf.NumField(); rcount++ {
		field := getField(strconv.Itoa(rcount), typeOf)
		fieldSettings := valueOf.FieldByName(field)

		//var Type types.BasicTypes

		switch fmt.Sprint(fieldSettings.Type()) {
		case "types.String":
			s := fieldSettings.Interface().(types.String)
			buf = append(buf, s.GetBytes()...)

			break
		case "types.Boolean":
			b := fieldSettings.Interface().(types.Boolean)
			buf = append(buf, b.GetBytes()...)

			break
		case "types.Short":
			s := fieldSettings.Interface().(types.Short)
			buf = append(buf, s.GetBytes()...)

			break
		case "types.UnsignedShort":
			us := fieldSettings.Interface().(types.UnsignedShort)
			buf = append(buf, us.GetBytes()...)

			break
		case "types.Int":
			i := fieldSettings.Interface().(types.Int)
			buf = append(buf, i.GetBytes()...)

			break
		case "types.Long":
			l := fieldSettings.Interface().(types.Long)
			buf = append(buf, l.GetBytes()...)

			break
		case "types.VarInt":
			vi := fieldSettings.Interface().(types.VarInt)
			buf = append(buf, vi.GetBytes()...)

			break
		case "types.VarLong":
			vl := fieldSettings.Interface().(types.VarLong)
			buf = append(buf, vl.GetBytes()...)

			break

			/*case "types.Boolean":
				Type = types.Boolean(fieldSettings.Bool())
				break
			case "types.Int":
				Type = types.Int(fieldSettings.Int())
				break
			case "types.Long":
				Type = types.Long(fieldSettings.Int())
				break
			case "types.Short":
				Type = types.Short(fieldSettings.Int())
				break
			case "types.VarInt":
				Type = types.VarInt(fieldSettings.Int())
				break
			case "types.VarLong":
				Type = types.VarLong(fieldSettings.Int())
				break*/
		default:
			return []byte{}, errors.New("Type of rcount" + strconv.Itoa(rcount) + " not supported")

		}

	}

	binary := make([]byte, 0)
	binary = append(binary, types.VarInt(len(buf)).GetBytes()...)
	buf = append(binary, buf...)
	return buf, nil
}
