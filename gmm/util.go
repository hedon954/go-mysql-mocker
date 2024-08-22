package gmm

import (
	"net"
	"os"
	"regexp"
	"strings"
	"unicode"
)

// splitSQLFile reads a SQL file and splits it into individual SQL statements.
func splitSQLFile(filePath string) ([]string, error) {
	bs, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	return splitSQLStatements(string(bs))
}

func splitSQLStatements(content string) ([]string, error) {
	// Remove SQL comments
	cleanedContent := removeSQLComments(content)

	// Split statements
	parts := strings.Split(cleanedContent, ";")
	var statements []string

	for _, part := range parts {
		// Clean whitespace and redundant spaces
		cleanedStatement := strings.TrimSpace(part)
		cleanedStatement = regexp.MustCompile(`\s+`).ReplaceAllString(cleanedStatement, " ")

		if cleanedStatement != "" {
			statements = append(statements, cleanedStatement)
		}
	}

	return statements, nil
}

// removeSQLComments removes SQL comments from the input.
func removeSQLComments(input string) string {
	input = cleanString(input)

	// Remove multiline comments
	multilineCommentRegex := regexp.MustCompile(`(?s)/\*.*?\*/`)
	result := multilineCommentRegex.ReplaceAllString(input, "")

	// Remove single line comments (both -- and #)
	singlelineCommentRegex := regexp.MustCompile(`(?m)^\s*(--|#).*\n?`)
	result = singlelineCommentRegex.ReplaceAllString(result, "")

	// Remove inline comments
	inlineCommentRegex := regexp.MustCompile(`--.*$`)
	result = inlineCommentRegex.ReplaceAllString(result, "")

	// Remove extra empty lines
	result = regexp.MustCompile(`(?m)^\s*$[\r\n]*|[\r\n]+\s+\z`).ReplaceAllString(result, "")

	return result
}

// cleanString removes invisible characters and whitespace from a string.
func cleanString(input string) string {
	// Remove ZERO WIDTH NO-BREAK SPACE
	cleaned := strings.ReplaceAll(input, "\uFEFF", "")

	// Remove other invisible Unicode characters
	cleaned = strings.Map(func(r rune) rune {
		if unicode.IsPrint(r) || unicode.IsSpace(r) {
			return r
		}
		return -1
	}, cleaned)

	return cleaned
}

// getFreePort returns a free port on the local machine
func getFreePort() (int, error) {
	listener, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		return 0, err
	}
	port := listener.Addr().(*net.TCPAddr).Port
	return port, listener.Close()
}
