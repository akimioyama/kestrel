package tcp

import (
	"encoding/binary"
	"errors"
)

// Encoder собирает фреймы протокола в байты.
type Encoder struct{}

// Ошибки, которые возвращает Encoder при сборке фрейма.
var (
	ErrIncompleteErrorFields = errors.New("error code and error message must be set together")
)

// Encode собирает фрейм целиком:
// [ фиксированный заголовок | HeaderLength | блок полей Header | Payload ].
func (encoder *Encoder) Encode(frame Frame) ([]byte, error) {
	var result []byte

	headerArray, err := encodeHeaderArray(frame.Header)
	if err != nil {
		return nil, err
	}

	header := encodeHeader(frame, headerArray)

	result = append(result, header...)
	result = binary.BigEndian.AppendUint32(result, uint32(len(headerArray)))
	result = append(result, headerArray...)
	result = append(result, frame.Payload...)

	return result, nil

}

// encodeHeaderArray собирает блок полей Header.
// Каждое поле: [ FieldID (1 байт) | FieldLength (2 байта) | Value (FieldLength байт) ]
func encodeHeaderArray(header Header) ([]byte, error) {
	var result []byte

	if header.CorrelationID != nil {
		result = append(result, fieldCorrelationID)
		result = binary.BigEndian.AppendUint16(result, correlationIDSize)
		result = binary.BigEndian.AppendUint64(result, *header.CorrelationID)
	}

	if header.Topic != nil {
		value := []byte(*header.Topic)

		result = append(result, fieldTopic)
		result = binary.BigEndian.AppendUint16(result, uint16(len(value)))
		result = append(result, value...)
	}

	if header.Key != nil {
		result = append(result, fieldKey)
		result = binary.BigEndian.AppendUint16(result, uint16(len(header.Key)))
		result = append(result, header.Key...)
	}

	if header.MessageID != nil {
		result = append(result, fieldMessageID)
		result = binary.BigEndian.AppendUint16(result, messageIDSize)
		result = binary.BigEndian.AppendUint64(result, *header.MessageID)
	}

	if (header.ErrorCode == nil) != (header.ErrorMessage == nil) {
		return nil, ErrIncompleteErrorFields
	}

	if header.ErrorCode != nil {
		result = append(result, fieldErrorCode)
		result = binary.BigEndian.AppendUint16(result, errorCodeSize)
		result = append(result, *header.ErrorCode)

		value := []byte(*header.ErrorMessage)
		result = append(result, fieldErrorMessage)
		result = binary.BigEndian.AppendUint16(result, uint16(len(value)))
		result = append(result, value...)
	}

	return result, nil
}

// encodeHeader собирает 8 байт фиксированного заголовка: Magic, Version, Type, Length.
// Length — длина всего тела: HeaderLength + блок полей Header + Payload.
func encodeHeader(frame Frame, headerArray []byte) []byte {
	buffer := make([]byte, fixedHeaderSize)

	binary.BigEndian.PutUint16(buffer[0:2], magicCode)

	buffer[2] = versionCode
	buffer[3] = frame.Type

	length := uint32(headerLengthSize + len(headerArray) + len(frame.Payload))
	binary.BigEndian.PutUint32(buffer[4:8], length)

	return buffer
}
