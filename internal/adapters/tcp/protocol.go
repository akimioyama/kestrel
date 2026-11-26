package tcp

// Формат фрейма.
const (
	magicCode   = 0xBEEF // первые 2 байта каждого фрейма
	versionCode = 1      // версия протокола

	fixedHeaderSize  = 8 // Magic (2) + Version (1) + Type (1) + Length (4)
	headerLengthSize = 4 // сколько байт занимает само поле HeaderLength в начале тела
	fieldPrefixSize  = 3 // FieldID (1 байт) + FieldLength (2 байта)

	minLength       = 4        // тело содержит хотя бы поле HeaderLength
	maxLength       = 10 << 20 // максимальная длина тела, 10 МБ
	maxHeaderLength = 64 << 10 // максимальная длина блока полей Header, 64 КБ
)

// Размеры значений у полей Header с фиксированной длиной.
const (
	correlationIDSize = 8 // uint64
	messageIDSize     = 8 // uint64
	errorCodeSize     = 1 // uint8
)

// FieldID полей Header.
const (
	fieldTopic         uint8 = 0x01
	fieldCorrelationID uint8 = 0x02
	fieldKey           uint8 = 0x03
	fieldErrorCode     uint8 = 0x04
	fieldErrorMessage  uint8 = 0x05
	fieldMessageID     uint8 = 0x06
)
