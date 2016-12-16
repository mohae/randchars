// Copyright © 2014, All rights reserved
// Joel Scoble, https://github.com/mohae/clitpl
//
// This is licensed under The MIT License. Please refer to the included
// LICENSE file for more information. If the LICENSE file has not been
// included, please refer to the url above.
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License

// randchars is a cli app that generates random characters.
package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/mohae/randchars"
	"github.com/mohae/randchars/crandchars"
)

var (
	c     bool
	out   = "stdout"
	chars = "base64"
)

func init() {
	flag.StringVar(&out, "-o", out, "output destination")
	flag.StringVar(&chars, "chars", chars, "charset: alphanum, alpha, lalphanum, lalpha, ualphanum, ualpha, base64")
	flag.BoolVar(&c, "c", false, "use a CSPRNG")
}

func main() {
	os.Exit(realMain())
}

func realMain() int {
	flag.Parse()
	args := flag.Args()
	if len(args) == 0 {
		fmt.Fprintln(os.Stderr, "at least 1 number equal to the number of random characters you wish to have generated must be passed")
		return 1
	}
	l := make([]int, len(args))
	var err error
	var n int
	for i, v := range args {
		l[i], err = strconv.Atoi(v)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			return 1
		}
		n += l[i]
	}
	// open output if not stdout
	var f *os.File
	if out == "stdout" {
		f = os.Stdout
	} else {
		f, err = os.OpenFile(out, os.O_CREATE|os.O_TRUNC|os.O_RDWR, 0664)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			return 1
		}
		defer f.Close()
	}
	g, err := NewGenerator(n, c, chars)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", err)
		return 1
	}

	n = 0
	for _, v := range l {
		b := g.GetChars(v)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error generating %d random chars: %s\n", v, err)
			return 1
		}
		i, err := f.Write(b)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error writing %q to %s\n", string(b), out)
			return 1
		}
		n += i
		_, err = f.Write([]byte("\n"))
		if err != nil {
			fmt.Fprintf(os.Stderr, "error writing a newline char to %s\n", out)
			return 1
		}
	}
	fmt.Printf("%d sets totalling %d random characters were generated and written to %s\n", len(l), n, out)
	return 0
}

// Generator handles the generation of random characters
type Generator struct {
	Gen      randchars.Generatorer
	GetChars func(n int) []byte
}

func NewGenerator(n int, c bool, chars string) (*Generator, error) {
	if n <= 0 {
		return nil, fmt.Errorf("%d: invalid character amount; must be > 0", n)
	}
	g := Generator{}
	if c {
		g.Gen = crandchars.NewGenerator(n)
	} else {
		g.Gen = randchars.NewGenerator()
	}
	switch strings.ToLower(chars) {
	case "alphanum":
		g.GetChars = g.Gen.AlphaNum
	case "alpha":
		g.GetChars = g.Gen.Alpha
	case "loweralphanum":
		g.GetChars = g.Gen.LowerAlphaNum
	case "loweralpha":
		g.GetChars = g.Gen.LowerAlpha
	case "upperalphanum":
		g.GetChars = g.Gen.UpperAlphaNum
	case "upperalpha":
		g.GetChars = g.Gen.UpperAlpha
	case "base64":
		g.GetChars = g.Gen.Base64
	default:
		return nil, fmt.Errorf("%q is not supported", chars)
	}
	return &g, nil
}
