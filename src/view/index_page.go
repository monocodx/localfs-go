// Copyright 2024 The localFS Authors.
// Use of this source code is governed by a GPL
// license that can be found in the LICENSE file.

package view

type IndexPageViewModel struct {
	Base64QRImage string
	Address       string
}

const IndexPageTmpl string = `<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="UTF-8">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <meta http-equiv="X-UA-Compatible" content="ie=edge">
  <title>localFS</title>
  <style>
    @media only screen and (max-width: 480px) {
      body {
        width: 86% !important;
        padding: .85rem !important;
      }
      div.center {
        padding: 1rem !important;
      }
      img {
        width: 200px !important;
        height: 200px !important;
      }
    }
    body {
      margin: auto;
      width: 60%;
      padding: 1.5rem;
      font-weight: 400;
      font-size: 1rem;
      line-height: 1rem;
      font-family: sans-serif;
    }
    div.center {
      display: block;
      padding: 1.5rem;
      text-align: center;
    }
    p {
      color: #607d8b;
      font-weight: 500;
      margin-bottom: 1rem;
      margin-block-start: 0rem !important;
      margin-block-end: 0rem !important;
      line-break: anywhere;
    }
    img {
      width: 240px;
      height: 240px;
    }
  </style>
</head>
<body>
  <div class="center">
    <p>Scan To Upload</p>
    <img src="data:image/png;base64, {{.Base64QRImage}}">
    <p>{{.Address}}</p>
  </div>
</body>
</html>
`
