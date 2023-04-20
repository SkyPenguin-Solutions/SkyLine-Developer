/////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//
//
//                              _____ _       __    _
//                             |   __| |_ _ _|  |  |_|___ ___
//                             |__   | '_| | |  |__| |   | -_|
//                             |_____|_,_|_  |_____|_|_|_|___|
//                                       |___|
//
// These sections are to help yopu better understand what each section is or what each file represents within the SkyLine programming language. These sections can also
//
// help seperate specific values so you can better understand what a specific section or specific set of values of functions is doing.
//
// These sections also give information on the file, project and status of the section
//
//
// :::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
//
// Filename      |  SkyLine_Script_Language_Backend_Models.go
// Project       |  SkyLine programming language
// Line Count    |  800+ active lines
// Status        |  Working and Active
// Package       |  SkyLine_Backend
//
//
// Defines       | This file defines all of the data structures, type structures, type aliases, constant lists, sub functiosn for data types, SLC_Object types, object methods
//                 flag types, settings, configuration files, token types, token keywords, regex patterns, signal patterns, byte/string/bin arrays, results etc for the SkyLine
//                 prime systems. These systems include but are not limited to the Parser, Scanner, Evaluator, Executor, Engine, Scripter, Writer, Reader, Filler, environment, etc.
//
//
package SkyLine_Backend

import (
	"flag"
	"strings"
	"unicode/utf8"
)

///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//
//
//																 ┏━┓
//																┃┃ ┃
//															━━━━┛
//
// This next section breaks down all of the tokenizer/scanner based data types or data structures which help the scanner categorize specific tokens or keywords
//
//
const (
	REGISTER        = "REGISTER"
	ILLEGAL         = "ILLEGAL"                                // Illegal character  | Implemented
	EOF             = "EOF"                                    // End Of File        | Implemented
	IDENT           = "IDENT"                                  // Identifier         | Implimented
	INT             = "INT"                                    // TYPE integer       | Implemented
	FLOAT           = "FLOAT"                                  // TYPE float         | Implemented
	STRING          = "STRING"                                 // TYPE string        | Implemented
	CONSTANT        = "CONST"                                  // Constant           | Implemented
	FUNCTION        = "FUNCTION"                               // Function           | Implemented
	LET             = "LET"                                    // let statement      | Implemented
	TRUE            = "TRUE"                                   // boolean type true  | Implemented
	FALSE           = "FALSE"                                  // boolean type false | Implemented
	IF              = "IF"                                     // If statement       | Implemented
	ELSE            = "ELSE"                                   // Else statement     | Implemented
	RETURN          = "RETURN"                                 // return statement   | Implemented
	SWITCH          = "SWITCH"                                 // Switch statement   | Implemented
	CASE            = "CASE"                                   // Case statement 	 | Implemented
	LTEQ            = "<="                                     // LT or equal to     | Implemented
	GTEQ            = ">="                                     // GT or equal to     | Implemented
	ASTERISK_EQUALS = "*="                                     // Multiply equals    | Implemented
	BANG            = "!"                                      // Boolean operator   | Implemented
	ASSIGN          = "="                                      // General assignment | Implemented
	PLUS            = "+"                                      // General operator   | Implemented
	MINUS           = "-"                                      // General operator   | Implemented
	ASTARISK        = "*"                                      // General operator   | Implemented
	SLASH           = "/"                                      // General operator   | Implemented
	LT              = "<"                                      // Boolean operator   | Implemented
	GT              = ">"                                      // Boolean operator   | Implemented
	EQ              = "=="                                     // Boolean operator   | Implemented
	MINUS_EQUALS    = "-="                                     // Minus equals       | Implemented
	NEQ             = "!="                                     // Boolean operator   | Implemented
	DIVEQ           = "/="                                     // Division operator  | Implemented
	PERIOD          = "."                                      // Method Call        | Implemented
	PLUS_EQUALS     = "+="                                     // Plus equals        | Implemented
	COMMA           = ","                                      // Seperation         | Implemented
	SEMICOLON       = ";"                                      // SemiColon          | Implemented
	COLON           = ":"                                      // Colon              | Implemented
	LPAREN          = "("                                      // Args start         | Implemented
	RPAREN          = ")"                                      // Args end           | Implemented
	LINE            = "|"                                      // Line con           | Implemented
	LBRACE          = "{"                                      // Open  f            | Implemented
	RBRACE          = "}"                                      // Close f            | Implemented
	LBRACKET        = "["                                      // Open               | Implemented
	RBRACKET        = "]"                                      // Close              | Implemented
	OROR            = "||"                                     // Condition or or    | Implemented
	ANDAND          = "&&"                                     // Boolean operator   | Implemented
	BACKTICK        = "`"                                      // Backtick           | Implemented
	POWEROF         = "**"                                     // General operator   | Implemented
	MODULO          = "%"                                      // General operator   | Implemented
	DEFAULT         = "DEFAULT"                                // Default keyword    | Implemented
	NEWLINE         = '\n'                                     // COND               | Implemented
	IMPORT          = "IMPORT"                                 // Import             | Implemented
	MODULE          = "module"                                 // Module             | Implemented
	KEYWORD_ENGINE  = "ENGINE"                                 // ENGINE             | Implemented
	ENGINE_TYPE     = "ENGINE::ENVIRONMENT_MODIFIER->CALL:::>" // ENGINE ENV MODIFY  | Implemented
	REGEXP          = "REGEXP"                                 // Regex Type         | Not implemented
	PLUS_PLUS       = "++"                                     // Plus Plus          | Not implemented
	QUESTION        = "?"                                      // Question que       | Not implemented
	DOTDOT          = ".."                                     // Range              | Not implemented
	CONTAINS        = "~="                                     // Contains           | Not implemented
	NOTCONTAIN      = "!~"                                     // Boolean operator   | Not implemented
	MINUS_MINUS     = "--"                                     // Minus minus        | Not Implemented

)

var keywords = map[string]Token_Type{
	"Func":     FUNCTION,       // Function
	"function": FUNCTION,       // Function
	"let":      LET,            // Variable declaration let
	"set":      LET,            // Variable declaration set
	"cause":    LET,            // Variable declaration cause
	"allow":    LET,            // Variable declaration allow
	"true":     TRUE,           // Boolean true
	"false":    FALSE,          // Boolean false
	"if":       IF,             // Conditional start
	"else":     ELSE,           // Conditional alternative
	"return":   RETURN,         // Return decl
	"ret":      RETURN,         // Return decl
	"const":    CONSTANT,       // Constant type
	"constant": CONSTANT,       // Constant type
	"switch":   SWITCH,         // Switch statement
	"sw":       SWITCH,         // Switch statement
	"case":     CASE,           // Case statement
	"cs":       CASE,           // Case statement
	"default":  DEFAULT,        // Switch alternative
	"df":       DEFAULT,        // Switch alternative
	"register": REGISTER,       // Register
	"ENGINE":   KEYWORD_ENGINE, // Engine caller
	"import":   IMPORT,         // Import data
	"for":      STD_FOR,        // For loop
	"STRING":   STRING,         // STRING data type
	"BOOLEANT": TRUE,           // Boolean
	"BOOLEANF": FALSE,          // Boolean

}

type Token_Type string

type Token struct {
	Token_Type Token_Type
	Literal    string
}

///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//
//
//																 ┏━┓
//																┃┃ ┃
//															━━━━┛
//
// This section is seperated to be specific for the SkyLine scanner or lexer, these structures also help implement and keep track of current tokens being parsed
//
//

type ScannerInterface interface {
	NT() Token
}

type ScannerStructure struct {
	CharInput string
	POS       int
	RPOS      int
	Char      byte
	Chars     []rune
	PrevTok   Token
	Prevch    byte
	CurLine   int // Current line
}

///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//
//
//																 ┏━┓
//																┃┃ ┃
//															━━━━┛
//
// This section of the models file is segmented to be specific for the skyline objects definitions. These structures and constant values help define the types and objects
//
// given that skyline is completely object orritented. This also defines all of the data type and keyword structures for the Objects such as switch, case, if, else, elif, func etc
//
type Type_Object string

// Data types
const (
	ImportingType   = "Import"
	IntegerType     = "Integer"     // integger
	FloatType       = "Float"       // float integer
	BooleanType     = "Boolean"     // bool true or bool false
	NilType         = "Nil"         // null
	ReturnValueType = "ReturnValue" // return
	ErrorType       = "Error"       // error
	FunctionType    = "Function"    // function
	StringType      = "String"      // string
	BuiltinType     = "Builtin"     // built in
	ArrayType       = "Array"       // array
	HashType        = "Hash"        // hash
	RegisterType    = "Registry"    // Register backend
)

type Switch struct {
	Token   Token
	Value   Expression
	Choices []*Case
}

type Case struct {
	Token Token
	Def   bool
	Expr  []Expression
	Block *BlockStatement
}

type SLC_Object interface {
	SL_RetrieveDataType() Type_Object
	SL_InspectObject() string
	InvokeMethod(method string, Env Environment_of_environment, args ...SLC_Object) SLC_Object
	ToInterface() interface{}
}

type HashKey struct {
	Type_Object Type_Object
	Value       uint64
}

type Hashable interface {
	HashKey() HashKey
}

type Integer struct {
	Value int64
}

type Float struct {
	Value float64
}

type Boolean_Object struct {
	Value bool
}

type String struct {
	Value  string
	Offset int
}

// Offset for string reset
func (str *String) Reset() {
	str.Offset = 0
}

type Nil struct{}

type ReturnValue struct {
	Value SLC_Object
}

type Error struct {
	Message string
}

type Function struct {
	Parameters []*Ident
	Body       *BlockStatement
	Env        *Environment_of_environment
}

type BuiltinFunction func(env *Environment_of_environment, args ...SLC_Object) SLC_Object

type Constant struct {
	Token Token
	Name  *Ident
	Value Expression
}

type Builtin struct {
	Fn BuiltinFunction
}

type Array struct {
	Elements []SLC_Object
	Offset   int
}

// Offset reset for array structure
func (array *Array) Reset() {
	array.Offset = 0
}

type HashPair struct {
	Key   SLC_Object
	Value SLC_Object
}

type Hash struct {
	Pairs  map[HashKey]HashPair
	Offset int
}

// Offset for hash reset
func (hash *Hash) Reset() {
	hash.Offset = 0
}

///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//
//
//																 ┏━┓
//																┃┃ ┃
//															━━━━┛
//
// This section defines the AST generat functions and interfaces, this allows us to create nodes, statements, expressions and other various types for the AST. This also defines
//
// functions and data type structures such as statement and expression type structures which the AST uses.
//
var U UserInterpretData

type Node interface {
	SL_ExtractNodeValue() string
	SL_ExtractStringValue() string
}

type Statement interface {
	Node
	SN()
}

type Expression interface {
	Node
	EN()
}

type Program struct {
	Statements []Statement
}

type LetStatement struct {
	Token Token
	Name  *Ident
	Value Expression
}

type AssignmentStatement struct {
	Token    Token
	Name     *Ident
	Value    Expression
	Operator string
}

type Ident struct {
	Token Token
	Value string
}

type ReturnStatement struct {
	Token       Token
	ReturnValue Expression
}

type ExpressionStatement struct {
	Token      Token
	Expression Expression
}

type IntegerLiteral struct {
	Token Token
	Value int64
}

type FloatLiteral struct {
	Token Token
	Value float64
}

type PrefixExpression struct {
	Token    Token
	Operator string
	Right    Expression
}

type InfixExpression struct {
	Token    Token
	Left     Expression
	Operator string
	Right    Expression
}

type Boolean_AST struct {
	Token Token
	Value bool
}

type ConditionalExpression struct {
	Token       Token
	Condition   Expression
	Consequence *BlockStatement
	Alternative *BlockStatement
}

type BlockStatement struct {
	Token      Token
	Statements []Statement
}

type FunctionLiteral struct {
	Token      Token
	Parameters []*Ident
	Body       *BlockStatement
}

type CallExpression struct {
	Token     Token      // the '(' token
	Function  Expression // Ident or FunctionLiteral
	Arguments []Expression
}

type StringLiteral struct {
	Token Token
	Value string
}

type ArrayLiteral struct {
	Token    Token // the '[' token
	Elements []Expression
}

type IndexExpression struct {
	Token Token // the '[' token
	Left  Expression
	Index Expression
}

type HashLiteral struct {
	Token Token // the '{' token
	Pairs map[Expression]Expression
}

type RegisterValue struct {
	Value SLC_Object
}

type Register struct {
	Token         Token
	RegistryValue Expression
}

type ENGINE_Value struct {
	Value SLC_Object
}

type ENGINE struct {
	Token       Token
	EngineValue Expression
}

type EngineCallValues struct {
	Name        string
	Version     string
	Require     []string
	Languages   string
	Description string
	SOS         string
	Prepped     bool // Engine has parsed data
}

type Module struct {
	Name  string
	Attrs SLC_Object
}

type ImportExpression struct {
	Token Token
	Name  Expression
}

///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//
//
//																 ┏━┓
//																┃┃ ┃
//															━━━━┛
//
// This section defines the environment structures for the skyline environment
//
//
//
type Environment_of_environment struct {
	Store  map[string]SLC_Object
	Outer  *Environment_of_environment
	ROM    map[string]bool // Read Only mode for constants
	permit []string
}

///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//
//
//																 ┏━┓
//																┃┃ ┃
//															━━━━┛
//
// This section defines the most basic evaluation models which are used
//
//
//
var (
	NilValue   = &Nil{}
	TrueValue  = &Boolean_Object{Value: true}
	FalseValue = &Boolean_Object{Value: false}
)

type ObjectCallExpression struct {
	Token      Token
	SLC_Object Expression
	Call       Expression
}

var FileCurrent FileCurrentWithinParserEnvironment

///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//
//
//																 ┏━┓
//																┃┃ ┃
//															━━━━┛
//
// This section defines all of the parser functions, types, constants, variables and maps which are used to communicate with other functions within Skyline. This also allows it to
//
// work with the tokens, types, prefix's, infix's and other various statements and expressions.
//
const (
	_ int = iota
	LOWEST
	ASSIGNMENT
	EQUALS
	LESSGREATER
	SUM
	PRODUCT
	PREFIX
	CALL
	INDEX
	GTEQS
	LTEQS
	POWER
	MOD
	DOT_DOT
	REGEXP_MATCH
	TERNARY
	COND
)

var Precedences = map[Token_Type]int{
	ANDAND:          COND,
	OROR:            COND,
	EQ:              EQUALS,
	GTEQ:            GTEQS,
	LTEQ:            LTEQS,
	NEQ:             EQUALS,
	LT:              LESSGREATER,
	GT:              LESSGREATER,
	PLUS:            SUM,
	PLUS_EQUALS:     SUM,
	MINUS:           SUM,
	MINUS_EQUALS:    SUM,
	SLASH:           PRODUCT,
	ASSIGN:          ASSIGNMENT,
	POWEROF:         POWER,
	QUESTION:        TERNARY,
	ASTARISK:        PRODUCT,
	ASTERISK_EQUALS: PRODUCT,
	DIVEQ:           PRODUCT,
	LPAREN:          CALL,
	PERIOD:          CALL,
	NOTCONTAIN:      REGEXP_MATCH,
	CONTAINS:        REGEXP_MATCH,
	DOTDOT:          DOT_DOT,
	LBRACKET:        INDEX,
	MODULECALL:      INDEX,
}

type (
	PrefixParseFn  func() Expression
	InfixParseFn   func(Expression) Expression
	PostfixParseFn func() Expression
)

type Parser struct {
	Lex             *ScannerStructure
	Errors          []string
	PreviousToken   Token
	CurrentToken    Token
	PeekToken       Token
	PrefixParseFns  map[Token_Type]PrefixParseFn
	InfixParseFns   map[Token_Type]InfixParseFn
	PostfixParseFns map[Token_Type]PostfixParseFn
}

//////////////////////////////////////////////////////////////////////////////////////////////////
////////////////////////// ENVIRONMENT MODIFIER
//////////////////////////////////////////////////////////////////////////////////////////////////

type Iterable interface {
	Reset()
	Next() (SLC_Object, SLC_Object, bool)
}

func (env *Environment_of_environment) Names(prefix string) []string {
	var ret []string
	for key := range env.Store {
		if strings.HasPrefix(key, prefix) {
			ret = append(ret, key)
		}
		if strings.HasPrefix(key, "object.") {
			ret = append(ret, key)
		}
	}
	return ret
}

///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//
//
//																 ┏━┓
//																┃┃ ┃
//															━━━━┛
//
// This section defines TOINTERFACE functions which are used to take a given key or type structure and convert it to an interface type.
//
//
//

func (ENGINE_VAL *ENGINE_Value) ToInterface() interface{} {
	return "<ENGINE_MODIFICATION : SkyLine External Environment Modifier>"
}

func (Arr *Array) ToInterface() interface{} {
	return "<ARRAY>"
}

func (hash *Hash) ToInterface() interface{} {
	return "<HASH>"
}

func (null *Nil) ToInterface() interface{} {
	return "<NULL>"
}

func (Func *Function) ToInterface() interface{} {
	return "<FUNCTION>"
}

func (Err *Error) ToInterface() interface{} {
	return "<ERROR>"
}

func (Mod *Module) ToInterface() interface{} {
	return "<Module>"
}

func (BuiltIn *Builtin) ToInterface() interface{} {
	return "<BUILT-IN-FUNCTION>"
}

func (Ret *ReturnValue) ToInterface() interface{} {
	return "<RETURN_VALUE>"
}

func (float *Float) ToInterface() interface{} {
	return float.Value
}

func (Int *Integer) ToInterface() interface{} {
	return Int.Value
}

func (Str *String) ToInterface() interface{} {
	return Str.Value
}

func (Boolean *Boolean_Object) ToInterface() interface{} {
	return Boolean.Value
}

func (RegisterValue *RegisterValue) ToInterface() interface{} {
	return "<REGISTRY>"
}

///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//
//
//																 ┏━┓
//																┃┃ ┃
//															━━━━┛
//
// The NEXT functions which is the next brick of code which defines functions to check and itterate over specfic elements or values of a specific data type that can be itterated
//
// over such as String, Hashes or arrays.
//
//

func (Arr *Array) Next() (SLC_Object, SLC_Object, bool) {
	if Arr.Offset < len(Arr.Elements) {
		Arr.Offset++
		element := Arr.Elements[Arr.Offset-1]
		return element, &Integer{Value: int64(Arr.Offset - 1)}, true
	}
	return nil, &Integer{Value: 0}, false
}

func (hash *Hash) Next() (SLC_Object, SLC_Object, bool) {
	if hash.Offset < len(hash.Pairs) {
		idx := 0
		for _, pair := range hash.Pairs {
			if hash.Offset == idx {
				hash.Offset++
				return pair.Key, pair.Value, true
			}
			idx++
		}
	}
	return nil, &Integer{Value: 0}, false
}

func (str *String) Next() (SLC_Object, SLC_Object, bool) {
	if str.Offset < utf8.RuneCountInString(str.Value) {
		str.Offset++
		chars := []rune(str.Value)
		value := &String{Value: string(chars[str.Offset-1])}
		return value, &Integer{Value: int64(str.Offset - 1)}, true
	}
	return nil, &Integer{Value: 0}, false
}

///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//
//
//																 ┏━┓
//																┃┃ ┃
//															━━━━┛
//
// The following code unerneath each of these sections defines all of the identifiers and libraries currently avalible for the SkyLine programming language. These maps and variables
//
// are used in the context of register which allows you to `register` all of the functions from one standard library into the environment of the programming language or the environment
//
// which the interpreter or SL engine has started. These also verify that they are able to be imported and are not currently unavalible.
//

var ConstantIdents = map[string]bool{
	"math.abs":                         true,
	"math.cos":                         true,
	"math.sin":                         true,
	"math.sqrt":                        true,
	"math.tan":                         true,
	"math.cbrt":                        true,
	"math.rand":                        true,
	"math.out":                         true,
	"crypt.hash":                       true,
	"io.input":                         true,
	"io.clear":                         true,
	"io.box":                           true,
	"io.restore":                       true,
	"io.listen":                        true,
	"forensics.new":                    true,
	"forensics.meta":                   true,
	"forensics.PngSettingsNew":         true,
	"forensics.InjectPNG":              true,
	"forensics.InjectImage":            true,
	"forensics.FindSigUnknownFile":     true,
	"forensics.CheckZIPSig":            true,
	"forensics.CreationNew":            true,
	"forensics.InjectRegular":          true,
	"forensics.CreateImage":            true,
	"http.Get":                         true,
	"http.MethodGet":                   true,
	"http.MethodDelete":                true,
	"http.MethodPut":                   true,
	"http.MethodPatch":                 true,
	"http.MethodOptions":               true,
	"http.MethodHead":                  true,
	"http.MethodPost":                  true,
	"http.MethodTrace":                 true,
	"env.getenv":                       true,
	"env.setenv":                       true,
	"env.environment":                  true,
	"Google.CastDeviceInfo":            true,
	"Google.CastDeviceReboot":          true,
	"Google.CastDeviceDescription":     true,
	"Google.CastDeviceWiFiForget":      true,
	"Google.CastDeviceWiFiScan":        true,
	"Google.CastDeviceWiFiScanResults": true,
	"Google.CastDeviceWiFiConfigured":  true,
	"Google.CastDeviceAlarms":          true,
	"Google.CastDeviceTimezones":       true,
	"Google.CastDeviceLegacyConf":      true,
	"Google.CastDeviceBleStat":         true,
	"Google.CastDeviceBlePaired":       true,
	"Google.CastDeviceSetName":         true,
	"Google.CastDeviceApplications":    true,
	"Roku.KeyPressHome":                true,
	"Roku.KeyPressPlay":                true,
	"Roku.KeyPressUp":                  true,
	"Roku.KeyPressDown":                true,
	"Roku.KeyPressLeft":                true,
	"Roku.KeyPressRight":               true,
	"Roku.KeyPressSelect":              true,
	"Roku.KeyPressRewind":              true,
	"Roku.KeyPressFFW":                 true,
	"Roku.KeyPressOptions":             true,
	"Roku.KeyPressPause":               true,
	"Roku.KeyPressBack":                true,
	"Roku.KeyPressPoweroff":            true,
	"Roku.KeyPressVolumeUp":            true,
	"Roku.KeyPressVolumeDown":          true,
	"Roku.KeyPressVolumeMute":          true,
	"Roku.KeyPressDeviceDown":          true,
	"Roku.KeyPressDeviceUp":            true,
	"Roku.KeyPressDevAppLaunch":        true,
	"Roku.KeyPressDevAppInstall":       true,
	"Roku.KeyPressDevDisableSGR":       true,
	"Roku.KeyPressDevEnableSGR":        true,
	"Roku.KeyPressDevTV":               true,
	"Roku.KeyPressDevSGNODES":          true,
	"Roku.KeyPressDevActiveTVS":        true,
	"Roku.KeyPressDevDial":             true,
	"Roku.DeviceBrowse":                true,
	"Roku.DeviceInformation":           true,
	"Roku.DeviceApplications":          true,
	"Roku.DeviceActiveApplications":    true,
	"Apple.AirPlayServerInfo":          true,
	"Apple.AirPlayPlayBackInfo":        true,
	"Apple.AirPlayScrubInfo":           true,
	"Apple.AirPlayStreamInfo":          true,
	"Apple.AirPlayInfo":                true,
	"Apple.DAAPMain":                   true,
	"Apple.DAAPLogin":                  true,
	"Apple.DAAPDatabase":               true,
	"Amazon.TvDeviceInformation":       true,
	"Amazon.TvDeviceDescription":       true,
}

var StandardLibNames = map[string]bool{
	"math":               true,
	"io":                 true,
	"forensics":          true,
	"os":                 true,
	"http":               true,
	"crypt":              true,
	"IoT/Apple/Database": true,
	"IoT/Roku/Database":  true,
	"json":               true,
	"net":                true,
	"xml":                true,
	"env":                true,
}

/*

 */

var RegisterStandard = map[string]func(){
	"io":                  RegisterIO,
	"math":                RegisterMath,
	"crypt":               RegisterCrypt,
	"forensics":           RegisterForensics,
	"http":                RegisterHTTP,
	"env":                 RegisterEnvironment,
	"IoT/Apple/Database":  RegisterAppleIoTDatabase,
	"IoT/Roku/Database":   RegisterRokuIoTDatabase,
	"IoT/Google/Database": RegisterGoogleIoTDatabase,
	"IoT/Amazon/Database": RegisterAmazonIoTDatabase,
}

var Datatypes = []string{
	"string.",
	"float.",
	"object.",
	"hash.",
	"array.",
}

var Static_String_Methods = []string{
	"length",         // Length of string
	"methods",        // Methods
	"ord",            // ord
	"to_i",           // to integer
	"to_f",           // to float
	"to_b",           // To byte
	"upper",          // To uppercase
	"lower",          // To lowercase
	"title",          // To title
	"split",          // Split
	"split_after",    // Split after
	"split_after_n",  // Split after n
	"trim",           // Trim
	"trim_space",     // Trim any spaces
	"UnlinkRegistry", // Unlink a libraries registration | This is a string because we unlink a library name

}

var Static_Integer_Methods = []string{
	"chr",
	"methods", // Grab all methods
}

var Static_Float_Methods = []string{
	"methods", // Grab all methods
}

var Static_Array_Methods = []string{
	"Reverse", // Reverse
	"Append",  // Append
	"Copy",    // Copy
	"Swap",    // Swap
	"Less",    // Less
	"Compare", // Compare
	"Typeof",  // Typeof
	"popR",    // Pop right
	"popL",    // Pop left
	"length",  // Length of the array
	"methods", // Grab all methods
}

var Static_Hash_Methods = []string{
	"keys",    // Dump all the hash's keys
	"methods", // Grab all methods
}

var Static_Boolean_Methods = []string{
	"methods", // Grab all methods
}

var Static_BuiltInFunction_Methods = []string{
	"methods", // Grab all methods
}

// All functions for each library
var MathLibFunctionsRegistered = []string{
	"math.tan",
	"math.cos",
	"math.sin",
	"math.rand",
	"math.abs",
	"math.cbrt",
	"math.sqrt",
}

var HTTPLibFunctionsRegistered = []string{
	"http.Get",
}

var IOLibFunctionsRegistered = []string{
	"io.box",
	"io.listen",
	"io.restore",
	"io.clear",
	"io.input",
}

var CryptLibFunctionsRegistered = []string{
	"crypt.hash",
}

var ForensicsFunctionsRegistered = []string{
	"forensics.new",
	"forensics.meta",
	"forensics.PngSettingsNew",
	"forensics.InjectPNG",
	"forensics.InjectImage",
	"forensics.FindSigUnknownFile",
	"forensics.CheckZIPSig",
	"forensics.CreationNew",
	"forensics.InjectRegular",
	"forensics.CreateImage",
}

var EnvironmentFunctionsRegistered = []string{
	"env.setenv",
	"env.getenv",
	"env.environment",
}

type UserInterpretData struct {
	OperatingSystem             string
	OperatingSystemArchitecture string
}

///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//
//
//																 ┏━┓
//																┃┃ ┃
//															━━━━┛
//
// This section defines all of the script information which works with the variables, the flags, command line details, frontend work, art work etc etc. These are all really
//
// variables to store flags, version numbers or art.
//

var (
	Version       = "0.0.5"
	Help          = flag.Bool("help", false, "Load help module")
	ErrorsTrace   = flag.Bool("trace", false, "Load tracer module for errors, or if script empty output panic or recovery")
	SourceFile    = flag.String("source", "", "Load source code file into SkyLine")
	Bnn           = flag.Bool("bout", false, "If true will output the SkyLine banner when running a code file")
	Server        = flag.Bool("server", false, "If true will load the SkyLine local server")
	CompileWithGo = flag.Bool("build", false, "Compile with the interpreter but rather take the input of a source code file and compile it with the embedded interpreter")
	RunRawC       = flag.String("e", "", "Run code without a file and without the REPL ( Read Eval Print Loop )")
)

var Stars = `
				*      *
			   * 		*
					*
				*				*
			*	     *
						*

				*			*
						*
`

type ScriptSettings struct {
	SL_Flag_OutputProg   bool // Output the program or script in a box format
	SL_Flag_EngineOutput bool // Tell the SL engine to check to parse the program to see if it works
	BannerOutput         bool // Output banner
	Verbose              bool // Show verbose debugging output
	Engine               bool // Run the engine normally without configuration
}

///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//
//
//																 ┏━┓
//																┃┃ ┃
//															━━━━┛
//
//
// This code section defines all models for the file system within SkyLine. Of course when using file system in this context we are not talking about direct FS (/, and \\) rather
//
// a system to manage files, allowed file types, file scanning, signature scanning, file verification, integrity checks and much more.
//

type FileCurrentWithinParserEnvironment struct {
	Filename      string
	FileLocation  string
	FileExtension string
	FileBasename  string
	IsDir         bool
}

///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//
//
//																 ┏━┓
//																┃┃ ┃
//															━━━━┛
//
//
// This section defines the colors and the specific hex code for the frontend, this is also in a seperate section due to it being colors and a constant list.
//
//
//

const (
	SUNRISE_HIGH_DEFINITION         = "\033[38;5;214m"
	SUNRISE_LIGHT_DEFINITION        = "\033[38;5;215m"
	SKYLINE_HIGH_DEFPURPLE          = "\033[38;5;57m"
	SKYLINE_HIGH_DEFAQUA            = "\033[38;5;51m"
	SKYLINE_HIGH_DEFRED             = "\033[38;5;196m"
	SKYLINE_SUNRISE_HIGH_DEF_YELLOW = "\033[38;5;190m"
	SKYLINE_HIGH_DEFWARN            = "\033[38;5;213m"
	SKYLINE_HIGH_FIXBLUE            = "\033[38;5;121m"
	SKYLINE_SICK_BLUE               = "\033[38;5;81m"
	SKYLINE_HIGH_RES_VIS_RED        = "\033[38;5;196m"
	SKYLINE_HIGH_RES_VIS_BLUE       = "\033[38;5;122m"
	SKYLINE_HIGH_RES_VIS_SUNSET     = "\033[38;5;217m"
	SKYLINE_RESTORE                 = "\033[39m"
)
