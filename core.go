package appCore

type CheckTicketResult struct {
	isValid bool
}

func CheckTicketCode(code string) CheckTicketResult {
	result := new(CheckTicketResult)
	result.isValid = code != ""

	return *result
}

func (res *CheckTicketResult) IsValid() bool {
	return res.isValid
}
