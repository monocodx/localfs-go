// Copyright 2024 The localFS Authors.
// Use of this source code is governed by a GPL
// license that can be found in the LICENSE file.

package view

type ErrorViewModel struct {
	Code    int
	Status  string
	Message string
}

const ErrorPageTmpl string = `
<!DOCTYPE html>
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
    }
    body {
      margin: auto;
      width: 60%;
      padding: 1rem;
      font-family: monospace;
    }
    div.head {
      display: block;
      padding: 1rem;
    }
    div.body {
      display: block;
      padding: 1rem;
    }
    p {
      color: #90a4ae;
      margin-block-start: 0rem !important;
      margin-block-end: 0rem !important;;
      text-align: center;
    }
    p.code {
      display: block;
      font-size: 6.5rem;
      font-weight: 500;
    }
    p.status {
      display: block;
      font-size: 1rem;
      font-weight: 500;
    }
    p.error {
      display: block;
      font-size: 1.5rem;
      font-weight: 500;
    }
  </style>
</head>
<body>
  <div class="head">
    <p class="code">{{.Code}}</p>
    <p class="status">{{.Status}}</p>
  </div>
  <div class="body">
    <p class="error">Oops! {{.Message}}</p>
  </div>
</body>
</html>
`
