package validators

import "gohub/pkg/verifycode"

func ValidatePasswordConfirm(password, passwordConfirm string, errs map[string][]string) map[string][]string {
	if password != passwordConfirm {
		errs["password_confirm"] = append(errs["password_confirm"], "Passwords do not match.")
	}
	return errs
}

func ValidateVerifyCode(key, answer string, errs map[string][]string) map[string][]string {
	if ok := verifycode.NewVerifyCode().CheckAnswer(key, answer); !ok {
		errs[key] = append(errs[key], "The verify code is invalid.")
	}
	return errs
}
