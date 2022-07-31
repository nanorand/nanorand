<!-- PROJECT SHIELDS -->
[![Go Report Card](https://goreportcard.com/badge/github.com/nanorand/nanorand)](https://goreportcard.com/report/github.com/nanorand/nanorand)
[![GoDoc](https://pkg.go.dev/badge/github.com/nanorand/nanorand?status.svg)](https://pkg.go.dev/github.com/nanorand/nanorand?tab=doc)
[![Go Version](https://img.shields.io/github/go-mod/go-version/nanorand/nanorand)](https://go.dev/)
[![Code Size](https://img.shields.io/github/languages/code-size/nanorand/nanorand)](https://github.com/nanorand/nanorand/blob/master/nanorand.go)
[![Repo License](https://img.shields.io/github/license/nanorand/nanorand)](https://github.com/nanorand/nanorand/blob/master/LICENSE.txt)
[![Release](https://img.shields.io/github/v/release/nanorand/nanorand)](https://github.com/nanorand/nanorand/releases)

<!-- PROJECT LOGO -->
<br />
<div align="center">
  <a href="https://github.com/nanorand/nanorand">
    <img src="https://raw.githubusercontent.com/nanorand/logo/master/logo.svg" alt="Logo" width="200" height="200">
  </a>

<h3 align="center">Nanorand</h3>

  <p align="center">
    Nanosecond pseudo-random generator
    <br />
    <a href="https://github.com/nanorand/nanorand/">View Demo</a>
    ·
    <a href="https://github.com/nanorand/nanorand/issues">Report Bug</a>
    ·
    <a href="https://github.com/nanorand/nanorand/issues">Request Feature</a>
  </p>
</div>



<!-- TABLE OF CONTENTS -->
<details>
  <summary>Table of Contents</summary>
  <ol>
    <li>
      <a href="#about-the-project">About The Project</a>
      <ul>
        <li><a href="#built-with">Built With</a></li>
      </ul>
    </li>
    <li>
      <a href="#getting-started">Getting Started</a>
      <ul>
        <li><a href="#installation">Installation</a></li>
      </ul>
    </li>
    <li><a href="#usage">Usage</a></li>
    <li><a href="#contributing">Contributing</a></li>
    <li><a href="#license">License</a></li>
    <li><a href="#contact">Contact</a></li>
  </ol>
</details>



<!-- ABOUT THE PROJECT -->
## About The Project

[![Nanorand Screen Shot][product-screenshot]](https://raw.githubusercontent.com/nanorand/logo/master/screenshot.png)

There are many great ways to generate "cryptographical-properly" 6-digit codes in Go; however, I didn't find one that really suited my needs. So I created this simplest one. I created a ONE line of code and few functions for creating array and long digit, for example: 127 length. It's so amazing. That be the best for you as soon as you try it -- I think this is it.

Here's why:
* Your time should be focused on creating something amazing. A function GenerateShort for digit up to length of 7 is the faster than other ways, because it use nanosecond of time on your machine. 
* You shouldn't be doing google "How to create 6-digit code" again and again. Just try it and you would love.
* You should implement DRY principles to the rest of your life :smile:

Of course, this project don't suite your purposes if you are creating Pentagon security system. In other way it works great.
Use following instruction to install and use it.

<p align="right">(<a href="#top">back to top</a>)</p>



<!-- GETTING STARTED -->
## Getting Started

This is an example of how you may give instructions on setting up your project locally.
To get a local copy up and running follow these simple example steps.

### Installation

   ```sh
    go get github.com/nanorand/nanorand
   ```

<p align="right">(<a href="#top">back to top</a>)</p>



<!-- USAGE EXAMPLES -->
## Usage

```
package main

import (
	"fmt"
	"github.com/nanorand/nanorand"
)

func main() {
	fmt.Println("6 digit code")

	code, err := nanorand.Gen(6)
	if err != nil {
		return
	}

	fmt.Println(code)
}
```

output:

```
798785
```

<p align="right">(<a href="#top">back to top</a>)</p>



<!-- CONTRIBUTING -->
## Contributing

Contributions are what make the open source community such an amazing place to learn, inspire, and create. Any contributions you make are **greatly appreciated**.

If you have a suggestion that would make this better, please fork the repo and create a pull request. You can also simply open an issue with the tag "enhancement".
Don't forget to give the project a star! Thanks again!

1. Fork the Project
2. Create your Feature Branch (`git checkout -b feature/AmazingFeature`)
3. Commit your Changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the Branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

<p align="right">(<a href="#top">back to top</a>)</p>



<!-- LICENSE -->
## License

Distributed under the MIT License. See `LICENSE.txt` for more information.

<p align="right">(<a href="#top">back to top</a>)</p>



<!-- CONTACT -->
## Contact

Kirill Bogomolov - [@manazoid](https://t.me/manazoid) - uralkir@gmail.com

Project Link: [https://github.com/nanorand/nanorand](https://github.com/nanorand/nanorand)

<p align="right">(<a href="#top">back to top</a>)</p>

<!-- MARKDOWN LINKS & IMAGES -->
<!-- https://www.markdownguide.org/basic-syntax/#reference-style-links -->
[product-screenshot]: https://raw.githubusercontent.com/nanorand/logo/master/screenshot.png
