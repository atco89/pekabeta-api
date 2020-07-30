package exception

import "pekabeta/internal/domain"

func Exception(code string, message ErrorMessage) *domain.Error {
	return &domain.Error{
		Code:    code,
		Message: string(message),
	}
}
