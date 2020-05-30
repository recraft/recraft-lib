package types

import (
	"bufio"
	"errors"
)

// BasicTypes Interface
type BasicTypes interface {
	GetBytes() []byte
	Read(buffer *bufio.Reader) error
}

// ReadBytes from buffer
func ReadBytes(reader *bufio.Reader, n int) ([]byte, error) {
	bytes := make([]byte, n)
	var err error
	for i := 0; i < n; i++ {

		bytes[i], err = reader.ReadByte()
		if err != nil {
			return []byte{}, err
		}
	}
	return bytes, nil
}

type (
	// String type
	String string
	// Boolean type. 0x01 if true, 0x00 if false.
	Boolean bool
	// Byte type
	Byte int8
	// UnsignedByte type
	UnsignedByte uint8
	// Short type
	Short int16
	// UnsignedShort type
	UnsignedShort uint16
	// Int type
	Int int32
	// Long type
	Long int64
	// Float type
	Float float32
	// Double type
	Double float64
	// PacketID - Internal type for PacketID enums
	PacketID []byte
	// VarInt special type
	VarInt int32
	// VarLong special type
	VarLong int64
)

// GetBytes converts String to byte slice
func (str *String) GetBytes() []byte {
	// create a new byte slice from the string
	stringBytes := []byte(*str)
	// return a byte slice that contains the length and the string bytes
	return append(VarInt(len(stringBytes)).GetBytes(), stringBytes...)
}

// Read a String from buffer
func (str *String) Read(buffer *bufio.Reader) error {
	// Get the string length from the buffer and check for errors
	length := new(VarInt)
	errorLength := length.Read(buffer)
	if errorLength != nil {
		return errorLength
	}

	// Get (length) bytes from buffer and check for errors
	stringBytes, errorRead := ReadBytes(buffer, int(*length))
	if errorRead != nil {
		return errorRead
	}

	// Convert bytes to string
	*str = String(string(stringBytes))
	// Return error nil
	return nil
}

// GetBytes converts Boolean to byte slice
func (b *Boolean) GetBytes() []byte {
	// If b is true, return 0x01, else if it's false return 0x00
	if *b {
		return []byte{0x01}
	}

	return []byte{0x00}
}

// Read a Boolean from buffer
func (b *Boolean) Read(buffer *bufio.Reader) error {
	// Read a byte from the buffer and check for errors
	bufByte, errorRead := buffer.ReadByte()

	if errorRead != nil {
		return errorRead
	}

	// Check if byte is a valid Boolean
	if !(bufByte == 0x00 || bufByte == 0x01) {
		return errors.New("invalid Boolean byte")
	}

	// Convert byte to boolean and assign it to b
	*b = bufByte != 0x0
	// Return error nil
	return nil
}

// GetBytes converts Short to byte slice
func (s *Short) GetBytes() []byte {
	// Divide int16 into two bytes, using right shift and casting the result to byte (same as uint8)
	return []byte{
		byte(*s >> 8), byte(*s),
	}
}

// Read a Short from buffer
func (s *Short) Read(buffer *bufio.Reader) error {
	// Read two bytes from buffer and check for errors
	bytesRead, errorRead := ReadBytes(buffer, 2)
	if errorRead != nil {
		return errorRead
	}

	// Convert two bytes to int16 and assign it to s
	*s = Short(int16(bytesRead[0])<<8 | int16(bytesRead[1]))
	// Return error nil
	return nil
}

// GetBytes converts Unsigned Short to byte slice
func (us *UnsignedShort) GetBytes() []byte {
	// Divide uint16 into two bytes, using right shift and casting the result to byte (same as uint8)
	return []byte{
		byte(*us >> 8), byte(*us),
	}
}

// Read an unsigned Short from buffer
func (us *UnsignedShort) Read(buffer *bufio.Reader) error {
	// Read two bytes from buffer and check for errors
	bytesRead, err := ReadBytes(buffer, 2)
	if err != nil {
		return err
	}

	// Convert two bytes to uint16 and assign it to us
	*us = UnsignedShort(uint16(bytesRead[0])<<8 | uint16(bytesRead[1]))
	// Return error nil
	return nil
}

// GetBytes converts Int to byte slice
func (i *Int) GetBytes() []byte {
	// Divide int32 into four bytes, using right shift and casting the result to byte (same as uint8)
	return []byte{
		byte(*i >> 24), byte(*i >> 16), byte(*i >> 8), byte(*i),
	}
}

// Read an Int from buffer
func (i *Int) Read(buffer *bufio.Reader) error {
	// Read four bytes from buffer and check for errors
	bytesRead, errorRead := ReadBytes(buffer, 4)
	if errorRead != nil {
		return errorRead
	}

	// Convert four bytes to int32 and assign it to i
	*i = Int(int32(bytesRead[0])<<24 | int32(bytesRead[1])<<16 | int32(bytesRead[2])<<8 | int32(bytesRead[3]))
	// Return error nil
	return nil
}

// GetBytes converts Long to byte slice
func (l *Long) GetBytes() []byte {
	// Divide int64 into four bytes, using right shift and casting the result to byte (same as uint8)
	return []byte{
		byte(*l >> 56), byte(*l >> 48), byte(*l >> 40), byte(*l >> 32),
		byte(*l >> 24), byte(*l >> 16), byte(*l >> 8), byte(*l),
	}
}

// Read a Long from buffer
func (l *Long) Read(buffer *bufio.Reader) error {
	// Read eight bytes from buffer and check for errors
	bytesRead, errorRead := ReadBytes(buffer, 8)
	if errorRead != nil {
		return errorRead
	}

	// Convert eight bytes to int64 and assign it to l
	*l = Long(int64(bytesRead[0])<<56 | int64(bytesRead[1])<<48 | int64(bytesRead[2])<<40 | int64(bytesRead[3])<<32 |
		int64(bytesRead[4])<<24 | int64(bytesRead[5])<<16 | int64(bytesRead[6])<<8 | int64(bytesRead[7]))
	// Return error nil
	return nil
}

// GetBytes converts VarInt to byte slice - vi is passed by value
func (vi VarInt) GetBytes() []byte {
	// Create an empty byte slice to fill
	result := make([]byte, 0)
	uvi := uint32(vi)

	// Fill the slice
	for {
		// temp: Take last byte of vi and make a bitwise AND with 7F (127)
		// &0x7F --> After this, the byte will have a value between 0 and 127 (2^7)
		temp := uvi & 0x7F
		// shift vi to the right by seven positions
		uvi >>= 7

		// If this isn't the last operation, make a bitwise OR of temp variable with 80.
		// In the previous step, we have seen that the number will be between 0 and 127, so certainly less than 0x80 (128).
		// = In this step we will add 128 (0x80) to the number.
		// Therefore, in the end, we will have a number between 128 and 255, else if we are in the last iteration,
		// it will be a number between 0 and 127.
		if uvi != 0 {
			temp |= 0x80
		}

		// Append the obtained byte to result
		result = append(result, byte(temp))

		// If there aren't anymore values, exit from cycle
		if uvi == 0 {
			break
		}
	}

	// Return byte slice result
	return result
}

// Read a VarInt from buffer
func (vi *VarInt) Read(buffer *bufio.Reader) error {
	// Create an empty uint32 variable to store the result
	var result uint32

	// Count the iterations number
	for i := 0; ; i++ {
		// Read a byte from buffer and check for errors
		byteRead, errorRead := buffer.ReadByte()
		if errorRead != nil {
			return errorRead
		}

		// Add to result the byte read ANDed with 7F and shifted left by (iteration number)*7 (reverse operation of GetBytes())
		result |= uint32(byteRead&0x7F) << uint32(7*i)

		// If VarInt is too big (max is 5 bytes), return an error
		if i >= 5 {
			return errors.New("VarInt is too big")
		}

		// If this is the last iteration (number is less than 128), exit from the cycle
		if byteRead&0x80 == 0 {
			break
		}
	}

	// Assign result to vi
	*vi = VarInt(result)
	// Return error nil
	return nil
}

// GetBytes converts VarLong to byte slice
func (vl VarLong) GetBytes() []byte {
	// Create an empty byte slice to fill
	result := make([]byte, 0)
	uvl := uint64(vl)

	// Fill slice
	for {
		// temp: Take last byte of vl and make a bitwise AND with 7F (127)
		temp := uvl & 0x7F
		// shift vl to the right by seven positions
		uvl >>= 7

		// If this isn't the last operation, make a bitwise OR on temp variable with 80
		if uvl != 0 {
			temp |= 0x80
		}

		// Append temp byte to result
		result = append(result, byte(temp))

		// If there aren't anymore values, exit from cycle
		if uvl == 0 {
			break
		}
	}

	// Return byte slice result
	return result
}

// Read a VarLong from buffer
func (vl *VarLong) Read(buffer *bufio.Reader) error {
	// Create an empty uint64 variable to store the result
	var result uint64

	// Count the iterations number
	for i := 0; ; i++ {
		// Read a byte from buffer and check for errors
		byteRead, errorRead := buffer.ReadByte()
		if errorRead != nil {
			return errorRead
		}

		// Add to result the byte read ANDed with 7F and shifted left by (iteration number)*7 (reverse operation of GetBytes())
		result |= uint64(byteRead&0x7F) << uint64(7*i)

		// If VarLong is too big (max is 10 bytes), return an error
		if i >= 10 {
			return errors.New("VarLong is too big")
		}

		// If this is the last iteration, exit from the cycle
		if byteRead&0x80 == 0 {
			break
		}
	}

	// Assign result to vl
	*vl = VarLong(result)
	// Return error nil
	return nil
}
