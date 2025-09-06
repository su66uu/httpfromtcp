package request

import (
	"errors"
	"fmt"
	"io"
	"strings"
)

type RequestLine struct {
	Method        string
	RequestTarget string
	HttpVersion   string
}

type Request struct {
	RequestLine RequestLine
}

var (
	ERROR_MALFORMED_START_LINE     = fmt.Errorf("malformed start line")
	ERROR_UNSUPPORTED_HTTP_VERSION = fmt.Errorf("unsupported HTTP version")
	INCOMPLETE_START_LINE          = fmt.Errorf("incomplete start line")
	ERROR_INVALID_HTTP_METHOD      = fmt.Errorf("invalid HTTP method")
	ERROR_INVALID_HTTP_VERSION     = fmt.Errorf("invalid http version")
	SEPARATOR                      = "\r\n"
)

func parseRequestLine(line string) (*RequestLine, string, error) {
	idx := strings.Index(line, SEPARATOR)
	if idx == -1 {
		return nil, line, nil
	}

	requestLine := line[:idx]
	restOfMsg := line[idx+len(SEPARATOR):]

	parts := strings.Split(requestLine, " ")
	if len(parts) != 3 {
		return nil, line, ERROR_MALFORMED_START_LINE
	}

	httpVersionParts := strings.Split(parts[2], "/")
	if httpVersionParts[0] != "HTTP" || httpVersionParts[1] != "1.1" {
		return nil, restOfMsg, ERROR_INVALID_HTTP_VERSION
	}

	rl := &RequestLine{
		Method:        parts[0],
		RequestTarget: parts[1],
		HttpVersion:   httpVersionParts[1],
	}

	return rl, restOfMsg, nil
}

func RequestFromReader(reader io.Reader) (*Request, error) {
	data, err := io.ReadAll(reader)
	if err != nil {
		return nil, errors.Join(
			fmt.Errorf("unable to io.ReadAll"),
			err,
		)
	}

	str := string(data)
	rl, _, err := parseRequestLine(str)
	if err != nil {
		return nil, err
	}

	return &Request{
		RequestLine: *rl,
	}, err
}
