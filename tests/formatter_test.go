package formatter_test

import (
	"bytes"
	"os"
	"regexp"
	"testing"

	formatter "github.com/mobigen/gologger"
	"github.com/sirupsen/logrus"
)

func ExampleFormatter_Format_default() {
	l := logrus.New()
	l.SetOutput(os.Stdout)
	l.SetLevel(logrus.DebugLevel)
	l.SetFormatter(&formatter.Formatter{
		TimestampFormat: "-",
	})

	l.Debug("test1")
	l.Info("test2")
	l.Warn("test3")
	l.Error("test4")

	// Output:
	// - [DEBU] test1
	// - [INFO] test2
	// - [WARN] test3
	// - [ERRO] test4
}

func ExampleFormatter_Format_full_level() {
	l := logrus.New()
	l.SetOutput(os.Stdout)
	l.SetLevel(logrus.DebugLevel)
	l.SetFormatter(&formatter.Formatter{
		TimestampFormat: "-",
		ShowFullLevel:   true,
	})

	l.Debug("test1")
	l.Info("test2")
	l.Warn("test3")
	l.Error("   test4")

	// Output:
	// - [DEBUG] test1
	// - [INFO] test2
	// - [WARNING] test3
	// - [ERROR]    test4
}
func ExampleFormatter_Format_show_fields() {
	l := logrus.New()
	l.SetOutput(os.Stdout)
	l.SetLevel(logrus.DebugLevel)
	l.SetFormatter(&formatter.Formatter{
		TimestampFormat: "-",
		ShowFields:      true,
	})

	ll := l.WithField("category", "rest")

	l.Info("test1")
	ll.Info("test2")

	// Output:
	// - [INFO] test1
	// - [INFO] [ category:rest ] test2
}

func ExampleFormatter_Format_no_uppercase_level() {
	l := logrus.New()
	l.SetOutput(os.Stdout)
	l.SetLevel(logrus.DebugLevel)
	l.SetFormatter(&formatter.Formatter{
		TimestampFormat:  "-",
		NoUppercaseLevel: true,
		ShowFields:       true,
		SortFields:       true,
	})

	ll := l.WithFields(map[string]interface{}{
		"component": "main",
	})
	lll := l.WithFields(map[string]interface{}{
		"category":  "rest",
		"component": "main",
	})
	llll := l.WithFields(map[string]interface{}{
		"category":  "other",
		"component": "main",
	})

	l.Debug("test1")
	ll.Info("test2")
	lll.Warn("test3")
	llll.Error("test4")

	// Output:
	// - [debu] test1
	// - [info] [ component:main ] test2
	// - [warn] [ category:rest, component:main ] test3
	// - [erro] [ category:other, component:main ] test4
}

func TestFormatter_Format_with_report_caller(t *testing.T) {
	output := bytes.NewBuffer([]byte{})

	l := logrus.New()
	l.SetOutput(output)
	l.SetLevel(logrus.DebugLevel)
	l.SetFormatter(&formatter.Formatter{
		TimestampFormat: "-",
	})
	l.SetReportCaller(true)

	l.Debug("test1")

	line, err := output.ReadString('\n')
	if err != nil {
		t.Errorf("Cannot read log output: %v", err)
	}

	expectedRegExp := "- \\[DEBU\\] \\[.+\\.go.+:.+[0-9]+\\] test1\n$"
	match, err := regexp.MatchString(
		expectedRegExp,
		line,
	)
	if err != nil {
		t.Errorf("Cannot check regexp: %v", err)
	} else if !match {
		t.Errorf(
			"logger.SetReportCaller(true) output doesn't match, expected: %s to find in: '%s'",
			expectedRegExp,
			line,
		)
	}
}
