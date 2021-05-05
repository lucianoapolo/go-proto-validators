package plugin

const (
	LangPtBr    = "pt_br"
	LangDefault = "default"
)

var errorIntGt = map[string]string{
	LangPtBr:    `ser maior que '%d'`,
	LangDefault: `be greater than '%d'`,
}

var errorIntLt = map[string]string{
	LangPtBr:    `ser menor que '%d'`,
	LangDefault: `be less than '%d'`,
}

var errorIntGte = map[string]string{
	LangPtBr:    `ser maior ou igual que '%d'`,
	LangDefault: `be greater or equal than '%d'`,
}

var errorIntLte = map[string]string{
	LangPtBr:    `ser menor ou igual que '%d'`,
	LangDefault: `be less or equal than '%d'`,
}

var errorIsInEnum = map[string]string{
	LangPtBr:    "ser um válido %s enumerador",
	LangDefault: "be a valid %s enumerator",
}

var errorLengthGt = map[string]string{
	LangPtBr:    `ter um comprimento maior que '%d'`,
	LangDefault: `have a length greater than '%d'`,
}

var errorLengthLt = map[string]string{
	LangPtBr:    `ter um comprimento menor que '%d'`,
	LangDefault: `have a length smaller than '%d'`,
}

var errorLengthEq = map[string]string{
	LangPtBr:    `ter um comprimento igual que '%d'`,
	LangDefault: `have a length equal than '%d'`,
}

var errorFloatGt = map[string]string{
	LangPtBr:    `ser estritamente maior que '%.2f'`,
	LangDefault: `be strictly greater than '%.2f'`,
}

var errorFloatGtEpsilon = map[string]string{
	LangPtBr:    ` com uma tolerância de '%.2f'`,
	LangDefault: ` with a tolerance of '%.2f'`,
}

var errorFloatGte = map[string]string{
	LangPtBr:    `ser maior ou igual a '%.2f'`,
	LangDefault: `be greater than or equal to '%.2f'`,
}

var errorFloatLt = map[string]string{
	LangPtBr:    `ser estritamente menor que '%.2f'`,
	LangDefault: `be strictly lower than '%.2f'`,
}

var errorFloatLtEpsilon = map[string]string{
	LangPtBr:    ` com uma tolerância de '%.2f'`,
	LangDefault: ` with a tolerance of '%.2f'`,
}

var errorFloatLte = map[string]string{
	LangPtBr:    `ser menor ou igual a '%.2f'`,
	LangDefault: `be lower than or equal to '%.2f'`,
}

var errorRegex = map[string]string{
	LangPtBr:    "estar em conformidade com regex ",
	LangDefault: "be a string conforming to regex ",
}

var errorStringNotEmpty = map[string]string{
	LangPtBr:    "deve ser preenchido",
	LangDefault: "must not be an empty string",
}

var errorTrimmedStringNotEmpty = map[string]string{
	LangPtBr:    "deve ser preenchido",
	LangDefault: "must not be an empty string",
}

var errorRepeatedCountMin = map[string]string{
	LangPtBr:    `conter pelo menos %v elementos`,
	LangDefault: `contain at least %v elements`,
}

var errorRepeatedCountMax = map[string]string{
	LangPtBr:    `conter no máximo %v elementos`,
	LangDefault: `contain at most %v elements`,
}

var errorMsgExists = map[string]string{
	LangPtBr:    `os dados devem ser preenchidos`,
	LangDefault: `message must exist`,
}

var errorMsgExistsIfAnotherNot = map[string]string{
	LangPtBr:    `os dados devem ser preenchidos se %s estiver vazio`,
	LangDefault: `message must exist if message %s is not exists`,
}

var errorString = map[string]string{
	LangPtBr:    `valor '%v' deve `,
	LangDefault: `value '%v' must `,
}

var errorOneofValidator = map[string]string{
	LangPtBr:    "um dos campos deve ser definido",
	LangDefault: "one of the fields must be set",
}
