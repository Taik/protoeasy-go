// Code generated by protoc-gen-go.
// source: pb/money/money.gen.proto
// DO NOT EDIT!

package pbmoney

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type CurrencyCode int32

const (
	CurrencyCode_CURRENCY_CODE_NONE CurrencyCode = 0
	CurrencyCode_CURRENCY_CODE_AED  CurrencyCode = 1
	CurrencyCode_CURRENCY_CODE_AFN  CurrencyCode = 2
	CurrencyCode_CURRENCY_CODE_ALL  CurrencyCode = 3
	CurrencyCode_CURRENCY_CODE_AMD  CurrencyCode = 4
	CurrencyCode_CURRENCY_CODE_ANG  CurrencyCode = 5
	CurrencyCode_CURRENCY_CODE_AOA  CurrencyCode = 6
	CurrencyCode_CURRENCY_CODE_ARS  CurrencyCode = 7
	CurrencyCode_CURRENCY_CODE_AUD  CurrencyCode = 8
	CurrencyCode_CURRENCY_CODE_AWG  CurrencyCode = 9
	CurrencyCode_CURRENCY_CODE_AZN  CurrencyCode = 10
	CurrencyCode_CURRENCY_CODE_BAM  CurrencyCode = 11
	CurrencyCode_CURRENCY_CODE_BBD  CurrencyCode = 12
	CurrencyCode_CURRENCY_CODE_BDT  CurrencyCode = 13
	CurrencyCode_CURRENCY_CODE_BGN  CurrencyCode = 14
	CurrencyCode_CURRENCY_CODE_BHD  CurrencyCode = 15
	CurrencyCode_CURRENCY_CODE_BIF  CurrencyCode = 16
	CurrencyCode_CURRENCY_CODE_BMD  CurrencyCode = 17
	CurrencyCode_CURRENCY_CODE_BND  CurrencyCode = 18
	CurrencyCode_CURRENCY_CODE_BOB  CurrencyCode = 19
	CurrencyCode_CURRENCY_CODE_BRL  CurrencyCode = 20
	CurrencyCode_CURRENCY_CODE_BSD  CurrencyCode = 21
	CurrencyCode_CURRENCY_CODE_BWP  CurrencyCode = 22
	CurrencyCode_CURRENCY_CODE_BYR  CurrencyCode = 23
	CurrencyCode_CURRENCY_CODE_BZD  CurrencyCode = 24
	CurrencyCode_CURRENCY_CODE_CAD  CurrencyCode = 25
	CurrencyCode_CURRENCY_CODE_CHF  CurrencyCode = 26
	CurrencyCode_CURRENCY_CODE_CLP  CurrencyCode = 27
	CurrencyCode_CURRENCY_CODE_CNY  CurrencyCode = 28
	CurrencyCode_CURRENCY_CODE_COP  CurrencyCode = 29
	CurrencyCode_CURRENCY_CODE_CRC  CurrencyCode = 30
	CurrencyCode_CURRENCY_CODE_CUP  CurrencyCode = 31
	CurrencyCode_CURRENCY_CODE_CVE  CurrencyCode = 32
	CurrencyCode_CURRENCY_CODE_CZK  CurrencyCode = 33
	CurrencyCode_CURRENCY_CODE_DJF  CurrencyCode = 34
	CurrencyCode_CURRENCY_CODE_DKK  CurrencyCode = 35
	CurrencyCode_CURRENCY_CODE_DOP  CurrencyCode = 36
	CurrencyCode_CURRENCY_CODE_DZD  CurrencyCode = 37
	CurrencyCode_CURRENCY_CODE_EGP  CurrencyCode = 38
	CurrencyCode_CURRENCY_CODE_ERN  CurrencyCode = 39
	CurrencyCode_CURRENCY_CODE_ETB  CurrencyCode = 40
	CurrencyCode_CURRENCY_CODE_EUR  CurrencyCode = 41
	CurrencyCode_CURRENCY_CODE_FJD  CurrencyCode = 42
	CurrencyCode_CURRENCY_CODE_FKP  CurrencyCode = 43
	CurrencyCode_CURRENCY_CODE_GBP  CurrencyCode = 44
	CurrencyCode_CURRENCY_CODE_GEL  CurrencyCode = 45
	CurrencyCode_CURRENCY_CODE_GHS  CurrencyCode = 46
	CurrencyCode_CURRENCY_CODE_GIP  CurrencyCode = 47
	CurrencyCode_CURRENCY_CODE_GMD  CurrencyCode = 48
	CurrencyCode_CURRENCY_CODE_GNF  CurrencyCode = 49
	CurrencyCode_CURRENCY_CODE_GTQ  CurrencyCode = 50
	CurrencyCode_CURRENCY_CODE_GYD  CurrencyCode = 51
	CurrencyCode_CURRENCY_CODE_HNL  CurrencyCode = 52
	CurrencyCode_CURRENCY_CODE_HRK  CurrencyCode = 53
	CurrencyCode_CURRENCY_CODE_HUF  CurrencyCode = 54
	CurrencyCode_CURRENCY_CODE_IDR  CurrencyCode = 55
	CurrencyCode_CURRENCY_CODE_ILS  CurrencyCode = 56
	CurrencyCode_CURRENCY_CODE_INR  CurrencyCode = 57
	CurrencyCode_CURRENCY_CODE_IQD  CurrencyCode = 58
	CurrencyCode_CURRENCY_CODE_IRR  CurrencyCode = 59
	CurrencyCode_CURRENCY_CODE_ISK  CurrencyCode = 60
	CurrencyCode_CURRENCY_CODE_JMD  CurrencyCode = 61
	CurrencyCode_CURRENCY_CODE_JOD  CurrencyCode = 62
	CurrencyCode_CURRENCY_CODE_JPY  CurrencyCode = 63
	CurrencyCode_CURRENCY_CODE_KES  CurrencyCode = 64
	CurrencyCode_CURRENCY_CODE_KGS  CurrencyCode = 65
	CurrencyCode_CURRENCY_CODE_KHR  CurrencyCode = 66
	CurrencyCode_CURRENCY_CODE_KMF  CurrencyCode = 67
	CurrencyCode_CURRENCY_CODE_KPW  CurrencyCode = 68
	CurrencyCode_CURRENCY_CODE_KRW  CurrencyCode = 69
	CurrencyCode_CURRENCY_CODE_KWD  CurrencyCode = 70
	CurrencyCode_CURRENCY_CODE_KYD  CurrencyCode = 71
	CurrencyCode_CURRENCY_CODE_KZT  CurrencyCode = 72
	CurrencyCode_CURRENCY_CODE_LAK  CurrencyCode = 73
	CurrencyCode_CURRENCY_CODE_LBP  CurrencyCode = 74
	CurrencyCode_CURRENCY_CODE_LKR  CurrencyCode = 75
	CurrencyCode_CURRENCY_CODE_LRD  CurrencyCode = 76
	CurrencyCode_CURRENCY_CODE_LYD  CurrencyCode = 77
	CurrencyCode_CURRENCY_CODE_MAD  CurrencyCode = 78
	CurrencyCode_CURRENCY_CODE_MDL  CurrencyCode = 79
	CurrencyCode_CURRENCY_CODE_MGA  CurrencyCode = 80
	CurrencyCode_CURRENCY_CODE_MKD  CurrencyCode = 81
	CurrencyCode_CURRENCY_CODE_MMK  CurrencyCode = 82
	CurrencyCode_CURRENCY_CODE_MNT  CurrencyCode = 83
	CurrencyCode_CURRENCY_CODE_MOP  CurrencyCode = 84
	CurrencyCode_CURRENCY_CODE_MRO  CurrencyCode = 85
	CurrencyCode_CURRENCY_CODE_MUR  CurrencyCode = 86
	CurrencyCode_CURRENCY_CODE_MVR  CurrencyCode = 87
	CurrencyCode_CURRENCY_CODE_MWK  CurrencyCode = 88
	CurrencyCode_CURRENCY_CODE_MXN  CurrencyCode = 89
	CurrencyCode_CURRENCY_CODE_MYR  CurrencyCode = 90
	CurrencyCode_CURRENCY_CODE_MZN  CurrencyCode = 91
	CurrencyCode_CURRENCY_CODE_NGN  CurrencyCode = 92
	CurrencyCode_CURRENCY_CODE_NIO  CurrencyCode = 93
	CurrencyCode_CURRENCY_CODE_NOK  CurrencyCode = 94
	CurrencyCode_CURRENCY_CODE_NPR  CurrencyCode = 95
	CurrencyCode_CURRENCY_CODE_NZD  CurrencyCode = 96
	CurrencyCode_CURRENCY_CODE_OMR  CurrencyCode = 97
	CurrencyCode_CURRENCY_CODE_PEN  CurrencyCode = 98
	CurrencyCode_CURRENCY_CODE_PGK  CurrencyCode = 99
	CurrencyCode_CURRENCY_CODE_PHP  CurrencyCode = 100
	CurrencyCode_CURRENCY_CODE_PKR  CurrencyCode = 101
	CurrencyCode_CURRENCY_CODE_PLN  CurrencyCode = 102
	CurrencyCode_CURRENCY_CODE_PYG  CurrencyCode = 103
	CurrencyCode_CURRENCY_CODE_QAR  CurrencyCode = 104
	CurrencyCode_CURRENCY_CODE_RON  CurrencyCode = 105
	CurrencyCode_CURRENCY_CODE_RSD  CurrencyCode = 106
	CurrencyCode_CURRENCY_CODE_RUB  CurrencyCode = 107
	CurrencyCode_CURRENCY_CODE_RWF  CurrencyCode = 108
	CurrencyCode_CURRENCY_CODE_SAR  CurrencyCode = 109
	CurrencyCode_CURRENCY_CODE_SBD  CurrencyCode = 110
	CurrencyCode_CURRENCY_CODE_SCR  CurrencyCode = 111
	CurrencyCode_CURRENCY_CODE_SDG  CurrencyCode = 112
	CurrencyCode_CURRENCY_CODE_SEK  CurrencyCode = 113
	CurrencyCode_CURRENCY_CODE_SGD  CurrencyCode = 114
	CurrencyCode_CURRENCY_CODE_SHP  CurrencyCode = 115
	CurrencyCode_CURRENCY_CODE_SLL  CurrencyCode = 116
	CurrencyCode_CURRENCY_CODE_SOS  CurrencyCode = 117
	CurrencyCode_CURRENCY_CODE_SRD  CurrencyCode = 118
	CurrencyCode_CURRENCY_CODE_SSP  CurrencyCode = 119
	CurrencyCode_CURRENCY_CODE_STD  CurrencyCode = 120
	CurrencyCode_CURRENCY_CODE_SYP  CurrencyCode = 121
	CurrencyCode_CURRENCY_CODE_SZL  CurrencyCode = 122
	CurrencyCode_CURRENCY_CODE_THB  CurrencyCode = 123
	CurrencyCode_CURRENCY_CODE_TJS  CurrencyCode = 124
	CurrencyCode_CURRENCY_CODE_TMT  CurrencyCode = 125
	CurrencyCode_CURRENCY_CODE_TND  CurrencyCode = 126
	CurrencyCode_CURRENCY_CODE_TOP  CurrencyCode = 127
	CurrencyCode_CURRENCY_CODE_TRY  CurrencyCode = 128
	CurrencyCode_CURRENCY_CODE_TTD  CurrencyCode = 129
	CurrencyCode_CURRENCY_CODE_TZS  CurrencyCode = 130
	CurrencyCode_CURRENCY_CODE_UAH  CurrencyCode = 131
	CurrencyCode_CURRENCY_CODE_UGX  CurrencyCode = 132
	CurrencyCode_CURRENCY_CODE_USD  CurrencyCode = 133
	CurrencyCode_CURRENCY_CODE_UYU  CurrencyCode = 134
	CurrencyCode_CURRENCY_CODE_UZS  CurrencyCode = 135
	CurrencyCode_CURRENCY_CODE_VEF  CurrencyCode = 136
	CurrencyCode_CURRENCY_CODE_VND  CurrencyCode = 137
	CurrencyCode_CURRENCY_CODE_VUV  CurrencyCode = 138
	CurrencyCode_CURRENCY_CODE_WST  CurrencyCode = 139
	CurrencyCode_CURRENCY_CODE_XAF  CurrencyCode = 140
	CurrencyCode_CURRENCY_CODE_XCD  CurrencyCode = 141
	CurrencyCode_CURRENCY_CODE_XOF  CurrencyCode = 142
	CurrencyCode_CURRENCY_CODE_XPF  CurrencyCode = 143
	CurrencyCode_CURRENCY_CODE_YER  CurrencyCode = 144
	CurrencyCode_CURRENCY_CODE_ZAR  CurrencyCode = 145
	CurrencyCode_CURRENCY_CODE_ZMW  CurrencyCode = 146
	CurrencyCode_CURRENCY_CODE_ZWL  CurrencyCode = 147
)

var CurrencyCode_name = map[int32]string{
	0:   "CURRENCY_CODE_NONE",
	1:   "CURRENCY_CODE_AED",
	2:   "CURRENCY_CODE_AFN",
	3:   "CURRENCY_CODE_ALL",
	4:   "CURRENCY_CODE_AMD",
	5:   "CURRENCY_CODE_ANG",
	6:   "CURRENCY_CODE_AOA",
	7:   "CURRENCY_CODE_ARS",
	8:   "CURRENCY_CODE_AUD",
	9:   "CURRENCY_CODE_AWG",
	10:  "CURRENCY_CODE_AZN",
	11:  "CURRENCY_CODE_BAM",
	12:  "CURRENCY_CODE_BBD",
	13:  "CURRENCY_CODE_BDT",
	14:  "CURRENCY_CODE_BGN",
	15:  "CURRENCY_CODE_BHD",
	16:  "CURRENCY_CODE_BIF",
	17:  "CURRENCY_CODE_BMD",
	18:  "CURRENCY_CODE_BND",
	19:  "CURRENCY_CODE_BOB",
	20:  "CURRENCY_CODE_BRL",
	21:  "CURRENCY_CODE_BSD",
	22:  "CURRENCY_CODE_BWP",
	23:  "CURRENCY_CODE_BYR",
	24:  "CURRENCY_CODE_BZD",
	25:  "CURRENCY_CODE_CAD",
	26:  "CURRENCY_CODE_CHF",
	27:  "CURRENCY_CODE_CLP",
	28:  "CURRENCY_CODE_CNY",
	29:  "CURRENCY_CODE_COP",
	30:  "CURRENCY_CODE_CRC",
	31:  "CURRENCY_CODE_CUP",
	32:  "CURRENCY_CODE_CVE",
	33:  "CURRENCY_CODE_CZK",
	34:  "CURRENCY_CODE_DJF",
	35:  "CURRENCY_CODE_DKK",
	36:  "CURRENCY_CODE_DOP",
	37:  "CURRENCY_CODE_DZD",
	38:  "CURRENCY_CODE_EGP",
	39:  "CURRENCY_CODE_ERN",
	40:  "CURRENCY_CODE_ETB",
	41:  "CURRENCY_CODE_EUR",
	42:  "CURRENCY_CODE_FJD",
	43:  "CURRENCY_CODE_FKP",
	44:  "CURRENCY_CODE_GBP",
	45:  "CURRENCY_CODE_GEL",
	46:  "CURRENCY_CODE_GHS",
	47:  "CURRENCY_CODE_GIP",
	48:  "CURRENCY_CODE_GMD",
	49:  "CURRENCY_CODE_GNF",
	50:  "CURRENCY_CODE_GTQ",
	51:  "CURRENCY_CODE_GYD",
	52:  "CURRENCY_CODE_HNL",
	53:  "CURRENCY_CODE_HRK",
	54:  "CURRENCY_CODE_HUF",
	55:  "CURRENCY_CODE_IDR",
	56:  "CURRENCY_CODE_ILS",
	57:  "CURRENCY_CODE_INR",
	58:  "CURRENCY_CODE_IQD",
	59:  "CURRENCY_CODE_IRR",
	60:  "CURRENCY_CODE_ISK",
	61:  "CURRENCY_CODE_JMD",
	62:  "CURRENCY_CODE_JOD",
	63:  "CURRENCY_CODE_JPY",
	64:  "CURRENCY_CODE_KES",
	65:  "CURRENCY_CODE_KGS",
	66:  "CURRENCY_CODE_KHR",
	67:  "CURRENCY_CODE_KMF",
	68:  "CURRENCY_CODE_KPW",
	69:  "CURRENCY_CODE_KRW",
	70:  "CURRENCY_CODE_KWD",
	71:  "CURRENCY_CODE_KYD",
	72:  "CURRENCY_CODE_KZT",
	73:  "CURRENCY_CODE_LAK",
	74:  "CURRENCY_CODE_LBP",
	75:  "CURRENCY_CODE_LKR",
	76:  "CURRENCY_CODE_LRD",
	77:  "CURRENCY_CODE_LYD",
	78:  "CURRENCY_CODE_MAD",
	79:  "CURRENCY_CODE_MDL",
	80:  "CURRENCY_CODE_MGA",
	81:  "CURRENCY_CODE_MKD",
	82:  "CURRENCY_CODE_MMK",
	83:  "CURRENCY_CODE_MNT",
	84:  "CURRENCY_CODE_MOP",
	85:  "CURRENCY_CODE_MRO",
	86:  "CURRENCY_CODE_MUR",
	87:  "CURRENCY_CODE_MVR",
	88:  "CURRENCY_CODE_MWK",
	89:  "CURRENCY_CODE_MXN",
	90:  "CURRENCY_CODE_MYR",
	91:  "CURRENCY_CODE_MZN",
	92:  "CURRENCY_CODE_NGN",
	93:  "CURRENCY_CODE_NIO",
	94:  "CURRENCY_CODE_NOK",
	95:  "CURRENCY_CODE_NPR",
	96:  "CURRENCY_CODE_NZD",
	97:  "CURRENCY_CODE_OMR",
	98:  "CURRENCY_CODE_PEN",
	99:  "CURRENCY_CODE_PGK",
	100: "CURRENCY_CODE_PHP",
	101: "CURRENCY_CODE_PKR",
	102: "CURRENCY_CODE_PLN",
	103: "CURRENCY_CODE_PYG",
	104: "CURRENCY_CODE_QAR",
	105: "CURRENCY_CODE_RON",
	106: "CURRENCY_CODE_RSD",
	107: "CURRENCY_CODE_RUB",
	108: "CURRENCY_CODE_RWF",
	109: "CURRENCY_CODE_SAR",
	110: "CURRENCY_CODE_SBD",
	111: "CURRENCY_CODE_SCR",
	112: "CURRENCY_CODE_SDG",
	113: "CURRENCY_CODE_SEK",
	114: "CURRENCY_CODE_SGD",
	115: "CURRENCY_CODE_SHP",
	116: "CURRENCY_CODE_SLL",
	117: "CURRENCY_CODE_SOS",
	118: "CURRENCY_CODE_SRD",
	119: "CURRENCY_CODE_SSP",
	120: "CURRENCY_CODE_STD",
	121: "CURRENCY_CODE_SYP",
	122: "CURRENCY_CODE_SZL",
	123: "CURRENCY_CODE_THB",
	124: "CURRENCY_CODE_TJS",
	125: "CURRENCY_CODE_TMT",
	126: "CURRENCY_CODE_TND",
	127: "CURRENCY_CODE_TOP",
	128: "CURRENCY_CODE_TRY",
	129: "CURRENCY_CODE_TTD",
	130: "CURRENCY_CODE_TZS",
	131: "CURRENCY_CODE_UAH",
	132: "CURRENCY_CODE_UGX",
	133: "CURRENCY_CODE_USD",
	134: "CURRENCY_CODE_UYU",
	135: "CURRENCY_CODE_UZS",
	136: "CURRENCY_CODE_VEF",
	137: "CURRENCY_CODE_VND",
	138: "CURRENCY_CODE_VUV",
	139: "CURRENCY_CODE_WST",
	140: "CURRENCY_CODE_XAF",
	141: "CURRENCY_CODE_XCD",
	142: "CURRENCY_CODE_XOF",
	143: "CURRENCY_CODE_XPF",
	144: "CURRENCY_CODE_YER",
	145: "CURRENCY_CODE_ZAR",
	146: "CURRENCY_CODE_ZMW",
	147: "CURRENCY_CODE_ZWL",
}
var CurrencyCode_value = map[string]int32{
	"CURRENCY_CODE_NONE": 0,
	"CURRENCY_CODE_AED":  1,
	"CURRENCY_CODE_AFN":  2,
	"CURRENCY_CODE_ALL":  3,
	"CURRENCY_CODE_AMD":  4,
	"CURRENCY_CODE_ANG":  5,
	"CURRENCY_CODE_AOA":  6,
	"CURRENCY_CODE_ARS":  7,
	"CURRENCY_CODE_AUD":  8,
	"CURRENCY_CODE_AWG":  9,
	"CURRENCY_CODE_AZN":  10,
	"CURRENCY_CODE_BAM":  11,
	"CURRENCY_CODE_BBD":  12,
	"CURRENCY_CODE_BDT":  13,
	"CURRENCY_CODE_BGN":  14,
	"CURRENCY_CODE_BHD":  15,
	"CURRENCY_CODE_BIF":  16,
	"CURRENCY_CODE_BMD":  17,
	"CURRENCY_CODE_BND":  18,
	"CURRENCY_CODE_BOB":  19,
	"CURRENCY_CODE_BRL":  20,
	"CURRENCY_CODE_BSD":  21,
	"CURRENCY_CODE_BWP":  22,
	"CURRENCY_CODE_BYR":  23,
	"CURRENCY_CODE_BZD":  24,
	"CURRENCY_CODE_CAD":  25,
	"CURRENCY_CODE_CHF":  26,
	"CURRENCY_CODE_CLP":  27,
	"CURRENCY_CODE_CNY":  28,
	"CURRENCY_CODE_COP":  29,
	"CURRENCY_CODE_CRC":  30,
	"CURRENCY_CODE_CUP":  31,
	"CURRENCY_CODE_CVE":  32,
	"CURRENCY_CODE_CZK":  33,
	"CURRENCY_CODE_DJF":  34,
	"CURRENCY_CODE_DKK":  35,
	"CURRENCY_CODE_DOP":  36,
	"CURRENCY_CODE_DZD":  37,
	"CURRENCY_CODE_EGP":  38,
	"CURRENCY_CODE_ERN":  39,
	"CURRENCY_CODE_ETB":  40,
	"CURRENCY_CODE_EUR":  41,
	"CURRENCY_CODE_FJD":  42,
	"CURRENCY_CODE_FKP":  43,
	"CURRENCY_CODE_GBP":  44,
	"CURRENCY_CODE_GEL":  45,
	"CURRENCY_CODE_GHS":  46,
	"CURRENCY_CODE_GIP":  47,
	"CURRENCY_CODE_GMD":  48,
	"CURRENCY_CODE_GNF":  49,
	"CURRENCY_CODE_GTQ":  50,
	"CURRENCY_CODE_GYD":  51,
	"CURRENCY_CODE_HNL":  52,
	"CURRENCY_CODE_HRK":  53,
	"CURRENCY_CODE_HUF":  54,
	"CURRENCY_CODE_IDR":  55,
	"CURRENCY_CODE_ILS":  56,
	"CURRENCY_CODE_INR":  57,
	"CURRENCY_CODE_IQD":  58,
	"CURRENCY_CODE_IRR":  59,
	"CURRENCY_CODE_ISK":  60,
	"CURRENCY_CODE_JMD":  61,
	"CURRENCY_CODE_JOD":  62,
	"CURRENCY_CODE_JPY":  63,
	"CURRENCY_CODE_KES":  64,
	"CURRENCY_CODE_KGS":  65,
	"CURRENCY_CODE_KHR":  66,
	"CURRENCY_CODE_KMF":  67,
	"CURRENCY_CODE_KPW":  68,
	"CURRENCY_CODE_KRW":  69,
	"CURRENCY_CODE_KWD":  70,
	"CURRENCY_CODE_KYD":  71,
	"CURRENCY_CODE_KZT":  72,
	"CURRENCY_CODE_LAK":  73,
	"CURRENCY_CODE_LBP":  74,
	"CURRENCY_CODE_LKR":  75,
	"CURRENCY_CODE_LRD":  76,
	"CURRENCY_CODE_LYD":  77,
	"CURRENCY_CODE_MAD":  78,
	"CURRENCY_CODE_MDL":  79,
	"CURRENCY_CODE_MGA":  80,
	"CURRENCY_CODE_MKD":  81,
	"CURRENCY_CODE_MMK":  82,
	"CURRENCY_CODE_MNT":  83,
	"CURRENCY_CODE_MOP":  84,
	"CURRENCY_CODE_MRO":  85,
	"CURRENCY_CODE_MUR":  86,
	"CURRENCY_CODE_MVR":  87,
	"CURRENCY_CODE_MWK":  88,
	"CURRENCY_CODE_MXN":  89,
	"CURRENCY_CODE_MYR":  90,
	"CURRENCY_CODE_MZN":  91,
	"CURRENCY_CODE_NGN":  92,
	"CURRENCY_CODE_NIO":  93,
	"CURRENCY_CODE_NOK":  94,
	"CURRENCY_CODE_NPR":  95,
	"CURRENCY_CODE_NZD":  96,
	"CURRENCY_CODE_OMR":  97,
	"CURRENCY_CODE_PEN":  98,
	"CURRENCY_CODE_PGK":  99,
	"CURRENCY_CODE_PHP":  100,
	"CURRENCY_CODE_PKR":  101,
	"CURRENCY_CODE_PLN":  102,
	"CURRENCY_CODE_PYG":  103,
	"CURRENCY_CODE_QAR":  104,
	"CURRENCY_CODE_RON":  105,
	"CURRENCY_CODE_RSD":  106,
	"CURRENCY_CODE_RUB":  107,
	"CURRENCY_CODE_RWF":  108,
	"CURRENCY_CODE_SAR":  109,
	"CURRENCY_CODE_SBD":  110,
	"CURRENCY_CODE_SCR":  111,
	"CURRENCY_CODE_SDG":  112,
	"CURRENCY_CODE_SEK":  113,
	"CURRENCY_CODE_SGD":  114,
	"CURRENCY_CODE_SHP":  115,
	"CURRENCY_CODE_SLL":  116,
	"CURRENCY_CODE_SOS":  117,
	"CURRENCY_CODE_SRD":  118,
	"CURRENCY_CODE_SSP":  119,
	"CURRENCY_CODE_STD":  120,
	"CURRENCY_CODE_SYP":  121,
	"CURRENCY_CODE_SZL":  122,
	"CURRENCY_CODE_THB":  123,
	"CURRENCY_CODE_TJS":  124,
	"CURRENCY_CODE_TMT":  125,
	"CURRENCY_CODE_TND":  126,
	"CURRENCY_CODE_TOP":  127,
	"CURRENCY_CODE_TRY":  128,
	"CURRENCY_CODE_TTD":  129,
	"CURRENCY_CODE_TZS":  130,
	"CURRENCY_CODE_UAH":  131,
	"CURRENCY_CODE_UGX":  132,
	"CURRENCY_CODE_USD":  133,
	"CURRENCY_CODE_UYU":  134,
	"CURRENCY_CODE_UZS":  135,
	"CURRENCY_CODE_VEF":  136,
	"CURRENCY_CODE_VND":  137,
	"CURRENCY_CODE_VUV":  138,
	"CURRENCY_CODE_WST":  139,
	"CURRENCY_CODE_XAF":  140,
	"CURRENCY_CODE_XCD":  141,
	"CURRENCY_CODE_XOF":  142,
	"CURRENCY_CODE_XPF":  143,
	"CURRENCY_CODE_YER":  144,
	"CURRENCY_CODE_ZAR":  145,
	"CURRENCY_CODE_ZMW":  146,
	"CURRENCY_CODE_ZWL":  147,
}

func (x CurrencyCode) String() string {
	return proto.EnumName(CurrencyCode_name, int32(x))
}
func (CurrencyCode) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func init() {
	proto.RegisterEnum("pb.money.CurrencyCode", CurrencyCode_name, CurrencyCode_value)
}

func init() { proto.RegisterFile("pb/money/money.gen.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 773 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x6c, 0xd7, 0x57, 0x76, 0xdd, 0x54,
	0x18, 0x47, 0x71, 0x42, 0x49, 0x31, 0x01, 0x76, 0x2e, 0xc4, 0x84, 0xd0, 0x3b, 0x04, 0x70, 0x80,
	0xd0, 0xbb, 0x74, 0x3f, 0x49, 0xd7, 0x96, 0x74, 0x74, 0x7c, 0x24, 0x59, 0x96, 0x28, 0x06, 0x27,
	0x26, 0xb4, 0xd8, 0xc6, 0x24, 0x80, 0xe9, 0xbd, 0x77, 0x06, 0xca, 0x14, 0x58, 0xcb, 0x3c, 0xb1,
	0xf4, 0x7f, 0xd9, 0x0f, 0xbf, 0x19, 0xec, 0xb9, 0x63, 0xdb, 0xeb, 0x27, 0xcf, 0x6d, 0x6d, 0x6e,
	0xec, 0xfe, 0xd7, 0x85, 0xb3, 0x1b, 0x9b, 0x0b, 0xdb, 0x3b, 0x5b, 0xe7, 0xb7, 0x26, 0x07, 0xb7,
	0xd7, 0x17, 0xf6, 0xec, 0xc4, 0x3f, 0xc7, 0xe7, 0x0e, 0x4f, 0x2f, 0xec, 0xec, 0x6c, 0x6c, 0x9e,
	0xde, 0x9d, 0x6e, 0x9d, 0xd9, 0x98, 0xcc, 0xcf, 0x4d, 0xa6, 0x6d, 0x08, 0x89, 0x9b, 0xf6, 0x6b,
	0xd3, 0xca, 0x92, 0x35, 0x57, 0xb9, 0x84, 0x8b, 0x26, 0x47, 0xe7, 0x8e, 0xfc, 0xdf, 0xa3, 0xc4,
	0xd8, 0x27, 0x38, 0x75, 0x5c, 0x2c, 0xb8, 0x28, 0xb8, 0x44, 0x70, 0x69, 0x5c, 0x2a, 0xd8, 0x65,
	0x5c, 0x26, 0xb8, 0x8a, 0xd8, 0x2f, 0x38, 0xd4, 0x1c, 0x10, 0xdc, 0x1a, 0x07, 0x05, 0x77, 0x19,
	0x87, 0x04, 0x0f, 0x8e, 0xb9, 0x31, 0xc7, 0x51, 0xc9, 0xe5, 0x82, 0x63, 0xe3, 0xb0, 0x60, 0x6b,
	0xb8, 0x42, 0x70, 0xe6, 0xb8, 0x52, 0xf0, 0xcc, 0xb8, 0x4a, 0xf0, 0x62, 0x0a, 0x82, 0x4b, 0xe3,
	0x88, 0x60, 0x67, 0x4c, 0x04, 0x57, 0x31, 0x57, 0x0b, 0x0e, 0x05, 0xd7, 0x08, 0xae, 0x8d, 0xa3,
	0x82, 0x3b, 0xcf, 0xbc, 0xe0, 0x3e, 0x70, 0xad, 0xe0, 0xc1, 0x38, 0x36, 0xe6, 0x69, 0x64, 0x5c,
	0x27, 0x78, 0x96, 0x72, 0x5c, 0x70, 0xe1, 0xb9, 0x5e, 0xb0, 0xeb, 0xb9, 0x41, 0x70, 0xe5, 0xb9,
	0x51, 0x70, 0x98, 0x72, 0x93, 0xe0, 0xd6, 0x73, 0xb3, 0xe0, 0x95, 0x84, 0x5b, 0x04, 0x0f, 0x39,
	0xb7, 0x8e, 0xd9, 0x96, 0x52, 0x6e, 0x13, 0x9c, 0xe7, 0xdc, 0x2e, 0xb8, 0xf2, 0xdc, 0x21, 0x78,
	0x30, 0xee, 0x1c, 0x73, 0x92, 0x79, 0xee, 0x12, 0x1c, 0x1c, 0x77, 0x0b, 0x6e, 0x62, 0xee, 0x11,
	0xdc, 0x06, 0xee, 0x1d, 0x73, 0xba, 0x64, 0x9c, 0x10, 0x9c, 0x7b, 0xee, 0x1b, 0x73, 0x16, 0x7b,
	0xee, 0x17, 0x9c, 0x14, 0x3c, 0x20, 0x78, 0x56, 0xb3, 0x20, 0x78, 0xd1, 0x73, 0x52, 0x70, 0x69,
	0x3c, 0x28, 0xd8, 0xa5, 0x3c, 0x24, 0xb8, 0x59, 0xe6, 0x61, 0xc1, 0xbd, 0x71, 0x6a, 0xcc, 0x33,
	0x57, 0xf0, 0x88, 0xe0, 0x90, 0xf3, 0xa8, 0xe0, 0x36, 0xe5, 0xb1, 0x31, 0x2f, 0x5a, 0xe0, 0x71,
	0xc1, 0x45, 0xcd, 0x13, 0x82, 0x5d, 0xe0, 0x49, 0xc1, 0xcb, 0xc6, 0x53, 0x82, 0x43, 0xe0, 0x69,
	0xc1, 0x75, 0xce, 0x33, 0x63, 0x5e, 0x2a, 0x8d, 0x67, 0x05, 0x57, 0xc6, 0x73, 0x82, 0x7d, 0xcf,
	0xf3, 0x63, 0xce, 0x93, 0x9a, 0x17, 0x04, 0x67, 0x35, 0x91, 0xe0, 0x59, 0x20, 0x16, 0x5c, 0xa6,
	0x4c, 0x05, 0xfb, 0x0e, 0x13, 0x1c, 0x3a, 0x12, 0xc1, 0x9d, 0x91, 0x0a, 0xee, 0x8d, 0x4c, 0xf0,
	0xd0, 0x30, 0x1b, 0x73, 0x11, 0xe5, 0x2c, 0x0a, 0x8e, 0x3d, 0x4b, 0x82, 0xf3, 0x40, 0x2e, 0x38,
	0x18, 0x85, 0xe0, 0xde, 0x28, 0xc7, 0x5c, 0x46, 0x86, 0x13, 0x6c, 0x05, 0x95, 0xe0, 0x2c, 0xc2,
	0x0b, 0xce, 0x8d, 0x65, 0xc1, 0x65, 0x4e, 0x10, 0xec, 0x1a, 0x6a, 0xc1, 0x95, 0xa7, 0x11, 0x1c,
	0x2a, 0x5a, 0xc1, 0x6d, 0x60, 0x45, 0xf0, 0x4a, 0xa0, 0x13, 0xdc, 0xe5, 0xac, 0x0a, 0x5e, 0x75,
	0xf4, 0x82, 0xfb, 0xc0, 0x20, 0x78, 0x70, 0xbc, 0x38, 0x66, 0x97, 0x39, 0x5e, 0x12, 0xbc, 0x58,
	0xf1, 0xb2, 0xe0, 0x2a, 0xe7, 0x15, 0xc1, 0x3e, 0xb0, 0x26, 0x78, 0x30, 0x5e, 0x1d, 0x73, 0x55,
	0x06, 0x5e, 0x1b, 0xb3, 0x4f, 0x1c, 0xeb, 0x82, 0xb3, 0x9c, 0xd3, 0x82, 0x67, 0x9e, 0x33, 0x82,
	0xf3, 0xc0, 0x86, 0xe0, 0xc2, 0xf1, 0xba, 0xe0, 0x3e, 0xe3, 0xec, 0x98, 0x97, 0xa3, 0xc0, 0x1b,
	0x63, 0x0e, 0x95, 0xe3, 0x4d, 0xc1, 0xb5, 0xf1, 0x96, 0xe0, 0x36, 0xe6, 0x6d, 0xc1, 0x5d, 0xca,
	0x3b, 0x63, 0xae, 0xa3, 0xc0, 0x39, 0xc1, 0xb1, 0xb1, 0x29, 0x78, 0x1a, 0xd8, 0x12, 0x6c, 0x19,
	0xdb, 0x82, 0x93, 0x9c, 0x77, 0x05, 0x67, 0xc6, 0x8e, 0xe0, 0x99, 0xe7, 0x3d, 0xc1, 0x45, 0xc1,
	0x79, 0xc1, 0x55, 0xcd, 0x05, 0xc1, 0xc1, 0x78, 0x5f, 0x70, 0xed, 0xf9, 0x40, 0x70, 0x63, 0x7c,
	0x28, 0xb8, 0xf7, 0xec, 0x0a, 0x1e, 0x0a, 0x3e, 0x1a, 0x73, 0x33, 0x8b, 0xf9, 0x58, 0xf0, 0x52,
	0xcd, 0x27, 0x82, 0xcb, 0x86, 0x4f, 0x05, 0x3b, 0xe3, 0x33, 0xc1, 0x95, 0xe7, 0xf3, 0xc9, 0xfc,
	0x88, 0x43, 0xcf, 0x17, 0xfb, 0x84, 0x37, 0xc6, 0x97, 0xca, 0x87, 0x9a, 0xaf, 0x84, 0xb7, 0xd1,
	0x8c, 0xaf, 0x95, 0x67, 0xab, 0x7c, 0xa3, 0xbc, 0x36, 0xbe, 0x55, 0xde, 0xb7, 0x7c, 0xa7, 0x7c,
	0xa8, 0xf9, 0x5e, 0xf8, 0x4a, 0x92, 0xf2, 0x83, 0x72, 0x67, 0xfc, 0xa8, 0xbc, 0x5d, 0xe1, 0x27,
	0xe1, 0x5d, 0xdd, 0xf0, 0xb3, 0xf0, 0xd5, 0x28, 0xe5, 0x17, 0xe5, 0x53, 0xe3, 0x57, 0xe5, 0x55,
	0xca, 0x6f, 0xca, 0x7d, 0xca, 0xef, 0xc2, 0xfb, 0x24, 0xf0, 0x87, 0xf0, 0x21, 0x0a, 0xfc, 0xa9,
	0xbc, 0xec, 0xf8, 0x4b, 0x79, 0x57, 0xf0, 0xf7, 0xbe, 0xf8, 0xd0, 0x70, 0x60, 0x7b, 0x7d, 0x6f,
	0xbe, 0xd6, 0xf7, 0xef, 0xdd, 0xd8, 0xa9, 0x7f, 0x03, 0x00, 0x00, 0xff, 0xff, 0x92, 0xad, 0x8e,
	0xff, 0xa9, 0x0d, 0x00, 0x00,
}