# SlugSpace
Real-time Parking Metrics and Analysis
<br>
---
## About
`SlugSpace` aims to provide U.C. Santa Cruz students real-time parking metrics in an effort to help curb parking issues campus-wide.

## Structure
`SlugSpace` is broken up into pieces in order to expand upon my knowledge of distributed systems, system administration, and full-stack development.

There are scripts found in `\cronjobs` that have further documentation within the folder. 

The `SlugSpaceAPI` is found in `\slugspaceapi` and has further documentation within the folder.

## Installation
`SlugSpace` has multiple moving parts. It is not a `go get` command and have it running. 

1. Run `go get github.com/colbyleiske/slugspace` to grab the project
2. Follow the `cronjobs` documentation and installation in order to have the basic infrastructure setup.
3. Follow the `slugspaceapi` documentation and installation before running the web-app.
4. Not Implemented Yet.

## Dependencies
`gorilla\mux` : [link](https://github.com/gorilla/mux)

## History
`SlugSpace` was first created January 2017 at CruzHacks, UCSC's yearly hackathon by Colby Leiske, Raul Lara, and JJ Serano. It placed first in IoT development, and third in the Google Cloud Platform category.

After the hackathon, `SlugSpace` was invited to the CEID Business competition. Public reception was all good, and every student believed this would benefit the community.

Post-CEID, `SlugSpace` was handed off to me, Colby, for further development as other team members had graduated.

Initially, `SlugSpace` was written in Python + Flask as the initial team had more Python developers. 
As that combo is not my specialty, I opted to rewrite everything in `Go` and use `gorilla\mux` for routing.
