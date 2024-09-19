package validation

import (
	"errors"
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/iancoleman/strcase"

	"github.com/hadroncorp/geck/systemerror"
)

var goPlaygroundFormatTypeSet = map[string]struct{}{
	"base64":                  {},
	"base64url":               {},
	"bcp47_language_tag":      {},
	"credit_card":             {},
	"cron":                    {},
	"e164":                    {},
	"email":                   {},
	"datetime":                {},
	"hexadecimal":             {},
	"hexcolor":                {},
	"hsl":                     {},
	"html":                    {},
	"html_encoded":            {},
	"isbn":                    {},
	"isbn10":                  {},
	"isbn13":                  {},
	"iso3166_1_alpha2":        {},
	"iso3166_1_alpha3":        {},
	"iso3166_1_alpha_numeric": {},
	"iso3166_2":               {},
	"iso4217":                 {},
	"json":                    {},
	"jwt":                     {},
	"latitude":                {},
	"longitude":               {},
	"luhn_checksum":           {},
	"rgb":                     {},
	"rgba":                    {},
	"timezone":                {},
	"uuid":                    {},
	"md5":                     {},
	"sha256":                  {},
	"sha512":                  {},
	"semver":                  {},
	"cidr":                    {},
	"cidrv4":                  {},
	"cidrv6":                  {},
	"datauri":                 {},
	"fqdn":                    {},
	"hostname":                {},
	"hostname_port":           {},
	"hostname_rfc1123":        {},
	"ip":                      {},
	"ip4_addr":                {},
	"ip6_addr":                {},
	"ip_addr":                 {},
	"mac":                     {},
	"tcp4_addr":               {},
	"tcp6_addr":               {},
	"tcp_addr":                {},
	"udp4_addr":               {},
	"udp6_addr":               {},
	"udp_addr":                {},
	"unix_addr":               {},
	"uri":                     {},
	"url":                     {},
	"http_url":                {},
	"url_encoded":             {},
	"urn_rfc2141":             {},
	"alpha":                   {},
	"alphanum":                {},
	"alphanumunicode":         {},
	"alphaunicode":            {},
	"ascii":                   {},
	"lowercase":               {},
	"number":                  {},
	"numeric":                 {},
	"uppercase":               {},
	"iscolor":                 {},
	"country_code":            {},
	"date":                    {},
}

var goPlaygroundOutOfRangeSet = map[string]struct{}{
	"min": {},
	"max": {},
	"len": {},
	"gt":  {},
	"gte": {},
	"lt":  {},
	"lte": {},
}

func adapterGoPlaygroundErrors(srcErr error) error {
	var validationErrs validator.ValidationErrors
	ok := errors.As(srcErr, &validationErrs)
	if !ok {
		return srcErr
	}

	errs := make([]error, 0, len(validationErrs))
	for _, err := range validationErrs {
		// using sets to avoid extra computation in switch statement
		_, isFormatErr := goPlaygroundFormatTypeSet[err.Tag()]
		field := err.StructNamespace()
		fieldSplit := strings.SplitN(field, ".", 2)
		if len(fieldSplit) > 1 {
			fieldSplit[1] = strcase.ToSnakeWithIgnore(fieldSplit[1], ".")
			field = strings.Join(fieldSplit, ".")
		}
		if isFormatErr {
			errs = append(errs, systemerror.NewInvalidFormatArgument(field, err.Tag()))
			continue
		}
		_, isOutOfRangeErr := goPlaygroundOutOfRangeSet[err.Tag()]
		if isOutOfRangeErr {
			errs = append(errs, systemerror.NewArgumentOutOfRangeSingle(field, err.Tag(), err.Param()))
			continue
		}

		switch err.Tag() {
		case "required":
			errs = append(errs, systemerror.NewMissingArgument(field))
			continue
		case "eq", "eq_ignore_case":
			errs = append(errs, systemerror.NewNotEqualsArgument(field, err.Param()))
			continue
		case "oneof":
			errs = append(errs, systemerror.NewArgumentNotOneOf(field, strings.Split(err.Param(), " ")...))
			continue
		case "endsnotwith":
			errs = append(errs, systemerror.NewInvalidNoSuffixArgument(field, err.Param()))
			continue
		case "endswith":
			errs = append(errs, systemerror.NewInvalidSuffixArgument(field, err.Param()))
			continue
		case "startsnotwith":
			errs = append(errs, systemerror.NewInvalidNoPrefixArgument(field, err.Param()))
			continue
		case "startswith":
			errs = append(errs, systemerror.NewInvalidPrefixArgument(field, err.Param()))
			continue
		default:
			errs = append(errs, systemerror.NewInvalidArgument(field, err.Param(),
				fmt.Sprintf("%v", err.Value())))
		}
	}
	return errors.Join(errs...)
}
