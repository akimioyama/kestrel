package tcp

import (
	"encoding/binary"
	"errors"
	"io"
)

// Decoder читает фреймы протокола из io.Reader.
type Decoder struct {
	reader io.Reader
}

const (
	fixedHeaderSize  = 8 // Magic + Version + Type + Length
	magicCode        = 0xBEEF
	versionCode      = 1
	minLength        = 4
	maxLength        = 10 << 20
	maxHeaderLength  = 64 << 10
	headerLengthSize = 4 // сколько байт занимает само поле HeaderLength в начале тела
	fieldPrefixSize  = 3 // FieldID (1 байт) + FieldLength (2 байта)
)

// Ошибки, которые возвращает Decoder при разборе фрейма.
var (
	ErrInvalidMagic        = errors.New("invalid magic")
	ErrInvalidVersion       = errors.New("invalid version")
	ErrLengthTooSmall       = errors.New("length too small")
	ErrLengthTooBig         = errors.New("length too big")
	ErrHeaderLengthTooBig   = errors.New("header length too big")
	ErrHeaderFieldTruncated = errors.New("header field truncated")
	ErrInvalidFieldSize     = errors.New("invalid header field size")
)

// NewDecoder создаёт Decoder, который читает из r.
func NewDecoder(r io.Reader) *Decoder {
	return &Decoder{
		reader: r,
	}
}

// Decode читает из потока один фрейм целиком и разбирает его.
func (d *Decoder) Decode() (Frame, error) {
	frameType, length, err := d.readFixedHeader()
	if err != nil {
		return Frame{}, err
	}

	headerBytes, payload, err := d.readBody(length)
	if err != nil {
		return Frame{}, err
	}

	header, err := parseHeader(headerBytes)
	if err != nil {
		return Frame{}, err
	}

	return Frame{
		Type:    frameType,
		Header:  header,
		Payload: payload,
	}, nil
}

// readFixedHeader читает 8 байт фиксированного заголовка,
// проверяет Magic и Version и возвращает Type и Length.
func (d *Decoder) readFixedHeader() (uint8, uint32, error) {
	buffer := make([]byte, fixedHeaderSize)

	_, err := io.ReadFull(d.reader, buffer)
	if err != nil {
		return 0, 0, err
	}

	magic := binary.BigEndian.Uint16(buffer[0:2])
	if magic != magicCode {
		return 0, 0, ErrInvalidMagic
	}

	version := buffer[2]
	if version != versionCode {
		return 0, 0, ErrInvalidVersion
	}

	frameType := buffer[3]
	length := binary.BigEndian.Uint32(buffer[4:8])

	return frameType, length, nil
}

// readBody проверяет Length, читает тело и делит его на Header и Payload.
func (d *Decoder) readBody(length uint32) ([]byte, []byte, error) {
	if length < minLength {
		return nil, nil, ErrLengthTooSmall
	}
	if length > maxLength {
		return nil, nil, ErrLengthTooBig
	}

	body := make([]byte, length)

	_, err := io.ReadFull(d.reader, body)
	if err != nil {
		return nil, nil, err
	}

	headerLength := binary.BigEndian.Uint32(body[0:headerLengthSize])
	if headerLength > maxHeaderLength {
		return nil, nil, ErrHeaderLengthTooBig
	}
	// Header должен влезть в тело после поля HeaderLength.
	if headerLength > length-headerLengthSize {
		return nil, nil, ErrHeaderLengthTooBig
	}

	// Тело: [ HeaderLength (4 байта) | Header (headerLength байт) | Payload (остаток) ]
	headerEnd := headerLengthSize + headerLength
	headerBytes := body[headerLengthSize:headerEnd]
	payload := body[headerEnd:]

	return headerBytes, payload, nil
}

// parseHeader разбирает Header на поля.
// Каждая запись: [ FieldID (1 байт) | FieldLength (2 байта) | Value (FieldLength байт) ]
func parseHeader(data []byte) (Header, error) {
	var header Header

	for len(data) > 0 {
		if len(data) < fieldPrefixSize {
			return Header{}, ErrHeaderFieldTruncated
		}

		fieldID := data[0]
		fieldLength := int(binary.BigEndian.Uint16(data[1:3]))
		data = data[fieldPrefixSize:]

		if len(data) < fieldLength {
			return Header{}, ErrHeaderFieldTruncated
		}

		value := data[:fieldLength]
		data = data[fieldLength:]

		err := parseField(&header, fieldID, value)
		if err != nil {
			return Header{}, err
		}
	}

	return header, nil
}

// parseField кладёт значение одного поля в header.
// Если поле пришло дважды, остаётся последнее значение.
func parseField(header *Header, fieldID uint8, value []byte) error {
	switch fieldID {
	case fieldTopic:
		topic := string(value)
		header.Topic = &topic

	case fieldCorrelationID:
		if len(value) != 8 {
			return ErrInvalidFieldSize
		}
		correlationID := binary.BigEndian.Uint64(value)
		header.CorrelationID = &correlationID

	case fieldKey:
		header.Key = value

	default:
		// Неизвестный FieldID по спеке пропускаем.
		// ErrorCode, ErrorMessage и MessageID тоже сюда: их шлёт только сервер.
	}

	return nil
}
