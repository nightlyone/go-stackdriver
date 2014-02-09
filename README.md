go-stackdriver
==============

Inofficial Go client for [stackdriver.com](http://stackdriver.com) API.

NOTE: Still work in progress, just a proof of concept, code heavily in flux. No tests and docs yet!

[![Build Status][1]][2]

[1]: https://secure.travis-ci.org/nightlyone/go-stackdriver.png
[2]: http://travis-ci.org/nightlyone/go-stackdriver


LICENSE
-------
BSD

documentation
-------------
package documentation at [![GoDoc][3]][4]
or [gowalker.org][5]

[3]: https://godoc.org/github.com/nightlyone/go-stackdriver/stackdriver?status.png
[4]: https://godoc.org/github.com/nightlyone/go-stackdriver/stackdriver
[5]: http://gowalker.org/github.com/nightlyone/go-stackdriver/stackdriver

quick usage
-----------
see the example.go for usage.

build and install
=================

install from source
-------------------

Install [Go 1][6], either [from source][7] or [with a prepackaged binary][8].

Then run

	go get github.com/nightlyone/go-stackdriver/stackdriver


Export your stackdriver API key as `STACKDRIVER_API_KEY` environment variable

Run the example:

    go run example.go


[6]: http://golang.org
[7]: http://golang.org/doc/install/source
[8]: http://golang.org/doc/install

LICENSE
-------
BSD

documentation
-------------

contributing
============

Contributions are welcome. Please open an issue or send me a pull request for a dedicated branch.
Make sure the git commit hooks show it works.

git commit hooks
-----------------------
enable commit hooks via

        cd .git ; rm -rf hooks; ln -s ../git-hooks hooks ; cd ..



[![Bitdeli Badge](https://d2weczhvl823v0.cloudfront.net/nightlyone/go-stackdriver/trend.png)](https://bitdeli.com/free "Bitdeli Badge")



[![Bitdeli Badge](https://d2weczhvl823v0.cloudfront.net/nightlyone/go-stackdriver/trend.png)](https://bitdeli.com/free "Bitdeli Badge")

