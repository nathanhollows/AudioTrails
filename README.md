[![DOI](https://zenodo.org/badge/415827273.svg)](https://zenodo.org/badge/latestdoi/415827273)

<p align="center">
  <a href="https://github.com/nathanhollows/AudioTrails">
    <img src="web/static/apple-touch-icon.png" alt="Logo" width="80" height="80">
  </a>

  <h3 align="center">QR Audio Trails</h3>

  <p align="center">
    <br />
    <a href="https://geo.trace.co.nz">View Demo</a>
    ·
    <a href="https://github.com/nathanhollows/AudioTrails/issues">Report Bug</a>
    ·
    <a href="https://github.com/nathanhollows/AudioTrails/issues">Request Feature</a>
  </p>
</p>



<!-- TABLE OF CONTENTS -->
## Table of Contents

* [About the Project](#about-the-project)
  * [Built With](#built-with)
* [Getting Started](#getting-started)
  * [Prerequisites](#prerequisites)
  * [Installation](#installation)
* [Roadmap](#roadmap)
* [Contributing](#contributing)
* [License](#license)
* [Contact](#contact)



<!-- ABOUT THE PROJECT -->
## About The Project

QR Audio Trails were built for the [Waitaki White Geopark](https://www.whitestonegeopark.nz/) as part of a [Centre for Science Communication](https://www.otago.ac.nz/science-communication/index.html) paper at the University of Otago.

This project was built in two parts. The first part is a set of audio stories (you can listen to these in the demo). The second is an online platform that enables visitors to each geosite to listen to the stories. The website presents itself as a sort of scavenger hunt, allowing visitors to explore the area to collect each site. Once visited, visitors may listen to past stories again.

Required features:

- Allow visitors to listen to audio
- Be accessible through QR code (their signage already has QR codes)
- Respect users data limits (resize images, compress audio etc.)
- Front end requirements
    - Link to Geopark website
    - Link to Geopark app
    - Show a unique page per site
    - Mimic the geopark website in terms of style
    - Show the visitor how many sites they have discovered
- Admin panel
    - Add new geosites
    - See how many people have scanned (my requirement: non-identifiable analytics)
    - Upload audio / cover images within 2 minutes
    - Sites should update live

Additional features:

- Tracked redirecting QR codes for printed media

### Built With

* [Go 1.16](https://golang.org)\
* [Wavesurfer.js](https://wavesurfer-js.org/)
* [Pico CSS](https://picocss.com/)
* [ImageMagick](https://imagemagick.org/index.php)

<!-- GETTING STARTED -->
## Getting Started

### Prerequisites

- Go 1.16
- ImageMagick

### Installation

1. Clone the repo
```sh
git clone https://github.com/nathanhollows/AudioTrails.git
```
2. Set environment variables
```sh
export GEOTRACE_SESSION_KEY=""
export GEOTRACE_SITEURL=""
export GEOTRACE_PORT=""
```

`GEOTRACE_SESSION_KEY` is used to encrypt session variables. Use <https://randomkeygen.com/> for generating secure keys. `GEOTRACE_SITEURL` is used to generate internal URLs. `GEOTRACE_PORT` is the port AudioTrails will run on locally.


<!-- ROADMAP -->
## Roadmap

See the [open issues](https://github.com/othneildrew/Best-README-Template/issues) for a list of proposed features (and known issues).

Short summary:

- Graphs in analytics
- Per site analytics


<!-- CONTRIBUTING -->
## Contributing

Contributions are what make the open source community such an amazing place to be learn, inspire, and create. Any contributions you make are **greatly appreciated**.

1. Fork the Project
2. Create your Feature Branch (`git checkout -b feature/AmazingFeature`)
3. Commit your Changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the Branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request


<!-- LICENSE -->
## License

Distributed under the MIT License. See `LICENSE` for more information.


<!-- CONTACT -->
## Contact

Your Name - me@example.nathanhollows.com

Project Link: [https://github.com/nathanhollows/AudioTrails](https://github.com/nathanhollows/AudioTrails)
